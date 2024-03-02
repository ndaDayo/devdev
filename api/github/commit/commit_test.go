package commit

import (
	"context"
	"encoding/json"
	"testing"
)

type ExpectedCommit struct {
	Commit GitCommit `json:"commit"`
}

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

	var expectedCommit ExpectedCommit
	if err := json.Unmarshal(res, &expectedCommit); err != nil {
		t.Fatalf("Failed to unmarshal expected commit: %v", err)
	}

	if commit.Commit.Message != expectedCommit.Commit.Message {
		t.Errorf("Expected commit message to be '%s', but got '%s'", expectedCommit.Commit.Message, commit.Commit.Message)
	}
}
