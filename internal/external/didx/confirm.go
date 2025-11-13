package didx

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) Confirm(ctx context.Context, body any) error {
	header := map[string]string{}
	token := c.cfg.Http.ApimDIDXApi.APIKey

	result, err := c.httpUtil.Post(ctx, c.cfg.Http.ApimDIDXApi.BaseUrl, body, token, header)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
