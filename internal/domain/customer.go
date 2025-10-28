package domain

import "time"

type Customer struct {
	OneAccountID     string    `json:"one_account_ID" db:"one_account_id"`
	FirstName        string    `json:"first_name" db:"first_name"`
	LastName         string    `json:"last_name" db:"last_name"`
	Gender           string    `json:"gender" db:"gender"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	Email            string    `json:"email" db:"email"`
	ID               string    `json:"id" db:"id"`
	HasjratID        string    `json:"hasjrat_id" db:"hasjrat_id"`
	IsNew            bool      `json:"is_new" db:"is_new"`
	IsMerge          bool      `json:"is_merge" db:"is_merge"`
	PrimaryUser      string    `json:"primary_user" db:"primary_user"`
	DealerCustomerID string    `json:"dealer_customer_ID" db:"dealer_customer_id"`
	IsValid          bool      `json:"is_valid" db:"is_valid"`
	IsOmnichannel    bool      `json:"is_omnichannel" db:"is_omnichannel"`
	LeadsInID        string    `json:"leads_in_id" db:"leads_in_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	CreatedBy        string    `json:"created_by" db:"created_by"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy        *string   `json:"updated_by" db:"updated_by"`
}

// TableName returns the database table name for the User model
func (u *Customer) TableName() string {
	return "dbo.tr_customer"
}

// Columns returns the list of database columns for the User model
func (u *Customer) Columns() []string {
	return []string{
		"one_account_id",
		"first_name",
		"last_name",
		"gender",
		"phone_number",
		"email",
		"id",
		"hasjrat_id",
		"is_new",
		"is_merge",
		"primary_user",
		"dealer_customer_id",
		"is_valid",
		"is_omnichannel",
		"leads_in_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	}
}

func (u *Customer) ToValues() []interface{} {
	return []interface{}{
		u.OneAccountID,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.PhoneNumber,
		u.Email,
		u.ID,
		u.HasjratID,
		u.IsNew,
		u.IsMerge,
		u.PrimaryUser,
		u.DealerCustomerID,
		u.IsValid,
		u.IsOmnichannel,
		u.LeadsInID,
		u.CreatedAt,
		u.CreatedBy,
		u.UpdatedAt,
		u.UpdatedBy,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Customer) SelectColumns() []string {
	return []string{
		"CAST(one_account_id AS NVARCHAR(36)) as one_account_id",
		"first_name",
		"last_name",
		"gender",
		"phone_number",
		"email",
		"CAST(id AS NVARCHAR(36)) as id",
		"CAST(hasjrat_id AS NVARCHAR(36)) as hasjrat_id",
		"is_new",
		"is_merge",
		"primary_user",
		"CAST(dealer_customer_id AS NVARCHAR(36)) as dealer_customer_id",
		"is_valid",
		"is_omnichannel",
		"CAST(leads_in_id AS NVARCHAR(36)) as leads_in_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	}
}

func (u *Customer) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.FirstName != "" {
		updateMap["first_name"] = u.FirstName
	}
	if u.LastName != "" {
		updateMap["last_name"] = u.LastName
	}
	if u.Gender != "" {
		updateMap["gender"] = u.Gender
	}
	if u.PhoneNumber != "" {
		updateMap["phone_number"] = u.PhoneNumber
	}
	if u.Email != "" {
		updateMap["email"] = u.Email
	}
	return updateMap
}
