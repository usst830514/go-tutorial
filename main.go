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
	client, err := gitlab.NewClient("_wL5Zkipsb9JxEW3LpVX", gitlab.WithBaseURL("http://gitlab.hupu.com/api/v3"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	listProjectsOptions := &gitlab.ListProjectsOptions{}
	projects, _, err := client.Projects.ListProjects(listProjectsOptions)
	p, _ := json.Marshal(projects)
	fmt.Println(string(p))
}

func githubTest() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	searchOptions := &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 1}}
	result, _, err := client.Search.Code(context.Background(), "test in:file", searchOptions)
	if err != nil {
		panic(err)
	}
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func main() {
	gitlabTest()
	//githubTest()
}
