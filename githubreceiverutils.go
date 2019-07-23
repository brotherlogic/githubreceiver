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

	if ping.Action == "closed" {
		s.Log(fmt.Sprintf("Deleting issue %v", ping.Issue))
		s.github.delete(ctx, &pbgh.Issue{Title: ping.Issue.Title, Url: ping.Issue.Url, Origin: pbgh.Issue_FROM_RECEIVER})
		return nil
	}

	if ping.Action == "synchronize" {
		return s.pullRequester.commitToPullRequest(ctx, ping.Url, ping.Head.Sha)
	}

	if ping.RefType == "branch" {
		s.Log(fmt.Sprintf("Building pull request for %v", ping.Ref))
		s.github.createPullRequest(ctx, ping.Repository.Name, ping.Ref, ping.Ref)
		return nil
	}

	if len(ping.Context) > 0 {
		return s.pullRequester.updatePullRequest(ctx, ping.Sha, ping.Name, ping.Context, ping.State == "success")
	}

	s.Log(fmt.Sprintf("Skipping processing of %v", ping))

	return nil
}
