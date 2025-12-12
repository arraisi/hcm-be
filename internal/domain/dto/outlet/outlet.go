package outlet

import "github.com/elgris/sqrl"

type GetOutletRequest struct {
	HasjratOutletID string
	TamOutletID     string
}

func (req GetOutletRequest) Apply(q *sqrl.SelectBuilder) {
	if req.HasjratOutletID != "" {
		q.Where(sqrl.Eq{"c_outlet": req.HasjratOutletID})
	}
	if req.TamOutletID != "" {
		q.Where(sqrl.Eq{"c_tamoutlet": req.TamOutletID})
	}
}
