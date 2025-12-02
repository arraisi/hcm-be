package domain

type LeadsTradeIn struct {
	ID                           string   `json:"id" db:"i_id"`
	LeadsID                      string   `json:"leads_id" db:"i_leads_id"`
	TradeInFlag                  bool     `json:"trade_in_flag" db:"b_trade_in_flag"`
	VIN                          string   `json:"vin" db:"c_vin"`
	PoliceNumber                 string   `json:"police_number" db:"c_police_number"`
	KatashikiSuffix              string   `json:"katashiki_suffix" db:"c_katashiki_suffix"`
	ColorCode                    string   `json:"color_code" db:"c_color_code"`
	UsedCarBrand                 string   `json:"used_car_brand" db:"c_used_car_brand"`
	UsedCarModel                 string   `json:"used_car_model" db:"c_used_car_model"`
	UsedCarVariant               string   `json:"used_car_variant" db:"c_used_car_variant"`
	UsedCarColor                 string   `json:"used_car_color" db:"n_used_car_color"`
	InitialEstimatedUsedCarPrice *float64 `json:"initial_estimated_used_car_price" db:"v_initial_estimated_used_car_price"`
	Mileage                      *float64 `json:"mileage" db:"v_mileage"`
	Province                     string   `json:"province" db:"n_province"`
	UsedCarYear                  *float64 `json:"used_car_year" db:"v_used_car_year"`
	UsedCarTransmission          string   `json:"used_car_transmission" db:"c_used_car_transmission"`
	UsedCarFuel                  string   `json:"used_car_fuel" db:"c_used_car_fuel"`
	UsedCarEngineCapacity        *float64 `json:"used_car_engine_capacity" db:"v_used_car_engine_capacity"`
	UsedCarMover                 string   `json:"used_car_mover" db:"c_used_car_mover"`
}

// TableName returns the database table name for the LeadsTradeIn model
func (l *LeadsTradeIn) TableName() string {
	return "dbo.tr_leads_trade_in"
}

// Columns returns the list of database columns for the LeadsTradeIn model
func (l *LeadsTradeIn) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"b_trade_in_flag",
		"c_vin",
		"c_police_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_used_car_brand",
		"c_used_car_model",
		"c_used_car_variant",
		"n_used_car_color",
		"v_initial_estimated_used_car_price",
		"v_mileage",
		"n_province",
		"v_used_car_year",
		"c_used_car_transmission",
		"c_used_car_fuel",
		"v_used_car_engine_capacity",
		"c_used_car_mover",
	}
}

// SelectColumns returns the list of columns to select in queries for the LeadsTradeIn model
func (l *LeadsTradeIn) SelectColumns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"b_trade_in_flag",
		"c_vin",
		"c_police_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_used_car_brand",
		"c_used_car_model",
		"c_used_car_variant",
		"n_used_car_color",
		"v_initial_estimated_used_car_price",
		"v_mileage",
		"n_province",
		"v_used_car_year",
		"c_used_car_transmission",
		"c_used_car_fuel",
		"v_used_car_engine_capacity",
		"c_used_car_mover",
	}
}

