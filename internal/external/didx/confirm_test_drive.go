package didx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
)

func (c *Client) ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error {
	url := fmt.Sprintf("%s/v1/test-drive", c.cfg.Http.MockDIDXApi.BaseUrl)
	header := map[string]string{}
	token := c.cfg.Http.MockDIDXApi.APIKey

	result, err := c.httpUtil.Post(ctx, url, request, token, header)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
