package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

// UsedCar domain model matching DDL tr_used_car
type UsedCar struct {
	ID                           int             `json:"id" db:"i_id"`
	CustomerID                   string          `json:"customer_id" db:"i_customer_id"`
	UsedCarBrand                 string          `json:"used_car_brand" db:"c_used_car_brand"`
	VIN                          string          `json:"vin" db:"c_vin"`
	PoliceNumber                 string          `json:"police_number" db:"c_police_number"`
	KatashikiSuffix              string          `json:"katashiki_suffix" db:"c_katashiki_suffix"`
	ColorCode                    string          `json:"color_code" db:"c_color_code"`
	UsedCarModel                 string          `json:"used_car_model" db:"c_used_car_model"`
	UsedCarVariant               string          `json:"used_car_variant" db:"c_used_car_variant"`
	UsedCarColor                 string          `json:"used_car_color" db:"c_used_car_color"`
	InitialEstimatedUsedCarPrice decimal.Decimal `json:"initial_estimated_used_car_price" db:"v_initial_estimated_price"`
	UsedCarYear                  int             `json:"used_car_year" db:"v_used_car_year"`
	UsedCarFuel                  string          `json:"used_car_fuel" db:"c_used_car_fuel"`
	UsedCarEngineCapacity        int             `json:"used_car_engine_capacity" db:"v_engine_capacity"`
	UsedCarTransmission          string          `json:"used_car_transmission" db:"c_used_car_transmission"`
	Mileage                      int             `json:"mileage" db:"v_mileage"`
	Province                     string          `json:"province" db:"c_province"`
	Mover                        string          `json:"mover" db:"c_mover"`
	StnkExpiryDate               time.Time       `json:"stnk_expiry_date" db:"d_stnk_expiry_date"`
	Stnk                         bool            `json:"stnk" db:"b_stnk"`
	Bpkb                         bool            `json:"bpkb" db:"b_bpkb"`
	Faktur                       bool            `json:"faktur" db:"b_faktur"`
	ServiceBook                  bool            `json:"service_book" db:"b_service_book"`
	AvailabilitySpareKey         bool            `json:"availability_spare_key" db:"b_availability_spare_key"`
	SpareTireCheck               bool            `json:"spare_tire_check" db:"b_spare_tire_check"`
	CustomerRequestedPrice       decimal.Decimal `json:"customer_requested_price" db:"v_customer_requested_price"`
	AdditionalNotes              string          `json:"additional_notes" db:"c_additional_notes"`
	CreatedAt                    time.Time       `json:"created_at" db:"d_created_at"`
}

//
// ────────────────────────────────────────────────────────────────
//   TABLE NAME
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) TableName() string {
	return "dbo.tm_used_car"
}

//
// ────────────────────────────────────────────────────────────────
//   COLUMNS
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) Columns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"c_used_car_brand",
		"c_vin",
		"c_police_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_used_car_model",
		"c_used_car_variant",
		"c_used_car_color",
		"v_initial_estimated_price",
		"v_used_car_year",
		"c_used_car_fuel",
		"v_engine_capacity",
		"c_used_car_transmission",
		"v_mileage",
		"c_province",
		"c_mover",
		"d_stnk_expiry_date",
		"b_stnk",
		"b_bpkb",
		"b_faktur",
		"b_service_book",
		"b_availability_spare_key",
		"b_spare_tire_check",
		"v_customer_requested_price",
		"c_additional_notes",
		"d_created_at",
	}
}

//
// ────────────────────────────────────────────────────────────────
//   SELECT COLUMNS
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) SelectColumns() []string {
	return []string{
		"CAST(i_id AS INT) as i_id",
		"i_customer_id",
		"c_used_car_brand",
		"c_vin",
		"c_police_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_used_car_model",
		"c_used_car_variant",
		"c_used_car_color",
		"v_initial_estimated_price",
		"v_used_car_year",
		"c_used_car_fuel",
		"v_engine_capacity",
		"c_used_car_transmission",
		"v_mileage",
		"c_province",
		"c_mover",
		"d_stnk_expiry_date",
		"b_stnk",
		"b_bpkb",
		"b_faktur",
		"b_service_book",
		"b_availability_spare_key",
		"b_spare_tire_check",
		"v_customer_requested_price",
		"c_additional_notes",
		"d_created_at",
	}
}

