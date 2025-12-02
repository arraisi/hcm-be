package domain

import (
	"time"
)

// UsedCar domain model matching DDL tr_used_car
type UsedCar struct {
	ID                           int       `json:"id" db:"i_id"`
	UsedCarBrand                 string    `json:"used_car_brand" db:"used_car_brand"`
	VIN                          string    `json:"vin" db:"vin"`
	PoliceNumber                 string    `json:"police_number" db:"police_number"`
	KatashikiSuffix              string    `json:"katashiki_suffix" db:"katashiki_suffix"`
	ColorCode                    string    `json:"color_code" db:"color_code"`
	UsedCarModel                 string    `json:"used_car_model" db:"used_car_model"`
	UsedCarVariant               string    `json:"used_car_variant" db:"used_car_variant"`
	UsedCarColor                 string    `json:"used_car_color" db:"used_car_color"`
	InitialEstimatedUsedCarPrice float64   `json:"initial_estimated_used_car_price" db:"initial_estimated_used_car_price"`
	UsedCarYear                  int       `json:"used_car_year" db:"used_car_year"`
	UsedCarFuel                  string    `json:"used_car_fuel" db:"used_car_fuel"`
	UsedCarEngineCapacity        int       `json:"used_car_engine_capacity" db:"used_car_engine_capacity"`
	UsedCarTransmission          string    `json:"used_car_transmission" db:"used_car_transmission"`
	Mileage                      int       `json:"mileage" db:"mileage"`
	Province                     string    `json:"province" db:"province"`
	Mover                        string    `json:"mover" db:"mover"`
	StnkExpiryDate               time.Time `json:"stnk_expiry_date" db:"stnk_expiry_date"`
	Stnk                         bool      `json:"stnk" db:"stnk"`
	Bpkb                         bool      `json:"bpkb" db:"bpkb"`
	Faktur                       bool      `json:"faktur" db:"faktur"`
	ServiceBook                  bool      `json:"service_book" db:"service_book"`
	AvailabilitySpareKey         bool      `json:"availability_spare_key" db:"availability_spare_key"`
	SpareTireCheck               bool      `json:"spare_tire_check" db:"spare_tire_check"`
	CustomerRequestedPrice       float64   `json:"customer_requested_price" db:"customer_requested_price"`
	AdditionalNotes              string    `json:"additional_notes" db:"additional_notes"`
	CreatedAt                    time.Time `json:"created_at" db:"created_at"`
}

//
// ────────────────────────────────────────────────────────────────
//   TABLE NAME
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) TableName() string {
	return "dbo.tr_used_car"
}

//
// ────────────────────────────────────────────────────────────────
//   COLUMNS
// ────────────────────────────────────────────────────────────────
//

