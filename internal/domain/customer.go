package domain

import (
	"time"
)

type Customer struct {
	ID                      string    `json:"id" db:"i_id"`
	OneAccountID            string    `json:"one_account_id" db:"i_one_account_id"`
	HasjratID               string    `json:"hasjrat_id" db:"i_hasjrat_id"`
	FirstName               string    `json:"first_name" db:"n_first_name"`
	LastName                string    `json:"last_name" db:"n_last_name"`
	Gender                  string    `json:"gender" db:"n_gender"`
	PhoneNumber             string    `json:"phone_number" db:"c_phone_number"`
	Email                   string    `json:"email" db:"e_email"`
	IsNew                   bool      `json:"is_new" db:"c_isnew"`
	IsMerge                 bool      `json:"is_merge" db:"c_ismerge"`
	PrimaryUser             *string   `json:"primary_user" db:"c_primary_user"`
	DealerCustomerID        string    `json:"dealer_customer_ID" db:"i_dealer_customer_id"`
	IsValid                 bool      `json:"is_valid" db:"c_isvalid"`
	IsOmnichannel           bool      `json:"is_omnichannel" db:"c_isomnichannel"`
	LeadsInID               string    `json:"leads_in_id" db:"i_leads_in_id"`
	CustomerCategory        string    `json:"customer_category" db:"c_customer_category"`
	KTPNumber               string    `json:"ktp_number" db:"c_ktp_number"`
	BirthDate               time.Time `json:"birth_date" db:"d_birth_date"`
	ResidenceAddress        string    `json:"residence_address" db:"e_residence_address"`
	ResidenceSubdistrict    string    `json:"residence_subdistrict" db:"e_residence_subdistrict"`
	ResidenceDistrict       string    `json:"residence_district" db:"e_residence_district"`
	ResidenceCity           string    `json:"residence_city" db:"e_residence_city"`
	ResidenceProvince       string    `json:"residence_province" db:"e_residence_province"`
	ResidencePostalCode     string    `json:"residence_postal_code" db:"e_residence_postal_code"`
	CustomerType            string    `json:"customer_type" db:"c_customer_type"`
	LeadsID                 string    `json:"leads_id" db:"i_leads_id"`
	Occupation              string    `json:"occupation" db:"n_occupation"`
	RegistrationChannel     string    `json:"registration_channel" db:"c_registration_channel"`
	RegistrationDatetime    time.Time `json:"registration_datetime" db:"d_registration_datetime"`
	ConsentGiven            bool      `json:"consent_given" db:"c_consent_given"`
	ConsentGivenAt          time.Time `json:"consent_given_at" db:"d_consent_given_at"`
	ConsentGivenDuring      string    `json:"consent_given_during" db:"e_consent_given_during"`
	AddressLabel            string    `json:"address_label" db:"e_address_label"`
	DetailAddress           string    `json:"detail_address" db:"e_detail_address"`
	ToyotaIDSingleStatus    string    `json:"toyota_id_single_status" db:"c_toyota_id_single_status"`
	PreferredContactChannel string    `json:"preferred_contact_channel" db:"c_preferred_contact_channel"`
	CreatedAt               time.Time `json:"created_at" db:"d_created_at"`
	CreatedBy               string    `json:"created_by" db:"c_created_by"`
	UpdatedAt               time.Time `json:"updated_at" db:"d_updated_at"`
	UpdatedBy               *string   `json:"updated_by" db:"c_updated_by"`
}

// TableName returns the database table name for the User model
func (u *Customer) TableName() string {
	return "dbo.tr_customer"
}

