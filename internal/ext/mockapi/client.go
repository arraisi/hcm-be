package mockapi

import (
	"context"
	"net/http"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/platform/httpclient"
)

type Client struct {
	httpClient config.HttpClientConfig
	http       *httpclient.Client
}

func New(httpClient config.HttpClientConfig) *Client {
	return &Client{
		httpClient: httpClient,
		http: httpclient.New(httpclient.Options{
			Headers: map[string]string{
				"X-API-Key": httpClient.APIKey,
			},
			Retries: httpClient.RetryCount,
			Timeout: httpClient.Timeout * time.Second,
		}),
	}
}

func (c *Client) GetUsers(ctx context.Context) ([]domain.User, error) {
	var resp []domain.User

	url := c.httpClient.BaseUrl + "/v1/users"
	err := c.http.DoJSON(ctx, http.MethodGet, url, nil, &resp)
	return resp, err
}

func (c *Client) GetUser(ctx context.Context) (domain.User, error) {
	var resp domain.User

	url := c.httpClient.BaseUrl + "/v1/users/1"
	err := c.http.DoJSON(ctx, http.MethodGet, url, nil, &resp)
	return resp, err
}

func (c *Client) CreatetUser(ctx context.Context) (domain.User, error) {
	var resp domain.User

	url := c.httpClient.BaseUrl + "/v1/users/"
	err := c.http.DoJSON(ctx, http.MethodPost, url, domain.User{
		Email:     "test@mail.com",
		Name:      "Me",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, &resp)
	return resp, err
}

func (c *Client) UpdateUser(ctx context.Context) (domain.User, error) {
	var resp domain.User

	url := c.httpClient.BaseUrl + "/v1/users/1"
	err := c.http.DoJSON(ctx, http.MethodPut, url, domain.User{
		Email:     "test@mail.com",
		Name:      "Me",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, &resp)

	return resp, err
}
