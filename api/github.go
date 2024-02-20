package api

import (
	"net/http"
	"time"
)

const (
	BASEURL = "https://api.github.com/"
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
