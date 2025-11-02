package mockapi

import (
	"context"

	"github.com/arraisi/hcm-be/pkg/utils"
)

func (c *Client) DeleteUser(ctx context.Context, userID int64) error {
	endPoint := "/v1/users/" + utils.ToString(userID)
	_, err := c.Delete(ctx, endPoint)
	if err != nil {
		return err
	}

	return nil
}
