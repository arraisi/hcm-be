package hmf

import (
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
