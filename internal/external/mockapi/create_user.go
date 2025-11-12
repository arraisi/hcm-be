package mockapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (c *Client) CreateUser(ctx context.Context, request domain.User) error {
	var user domain.User

	header := map[string]string{}
	token := c.cfg.Http.ApimDIDXApi.APIKey

	result, err := c.httpUtil.Post(ctx, c.cfg.Http.MockApi.BaseUrl, request, token, header)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(result, &user); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
