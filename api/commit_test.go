package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCommitService_Get(t *testing.T) {
	resdata, err := ioutil.ReadFile("testdata/commit.json")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	mockHttpClient := &MockHttpClient{
		ResponseBody: string(resdata),
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
