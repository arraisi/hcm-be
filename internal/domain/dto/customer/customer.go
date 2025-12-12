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
	HasjratID            *string `json:"hasjrat_id"`
}

func NewOneAccountRequest(customer domain.Customer) OneAccountRequest {
	return OneAccountRequest{
		OneAccountID:         utils.ToValue(customer.OneAccountID),
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
		HasjratID:            customer.HasjratID,
	}
}

// GetCustomerRequest represents the request parameters for getting users
type GetCustomerRequest struct {
	Limit            int
	Offset           int
	Search           string
	SortBy           string
	Order            string
	OneAccountID     string
	CustomerID       string
	KTPNumber        string
	PhoneNumber      string
	Page             int
	PageSize         int
	DealerCustomerID string
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

	if req.DealerCustomerID != "" {
		q.Where(sqrl.Eq{"i_dealer_customer_id": req.DealerCustomerID})
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
		OneAccountID:         utils.ToPointer(be.OneAccountID),
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
		HasjratID:            be.HasjratID,
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

type CreateCustomerRequest struct {
	DealerCustomerID     string `json:"dealer_customer_ID"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	PhoneNumber          string `json:"phone_number"`
	Email                string `json:"email"`
	BirthDate            string `json:"birth_date"`
	VerificationChannel  string `json:"verification_channel"`
	KtpNumber            string `json:"ktp_number"`
	Occupation           string `json:"occupation"`
	Gender               string `json:"gender"`
	RegistrationChannel  string `json:"registration_channel"`
	RegistrationDatetime int    `json:"registration_datetime"`
	ConsentGiven         bool   `json:"consent_given"`
	ConsentGivenAt       int    `json:"consent_given_at"`
	ConsentGivenDuring   string `json:"consent_given_during"`
	AddressLabel         string `json:"address_label"`
	ResidenceAddress     string `json:"residence_address"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	District             string `json:"district"`
	Subdistrict          string `json:"subdistrict"`
	PostalCode           string `json:"postal_code"`
	OutletID             string `json:"outlet_ID"`
	DetailAddress        string `json:"detail_address"`
	HasjratID            string `json:"hasjrat_id"`
	CustomerType         string `json:"customer_type"`
}

type CreateCustomerResponse struct {
	HasjratID string `json:"hasjrat_id"`
}

func (req CreateCustomerRequest) ToDomain() domain.Customer {
	now := time.Now().UTC()

	entity := domain.Customer{
		HasjratID:            utils.ToPointer(req.HasjratID),
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Gender:               utils.ToPointer(req.Gender),
		PhoneNumber:          req.PhoneNumber,
		Email:                req.Email,
		IsNew:                true,  // business rule guess: first time we see this event
		IsMerge:              false, // default false
		DealerCustomerID:     req.DealerCustomerID,
		IsValid:              true,  // you may flip this if validation failed elsewhere
		IsOmnichannel:        false, // not provided
		LeadsInID:            "",    // not provided
		CustomerCategory:     "",    // not provided
		KTPNumber:            req.KtpNumber,
		ResidenceAddress:     req.ResidenceAddress,
		ResidenceSubdistrict: req.Subdistrict,
		ResidenceDistrict:    req.District,
		ResidenceCity:        req.City,
		ResidenceProvince:    req.Province,
		ResidencePostalCode:  req.PostalCode,
		CustomerType:         req.CustomerType, // not provided
		LeadsID:              "",               // not provided
		Occupation:           req.Occupation,
		RegistrationChannel:  req.RegistrationChannel,
		ConsentGiven:         req.ConsentGiven,
		ConsentGivenDuring:   req.ConsentGivenDuring,
		AddressLabel:         req.AddressLabel,
		DetailAddress:        req.DetailAddress,
		ToyotaIDSingleStatus: "", // not provided
		CreatedAt:            now,
		CreatedBy:            constants.System,
		UpdatedAt:            now,
		UpdatedBy:            utils.ToPointer(constants.System),
		OutletID:             utils.ToPointer(req.OutletID),
	}

	if req.BirthDate != "" {
		birthDate, err := utils.ParseDateString(req.BirthDate)
		if err != nil {
			return domain.Customer{}
		}
		entity.BirthDate = birthDate
	}

	entity.RegistrationDatetime = now

	return entity
}
