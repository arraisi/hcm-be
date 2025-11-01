package customervehicle

import "github.com/elgris/sqrl"

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
