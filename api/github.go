package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func NewClient() *Client {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("GITHUB_TOKEN")

	c := &Client{
		token:    token,
		endpoint: baseUrl,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	c.Commits = &CommitsService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	fullUrl := c.endpoint + urlStr
	fmt.Println(fullUrl)
	req, err := http.NewRequest(method, fullUrl, nil)
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
