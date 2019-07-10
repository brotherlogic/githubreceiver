package main

import (
	"testing"

	pbgh "github.com/brotherlogic/githubcard/proto"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/githubreceiver/proto"
)

type testGithub struct {
	issues       int
	pullRequests int
}

type testBuilder struct {
	builds int
}

func (t *testGithub) add(ctx context.Context, issue *pbgh.Issue) error {
	t.issues++
	return nil
}
func (t *testGithub) delete(ctx context.Context, issue *pbgh.Issue) error {
	t.issues--
	return nil
}

func (t *testGithub) createPullRequest(ctx context.Context, job, branch string) error {
	t.pullRequests++
	return nil
}

func (t *testBuilder) build(ctx context.Context, name, fullName string) error {
	t.builds++
	return nil
}

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")

	return s
}

func TestBasicPing(t *testing.T) {
	s := InitTestServer()

	err := s.processPing(context.Background(), &pb.Ping{})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}
}

func TestBasicBuildPing(t *testing.T) {
	s := InitTestServer()
	tb := &testBuilder{}
	s.builder = tb

	err := s.processPing(context.Background(), &pb.Ping{Ref: "refs/heads/master", Repository: &pb.Repository{Name: "test"}})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}

	if tb.builds != 1 {
		t.Errorf("Did not start a build")
	}
}

func TestBasicIssuePing(t *testing.T) {
	s := InitTestServer()
	tgh := &testGithub{}
	s.github = tgh

	err := s.processPing(context.Background(), &pb.Ping{Action: "opened", Issue: &pb.Issue{}})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}

	if tgh.issues != 1 {
		t.Errorf("Did not start a build")
	}

	err = s.processPing(context.Background(), &pb.Ping{Action: "closed", Issue: &pb.Issue{}})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}

	if tgh.issues != 0 {
		t.Errorf("Did not delete an issue")
	}

}
