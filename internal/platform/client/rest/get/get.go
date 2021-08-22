package get

import (
	"errors"
	"net/http"
)

type get struct {
	request Request
}

func NewClient() *get {
	get := get{
		request: newRequest(),
	}
	return &get
}

func (g *get) Get(
	url string,
	pathParams map[string]string,
	queryParams map[string]string,
	headers map[string]string,
	cookies []*http.Cookie) (*http.Response, error) {

	if url == "" {
		return nil, errors.New("empty url")
	}

	request := g.request.R()
	if queryParams != nil && len(queryParams) > 0 {
		request = request.SetQueryParams(queryParams)
	}
	if pathParams != nil && len(pathParams) > 0 {
		request = request.SetPathParams(pathParams)
	}
	if headers != nil && len(headers) > 0 {
		request = request.SetHeaders(headers)
	}
	if cookies != nil && len(cookies) > 0 {
		request = request.SetCookies(cookies)
	}

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
