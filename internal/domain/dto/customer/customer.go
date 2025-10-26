package customer

import (
	"fmt"

	"github.com/elgris/sqrl"
)

// OneAccountRequest represents the one account information from the webhook
type OneAccountRequest struct {
	OneAccountID string `json:"one_account_ID" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Gender       string `json:"gender" validate:"required,oneof=MALE FEMALE"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Email        string `json:"email" validate:"omitempty,email"`
}

// GetCustomerRequest represents the request parameters for getting users
type GetCustomerRequest struct {
	Limit        int
	Offset       int
	Search       string
	SortBy       string
	Order        string
	OneAccountID string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerRequest) Apply(q *sqrl.SelectBuilder) {
	if req.OneAccountID != "" {
		q.Where(sqrl.Eq{"one_account_ID": req.OneAccountID})
	}

	if req.Limit > 0 {
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit))
	}
}
