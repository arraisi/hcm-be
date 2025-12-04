package dmssales

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *client) GetOfferRequest(ctx context.Context, body any) error {
	header := map[string]string{}
	token := c.cfg.Http.DMSApi.APIKey

	p, _ := json.Marshal(body)
	fmt.Println(string(p))

	url := fmt.Sprintf("%s/didx/addgetoffer", c.cfg.Http.DMSApi.BaseUrl)

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
