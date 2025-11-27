package spk

import "github.com/elgris/sqrl"

type GetSpkRequest struct {
	SpkID     *string `json:"spk_id"`
	SpkNumber *string `json:"spk_number"`
	OutletID  *string `json:"outlet_id"`
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetSpkRequest) Apply(q *sqrl.SelectBuilder) {
	if req.SpkID != nil {
		q.Where(sqrl.Eq{"i_spk_id": req.SpkID})
	}
	if req.SpkNumber != nil {
		q.Where(sqrl.Eq{"c_spk_number": req.SpkNumber})
	}
	if req.OutletID != nil {
		q.Where(sqrl.Eq{"i_outlet_id": req.OutletID})
	}
}
