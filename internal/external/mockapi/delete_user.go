package mockapi

import (
	"context"
	"fmt"
)

func (c *Client) DeleteUser(ctx context.Context, userID int64) error {
	url := fmt.Sprintf(c.cfg.Http.MockApi.BaseUrl, userID)
	token := c.cfg.Http.MockApi.APIKey
	header := map[string]string{}

	_, err := c.httpUtil.Delete(ctx, url, token, header)
	if err != nil {
		return err
	}

	return nil
}
