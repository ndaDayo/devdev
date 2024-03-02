package github

import (
	"context"
	"encoding/json"
	"net/http"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
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
