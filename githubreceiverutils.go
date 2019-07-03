package main

import (
	"fmt"

	pb "github.com/brotherlogic/githubreceiver/proto"
	"golang.org/x/net/context"
)

func (s *Server) processPing(ctx context.Context, ping *pb.Ping) error {
	if ping.Ref == "/refs/heads/master" {
		s.Log(fmt.Sprintf("Starting build for %v", ping.Repository.Name))
		s.builder.build(ctx, ping.Repository.Name, ping.Repository.FullName)
	}

	return nil
}
