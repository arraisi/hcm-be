package mockapi

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (c *Client) CreateUser(ctx context.Context, request domain.User) error {
	resp, err := c.Post(ctx, request, "/v1/users")
	if err != nil {
		return err
	}

	fmt.Printf("CreateUser response: %s\n", string(resp))

	return nil
}
