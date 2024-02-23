package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type MockHttpClient struct {
	ResponseBody string
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(m.ResponseBody)),
		Header:     make(http.Header),
	}
	return response, nil
}
