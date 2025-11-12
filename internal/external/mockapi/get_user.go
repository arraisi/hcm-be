package mockapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (c *Client) GetUser(ctx context.Context, userID int64) (domain.User, error) {
	var user domain.User

	url := fmt.Sprintf("%s/v1/test-drive/%d", c.cfg.Http.MockApi.BaseUrl, userID)
	header := map[string]string{}
	token := c.cfg.Http.MockApi.APIKey

	result, err := c.httpUtil.Get(ctx, url, token, header)
	if err != nil {
		return domain.User{}, err
	}

	if err := json.Unmarshal(result, &user); err != nil {
		return user, fmt.Errorf("parse response: %w", err)
	}

	return user, nil
}
