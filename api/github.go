package api

import "net/http"

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
