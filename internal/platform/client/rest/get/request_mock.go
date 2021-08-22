package get

import (
	"errors"
	"net/http"
)

type MockRequest struct {
	queryParams map[string]string
	pathParams  map[string]string
	headers     map[string]string
	cookies     []*http.Cookie
	response    *http.Response
	message     string
}

func (m MockRequest) R() Request {
	return m
}

func (m MockRequest) SetQueryParams(params map[string]string) Request {
	m.queryParams = params
	return m
}

func (m MockRequest) SetPathParams(params map[string]string) Request {
	m.pathParams = params
	return m
}

func (m MockRequest) SetHeaders(headers map[string]string) Request {
	m.headers = headers
	return m
}

func (m MockRequest) SetCookies(cookies []*http.Cookie) Request {
	m.cookies = cookies
	return m
}

func (m MockRequest) Get(url string) (*http.Response, error) {
	if m.response != nil {
		m.response.Status = ""
		if m.queryParams != nil && len(m.queryParams) > 0 {
			m.response.Status = m.response.Status + "query,"
		}
		if m.pathParams != nil && len(m.pathParams) > 0 {
			m.response.Status = m.response.Status + "path,"
		}
		if m.headers != nil && len(m.headers) > 0 {
			m.response.Status = m.response.Status + "headers,"
		}
		if m.cookies != nil && len(m.cookies) > 0 {
			m.response.Status = m.response.Status + "cookies,"
		}
	}
	if m.message != "" {
		return m.response, errors.New(m.message)
	}
	return m.response, nil
}

func newRequestMock(response *http.Response, message string) Request {
	return MockRequest{
		response: response,
		message:  message,
	}
}
