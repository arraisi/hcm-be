package customer

import (
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

// OneAccountRequest represents the one account information from the webhook
type OneAccountRequest struct {
	OneAccountID string `json:"one_account_ID" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Gender       string `json:"gender" validate:"required,oneof=MALE FEMALE OTHER"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Email        string `json:"email" validate:"omitempty,email"`
}

func NewOneAccountRequest(customer domain.Customer) OneAccountRequest {
	return OneAccountRequest{
		OneAccountID: customer.OneAccountID,
		FirstName:    customer.FirstName,
		LastName:     customer.LastName,
		Gender:       utils.ToValue(customer.Gender),
		PhoneNumber:  customer.PhoneNumber,
		Email:        customer.Email,
	}
}

// GetCustomerRequest represents the request parameters for getting users
type GetCustomerRequest struct {
	Limit        int
	Offset       int
	Search       string
	SortBy       string
	Order        string
	OneAccountID string
	CustomerID   string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerRequest) Apply(q *sqrl.SelectBuilder) {
	if req.OneAccountID != "" {
		q.Where(sqrl.Eq{"i_one_account_id": req.OneAccountID})
	}

	if req.Limit > 0 {
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit))
	}
}

type UpdateCustomerRequest struct {
	ID           string
	OneAccountID string
	FirstName    *string
	LastName     *string
	Gender       *string
	PhoneNumber  *string
	Email        *string
}

func (req UpdateCustomerRequest) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if req.FirstName != nil {
		updateMap["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updateMap["last_name"] = *req.LastName
	}
	if req.Gender != nil {
		updateMap["gender"] = *req.Gender
	}
	if req.PhoneNumber != nil {
		updateMap["phone_number"] = *req.PhoneNumber
	}
	if req.Email != nil {
		updateMap["email"] = *req.Email
	}

	return updateMap
}

func (be *OneAccountRequest) ToDomain() domain.Customer {
	return domain.Customer{
		OneAccountID: be.OneAccountID,
		FirstName:    be.FirstName,
		LastName:     be.LastName,
		Email:        be.Email,
		PhoneNumber:  be.PhoneNumber,
		Gender:       utils.ToPointer(be.Gender),
		CreatedAt:    time.Now(),
		CreatedBy:    constants.System,
		UpdatedAt:    time.Now(),
		UpdatedBy:    utils.ToPointer(constants.System),
	}
}
