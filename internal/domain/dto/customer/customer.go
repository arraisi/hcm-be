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
	OneAccountID         string  `json:"one_account_ID" validate:"required"`
	DealerCustomerID     *string `json:"dealer_customer_ID"`
	CustomerCategory     *string `json:"customer_category"`
	FirstName            string  `json:"first_name" validate:"required"`
	LastName             string  `json:"last_name" validate:"required"`
	KTPNumber            *string `json:"ktp_number"`
	BirthDate            *string `json:"birth_date"`
	ResidenceAddress     *string `json:"residence_address"`
	ResidenceSubdistrict *string `json:"residence_subdistrict"`
	ResidenceDistrict    *string `json:"residence_district"`
	ResidenceCity        *string `json:"residence_city"`
	ResidenceProvince    *string `json:"residence_province"`
	ResidencePostalCode  *string `json:"residence_postal_code"`
	Email                string  `json:"email" validate:"omitempty,email"`
	PhoneNumber          string  `json:"phone_number" validate:"required"`
	CustomerType         *string `json:"customer_type"`
	Gender               *string `json:"gender"`
}

func NewOneAccountRequest(customer domain.Customer) OneAccountRequest {
	return OneAccountRequest{
		OneAccountID:         customer.OneAccountID,
		DealerCustomerID:     utils.ToPointer(customer.DealerCustomerID),
		CustomerCategory:     utils.ToPointer(customer.CustomerCategory),
		FirstName:            customer.FirstName,
		LastName:             customer.LastName,
		KTPNumber:            utils.ToPointer(customer.KTPNumber),
		BirthDate:            utils.ToPointer(customer.BirthDate.Format("2006-01-02")),
		ResidenceAddress:     utils.ToPointer(customer.ResidenceAddress),
		ResidenceSubdistrict: utils.ToPointer(customer.ResidenceSubdistrict),
		ResidenceDistrict:    utils.ToPointer(customer.ResidenceDistrict),
		ResidenceCity:        utils.ToPointer(customer.ResidenceCity),
		ResidenceProvince:    utils.ToPointer(customer.ResidenceProvince),
		ResidencePostalCode:  utils.ToPointer(customer.ResidencePostalCode),
		Email:                customer.Email,
		PhoneNumber:          customer.PhoneNumber,
		CustomerType:         utils.ToPointer(customer.CustomerType),
		Gender:               customer.Gender,
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
	KTPNumber    string
	PhoneNumber  string
	Page         int
	PageSize     int
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerRequest) Apply(q *sqrl.SelectBuilder) {
	if req.OneAccountID != "" {
		q.Where(sqrl.Eq{"i_one_account_id": req.OneAccountID})
	}

	if req.KTPNumber != "" {
		q.Where(sqrl.Eq{"c_ktp_number": req.KTPNumber})
	}

	if req.PhoneNumber != "" {
		q.Where(sqrl.Eq{"c_phone_number": req.PhoneNumber})
	}

	if req.PageSize > 0 {
		// Calculate offset: (page - 1) * pageSize
		offset := 0
		if req.Page > 1 {
			offset = (req.Page - 1) * req.PageSize
		}
		// Use pageSize + 1 to detect if there's a next page
		limit := req.PageSize + 1
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, limit))
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
	birthDate := time.Time{}
	if be.BirthDate != nil {
		birthDate, _ = time.Parse("2006-01-02", *be.BirthDate)
	}

	return domain.Customer{
		OneAccountID:         be.OneAccountID,
		DealerCustomerID:     utils.ToValue(be.DealerCustomerID),
		CustomerCategory:     utils.ToValue(be.CustomerCategory),
		FirstName:            be.FirstName,
		LastName:             be.LastName,
		KTPNumber:            utils.ToValue(be.KTPNumber),
		BirthDate:            birthDate,
		ResidenceAddress:     utils.ToValue(be.ResidenceAddress),
		ResidenceSubdistrict: utils.ToValue(be.ResidenceSubdistrict),
		ResidenceDistrict:    utils.ToValue(be.ResidenceDistrict),
		ResidenceCity:        utils.ToValue(be.ResidenceCity),
		ResidenceProvince:    utils.ToValue(be.ResidenceProvince),
		ResidencePostalCode:  utils.ToValue(be.ResidencePostalCode),
		Email:                be.Email,
		PhoneNumber:          be.PhoneNumber,
		CustomerType:         utils.ToValue(be.CustomerType),
		Gender:               be.Gender,
		CreatedAt:            time.Now(),
		CreatedBy:            constants.System,
		UpdatedAt:            time.Now(),
		UpdatedBy:            utils.ToPointer(constants.System),
	}
}

type CustomerInquiryRequest struct {
	NIK      *string `json:"nik" validate:"omitempty,min=16,max=16"`
	NoHp     *string `json:"nohp" validate:"omitempty,min=10,max=13"`
	FlagNoHp *bool   `json:"flag_nohp" validate:"required"`
}

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

type GetCustomersResponse struct {
	Data       []domain.Customer `json:"data"`
	Pagination Pagination        `json:"pagination"`
}
