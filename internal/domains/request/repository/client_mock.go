package repository

import (
	"errors"
	"net/http"
)

type GetClientMock struct {
	response *http.Response
	message  string
}

func (g GetClientMock) Get(url string, pathParams map[string]string, queryParams map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error) {
	if g.message != "" {
		return g.response, errors.New(g.message)
	}
	return g.response, nil
}

func newClientMock(response *http.Response, message string) GetClient {
	return GetClientMock{
		response: response,
		message:  message,
	}
}