//
// ────────────────────────────────────────────────────────────────
//   INSERT MAPPER
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) ToCreateMap() (columns []string, values []interface{}) {
	columns = []string{}
	values = []interface{}{}

	if u.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, u.CustomerID)
	}

	if u.UsedCarBrand != "" {
		columns = append(columns, "c_used_car_brand")
		values = append(values, u.UsedCarBrand)
	}
	if u.VIN != "" {
		columns = append(columns, "c_vin")
		values = append(values, u.VIN)
	}
	if u.PoliceNumber != "" {
		columns = append(columns, "c_police_number")
		values = append(values, u.PoliceNumber)
	}
	if u.KatashikiSuffix != "" {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, u.KatashikiSuffix)
	}
	if u.ColorCode != "" {
		columns = append(columns, "c_color_code")
		values = append(values, u.ColorCode)
	}
	if u.UsedCarModel != "" {
		columns = append(columns, "c_used_car_model")
		values = append(values, u.UsedCarModel)
	}
	if u.UsedCarVariant != "" {
		columns = append(columns, "c_used_car_variant")
		values = append(values, u.UsedCarVariant)
	}
	if u.UsedCarColor != "" {
		columns = append(columns, "c_used_car_color")
		values = append(values, u.UsedCarColor)
	}
	if !u.InitialEstimatedUsedCarPrice.IsZero() {
		columns = append(columns, "v_initial_estimated_price")
		values = append(values, u.InitialEstimatedUsedCarPrice)
	}
	if u.UsedCarYear != 0 {
		columns = append(columns, "v_used_car_year")
		values = append(values, u.UsedCarYear)
	}
	if u.UsedCarFuel != "" {
		columns = append(columns, "c_used_car_fuel")
		values = append(values, u.UsedCarFuel)
	}
	if u.UsedCarEngineCapacity != 0 {
		columns = append(columns, "v_engine_capacity")
		values = append(values, u.UsedCarEngineCapacity)
	}
	if u.UsedCarTransmission != "" {
		columns = append(columns, "c_used_car_transmission")
		values = append(values, u.UsedCarTransmission)
	}
	if u.Mileage != 0 {
		columns = append(columns, "v_mileage")
		values = append(values, u.Mileage)
	}
	if u.Province != "" {
		columns = append(columns, "c_province")
		values = append(values, u.Province)
	}
	if u.Mover != "" {
		columns = append(columns, "c_mover")
		values = append(values, u.Mover)
	}
	if !u.StnkExpiryDate.IsZero() {
		columns = append(columns, "d_stnk_expiry_date")
		values = append(values, u.StnkExpiryDate)
	}

	columns = append(columns, "b_stnk")
	values = append(values, u.Stnk)

	columns = append(columns, "b_bpkb")
	values = append(values, u.Bpkb)

	columns = append(columns, "b_faktur")
	values = append(values, u.Faktur)

	columns = append(columns, "b_service_book")
	values = append(values, u.ServiceBook)

	columns = append(columns, "b_availability_spare_key")
	values = append(values, u.AvailabilitySpareKey)

	columns = append(columns, "b_spare_tire_check")
	values = append(values, u.SpareTireCheck)

	if !u.CustomerRequestedPrice.IsZero() {
		columns = append(columns, "v_customer_requested_price")
		values = append(values, u.CustomerRequestedPrice)
	}

	if u.AdditionalNotes != "" {
		columns = append(columns, "c_additional_notes")
		values = append(values, u.AdditionalNotes)
	}

	columns = append(columns, "d_created_at")
	values = append(values, u.CreatedAt)

	return
}

//
// ────────────────────────────────────────────────────────────────
//   UPDATE MAPPER
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) ToUpdateMap() map[string]interface{} {
	m := map[string]interface{}{}

	if u.CustomerID != "" {
		m["i_customer_id"] = u.CustomerID
	}

	if u.UsedCarBrand != "" {
		m["c_used_car_brand"] = u.UsedCarBrand
	}
	if u.VIN != "" {
		m["c_vin"] = u.VIN
	}
	if u.PoliceNumber != "" {
		m["c_police_number"] = u.PoliceNumber
	}
	if u.KatashikiSuffix != "" {
		m["c_katashiki_suffix"] = u.KatashikiSuffix
	}
	if u.ColorCode != "" {
		m["c_color_code"] = u.ColorCode
	}
	if u.UsedCarModel != "" {
		m["c_used_car_model"] = u.UsedCarModel
	}
	if u.UsedCarVariant != "" {
		m["c_used_car_variant"] = u.UsedCarVariant
	}
	if u.UsedCarColor != "" {
		m["c_used_car_color"] = u.UsedCarColor
	}
	if !u.InitialEstimatedUsedCarPrice.IsZero() {
		m["v_initial_estimated_price"] = u.InitialEstimatedUsedCarPrice
	}
	if u.UsedCarYear != 0 {
		m["v_used_car_year"] = u.UsedCarYear
	}
	if u.UsedCarFuel != "" {
		m["c_used_car_fuel"] = u.UsedCarFuel
	}
	if u.UsedCarEngineCapacity != 0 {
		m["v_engine_capacity"] = u.UsedCarEngineCapacity
	}
	if u.UsedCarTransmission != "" {
		m["c_used_car_transmission"] = u.UsedCarTransmission
	}
	if u.Mileage != 0 {
		m["v_mileage"] = u.Mileage
	}
	if u.Province != "" {
		m["c_province"] = u.Province
	}
	if u.Mover != "" {
		m["c_mover"] = u.Mover
	}
	if !u.StnkExpiryDate.IsZero() {
		m["d_stnk_expiry_date"] = u.StnkExpiryDate
	}

	m["b_stnk"] = u.Stnk
	m["b_bpkb"] = u.Bpkb
	m["b_faktur"] = u.Faktur
	m["b_service_book"] = u.ServiceBook
	m["b_availability_spare_key"] = u.AvailabilitySpareKey
	m["b_spare_tire_check"] = u.SpareTireCheck

	if !u.CustomerRequestedPrice.IsZero() {
		m["v_customer_requested_price"] = u.CustomerRequestedPrice
	}
	if u.AdditionalNotes != "" {
		m["c_additional_notes"] = u.AdditionalNotes
	}

	return m
}
