package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/githubreceiver/proto"
)

func getFromUrl(url string) (int32, string) {
	elems := strings.Split(url, "/")
	num, _ := strconv.ParseInt(elems[7], 10, 32)
	service := elems[5]
	return int32(num), service
}

func (s *Server) processPing(ctx context.Context, ping *pb.Ping) error {
	if ping.Ref == "refs/heads/master" || ping.Ref == "refs/heads/main" {
		s.CtxLog(ctx, fmt.Sprintf("Starting build for %v", ping.Repository.Name))
		err := s.builder.build(ctx, ping.Repository.Name, ping.Repository.FullName)
		return err
	}

	if ping.Action == "opened" && ping.Number == 0 && ping.GetIssue().GetPullRequest().GetUrl() == "" {
		pieces := strings.Split(ping.Issue.Url, "/")
		s.CtxLog(ctx, fmt.Sprintf("Adding issue %v (%v and %v)", ping.Issue, pieces[5], pieces[7]))
		num, _ := strconv.ParseInt(pieces[7], 10, 32)
		s.github.add(ctx, &pbgh.Issue{Title: ping.Issue.Title, Number: int32(num), Service: pieces[5], Url: ping.Issue.Url, Origin: pbgh.Issue_FROM_RECEIVER})
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
			number, service := getFromUrl(ping.Issue.Url)
			s.github.delete(ctx, &pbgh.Issue{
				Title:   ping.Issue.Title,
				Url:     ping.Issue.Url,
				Origin:  pbgh.Issue_FROM_RECEIVER,
				Number:  number,
				Service: service,
			})
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
