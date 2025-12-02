package usedcar

import (
	"github.com/elgris/sqrl"
)

// GetUsedCarRequest represents the request parameters for getting used car
type GetUsedCarRequest struct {
	VIN          string
	PoliceNumber string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetUsedCarRequest) Apply(q *sqrl.SelectBuilder) {
	if req.VIN != "" {
		q.Where(sqrl.Eq{"i_one_account_id": req.VIN})
	}

	if req.PoliceNumber != "" {
		q.Where(sqrl.Eq{"c_ktp_number": req.PoliceNumber})
	}
}
