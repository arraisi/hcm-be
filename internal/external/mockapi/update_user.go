package mockapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (c *Client) UpdateUser(ctx context.Context, userID int64, request domain.User) error {
	var user domain.User

	url := fmt.Sprintf("%s/v1/test-drive/%d", c.cfg.Http.MockDIDXApi.BaseUrl, userID)
	header := map[string]string{}
	token := c.cfg.Http.MockApi.APIKey

	result, err := c.httpUtil.Put(ctx, url, request, token, header)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(result, &user); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