// Columns returns the list of database columns for the User model
func (u *Customer) Columns() []string {
	return []string{
		"i_id",
		"i_one_account_id",
		"i_hasjrat_id",
		"n_first_name",
		"n_last_name",
		"n_gender",
		"c_phone_number",
		"e_email",
		"c_isnew",
		"c_ismerge",
		"c_primary_user",
		"i_dealer_customer_id",
		"c_isvalid",
		"c_isomnichannel",
		"i_leads_in_id",
		"c_customer_category",
		"c_ktp_number",
		"d_birth_date",
		"e_residence_address",
		"e_residence_subdistrict",
		"e_residence_district",
		"e_residence_city",
		"e_residence_province",
		"e_residence_postal_code",
		"c_customer_type",
		"i_leads_id",
		"n_occupation",
		"c_registration_channel",
		"d_registration_datetime",
		"c_consent_given",
		"d_consent_given_at",
		"e_consent_given_during",
		"e_address_label",
		"e_detail_address",
		"c_toyota_id_single_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Customer) SelectColumns() []string {
	return []string{
		"i_id",
		"i_one_account_id",
		"n_first_name",
		"n_last_name",
		"n_gender",
		"c_phone_number",
		"e_email",
		"c_primary_user",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (u *Customer) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(u.Columns()))
	values = make([]interface{}, 0, len(u.Columns()))

	if u.OneAccountID != "" {
		columns = append(columns, "i_one_account_id")
		values = append(values, u.OneAccountID)
	}
	if u.HasjratID != "" {
		columns = append(columns, "i_hasjrat_id")
		values = append(values, u.HasjratID)
	}
	if u.FirstName != "" {
		columns = append(columns, "n_first_name")
		values = append(values, u.FirstName)
	}
	if u.LastName != "" {
		columns = append(columns, "n_last_name")
		values = append(values, u.LastName)
	}
	if u.Gender != "" {
		columns = append(columns, "n_gender")
		values = append(values, u.Gender)
	}
	if u.PhoneNumber != "" {
		columns = append(columns, "c_phone_number")
		values = append(values, u.PhoneNumber)
	}
	if u.Email != "" {
		columns = append(columns, "e_email")
		values = append(values, u.Email)
	}
	if u.PrimaryUser != nil {
		columns = append(columns, "c_primary_user")
		values = append(values, u.PrimaryUser)
	}
	if u.DealerCustomerID != "" {
		columns = append(columns, "i_dealer_customer_id")
		values = append(values, u.DealerCustomerID)
	}
	if u.LeadsInID != "" {
		columns = append(columns, "i_leads_in_id")
		values = append(values, u.LeadsInID)
	}
	if u.CustomerCategory != "" {
		columns = append(columns, "c_customer_category")
		values = append(values, u.CustomerCategory)
	}
	if u.KTPNumber != "" {
		columns = append(columns, "c_ktp_number")
		values = append(values, u.KTPNumber)
	}
	if !u.BirthDate.IsZero() {
		columns = append(columns, "d_birth_date")
		values = append(values, u.BirthDate.UTC())
	}
	if u.ResidenceAddress != "" {
		columns = append(columns, "e_residence_address")
		values = append(values, u.ResidenceAddress)
	}
	if u.ResidenceSubdistrict != "" {
		columns = append(columns, "e_residence_subdistrict")
		values = append(values, u.ResidenceSubdistrict)
	}
	if u.ResidenceDistrict != "" {
		columns = append(columns, "e_residence_district")
		values = append(values, u.ResidenceDistrict)
	}
	if u.ResidenceCity != "" {
		columns = append(columns, "e_residence_city")
		values = append(values, u.ResidenceCity)
	}
	if u.ResidenceProvince != "" {
		columns = append(columns, "e_residence_province")
		values = append(values, u.ResidenceProvince)
	}
	if u.ResidencePostalCode != "" {
		columns = append(columns, "e_residence_postal_code")
		values = append(values, u.ResidencePostalCode)
	}
	if u.CustomerType != "" {
		columns = append(columns, "c_customer_type")
		values = append(values, u.CustomerType)
	}
	if u.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, u.LeadsID)
	}
	if u.Occupation != "" {
		columns = append(columns, "n_occupation")
		values = append(values, u.Occupation)
	}
	if u.RegistrationChannel != "" {
		columns = append(columns, "c_registration_channel")
		values = append(values, u.RegistrationChannel)
	}
	if !u.RegistrationDatetime.IsZero() {
		columns = append(columns, "d_registration_datetime")
		values = append(values, u.RegistrationDatetime.UTC())
	}
	if !u.ConsentGivenAt.IsZero() {
		columns = append(columns, "d_consent_given_at")
		values = append(values, u.ConsentGivenAt.UTC())
	}
	if u.ConsentGivenDuring != "" {
		columns = append(columns, "e_consent_given_during")
		values = append(values, u.ConsentGivenDuring)
	}
	if u.AddressLabel != "" {
		columns = append(columns, "e_address_label")
		values = append(values, u.AddressLabel)
	}
	if u.DetailAddress != "" {
		columns = append(columns, "e_detail_address")
		values = append(values, u.DetailAddress)
	}
	if u.ToyotaIDSingleStatus != "" {
		columns = append(columns, "c_toyota_id_single_status")
		values = append(values, u.ToyotaIDSingleStatus)
	}
	columns = append(columns, "c_isnew")
	values = append(values, u.IsNew)
	columns = append(columns, "c_ismerge")
	values = append(values, u.IsMerge)
	columns = append(columns, "c_isvalid")
	values = append(values, u.IsValid)
	columns = append(columns, "c_isomnichannel")
	values = append(values, u.IsOmnichannel)
	columns = append(columns, "c_consent_given")
	values = append(values, u.ConsentGiven)
	columns = append(columns, "d_created_at")
	values = append(values, u.CreatedAt.UTC())
	columns = append(columns, "c_created_by")
	values = append(values, u.CreatedBy)
	columns = append(columns, "d_updated_at")
	values = append(values, u.UpdatedAt.UTC())
	columns = append(columns, "c_updated_by")
	values = append(values, u.UpdatedBy)

	return columns, values
}

func (u *Customer) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.FirstName != "" {
		updateMap["n_first_name"] = u.FirstName
	}
	if u.LastName != "" {
		updateMap["n_last_name"] = u.LastName
	}
	if u.Gender != "" {
		updateMap["n_gender"] = u.Gender
	}
	if u.PhoneNumber != "" {
		updateMap["c_phone_number"] = u.PhoneNumber
	}
	if u.Email != "" {
		updateMap["e_email"] = u.Email
	}
	updateMap["d_updated_at"] = u.UpdatedAt.UTC()
	updateMap["c_updated_by"] = u.UpdatedBy
	return updateMap
}
