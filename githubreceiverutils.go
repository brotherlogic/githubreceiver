package main

import (
	"fmt"

	"golang.org/x/net/context"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/githubreceiver/proto"
)

func (s *Server) processPing(ctx context.Context, ping *pb.Ping) error {
	if ping.Ref == "refs/heads/master" {
		s.Log(fmt.Sprintf("Starting build for %v", ping.Repository.Name))
		s.builder.build(ctx, ping.Repository.Name, ping.Repository.FullName)
		return nil
	}

	if ping.Action == "opened" {
		s.Log(fmt.Sprintf("Adding issue %v", ping.Issue))
		s.github.add(ctx, &pbgh.Issue{Title: ping.Issue.Title, Url: ping.Issue.Url, Origin: pbgh.Issue_FROM_RECEIVER})
		return nil
	}

	s.Log(fmt.Sprintf("Skipping processing of %v", ping))

	return nil
}
