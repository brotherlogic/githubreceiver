package main

import (
	"testing"

	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/githubreceiver/proto"
)

type testBuilder struct {
	builds int
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

	err := s.processPing(context.Background(), &pb.Ping{Ref: "/refs/heads/master", Repository: &pb.Repository{Name: "test"}})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}

	if tb.builds != 1 {
		t.Errorf("Did not start a build")
	}
}
