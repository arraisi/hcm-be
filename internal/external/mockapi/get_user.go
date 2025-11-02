package mockapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (c *Client) GetUser(ctx context.Context, userID int64) (domain.User, error) {
	var user domain.User

	endPoint := "/v1/users/" + utils.ToString(userID)
	result, err := c.Get(ctx, endPoint)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(result, &user); err != nil {
		return user, fmt.Errorf("parse response: %w", err)
	}

	return user, nil
}
