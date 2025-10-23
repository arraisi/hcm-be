package user

import (
	"fmt"
	"time"

	"github.com/elgris/sqrl"
)

// GetUserRequest represents the request parameters for getting users
type GetUserRequest struct {
	Limit  int
	Offset int
	Search string
	SortBy string
	Order  string
	ID     string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetUserRequest) Apply(q *sqrl.SelectBuilder) {
	if req.Limit > 0 {
		q = q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit))
	}

	if req.Search != "" {
		searchTerm := "%" + req.Search + "%"
		q.Where(sqrl.Or{
			sqrl.Expr("name ILIKE ?", searchTerm),
			sqrl.Expr("email ILIKE ?", searchTerm),
		})
	}
}

// CreateUserRequest represents the request for creating a user
type CreateUserRequest struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
}

// UpdateUserRequest represents the request for updating a user
type UpdateUserRequest struct {
	Email *string
	Name  *string
}

// MapToUpdateBuilder maps the UpdateUserRequest fields to a map for updating
func (req UpdateUserRequest) MapToUpdateBuilder() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if req.Email != nil {
		updateMap["email"] = *req.Email
	}
	if req.Name != nil {
		updateMap["name"] = *req.Name
	}

	return updateMap
}
