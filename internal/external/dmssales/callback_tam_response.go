package dmssales

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *client) CallbackTamResponse(ctx context.Context, body any) error {
	header := map[string]string{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6IkRNU0FQSSIsIm5iZiI6MTY1NzY4Njk0OCwiZXhwIjoxNjU3NzczMzQ4LCJpYXQiOjE2NTc2ODY5NDh9.bUYAvpDlEGnQl386hYhkHTaOp2msMX2jtQYcKma2JJQ",
	}
	token := c.cfg.Http.DMSApi.APIKey

	p, _ := json.Marshal(body)
	fmt.Println(string(p))

	url := fmt.Sprintf("%s/didx/response", c.cfg.Http.DMSApi.BaseUrl)

	result, err := c.httpUtil.Post(ctx, url, body, token, header)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
