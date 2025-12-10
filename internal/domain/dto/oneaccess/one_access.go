package oneaccess

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

type Request struct {
	Process   string      `json:"process" validate:"required"`
	EventID   string      `json:"event_ID" validate:"required,uuid4"` // UUID v4
	Timestamp int64       `json:"timestamp" validate:"required"`
	Data      RequestData `json:"data" validate:"required"`
}

type RequestData struct {
	OneAccount           OneAccount            `json:"one_account" validate:"required"`
	PICAssignmentRequest *PICAssignmentRequest `json:"pic_assignment,omitempty"`
}

type PICAssignmentRequest struct {
	EmployeeID string `json:"employee_ID"`
	FirstName  string `json:"first_name" `
	LastName   string `json:"last_name"`

	NIK string `json:"nik"`
}

type OneAccount struct {
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
	DetailAddress       string `json:"detail_address" validate:"omitempty,max=256"`
	OutletID            string `json:"outlet_id" validate:"omitempty,max=64"` // N
}

func (oa *OneAccount) ToCustomerModel() (domain.Customer, error) {
	now := time.Now().UTC()

	entity := domain.Customer{
		OneAccountID:         oa.OneAccountID,
		HasjratID:            "", // not present in payload
		FirstName:            oa.FirstName,
		LastName:             oa.LastName,
		Gender:               utils.ToPointer(oa.Gender),
		PhoneNumber:          oa.PhoneNumber,
		Email:                oa.Email,
		IsNew:                true,  // business rule guess: first time we see this event
		IsMerge:              false, // default false
		DealerCustomerID:     oa.DealerCustomerID,
		IsValid:              true,  // you may flip this if validation failed elsewhere
		IsOmnichannel:        false, // not provided
		LeadsInID:            "",    // not provided
		CustomerCategory:     "",    // not provided
		KTPNumber:            oa.KtpNumber,
		ResidenceAddress:     oa.ResidenceAddress,
		ResidenceSubdistrict: oa.Subdistrict,
		ResidenceDistrict:    oa.District,
		ResidenceCity:        oa.City,
		ResidenceProvince:    oa.Province,
		ResidencePostalCode:  oa.PostalCode,
		CustomerType:         "", // not provided
		LeadsID:              "", // not provided
		Occupation:           oa.Occupation,
		RegistrationChannel:  oa.RegistrationChannel,
		ConsentGiven:         oa.ConsentGiven,
		ConsentGivenDuring:   oa.ConsentGivenDuring,
		AddressLabel:         oa.AddressLabel,
		DetailAddress:        oa.DetailAddress,
		ToyotaIDSingleStatus: "", // not provided
		CreatedAt:            now,
		CreatedBy:            constants.System,
		UpdatedAt:            now,
		UpdatedBy:            utils.ToPointer(constants.System),
		OutletID:             utils.ToPointer(oa.OutletID),
	}

	if oa.BirthDate != "" {
		birthDate, err := utils.ParseDateString(oa.BirthDate)
		if err != nil {
			return domain.Customer{}, err
		}
		entity.BirthDate = birthDate
	}

	if oa.RegistrationDate != 0 {
		entity.RegistrationDatetime = utils.GetTimeUnix(oa.RegistrationDate)
	}

	if oa.ConsentGivenAt != 0 {
		entity.ConsentGivenAt = utils.GetTimeUnix(oa.ConsentGivenAt)
	}

	return entity, nil
}
