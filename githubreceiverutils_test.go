package main

import (
	"testing"

	pbgh "github.com/brotherlogic/githubcard/proto"
	keystoreclient "github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/githubreceiver/proto"
)

type testPullRequester struct {
	updates int
	commits int
}

func (p *testPullRequester) commitToPullRequest(ctx context.Context, url, sha, name string) error {
	p.commits++
	return nil
}

func (p *testPullRequester) updatePullRequest(ctx context.Context, url, name, checkName string, pass bool) error {
	p.updates++
	return nil
}

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

func (t *testGithub) createPullRequest(ctx context.Context, job, branch, title string) error {
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

func TestAddPullRequest(t *testing.T) {
	s := InitTestServer()
	tgh := &testGithub{}
	s.github = tgh
	s.pullRequester = &testPullRequester{}

	err := s.processPing(context.Background(), &pb.Ping{Action: "opened", Number: 20, Issue: &pb.Issue{}, PullRequest: &pb.PullRequest{Url: "blahurl", Head: &pb.Head{Sha: "blah"}}})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}
}
