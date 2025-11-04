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
		Gender:       customer.Gender,
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
		Gender:       be.Gender,
		CreatedAt:    time.Now(),
		CreatedBy:    constants.System,
		UpdatedAt:    time.Now(),
		UpdatedBy:    utils.ToPointer(constants.System),
	}
}

type OneAccountCreate struct {
	Process   string               `json:"process" validate:"required"`
	EventID   string               `json:"event_ID" validate:"required,uuid4"` // UUID v4
	Timestamp int64                `json:"timestamp" validate:"required"`
	Data      OneAccountCreateData `json:"data" validate:"required"`
}

type OneAccountCreateData struct {
	OneAccount OneAccountCreateRequest `json:"one_account" validate:"required"`
}

type OneAccountCreateRequest struct {
	OneAccountID        string `json:"one_account_ID" validate:"omitempty,max=32"`                                                 // Y Conditional
	DealerCustomerID    string `json:"dealer_customer_ID" validate:"required,max=32"`                                              // Y
	FirstName           string `json:"first_name" validate:"required,max=64"`                                                      // Y
	LastName            string `json:"last_name" validate:"required,max=64"`                                                       // Y
	PhoneNumber         string `json:"phone_number" validate:"required,max=16"`                                                    // Y
	Email               string `json:"email" validate:"omitempty,email,max=64"`                                                    // Y Conditional
	BirthDate           string `json:"birth_date" validate:"omitempty,max=32"`                                                     // N (format YYYY-MM-DD before tokenization)
	VerificationChannel string `json:"verification_channel" validate:"required,oneof=SMS WHATSAPP EMAIL,max=8"`                    // Y
	KtpNumber           string `json:"ktp_number" validate:"required,max=16"`                                                      // Y
	Occupation          string `json:"occupation" validate:"omitempty,max=64"`                                                     // N
	Gender              string `json:"gender" validate:"omitempty,oneof=FEMALE MALE OTHER,max=6"`                                  // N
	RegistrationChannel string `json:"registration_channel" validate:"required,oneof=MTOYOTA DXMI DEALER_SYSTEM,max=13"`           // Y
	RegistrationDate    int64  `json:"registration_datetime" validate:"required"`                                                  // Y (UNIX TIMESTAMP)
	ConsentGiven        bool   `json:"consent_given" validate:"required"`                                                          // Y (BOOLEAN)
	ConsentGivenAt      int64  `json:"consent_given_at" validate:"required"`                                                       // Y (UNIX TIMESTAMP)
	ConsentGivenDuring  string `json:"consent_given_during" validate:"required,oneof=REGISTRATION SPK SERVICE ADD_VEHICLE,max=32"` // Y
	AddressLabel        string `json:"address_label" validate:"omitempty,max=256"`                                                 // N
	ResidenceAddress    string `json:"residence_address" validate:"omitempty,max=256"`                                             // N
	Province            string `json:"province" validate:"omitempty,max=64"`                                                       // N
	City                string `json:"city" validate:"omitempty,max=64"`                                                           // N
	District            string `json:"district" validate:"omitempty,max=64"`                                                       // N
	Subdistrict         string `json:"subdistrict" validate:"omitempty,max=64"`                                                    // N
	PostalCode          string `json:"postal_code" validate:"omitempty,len=5,numeric"`                                             // N (VARCHAR(5), treating as numeric postal code)
	DetailAddress       string `json:"detail_address" validate:"omitempty,max=256"`                                                // N
}

func (oa *OneAccountCreateRequest) GetTimeRegistrationDate() time.Time {
	if oa.RegistrationDate == 0 {
		return time.Time{}
	}
	return time.Unix(oa.RegistrationDate, 0)
}

func (oa *OneAccountCreateRequest) GetTimeConsentGivenAt() time.Time {
	if oa.ConsentGivenAt == 0 {
		return time.Time{}
	}
	return time.Unix(oa.ConsentGivenAt, 0)
}

func (oa *OneAccountCreateRequest) GetTimeBirthDate() time.Time {
	if oa.BirthDate == "" {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", oa.BirthDate)
	if err != nil {
		return time.Time{}
	}
	return t
}

func (oa *OneAccountCreate) ToCustomerModel() domain.Customer {
	src := oa.Data.OneAccount
	systemName := "webhook_one_access"
	now := time.Now().UTC()

	return domain.Customer{
		OneAccountID:         src.OneAccountID,
		HasjratID:            "", // not present in payload
		FirstName:            src.FirstName,
		LastName:             src.LastName,
		Gender:               src.Gender,
		PhoneNumber:          src.PhoneNumber,
		Email:                src.Email,
		IsNew:                true,  // business rule guess: first time we see this event
		IsMerge:              false, // default false
		DealerCustomerID:     src.DealerCustomerID,
		IsValid:              true,  // you may flip this if validation failed elsewhere
		IsOmnichannel:        false, // not provided
		LeadsInID:            "",    // not provided
		CustomerCategory:     "",    // not provided
		KTPNumber:            src.KtpNumber,
		BirthDate:            src.GetTimeBirthDate(),
		ResidenceAddress:     src.ResidenceAddress,
		ResidenceSubdistrict: src.Subdistrict,
		ResidenceDistrict:    src.District,
		ResidenceCity:        src.City,
		ResidenceProvince:    src.Province,
		ResidencePostalCode:  src.PostalCode,
		CustomerType:         "", // not provided
		LeadsID:              "", // not provided
		Occupation:           src.Occupation,
		RegistrationChannel:  src.RegistrationChannel,
		RegistrationDatetime: src.GetTimeRegistrationDate(),
		ConsentGiven:         src.ConsentGiven,
		ConsentGivenAt:       src.GetTimeConsentGivenAt(),
		ConsentGivenDuring:   src.ConsentGivenDuring,
		AddressLabel:         src.AddressLabel,
		DetailAddress:        src.DetailAddress,
		ToyotaIDSingleStatus: "", // not provided
		CreatedAt:            now,
		CreatedBy:            systemName,
		UpdatedAt:            now,
		UpdatedBy:            &systemName,
	}
}
