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
		q.Where(sqrl.Eq{"c_vin": req.VIN})
	}

	if req.PoliceNumber != "" {
		q.Where(sqrl.Eq{"c_police_number": req.PoliceNumber})
	}
}
