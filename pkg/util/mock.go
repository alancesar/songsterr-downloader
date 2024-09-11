package util

import (
	"bytes"
	"io"
	"net/http"
)

type (
	fakeRoundTripper struct {
		response   []byte
		statusCode int
	}
)

func (r fakeRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return &http.Response{
		Status:     http.StatusText(r.statusCode),
		StatusCode: r.statusCode,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body:          io.NopCloser(bytes.NewReader(r.response)),
		ContentLength: int64(len(r.response)),
		Request:       request,
	}, nil
}

func NewFakeHTTPClient(response []byte, statusCode int) *http.Client {
	return &http.Client{
		Transport: fakeRoundTripper{
			response:   response,
			statusCode: statusCode,
		},
	}
}
