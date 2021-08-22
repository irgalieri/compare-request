package get

import (
	resty "github.com/go-resty/resty/v2"
	"net/http"
)

type Request interface {
	R() Request
	SetQueryParams(params map[string]string) Request
	SetPathParams(params map[string]string) Request
	SetHeaders(headers map[string]string) Request
	SetCookies(cookies []*http.Cookie) Request
	Get(url string) (*http.Response, error)
}

type RequestWrapper struct {
	request *resty.Request
}

func (rr *RequestWrapper) R() Request {
	return rr
}

func (rr *RequestWrapper) SetQueryParams(params map[string]string) Request {
	rr.request = rr.request.SetQueryParams(params)
	return rr
}

func (rr *RequestWrapper) SetPathParams(params map[string]string) Request {
	rr.request = rr.request.SetPathParams(params)
	return rr
}

func (rr *RequestWrapper) SetHeaders(headers map[string]string) Request {
	rr.request = rr.request.SetHeaders(headers)
	return rr
}

func (rr *RequestWrapper) SetCookies(cookies []*http.Cookie) Request {
	rr.request = rr.request.SetCookies(cookies)
	return rr
}

func (rr *RequestWrapper) Get(url string) (*http.Response, error) {
	response, err := rr.request.Get(url)
	return response.RawResponse, err
}

func newRequest() Request {
	return &RequestWrapper{
		request: resty.New().R(),
	}
}
