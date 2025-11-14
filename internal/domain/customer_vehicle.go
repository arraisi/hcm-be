package domain

import (
	"time"
)

type CustomerVehicle struct {
	ID                     string     `db:"i_id"`
	CustomerID             string     `db:"i_customer_id"`
	OneAccountID           string     `db:"i_one_account_id"`
	Vin                    string     `db:"c_vin"`
	KatashikiSuffix        string     `db:"c_katashiki_suffix"`
	ColorCode              string     `db:"c_color_code"`
	Model                  string     `db:"c_model"`
	Variant                string     `db:"c_variant"`
	Color                  string     `db:"n_color"`
	PoliceNumber           string     `db:"c_police_number"`
	ActualMileage          int32      `db:"v_actual_mileage"`
	VehicleCategory        *string    `db:"c_vehicle_category"`
	StnkNumber             *string    `db:"c_stnk_number"`
	StnkName               *string    `db:"n_stnk_name"`
	StnkExpiryDate         *time.Time `db:"d_stnk_expiry_date"`
	StnkAddress            *string    `db:"e_stnk_address"`
	StnkImage              []byte     `db:"e_stnk_image"`
	InvoiceFile            []byte     `db:"e_invoice_file"`
	MainVehicleFlag        *bool      `db:"b_main_vehicle_flag"`
	DistanceTravelled      *int32     `db:"v_distance_travelled"`
	CustomerType           *string    `db:"c_customer_type"`
	PrimaryUser            string     `db:"c_primary_user"`
	DigitalCatalogFlag     *bool      `db:"c_digital_catalog_flag"`
	OrderNumberTCO         string     `db:"c_order_number_tco"`
	CustomerInterestToCare string     `db:"c_customer_interest_to_care"`
	CustomerInterestToBuy  string     `db:"c_customer_interest_to_buy"`
	CreatedAt              time.Time  `db:"d_created_at"`
	CreatedBy              string     `db:"c_created_by"`
	UpdatedAt              time.Time  `db:"d_updated_at"`
	UpdatedBy              string     `db:"c_updated_by"`
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
		"c_police_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_model",
		"c_variant",
		"n_color",
		"v_actual_mileage",
		"c_vehicle_category",
		"c_stnk_number",
		"n_stnk_name",
		"d_stnk_expiry_date",
		"e_stnk_address",
		"e_stnk_image",
		"e_invoice_file",
		"b_main_vehicle_flag",
		"v_distance_travelled",
		"c_customer_type",
		"c_primary_user",
		"c_digital_catalog_flag",
		"c_order_number_tco",
		"c_customer_interest_to_care",
		"c_customer_interest_to_buy",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the CustomerVehicle model
func (cv *CustomerVehicle) SelectColumns() []string {
	// for now we just select everything; if you want a lighter projection
	// you can trim this list
	return cv.Columns()
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
	if cv.PoliceNumber != "" {
		columns = append(columns, "c_police_number")
		values = append(values, cv.PoliceNumber)
	}
	if cv.KatashikiSuffix != "" {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, cv.KatashikiSuffix)
	}
	if cv.ColorCode != "" {
		columns = append(columns, "c_color_code")
		values = append(values, cv.ColorCode)
	}
	if cv.Model != "" {
		columns = append(columns, "c_model")
		values = append(values, cv.Model)
	}
	if cv.Variant != "" {
		columns = append(columns, "c_variant")
		values = append(values, cv.Variant)
	}
	if cv.Color != "" {
		columns = append(columns, "n_color")
		values = append(values, cv.Color)
	}
	if cv.ActualMileage != 0 {
		columns = append(columns, "v_actual_mileage")
		values = append(values, cv.ActualMileage)
	}
	if cv.VehicleCategory != nil {
		columns = append(columns, "c_vehicle_category")
		values = append(values, cv.VehicleCategory)
	}
	if cv.StnkNumber != nil {
		columns = append(columns, "c_stnk_number")
		values = append(values, cv.StnkNumber)
	}
	if cv.StnkName != nil {
		columns = append(columns, "n_stnk_name")
		values = append(values, cv.StnkName)
	}
	if cv.StnkExpiryDate != nil {
		columns = append(columns, "d_stnk_expiry_date")
		values = append(values, *cv.StnkExpiryDate)
	}
	if cv.StnkAddress != nil {
		columns = append(columns, "e_stnk_address")
		values = append(values, cv.StnkAddress)
	}
	if len(cv.StnkImage) > 0 {
		columns = append(columns, "e_stnk_image")
		values = append(values, cv.StnkImage)
	}
	if len(cv.InvoiceFile) > 0 {
		columns = append(columns, "e_invoice_file")
		values = append(values, cv.InvoiceFile)
	}
	if cv.MainVehicleFlag != nil {
		columns = append(columns, "b_main_vehicle_flag")
		values = append(values, *cv.MainVehicleFlag)
	}
	if cv.DistanceTravelled != nil {
		columns = append(columns, "v_distance_travelled")
		values = append(values, *cv.DistanceTravelled)
	}
	if cv.CustomerType != nil {
		columns = append(columns, "c_customer_type")
		values = append(values, cv.CustomerType)
	}
	if cv.PrimaryUser != "" {
		columns = append(columns, "c_primary_user")
		values = append(values, cv.PrimaryUser)
	}
	if cv.DigitalCatalogFlag != nil {
		columns = append(columns, "c_digital_catalog_flag")
		values = append(values, *cv.DigitalCatalogFlag)
	}
	if cv.OrderNumberTCO != "" {
		columns = append(columns, "c_order_number_tco")
		values = append(values, cv.OrderNumberTCO)
	}
	if cv.CustomerInterestToCare != "" {
		columns = append(columns, "c_customer_interest_to_care")
		values = append(values, cv.CustomerInterestToCare)
	}
	if cv.CustomerInterestToBuy != "" {
		columns = append(columns, "c_customer_interest_to_buy")
		values = append(values, cv.CustomerInterestToBuy)
	}

	// audit fields (created_by / updated_by);
	// d_created_at / d_updated_at rely on DB defaults
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
	if cv.PoliceNumber != "" {
		updateMap["c_police_number"] = cv.PoliceNumber
	}
	if cv.KatashikiSuffix != "" {
		updateMap["c_katashiki_suffix"] = cv.KatashikiSuffix
	}
	if cv.ColorCode != "" {
		updateMap["c_color_code"] = cv.ColorCode
	}
	if cv.Model != "" {
		updateMap["c_model"] = cv.Model
	}
	if cv.Variant != "" {
		updateMap["c_variant"] = cv.Variant
	}
	if cv.Color != "" {
		updateMap["n_color"] = cv.Color
	}
	if cv.ActualMileage != 0 {
		updateMap["v_actual_mileage"] = cv.ActualMileage
	}
	if cv.VehicleCategory != nil {
		updateMap["c_vehicle_category"] = cv.VehicleCategory
	}
	if cv.StnkNumber != nil {
		updateMap["c_stnk_number"] = cv.StnkNumber
	}
	if cv.StnkName != nil {
		updateMap["n_stnk_name"] = cv.StnkName
	}
	if cv.StnkExpiryDate != nil {
		updateMap["d_stnk_expiry_date"] = *cv.StnkExpiryDate
	}
	if cv.StnkAddress != nil {
		updateMap["e_stnk_address"] = cv.StnkAddress
	}
	if len(cv.StnkImage) > 0 {
		updateMap["e_stnk_image"] = cv.StnkImage
	}
	if len(cv.InvoiceFile) > 0 {
		updateMap["e_invoice_file"] = cv.InvoiceFile
	}
	if cv.MainVehicleFlag != nil {
		updateMap["b_main_vehicle_flag"] = *cv.MainVehicleFlag
	}
	if cv.DistanceTravelled != nil {
		updateMap["v_distance_travelled"] = *cv.DistanceTravelled
	}
	if cv.CustomerType != nil {
		updateMap["c_customer_type"] = cv.CustomerType
	}
	if cv.PrimaryUser != "" {
		updateMap["c_primary_user"] = cv.PrimaryUser
	}
	if cv.DigitalCatalogFlag != nil {
		updateMap["c_digital_catalog_flag"] = *cv.DigitalCatalogFlag
	}
	if cv.OrderNumberTCO != "" {
		updateMap["c_order_number_tco"] = cv.OrderNumberTCO
	}
	if cv.CustomerInterestToCare != "" {
		updateMap["c_customer_interest_to_care"] = cv.CustomerInterestToCare
	}
	if cv.CustomerInterestToBuy != "" {
		updateMap["c_customer_interest_to_buy"] = cv.CustomerInterestToBuy
	}

	updateMap["d_updated_at"] = time.Now()
	updateMap["c_updated_by"] = cv.UpdatedBy

	return updateMap
}
