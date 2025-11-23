package domain

import (
	"time"

	"github.com/arraisi/hcm-be/pkg/utils"
)

type RoLeads struct {
	ID                      string     `json:"i_id"`
	CustomerVehicleID       string     `json:"i_customer_vehicle_id"`
	CarAge                  int        `json:"i_car_age"`
	CarAgeScore             int        `json:"i_car_age_score"`
	CarPaymentStatusScore   int        `json:"i_car_payment_status_score"`
	CarServiceActivityScore int        `json:"i_car_service_activity_score"`
	CarServiceScore         int        `json:"i_car_service_score"`
	RoScore                 int        `json:"i_ro_score"`
	CustomerResponse        string     `json:"c_customer_response"`
	LeadsOutlet             string     `json:"c_leads_outlet"`
	LeadsSalesNik           string     `json:"c_leads_sales_nik"`
	CreatedAt               time.Time  `json:"d_created_at"`
	UpdatedAt               *time.Time `json:"d_updated_at"`
}

var (
	RoScoreWeight = map[string]float64{
		"car_age":              0.37,
		"car_payment_status":   0.19,
		"car_service_activity": 0.19,
		"car_service":          0.25,
	}
)

// Columns returns the list of columns in the RoLeads model
func (rl *RoLeads) Columns() []string {
	return []string{
		"i_id",
		"i_customer_vehicle_id",
		"i_car_age",
		"i_car_age_score",
		"i_car_payment_status_score",
		"i_car_service_activity_score",
		"i_car_service_score",
		"i_ro_score",
		"c_customer_response",
		"c_leads_outlet",
		"c_leads_sales_nik",
		"d_created_at",
		"d_updated_at",
	}
}

func (rl *RoLeads) TableName() string {
	return "tm_ro_leads"
}

func (rl *RoLeads) ToCreateMap() ([]string, []interface{}) {
	now := time.Now()
	columns := []string{
		"i_customer_vehicle_id",
		"i_car_age",
		"i_car_age_score",
		"i_car_payment_status_score",
		"i_car_service_activity_score",
		"i_car_service_score",
		"i_ro_score",
		"c_customer_response",
		"c_leads_outlet",
		"c_leads_sales_nik",
		"d_created_at",
		"d_updated_at",
	}

	values := []interface{}{
		rl.CustomerVehicleID,
		rl.CarAge,
		rl.CarAgeScore,
		rl.CarPaymentStatusScore,
		rl.CarServiceActivityScore,
		rl.CarServiceScore,
		rl.RoScore,
		rl.CustomerResponse,
		rl.LeadsOutlet,
		rl.LeadsSalesNik,
		now,
		utils.ToPointer(now),
	}

	return columns, values
}

func (rl *RoLeads) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if rl.CustomerVehicleID != "" {
		updateMap["i_customer_vehicle_id"] = rl.CustomerVehicleID
	}
	if rl.CarAge != 0 {
		updateMap["i_car_age"] = rl.CarAge
	}
	if rl.CarAgeScore != 0 {
		updateMap["i_car_age_score"] = rl.CarAgeScore
	}
	if rl.CarPaymentStatusScore != 0 {
		updateMap["i_car_payment_status_score"] = rl.CarPaymentStatusScore
	}
	if rl.CarServiceActivityScore != 0 {
		updateMap["i_car_service_activity_score"] = rl.CarServiceActivityScore
	}
	if rl.CarServiceScore != 0 {
		updateMap["i_car_service_score"] = rl.CarServiceScore
	}
	if rl.RoScore != 0 {
		updateMap["i_ro_score"] = rl.RoScore
	}
	if rl.CustomerResponse != "" {
		updateMap["c_customer_response"] = rl.CustomerResponse
	}
	if rl.LeadsOutlet != "" {
		updateMap["c_leads_outlet"] = rl.LeadsOutlet
	}
	if rl.LeadsSalesNik != "" {
		updateMap["c_leads_sales_nik"] = rl.LeadsSalesNik
	}
	updateMap["d_updated_at"] = time.Now()
	return updateMap
}
