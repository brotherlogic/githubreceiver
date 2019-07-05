package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbbs "github.com/brotherlogic/buildserver/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/githubreceiver/proto"
	pbgbs "github.com/brotherlogic/gobuildslave/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
)

type github interface {
	add(ctx context.Context, issue *pbgh.Issue) error
	delete(ctx context.Context, issue *pbgh.Issue) error
}

type builder interface {
	build(ctx context.Context, name, fullName string) error
}

type prodGithub struct {
	dial func(server string) (*grpc.ClientConn, error)
}

func (p *prodGithub) add(ctx context.Context, issue *pbgh.Issue) error {
	conn, err := p.dial("githubcard")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.AddIssue(ctx, issue)

	return err

}
func (p *prodGithub) delete(ctx context.Context, issue *pbgh.Issue) error {
	conn, err := p.dial("githubcard")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.DeleteIssue(ctx, &pbgh.DeleteRequest{Issue: issue})

	return err

}

type prodBuilder struct {
	dial func(server string) (*grpc.ClientConn, error)
}

func (p *prodBuilder) build(ctx context.Context, name, fullName string) error {
	conn, err := p.dial("buildserver")
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
	config       *pb.Config
	webhookcount int64
	builder      builder
	github       github
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
		config:   &pb.Config{TimeBetweenQueueProcess: 60},
	}
	s.builder = &prodBuilder{dial: s.DialMaster}
	s.github = &prodGithub{dial: s.DialMaster}
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
	return []*pbg.State{
		&pbg.State{Key: "web_hooks", Value: s.webhookcount},
	}
}

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	s.Log(fmt.Sprintf("Received web ping"))
	s.webhookcount++

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		s.Log(fmt.Sprintf("Error reading body: %v", err))
		return
	}

	s.Log(fmt.Sprintf("%v", string(body)))

	var ping *pb.Ping
	err = json.Unmarshal([]byte(body), &ping)
	if err != nil {
		s.Log(fmt.Sprintf("Error unmarshalling JSON: %v", err))
		return
	}

	ctx, cancel := utils.BuildContext("githubreceiver", "pingprocess")
	defer cancel()
	s.processPing(ctx, ping)
}

func (s *Server) serveUp(port int32) {
	http.HandleFunc("/githubwebhook", s.githubwebhook)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	server.PrepServer()
	server.Register = server
	server.RegisterServer("githubreceiver", false)

	// Handle web requests
	go server.serveUp(server.Registry.Port - 1)

	fmt.Printf("%v", server.Serve())
}
