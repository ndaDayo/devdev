package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	ResponseBody string
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(m.ResponseBody)),
		Header:     make(http.Header),
	}
	return response, nil
}

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
