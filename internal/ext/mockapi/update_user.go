package mockapi

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (c *Client) UpdateUser(ctx context.Context, userID int64, request domain.User) error {
	endPoint := "/v1/users/" + utils.ToString(userID)
	_, err := c.Put(ctx, request, endPoint)
	if err != nil {
		return err
	}

	return nil
}
