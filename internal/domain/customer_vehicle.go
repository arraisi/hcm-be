package domain

import (
	"time"
)

type CustomerVehicle struct {
	ID              string `db:"i_id"`
	CustomerID      string `db:"i_customer_id"`
	OneAccountID    string `db:"i_one_account_id"`
	Vin             string `db:"c_vin"`
	KatashikiSuffix string `db:"c_katashiki_suffix"`
	Model           string `db:"c_model"`
	ColorCode       string `db:"c_color_code"`
	Variant         string `db:"c_variant"`
	Color           string `db:"n_color"`
	PoliceNumber    string `db:"c_police_number"`
	ActualMileage   int64  `db:"v_actual_mileage"`
	CreatedAt       string `db:"d_created_at"`
	CreatedBy       string `db:"c_created_by"`
	UpdatedAt       string `db:"d_updated_at"`
	UpdatedBy       string `db:"c_updated_by"`
}

// TableName returns the database table name for the CustomerVehicle model
func (cv *CustomerVehicle) TableName() string {
	return "dbo.tm_customer_vehicle"
}

// Columns returns the list of database columns for the CustomerVehicle model
func (cv *CustomerVehicle) Columns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"i_one_account_id",
		"c_vin",
		"c_model",
		"c_katashiki_suffix",
		"c_color_code",
		"c_variant",
		"n_color",
		"c_police_number",
		"v_actual_mileage",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the CustomerVehicle model
func (cv *CustomerVehicle) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_customer_id AS NVARCHAR(36)) as i_customer_id",
		"CAST(i_one_account_id AS NVARCHAR(36)) as i_one_account_id",
		"c_vin",
		"c_model",
		"c_katashiki_suffix",
		"c_color_code",
		"c_variant",
		"n_color",
		"c_police_number",
		"v_actual_mileage",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting a new CustomerVehicle record
func (cv *CustomerVehicle) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(cv.Columns()))
	values := make([]interface{}, 0, len(cv.Columns()))

	if cv.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, cv.CustomerID)
	}
	if cv.OneAccountID != "" {
		columns = append(columns, "i_one_account_id")
		values = append(values, cv.OneAccountID)
	}
	if cv.Vin != "" {
		columns = append(columns, "c_vin")
		values = append(values, cv.Vin)
	}
	if cv.Model != "" {
		columns = append(columns, "c_model")
		values = append(values, cv.Model)
	}
	if cv.KatashikiSuffix != "" {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, cv.KatashikiSuffix)
	}
	if cv.ColorCode != "" {
		columns = append(columns, "c_color_code")
		values = append(values, cv.ColorCode)
	}
	if cv.Variant != "" {
		columns = append(columns, "c_variant")
		values = append(values, cv.Variant)
	}
	if cv.Color != "" {
		columns = append(columns, "n_color")
		values = append(values, cv.Color)
	}
	if cv.PoliceNumber != "" {
		columns = append(columns, "c_police_number")
		values = append(values, cv.PoliceNumber)
	}
	if cv.ActualMileage != 0 {
		columns = append(columns, "v_actual_mileage")
		values = append(values, cv.ActualMileage)
	}
	columns = append(columns, "c_created_by")
	values = append(values, cv.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, cv.CreatedBy)

	return columns, values
}

// ToUpdateMap prepares the columns and values for updating an existing CustomerVehicle record
func (cv *CustomerVehicle) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if cv.CustomerID != "" {
		updateMap["i_customer_id"] = cv.CustomerID
	}
	if cv.OneAccountID != "" {
		updateMap["i_one_account_id"] = cv.OneAccountID
	}
	if cv.Vin != "" {
		updateMap["c_vin"] = cv.Vin
	}
	if cv.Model != "" {
		updateMap["c_model"] = cv.Model
	}
	if cv.KatashikiSuffix != "" {
		updateMap["c_katashiki_suffix"] = cv.KatashikiSuffix
	}
	if cv.ColorCode != "" {
		updateMap["c_color_code"] = cv.ColorCode
	}
	if cv.Variant != "" {
		updateMap["c_variant"] = cv.Variant
	}
	if cv.Color != "" {
		updateMap["n_color"] = cv.Color
	}
	if cv.PoliceNumber != "" {
		updateMap["c_police_number"] = cv.PoliceNumber
	}
	if cv.ActualMileage != 0 {
		updateMap["v_actual_mileage"] = cv.ActualMileage
	}
	updateMap["updated_at"] = time.Now()
	updateMap["updated_by"] = cv.UpdatedBy

	return updateMap
}