// ToCreateMap converts the model to columns and values for insert operation
func (l *LeadsTradeIn) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(l.Columns()))
	values = make([]interface{}, 0, len(l.Columns()))

	if l.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, l.LeadsID)
	}
	columns = append(columns, "b_trade_in_flag")
	values = append(values, l.TradeInFlag)
	if l.VIN != "" {
		columns = append(columns, "c_vin")
		values = append(values, l.VIN)
	}
	if l.PoliceNumber != "" {
		columns = append(columns, "c_police_number")
		values = append(values, l.PoliceNumber)
	}
	if l.KatashikiSuffix != "" {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, l.KatashikiSuffix)
	}
	if l.ColorCode != "" {
		columns = append(columns, "c_color_code")
		values = append(values, l.ColorCode)
	}
	if l.UsedCarBrand != "" {
		columns = append(columns, "c_used_car_brand")
		values = append(values, l.UsedCarBrand)
	}
	if l.UsedCarModel != "" {
		columns = append(columns, "c_used_car_model")
		values = append(values, l.UsedCarModel)
	}
	if l.UsedCarVariant != "" {
		columns = append(columns, "c_used_car_variant")
		values = append(values, l.UsedCarVariant)
	}
	if l.UsedCarColor != "" {
		columns = append(columns, "n_used_car_color")
		values = append(values, l.UsedCarColor)
	}
	if l.InitialEstimatedUsedCarPrice != nil {
		columns = append(columns, "v_initial_estimated_used_car_price")
		values = append(values, l.InitialEstimatedUsedCarPrice)
	}
	if l.Mileage != nil {
		columns = append(columns, "v_mileage")
		values = append(values, l.Mileage)
	}
	if l.Province != "" {
		columns = append(columns, "n_province")
		values = append(values, l.Province)
	}
	if l.UsedCarYear != nil {
		columns = append(columns, "v_used_car_year")
		values = append(values, l.UsedCarYear)
	}
	if l.UsedCarTransmission != "" {
		columns = append(columns, "c_used_car_transmission")
		values = append(values, l.UsedCarTransmission)
	}
	if l.UsedCarFuel != "" {
		columns = append(columns, "c_used_car_fuel")
		values = append(values, l.UsedCarFuel)
	}
	if l.UsedCarEngineCapacity != nil {
		columns = append(columns, "v_used_car_engine_capacity")
		values = append(values, l.UsedCarEngineCapacity)
	}
	if l.UsedCarMover != "" {
		columns = append(columns, "c_used_car_mover")
		values = append(values, l.UsedCarMover)
	}

	return columns, values
}

// ToUpdateMap converts the model to a map for update operation
func (l *LeadsTradeIn) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	updateMap["b_trade_in_flag"] = l.TradeInFlag
	if l.VIN != "" {
		updateMap["c_vin"] = l.VIN
	}
	if l.PoliceNumber != "" {
		updateMap["c_police_number"] = l.PoliceNumber
	}
	if l.KatashikiSuffix != "" {
		updateMap["c_katashiki_suffix"] = l.KatashikiSuffix
	}
	if l.ColorCode != "" {
		updateMap["c_color_code"] = l.ColorCode
	}
	if l.UsedCarBrand != "" {
		updateMap["c_used_car_brand"] = l.UsedCarBrand
	}
	if l.UsedCarModel != "" {
		updateMap["c_used_car_model"] = l.UsedCarModel
	}
	if l.UsedCarVariant != "" {
		updateMap["c_used_car_variant"] = l.UsedCarVariant
	}
	if l.UsedCarColor != "" {
		updateMap["n_used_car_color"] = l.UsedCarColor
	}
	if l.InitialEstimatedUsedCarPrice != nil {
		updateMap["v_initial_estimated_used_car_price"] = l.InitialEstimatedUsedCarPrice
	}
	if l.Mileage != nil {
		updateMap["v_mileage"] = l.Mileage
	}
	if l.Province != "" {
		updateMap["n_province"] = l.Province
	}
	if l.UsedCarYear != nil {
		updateMap["v_used_car_year"] = l.UsedCarYear
	}
	if l.UsedCarTransmission != "" {
		updateMap["c_used_car_transmission"] = l.UsedCarTransmission
	}
	if l.UsedCarFuel != "" {
		updateMap["c_used_car_fuel"] = l.UsedCarFuel
	}
	if l.UsedCarEngineCapacity != nil {
		updateMap["v_used_car_engine_capacity"] = l.UsedCarEngineCapacity
	}
	if l.UsedCarMover != "" {
		updateMap["c_used_car_mover"] = l.UsedCarMover
	}

	return updateMap
}
