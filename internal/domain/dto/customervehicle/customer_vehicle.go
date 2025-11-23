package customervehicle

import (
	"fmt"

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

type GetCustomerVehiclePaginatedRequest struct {
	Limit                   int
	Offset                  int
	Search                  string
	SortBy                  string
	Order                   string
	ID                      string
	DecDateNotNull          bool
	CarPaymentStatusNotNull bool
	OutletCodeNotNull       bool
	SalesNikNotNull         bool
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerVehiclePaginatedRequest) Apply(q *sqrl.SelectBuilder) {
	if req.Search != "" {
		// Example search implementation - adjust fields as necessary
		q.Where(sqrl.Or{
			sqrl.Expr("c_vin LIKE ?", "%"+req.Search+"%"),
			sqrl.Expr("c_police_number LIKE ?", "%"+req.Search+"%"),
			sqrl.Expr("c_model LIKE ?", "%"+req.Search+"%"),
		})
	}

	if req.DecDateNotNull {
		q.Where("d_dec_date IS NOT NULL")
	}
	if req.CarPaymentStatusNotNull {
		q.Where("c_car_payment_status IS NOT NULL")
	}
	if req.OutletCodeNotNull {
		q.Where("c_outlet_code IS NOT NULL")
	}
	if req.SalesNikNotNull {
		q.Where("c_sales_nik IS NOT NULL")
	}

	if req.Limit > 0 {
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit+1))
	}
}

type CustomerVehicleRequest struct {
	Vin             string `json:"vin"`
	KatashikiSuffix string `json:"katashiki_suffix"`
	ColorCode       string `json:"color_code"`
	Model           string `json:"model"`
	Variant         string `json:"variant"`
	Color           string `json:"color"`
	PoliceNumber    string `json:"police_number" validate:"required"`
	ActualMileage   int32  `json:"actual_mileage"`
}

func NewCustomerVehicleRequest(domain domain.CustomerVehicle) CustomerVehicleRequest {
	return CustomerVehicleRequest{
		Vin:             domain.Vin,
		KatashikiSuffix: domain.KatashikiSuffix,
		ColorCode:       domain.ColorCode,
		Model:           domain.Model,
		Variant:         domain.Variant,
		Color:           domain.Color,
		PoliceNumber:    domain.PoliceNumber,
		ActualMileage:   *domain.ActualMileage,
	}
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
		ActualMileage:   &req.ActualMileage,
		CreatedBy:       constants.System,
		UpdatedBy:       constants.System,
	}
}
