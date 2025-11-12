package didx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
)

func (c *Client) ConfirmServiceBooking(ctx context.Context, request servicebooking.ServiceBookingEvent) error {
	header := map[string]string{}
	token := c.cfg.Http.ApimDIDXApi.APIKey

	result, err := c.httpUtil.Post(ctx, c.cfg.Http.ApimDIDXApi.BaseUrl, request, token, header)
	if err != nil {
		return err
	}

	var resp any
	if err := json.Unmarshal(result, &resp); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	return nil
}
