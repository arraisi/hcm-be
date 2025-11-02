package mockapi

import (
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
