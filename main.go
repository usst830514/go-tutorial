package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v35/github"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"log"
)

func gitlabTest() {
	git, err := gitlab.NewClient("_wL5Zkipsb9JxEW3LpVX", gitlab.WithBaseURL("http://gitlab.hupu.io/api/v4"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	//u, err := json.Marshal(users)
	fmt.Println(len(users))
}

func githubTest() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	result, _, err := client.Search.Code(context.Background(), "star:>100", &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 2}})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func main() {
	//gitlabTest()
	githubTest()
}
