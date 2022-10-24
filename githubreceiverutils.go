package main

import (
	"fmt"

	"golang.org/x/net/context"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/githubreceiver/proto"
)

func (s *Server) processPing(ctx context.Context, ping *pb.Ping) error {
	if ping.Ref == "refs/heads/master" || ping.Ref == "refs/heads/main" {
		s.CtxLog(ctx, fmt.Sprintf("Starting build for %v", ping.Repository.Name))
		err := s.builder.build(ctx, ping.Repository.Name, ping.Repository.FullName)
		return err
	}

	if ping.Action == "opened" && ping.Number == 0 {
		s.CtxLog(ctx, fmt.Sprintf("Adding issue %v", ping.Issue))
		s.github.add(ctx, &pbgh.Issue{Title: ping.Issue.Title, Url: ping.Issue.Url, Origin: pbgh.Issue_FROM_RECEIVER})
		return nil
	}

	/*
		if ping.Action == "opened" && ping.Number > 0 {
			s.CtxLog(ctx, fmt.Sprintf("Opening PR %v", ping))
			s.pqueue = append(s.pqueue, pull{ping.PullRequest.Url, ping.PullRequest.Head.Sha, ping.PullRequest.Title})
			s.pullRequester.commitToPullRequest(ctx, ping.PullRequest.Url, ping.PullRequest.Head.Sha, ping.PullRequest.Title)

			return nil
		}*/

	if ping.Action == "closed" {
		s.CtxLog(ctx, fmt.Sprintf("Deleting issue %v", ping.Issue))
		if ping != nil && ping.Issue != nil {
			s.github.delete(ctx, &pbgh.Issue{Title: ping.Issue.Title, Url: ping.Issue.Url, Origin: pbgh.Issue_FROM_RECEIVER})
		}
		return nil
	}

	/*
		if ping.Action == "synchronize" {
			s.CtxLog(ctx, fmt.Sprintf("Commiting to  PR %v", ping))
			s.pqueue = append(s.pqueue, pull{ping.PullRequest.Url, ping.PullRequest.Head.Sha, ping.PullRequest.Title})
			s.pullRequester.commitToPullRequest(ctx, ping.PullRequest.Url, ping.PullRequest.Head.Sha, ping.PullRequest.Title)
			return nil
		}*/

	if ping.RefType == "branch" {
		err := s.github.createPullRequest(ctx, ping.Repository.Name, ping.Ref, ping.Ref)
		s.CtxLog(ctx, fmt.Sprintf("Building pull request for %v -> %v", ping.Ref, err))
		return nil
	}

	/*
		if len(ping.Context) > 0 {
			s.CtxLog(ctx, fmt.Sprintf("Updating PR %v", ping))
			return s.pullRequester.updatePullRequest(ctx, ping.Sha, "", ping.Context, ping.State == "success")
		}

		if len(ping.GetCheckRun().GetName()) > 0 {
			s.CtxLog(ctx, fmt.Sprintf("Updating CheckRunPR %v", ping))
			return s.pullRequester.updatePullRequest(ctx, ping.GetCheckRun().GetHeadSha(), "", ping.GetCheckRun().GetName(), ping.GetCheckRun().GetConclusion() == "success")
		}*/

	s.CtxLog(ctx, fmt.Sprintf("Skipping processing of %v", ping))

	return nil
}
