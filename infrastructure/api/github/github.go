package github

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

const (
	baseUrl = "https://api.github.com"

	mediaTypeV3       = "application/vnd.github.v3+json"
	headerAPIVersion  = "X-GitHub-Api-Version"
	defaultAPIVersion = "2022-11-28"
)

type Client struct {
	token      string
	httpClient httpClient

	PullRequests *PullRequestsService
}

type service struct {
	client *Client
}

type ClientOption func(*Client)

func GetResource(resource interface{}) (interface{}, error) {
	client := NewClient(WithToken())
	ctx := context.Background()

	switch r := resource.(type) {
	case PullRequestsParam:
		pr, _, err := client.PullRequests.Get(ctx, r)
		if err != nil {
			return nil, err
		}
		return pr, nil
	default:
		return nil, nil
	}
}

func NewClient(options ...ClientOption) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	for _, option := range options {
		option(c)
	}

	c.initialize()

	return c
}

func (c *Client) initialize() {
	c.PullRequests = &PullRequestsService{client: c}
}

func WithNoToken() ClientOption {
	return func(c *Client) {
		c.token = ""
	}
}

func WithToken() ClientOption {
	return func(c *Client) {
		err := gotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error  loading .env file")
		}

		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN is not set in the environment variables")
		}
		c.token = token
	}
}

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	endpoint, err := url.JoinPath(baseUrl, path)
	end, err := url.QueryUnescape(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, end, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaTypeV3)
	req.Header.Set(headerAPIVersion, defaultAPIVersion)
	req.Header.Set("Authorization", "Bearer "+c.token)

	return req, nil
}
