package roleads

import (
	"time"

	"github.com/elgris/sqrl"
)

type GetRoLeadsRequest struct {
	ID                *string
	CustomerVehicleID *string
	CreatedThisMonth  bool
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetRoLeadsRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.CustomerVehicleID != nil {
		q.Where(sqrl.Eq{"i_customer_vehicle_id": req.CustomerVehicleID})
	}
	if req.CreatedThisMonth {
		q.Where(sqrl.Expr("MONTH(d_created_at) = ?", time.Now().Month()))
		q.Where(sqrl.Expr("YEAR(d_created_at) = ?", time.Now().Year()))
	}
}
