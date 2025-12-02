package domain

import "time"

type Outlet struct {
	ID            string    `json:"id" db:"i_id"`
	BranchID      *string   `json:"branch_id" db:"i_idbranch"`
	OutletCode    string    `json:"outlet_code" db:"c_outlet"`
	OutletName    string    `json:"outlet_name" db:"n_outlet"`
	CreatedDate   time.Time `json:"created_date" db:"d_createdate"`
	HCMCustomerID string    `json:"hcm_customer_id" db:"c_idhcmcustomer"`
	Type          *string   `json:"type" db:"c_tipe"` // 1=dealer, 2=bengkel, 3=service point
	TAMOutletCode *string   `json:"tam_outlet_code" db:"c_tamoutlet"`
}

// TableName returns the database table name for the Outlet model
func (o *Outlet) TableName() string {
	return "dbo.tr_outlet"
}

// Columns returns the full list of database columns for the Outlet model
func (o *Outlet) Columns() []string {
	return []string{
		"i_id",
		"i_idbranch",
		"c_outlet",
		"n_outlet",
		"d_createdate",
		"c_idhcmcustomer",
		"c_tipe",
		"c_tamoutlet",
	}
}

// SelectColumns returns the list of columns to select in queries for the Outlet model
func (o *Outlet) SelectColumns() []string {
	return []string{
		"i_id",
		"i_idbranch",
		"c_outlet",
		"n_outlet",
		"d_createdate",
		"c_idhcmcustomer",
		"c_tipe",
		"c_tamoutlet",
	}
}

// ToCreateMap builds the columns & values slice for INSERT
func (o *Outlet) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(o.Columns()))
	values = make([]interface{}, 0, len(o.Columns()))

	// i_id: DB has default NEWID(); only send if you explicitly set it
	if o.ID != "" {
		columns = append(columns, "i_id")
		values = append(values, o.ID)
	}

	if o.BranchID != nil && *o.BranchID != "" {
		columns = append(columns, "i_idbranch")
		values = append(values, *o.BranchID)
	}

	// Required fields
	if o.OutletCode != "" {
		columns = append(columns, "c_outlet")
		values = append(values, o.OutletCode)
	}
	if o.OutletName != "" {
		columns = append(columns, "n_outlet")
		values = append(values, o.OutletName)
	}

	// d_createdate: DB has default GETDATE(); send if you want control
	if !o.CreatedDate.IsZero() {
		columns = append(columns, "d_createdate")
		values = append(values, o.CreatedDate.UTC())
	}

	// c_idhcmcustomer: default '-1', but you may want to set explicit value
	if o.HCMCustomerID != "" {
		columns = append(columns, "c_idhcmcustomer")
		values = append(values, o.HCMCustomerID)
	}

	if o.Type != nil && *o.Type != "" {
		columns = append(columns, "c_tipe")
		values = append(values, *o.Type)
	}

	if o.TAMOutletCode != nil && *o.TAMOutletCode != "" {
		columns = append(columns, "c_tamoutlet")
		values = append(values, *o.TAMOutletCode)
	}

	return columns, values
}

// ToUpdateMap builds the map for UPDATE SET
func (o *Outlet) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if o.BranchID != nil {
		updateMap["i_idbranch"] = *o.BranchID
	}
	if o.OutletCode != "" {
		updateMap["c_outlet"] = o.OutletCode
	}
	if o.OutletName != "" {
		updateMap["n_outlet"] = o.OutletName
	}
	if o.HCMCustomerID != "" {
		updateMap["c_idhcmcustomer"] = o.HCMCustomerID
	}
	if o.Type != nil {
		updateMap["c_tipe"] = *o.Type
	}
	if o.TAMOutletCode != nil {
		updateMap["c_tamoutlet"] = *o.TAMOutletCode
	}

	// d_createdate biasanya tidak di-update, jadi tidak diset di sini.

	return updateMap
}
