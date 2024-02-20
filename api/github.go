package api

import (
	"net/http"
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
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func NewClient(token, endpoint string) *Client {
	c := &Client{
		token:    token,
		endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	return c
}
