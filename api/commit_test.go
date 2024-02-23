package api

import (
	"context"
	"fmt"
	"testing"
)

func TestCommitService_Get(t *testing.T) {
	res := loadTestData(t, "commit")

	mockHttpClient := &MockHttpClient{
		ResponseBody: string(res),
	}
	client := NewClient(WithNoToken())
	client.httpClient = mockHttpClient

	cs := &CommitService{client: client}
	commit, _, err := cs.Get(context.Background(), CommitParam{
		Owner: "owner",
		Repo:  "repo",
		Ref:   "ref",
	})

	if err != nil {
		t.Fatalf("CommitService.Get returned an error: %v", err)
	}

	fmt.Println(commit.Commit.Message)
}
