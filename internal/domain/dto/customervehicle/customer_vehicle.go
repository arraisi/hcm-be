package customervehicle

import (
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/elgris/sqrl"
)

type GetCustomerVehicleRequest struct {
	ID           *string
	CustomerID   *string
	OneAccountID *string
	Vin          *string
	PoliceNumber *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerVehicleRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.CustomerID != nil {
		q.Where(sqrl.Eq{"i_customer_id": req.CustomerID})
	}
	if req.OneAccountID != nil {
		q.Where(sqrl.Eq{"i_one_account_id": req.OneAccountID})
	}
	if req.Vin != nil {
		q.Where(sqrl.Eq{"c_vin": req.Vin})
	}
	if req.PoliceNumber != nil {
		q.Where(sqrl.Eq{"c_police_number": req.PoliceNumber})
	}
}

type CustomerVehicleRequest struct {
	Vin             string `json:"vin"`
	KatashikiSuffix string `json:"katashiki_suffix"`
	ColorCode       string `json:"color_code"`
	Model           string `json:"model"`
	Variant         string `json:"variant"`
	Color           string `json:"color"`
	PoliceNumber    string `json:"police_number"`
	ActualMileage   int32  `json:"actual_mileage"`
}

func (req CustomerVehicleRequest) ToDomain(customerID, oneAccountID string) domain.CustomerVehicle {
	return domain.CustomerVehicle{
		CustomerID:      customerID,
		OneAccountID:    oneAccountID,
		Vin:             req.Vin,
		KatashikiSuffix: req.KatashikiSuffix,
		ColorCode:       req.ColorCode,
		Model:           req.Model,
		Variant:         req.Variant,
		Color:           req.Color,
		PoliceNumber:    req.PoliceNumber,
		ActualMileage:   req.ActualMileage,
		CreatedBy:       constants.System,
		UpdatedBy:       constants.System,
	}
}
