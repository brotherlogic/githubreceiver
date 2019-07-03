package main

import (
	"testing"

	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/githubreceiver/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")

	return s
}

func TestBasicPing(t *testing.T) {
	s := InitTestServer()

	err := s.processPing(&pb.Ping{})

	if err != nil {
		t.Errorf("Process has failed: %v", err)
	}
}
