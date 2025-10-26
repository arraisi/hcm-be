package domain

type Customer struct {
	OneAccountID     string `json:"one_account_ID" db:"one_account_ID"`
	FirstName        string `json:"first_name" db:"first_name"`
	LastName         string `json:"last_name" db:"last_name"`
	Gender           string `json:"gender" db:"gender"`
	PhoneNumber      string `json:"phone_number" db:"phone_number"`
	Email            string `json:"email" db:"email"`
	IID              string `json:"i_id" db:"i_id"`
	IHasjratid       string `json:"i_hasjratid" db:"i_hasjratid"`
	CIsNew           int    `json:"c_isnew" db:"c_isnew"`
	CIsMerge         int    `json:"c_ismerge" db:"c_ismerge"`
	PrimaryUser      string `json:"primary_user" db:"primary_user"`
	DealerCustomerID string `json:"dealer_customer_ID" db:"dealer_customer_ID"`
	CIsValid         bool   `json:"c_isvalid" db:"c_isvalid"`
	CIsOmnichannel   bool   `json:"c_isomnichannel" db:"c_isomnichannel"`
	IIdleadsin       string `json:"i_idleadsin" db:"i_idleadsin"`
}

// TableName returns the database table name for the User model
func (u *Customer) TableName() string {
	return "dbo.tr_hcmcustomer"
}

// Columns returns the list of database columns for the User model
func (u *Customer) Columns() []string {
	return []string{
		"one_account_ID",
		"first_name",
		"last_name",
		"gender",
		"phone_number",
		"email",
		"i_id",
		"i_hasjratid",
		"c_isnew",
		"c_ismerge",
		"primary_user",
		"dealer_customer_ID",
		"c_isvalid",
		"c_isomnichannel",
		"i_idleadsin",
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
		u.IID,
		u.IHasjratid,
		u.CIsNew,
		u.CIsMerge,
		u.PrimaryUser,
		u.DealerCustomerID,
		u.CIsValid,
		u.CIsOmnichannel,
		u.IIdleadsin,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Customer) SelectColumns() []string {
	return []string{
		"CAST(one_account_ID AS NVARCHAR(36)) as one_account_ID",
		"first_name",
		"last_name",
		"gender",
		"phone_number",
		"email",
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_hasjratid AS NVARCHAR(36)) as i_hasjratid",
		"c_isnew",
		"c_ismerge",
		"primary_user",
		"CAST(dealer_customer_ID AS NVARCHAR(36)) as dealer_customer_ID",
		"c_isvalid",
		"c_isomnichannel",
		"CAST(i_idleadsin AS NVARCHAR(36)) as i_idleadsin",
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
