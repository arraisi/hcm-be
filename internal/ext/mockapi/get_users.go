package mockapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/user"
)

func (c *Client) GetUsers(ctx context.Context, request user.GetUserRequest) ([]domain.User, error) {
	var users []domain.User

	// Build query params
	query := url.Values{}

	if request.Limit > 0 {
		query.Add("limit", strconv.Itoa(request.Limit))
	}
	if request.Offset > 0 {
		query.Add("page", strconv.Itoa(request.Offset))
	}

	// Build endpoint with query
	endPoint := "/v1/users"
	endPointWithQuery := endPoint + "?" + query.Encode()

	// Make request
	result, err := c.Get(ctx, endPointWithQuery)
	if err != nil {
		return users, err
	}

	// Parse JSON response
	if err := json.Unmarshal(result, &users); err != nil {
		return users, fmt.Errorf("parse response: %w", err)
	}

	return users, nil
}
