package dmsaftersales

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// Client handles test drive operations via external API
type client struct {
	cfg      *config.Config
	httpUtil utils.HttpUtil
}

// New creates a new DIDX client with HttpUtil
func New(cfg *config.Config, httpUtil utils.HttpUtil) *client {
	return &client{
		cfg:      cfg,
		httpUtil: httpUtil,
	}
}

func (c *client) ServiceBookingRequest(ctx context.Context, body any) error {
	// Set custom headers as shown in the curl request
	header := map[string]string{
		"token": c.cfg.Http.ApimDMSAfterSalesApi.APIKey,
		"str":   "test", // TODO: Update this value based on actual requirements
	}

	p, _ := json.Marshal(body)
	fmt.Println(string(p))

	url := fmt.Sprintf("%s/webhook/after-sales", c.cfg.Http.ApimDMSAfterSalesApi.BaseUrl)

	// Pass empty string for token since we're using custom "token" header
	result, err := c.httpUtil.Post(ctx, url, body, "", header)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
