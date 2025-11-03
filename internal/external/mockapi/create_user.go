package mockapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (c *Client) CreateUser(ctx context.Context, request domain.User) error {
	var user domain.User

	url := fmt.Sprintf("%s/v1/users", c.cfg.Http.MockDIDXApi.BaseUrl)
	header := map[string]string{}
	token := c.cfg.Http.MockApi.APIKey

	result, err := c.httpUtil.Post(ctx, url, request, token, header)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(result, &user); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
