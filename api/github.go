package api

import (
	"context"
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

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaTypeV3)
	req.Header.Set(headerAPIVersion, defaultAPIVersion)

	return req, nil
}

type Response struct {
	*http.Response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &Response{Response: resp}, nil
}
