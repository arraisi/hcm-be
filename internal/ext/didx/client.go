package didx

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// Client handles test drive operations via external API
type Client struct {
	cfg      *config.Config
	httpUtil utils.HttpUtil
}

// New creates a new DIDX client with HttpUtil
func New(cfg *config.Config, httpUtil utils.HttpUtil) *Client {
	return &Client{
		cfg:      cfg,
		httpUtil: httpUtil,
	}
}

func (c *Client) Post(ctx context.Context, request interface{}, path string) (response []byte, err error) {

	url := c.cfg.Http.MockDIDXApi.BaseUrl + path
	header := map[string]string{}

	token := c.cfg.Http.MockDIDXApi.APIKey

	respBody, err := c.httpUtil.Post(ctx, url, request, token, header)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

func (c *Client) Put(ctx context.Context, request interface{}, path string) (response []byte, err error) {
	url := c.cfg.Http.MockDIDXApi.BaseUrl + path

	header := map[string]string{}

	token := c.cfg.Http.MockDIDXApi.APIKey

	respBody, err := c.httpUtil.Put(ctx, url, request, token, header)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

func (c *Client) Get(ctx context.Context, path string) (response []byte, err error) {

	url := c.cfg.Http.MockDIDXApi.BaseUrl + path
	header := map[string]string{}

	token := c.cfg.Http.MockDIDXApi.APIKey

	respBody, err := c.httpUtil.Get(ctx, url, token, header)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

func (c *Client) Delete(ctx context.Context, path string) (response []byte, err error) {

	url := c.cfg.Http.MockDIDXApi.BaseUrl + path
	header := map[string]string{}

	token := c.cfg.Http.MockDIDXApi.APIKey

	respBody, err := c.httpUtil.Delete(ctx, url, token, header)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
