package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type service struct {
	client *Client
}

type ClientOption func(*Client)

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
