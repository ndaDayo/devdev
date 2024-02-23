package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

type MockHttpClient struct {
	ResponseBody string
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(m.ResponseBody)),
		Header:     make(http.Header),
	}
	return response, nil
}

func loadTestData(t *testing.T, resource string) []byte {
	t.Helper()

	path := fmt.Sprintf("testdata/%s.json", resource)
	data, err := os.ReadFile(path)

	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	return data
}
