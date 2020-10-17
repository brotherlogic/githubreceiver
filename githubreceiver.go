package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/goserver/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbbs "github.com/brotherlogic/buildserver/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/githubreceiver/proto"
	pbgbs "github.com/brotherlogic/gobuildslave/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	pbpr "github.com/brotherlogic/pullrequester/proto"
)

type pullRequester interface {
	updatePullRequest(ctx context.Context, url, name, checkName string, pass bool) error
	commitToPullRequest(ctx context.Context, url, sha, name string) error
}

type prodPullRequester struct {
	dial       func(ctx context.Context, server string) (*grpc.ClientConn, error)
	RaiseIssue func(title, body string)
}

func (p *prodPullRequester) updatePullRequest(ctx context.Context, sha, name, checkName string, pass bool) error {
	conn, err := p.dial(ctx, "pullrequester")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbpr.NewPullRequesterServiceClient(conn)
	passV := pbpr.PullRequest_Check_FAIL
	if pass {
		passV = pbpr.PullRequest_Check_PASS
	}
	_, err = client.UpdatePullRequest(ctx, &pbpr.UpdateRequest{Update: &pbpr.PullRequest{Name: name, Shas: []string{sha}, Checks: []*pbpr.PullRequest_Check{&pbpr.PullRequest_Check{Source: checkName, Pass: passV}}}})
	return err
}

type pull struct {
	url  string
	sha  string
	name string
}

func (s *Server) runQueue(ctx context.Context) error {
	if len(s.pqueue) > 0 {
		err := s.pullRequester.commitToPullRequest(ctx, s.pqueue[0].url, s.pqueue[0].sha, s.pqueue[0].name)
		if err == nil {
			s.backends[s.pqueue[0].sha] = 1
			s.pqueue = s.pqueue[1:]
			return nil
		}
		s.pullFails++
	}
	return nil
}

func (p *prodPullRequester) commitToPullRequest(ctx context.Context, url, sha, name string) error {
	conn, err := p.dial(ctx, "pullrequester")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbpr.NewPullRequesterServiceClient(conn)
	_, err = client.UpdatePullRequest(ctx, &pbpr.UpdateRequest{Update: &pbpr.PullRequest{Url: url, Shas: []string{sha}, Name: name}})
	if err != nil {
		p.RaiseIssue("PR problem", fmt.Sprintf("Error creating a pull request: %v", err))
	}
	return err
}

type github interface {
	add(ctx context.Context, issue *pbgh.Issue) error
	delete(ctx context.Context, issue *pbgh.Issue) error
	createPullRequest(ctx context.Context, job, branch, title string) error
}

type builder interface {
	build(ctx context.Context, name, fullName string) error
}

type prodGithub struct {
	dial func(ctx context.Context, server string) (*grpc.ClientConn, error)
}

func (p *prodGithub) add(ctx context.Context, issue *pbgh.Issue) error {
	conn, err := p.dial(ctx, "githubcard")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.AddIssue(ctx, issue)

	return err

}
func (p *prodGithub) delete(ctx context.Context, issue *pbgh.Issue) error {
	conn, err := p.dial(ctx, "githubcard")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.DeleteIssue(ctx, &pbgh.DeleteRequest{Issue: issue})

	return err

}

func (p *prodGithub) createPullRequest(ctx context.Context, job, branch, title string) error {
	conn, err := p.dial(ctx, "githubcard")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.CreatePullRequest(ctx, &pbgh.PullRequest{Job: job, Branch: branch, Title: title})

	return err

}

type prodBuilder struct {
	dial func(ctx context.Context, server string) (*grpc.ClientConn, error)
}

func (p *prodBuilder) build(ctx context.Context, name, fullName string) error {
	conn, err := p.dial(ctx, "buildserver")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbbs.NewBuildServiceClient(conn)
	_, err = client.Build(ctx, &pbbs.BuildRequest{Job: &pbgbs.Job{Name: name, GoPath: "github.com/" + fullName}})

	return err

}

const (
	// KEY - where we store sale info
	KEY = "/github.com/brotherlogic/githubreceiver/config"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	config        *pb.Config
	webhookcount  int64
	webhookfail   int64
	builder       builder
	github        github
	pullRequester pullRequester
	pqueue        []pull
	pullFails     int64
	backends      map[string]int
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
		config:   &pb.Config{TimeBetweenQueueProcess: 60},
		backends: make(map[string]int),
	}
	s.builder = &prodBuilder{dial: s.FDialServer}
	s.github = &prodGithub{dial: s.FDialServer}
	s.pullRequester = &prodPullRequester{dial: s.FDialServer, RaiseIssue: s.RaiseIssue}
	s.pqueue = make([]pull, 0)
	return s
}

func (s *Server) save(ctx context.Context) {
	s.KSclient.Save(ctx, KEY, s.config)
}

func (s *Server) load(ctx context.Context) error {
	config := &pb.Config{}
	data, _, err := s.KSclient.Read(ctx, KEY, config)

	if err != nil {
		return err
	}

	s.config = data.(*pb.Config)
	return nil
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {

}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.save(ctx)
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	if master {
		err := s.load(ctx)
		return err
	}

	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{}
}

var (
	hook = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubreceiver_hook",
		Help: "Push Size",
	}, []string{"type", "error"})
)

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	s.webhookcount++

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if strings.HasPrefix(string(body), "payload") {
		hook.With(prometheus.Labels{"type": "unknown", "error": "payload"}).Inc()
		s.Log(fmt.Sprintf("Payload issue with http body"))
		return
	}

	if err != nil {
		hook.With(prometheus.Labels{"type": "unknown", "error": "bodyread"}).Inc()
		s.Log(fmt.Sprintf("Error reading body: %v", err))
		return
	}

	var ping *pb.Ping
	err = json.Unmarshal([]byte(body), &ping)
	if err != nil {
		hook.With(prometheus.Labels{"type": "unknown", "error": "unmarshal"}).Inc()
		s.Log(fmt.Sprintf("Error unmarshalling JSON: %v -> %v", err, string(body)))
		return
	}

	ctx, cancel := utils.ManualContext("githubreceiver", "pingprocess", time.Minute, true)
	defer cancel()
	err = s.processPing(ctx, ping)
	if err != nil {
		hook.With(prometheus.Labels{"type": "unknown", "error": "ping"}).Inc()
		s.Log(fmt.Sprintf("Error processing ping: %v", err))
		s.webhookfail++
	}
}

func (s *Server) serveUp(port int32) {
	http.HandleFunc("/githubwebhook", s.githubwebhook)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) become(ctx context.Context) error {
	return nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var init = flag.Bool("init", false, "Init the system")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer()
	server.Register = server
	err := server.RegisterServerV2("githubreceiver", false, true)
	if err != nil {
		return
	}

	if *init {
		ctx, cancel := utils.BuildContext("githubreceiver", "githubreceiver")
		defer cancel()
		server.config.Processed++
		server.save(ctx)
		return
	}

	// Handle web requests
	go server.serveUp(server.Registry.Port - 1)

	fmt.Printf("%v", server.Serve())
}
