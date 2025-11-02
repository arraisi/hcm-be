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

	baseUrl := fmt.Sprintf("%s/v1/users", c.cfg.Http.MockDIDXApi.BaseUrl)
	urlWithQuery := baseUrl + "?" + query.Encode()

	header := map[string]string{}
	token := c.cfg.Http.MockApi.APIKey

	result, err := c.httpUtil.Get(ctx, urlWithQuery, token, header)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	if err := json.Unmarshal(result, &users); err != nil {
		return users, fmt.Errorf("parse response: %w", err)
	}

	return users, nil
}
