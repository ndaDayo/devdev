package api

import (
	"context"
	"encoding/json"
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

	Commits *CommitsService
	Commit  *CommitService
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type service struct {
	client *Client
}

func NewClient() *Client {
	c := &Client{
		token: token(),
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	c.initialize()

	return c
}

func (c *Client) initialize() {
	c.Commits = &CommitsService{client: c}
	c.Commit = &CommitService{client: c}
}

func token() string {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	return token
}

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	endpoint, err := url.JoinPath(baseUrl, path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaTypeV3)
	req.Header.Set(headerAPIVersion, defaultAPIVersion)
	req.Header.Set("Authorization", "Bearer "+c.token)

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

	decErr := json.NewDecoder(resp.Body).Decode(v)
	if decErr != nil {
		err = decErr
	}

	return &Response{Response: resp}, nil
}