func (u *UsedCar) Columns() []string {
	return []string{
		"i_id",
		"used_car_brand",
		"vin",
		"police_number",
		"katashiki_suffix",
		"color_code",
		"used_car_model",
		"used_car_variant",
		"used_car_color",
		"initial_estimated_used_car_price",
		"used_car_year",
		"used_car_fuel",
		"used_car_engine_capacity",
		"used_car_transmission",
		"mileage",
		"province",
		"mover",
		"stnk_expiry_date",
		"stnk",
		"bpkb",
		"faktur",
		"service_book",
		"availability_spare_key",
		"spare_tire_check",
		"customer_requested_price",
		"additional_notes",
		"created_at",
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
		"used_car_brand",
		"vin",
		"police_number",
		"katashiki_suffix",
		"color_code",
		"used_car_model",
		"used_car_variant",
		"used_car_color",
		"initial_estimated_used_car_price",
		"used_car_year",
		"used_car_fuel",
		"used_car_engine_capacity",
		"used_car_transmission",
		"mileage",
		"province",
		"mover",
		"stnk_expiry_date",
		"stnk",
		"bpkb",
		"faktur",
		"service_book",
		"availability_spare_key",
		"spare_tire_check",
		"customer_requested_price",
		"additional_notes",
		"created_at",
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

	if u.UsedCarBrand != "" {
		columns = append(columns, "used_car_brand")
		values = append(values, u.UsedCarBrand)
	}
	if u.VIN != "" {
		columns = append(columns, "vin")
		values = append(values, u.VIN)
	}
	if u.PoliceNumber != "" {
		columns = append(columns, "police_number")
		values = append(values, u.PoliceNumber)
	}
	if u.KatashikiSuffix != "" {
		columns = append(columns, "katashiki_suffix")
		values = append(values, u.KatashikiSuffix)
	}
	if u.ColorCode != "" {
		columns = append(columns, "color_code")
		values = append(values, u.ColorCode)
	}
	if u.UsedCarModel != "" {
		columns = append(columns, "used_car_model")
		values = append(values, u.UsedCarModel)
	}
	if u.UsedCarVariant != "" {
		columns = append(columns, "used_car_variant")
		values = append(values, u.UsedCarVariant)
	}
	if u.UsedCarColor != "" {
		columns = append(columns, "used_car_color")
		values = append(values, u.UsedCarColor)
	}
	if u.InitialEstimatedUsedCarPrice != 0 {
		columns = append(columns, "initial_estimated_used_car_price")
		values = append(values, u.InitialEstimatedUsedCarPrice)
	}
	if u.UsedCarYear != 0 {
		columns = append(columns, "used_car_year")
		values = append(values, u.UsedCarYear)
	}
	if u.UsedCarFuel != "" {
		columns = append(columns, "used_car_fuel")
		values = append(values, u.UsedCarFuel)
	}
	if u.UsedCarEngineCapacity != 0 {
		columns = append(columns, "used_car_engine_capacity")
		values = append(values, u.UsedCarEngineCapacity)
	}
	if u.UsedCarTransmission != "" {
		columns = append(columns, "used_car_transmission")
		values = append(values, u.UsedCarTransmission)
	}
	if u.Mileage != 0 {
		columns = append(columns, "mileage")
		values = append(values, u.Mileage)
	}
	if u.Province != "" {
		columns = append(columns, "province")
		values = append(values, u.Province)
	}
	if u.Mover != "" {
		columns = append(columns, "mover")
		values = append(values, u.Mover)
	}
	if !u.StnkExpiryDate.IsZero() {
		columns = append(columns, "stnk_expiry_date")
		values = append(values, u.StnkExpiryDate)
	}

	columns = append(columns, "stnk")
	values = append(values, u.Stnk)

	columns = append(columns, "bpkb")
	values = append(values, u.Bpkb)

	columns = append(columns, "faktur")
	values = append(values, u.Faktur)

	columns = append(columns, "service_book")
	values = append(values, u.ServiceBook)

	columns = append(columns, "availability_spare_key")
	values = append(values, u.AvailabilitySpareKey)

	columns = append(columns, "spare_tire_check")
	values = append(values, u.SpareTireCheck)

	if u.CustomerRequestedPrice != 0 {
		columns = append(columns, "customer_requested_price")
		values = append(values, u.CustomerRequestedPrice)
	}

	if u.AdditionalNotes != "" {
		columns = append(columns, "additional_notes")
		values = append(values, u.AdditionalNotes)
	}

	columns = append(columns, "created_at")
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

	if u.UsedCarBrand != "" {
		m["used_car_brand"] = u.UsedCarBrand
	}
	if u.VIN != "" {
		m["vin"] = u.VIN
	}
	if u.PoliceNumber != "" {
		m["police_number"] = u.PoliceNumber
	}
	if u.KatashikiSuffix != "" {
		m["katashiki_suffix"] = u.KatashikiSuffix
	}
	if u.ColorCode != "" {
		m["color_code"] = u.ColorCode
	}
	if u.UsedCarModel != "" {
		m["used_car_model"] = u.UsedCarModel
	}
	if u.UsedCarVariant != "" {
		m["used_car_variant"] = u.UsedCarVariant
	}
	if u.UsedCarColor != "" {
		m["used_car_color"] = u.UsedCarColor
	}
	if u.InitialEstimatedUsedCarPrice != 0 {
		m["initial_estimated_used_car_price"] = u.InitialEstimatedUsedCarPrice
	}
	if u.UsedCarYear != 0 {
		m["used_car_year"] = u.UsedCarYear
	}
	if u.UsedCarFuel != "" {
		m["used_car_fuel"] = u.UsedCarFuel
	}
	if u.UsedCarEngineCapacity != 0 {
		m["used_car_engine_capacity"] = u.UsedCarEngineCapacity
	}
	if u.UsedCarTransmission != "" {
		m["used_car_transmission"] = u.UsedCarTransmission
	}
	if u.Mileage != 0 {
		m["mileage"] = u.Mileage
	}
	if u.Province != "" {
		m["province"] = u.Province
	}
	if u.Mover != "" {
		m["mover"] = u.Mover
	}
	if !u.StnkExpiryDate.IsZero() {
		m["stnk_expiry_date"] = u.StnkExpiryDate
	}

	m["stnk"] = u.Stnk
	m["bpkb"] = u.Bpkb
	m["faktur"] = u.Faktur
	m["service_book"] = u.ServiceBook
	m["availability_spare_key"] = u.AvailabilitySpareKey
	m["spare_tire_check"] = u.SpareTireCheck

	if u.CustomerRequestedPrice != 0 {
		m["customer_requested_price"] = u.CustomerRequestedPrice
	}
	if u.AdditionalNotes != "" {
		m["additional_notes"] = u.AdditionalNotes
	}

	return m
}
