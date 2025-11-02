package didx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (c *Client) ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error {
	endPoint := "/v1/test-drive/" + utils.ToString(1)
	result, err := c.Put(ctx, request, endPoint)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
