package repository

import (
	"errors"
	"net/http"
)

type GetClient interface {
	Get(url string, pathParams map[string]string, queryParams map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Response, error)
}

type Request interface {
	GetUrl() string
	GetPathParams() map[string]string
	GetQueryParams() map[string]string
	GetHeaders() map[string]string
	GetCookies() []*http.Cookie
}

type Repository struct {
	client GetClient
}

func (r *Repository) GetResponse(request Request) (*http.Response, error) {
	if request == nil {
		return nil, errors.New("request was empty")
	}
	response, err := r.client.Get(
		request.GetUrl(),
		request.GetPathParams(),
		request.GetQueryParams(),
		request.GetHeaders(),
		request.GetCookies(),
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewRepository(client GetClient) *Repository {
	return &Repository{
		client: client,
	}
}
