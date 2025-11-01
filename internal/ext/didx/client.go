package mockapi

import (
	"context"
	"net/http"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
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

func (c *Client) ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error {
	type Response struct {
		RequestID string `json:"requestId"`
	}
	var resp Response
	url := c.httpClient.BaseUrl + "/v1/test-drive/1"
	err := c.http.DoJSON(ctx, http.MethodPut, url, request, &resp)
	if err != nil {
		return err
	}

	return nil
}
