package mockapi

import (
	"context"
	"fmt"
)

func (c *Client) DeleteUser(ctx context.Context, userID int64) error {
	url := fmt.Sprintf("%s/v1/users/%d", c.cfg.Http.MockDIDXApi.BaseUrl, userID)
	token := c.cfg.Http.MockApi.APIKey
	header := map[string]string{}

	_, err := c.httpUtil.Delete(ctx, url, token, header)
	if err != nil {
		return err
	}

	return nil
}
