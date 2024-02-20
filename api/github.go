package api

import (
	"net/http"
	"os"
	"time"
)

const (
	baseUrl = "https://api.github.com/"

	mediaTypeV3       = "application/vnd.github.v3+json"
	headerAPIVersion  = "X-GitHub-Api-Version"
	defaultAPIVersion = "2022-11-28"
)

type Client struct {
	token      string
	endpoint   string
	httpClient httpClient

	Commits *CommitsService
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type service struct {
	client *Client
}

func NewClient(endpoint string) *Client {
	c := &Client{
		token:    os.Getenv("GITHUB_TOKEN"),
		endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	return c
}

type Response struct {
	*http.Response
}
