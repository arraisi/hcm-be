package toyotaid

type Request struct {
	Process   string `json:"process" validate:"required"`
	EventID   string `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64  `json:"timestamp" validate:"required"`
	Data      Data   `json:"data" validate:"required"`
}

type Data struct {
	OneAccount      OneAccount      `json:"one_account" validate:"required"`
	CustomerVehicle CustomerVehicle `json:"customer_vehicle" validate:"required"`
}

// OneAccount represents customer information
type OneAccount struct {
	OneAccountID         string `json:"one_account_ID" validate:"required"`
	DealerCustomerID     string `json:"dealer_customer_ID" validate:"required"`
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name" validate:"required"`
	PhoneNumber          string `json:"phone_number" validate:"required"`
	Email                string `json:"email" validate:"omitempty,email"`
	BirthDate            string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	VerificationChannel  string `json:"verification_channel" validate:"required,oneof=SMS WHATSAPP EMAIL"`
	KTPNumber            string `json:"ktp_number" validate:"required"`
	Occupation           string `json:"occupation" validate:"omitempty"`
	Gender               string `json:"gender" validate:"omitempty,oneof=FEMALE MALE OTHER"`
	AddressLabel         string `json:"address_label" validate:"omitempty"`
	ResidenceAddress     string `json:"residence_address" validate:"omitempty"`
	Province             string `json:"province" validate:"omitempty"`
	City                 string `json:"city" validate:"omitempty"`
	District             string `json:"district" validate:"omitempty"`
	Subdistrict          string `json:"subdistrict" validate:"omitempty"`
	PostalCode           string `json:"postal_code" validate:"omitempty"`
	DetailAddress        string `json:"detail_address" validate:"omitempty"`
	RegistrationChannel  string `json:"registration_channel" validate:"omitempty,oneof=MTOYOTA DXMI DEALER_SYSTEM"`
	RegistrationDatetime int64  `json:"registration_datetime" validate:"omitempty"`
	ConsentGiven         bool   `json:"consent_given" validate:"omitempty"`
	ConsentGivenAt       int64  `json:"consent_given_at" validate:"omitempty"`
	ConsentGivenDuring   string `json:"consent_given_during" validate:"omitempty,oneof=REGISTRATION SPK SERVICE ADD_VEHICLE"`
	ToyotaSingleIDStatus string `json:"toyota_single_ID_status" validate:"omitempty,oneof=ACTIVE INACTIVE DELETED"`
	CustomerCategory     string `json:"customer_category" validate:"omitempty,oneof=INDIVIDUAL COMPANY"`
	KTPImage             string `json:"ktp_image" validate:"omitempty,base64"`
}

// CustomerVehicle represents vehicle information for the OneAccount
type CustomerVehicle struct {
	PrimaryUser     string   `json:"primary_user" validate:"omitempty,oneof=MASTER MEMBER"`
	VIN             string   `json:"vin" validate:"required"`
	PoliceNumber    string   `json:"police_number" validate:"required"`
	KatashikiSuffix string   `json:"katashiki_suffix" validate:"required"`
	ColorCode       string   `json:"color_code" validate:"required"`
	Model           string   `json:"model" validate:"required"`
	Variant         string   `json:"variant" validate:"required"`
	Color           string   `json:"color" validate:"required"`
	STNKNumber      string   `json:"stnk_number" validate:"omitempty"`
	STNKName        string   `json:"stnk_name" validate:"omitempty"`
	STNKExpiryDate  int64    `json:"stnk_expiry_date" validate:"omitempty"`
	STNKAddress     string   `json:"stnk_address" validate:"omitempty"`
	CustomerType    []string `json:"customer_type" validate:"required,dive,oneof=OWNER BUYER USER"`
	VehicleCategory string   `json:"vehicle_category" validate:"required,oneof=RETAIL FLEET"`
	STNKImage       string   `json:"stnk_image" validate:"omitempty,base64"`
}
