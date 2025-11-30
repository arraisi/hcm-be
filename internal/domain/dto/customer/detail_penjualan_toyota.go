package customer

import (
	"fmt"

	"github.com/elgris/sqrl"
)

type GetDetailPenjualanToyotaRequest struct {
	CustomerID *string `json:"customer_id"`
	Phone      *string `json:"phone"`
	Chasis     *string `json:"chasis"`
	Engine     *string `json:"engine"`
	Limit      int     `json:"limit"`
	Offset     int     `json:"offset"`
}

func (req GetDetailPenjualanToyotaRequest) Apply(q *sqrl.SelectBuilder) {
	if req.CustomerID != nil {
		q.Where(sqrl.Eq{"Customer Id": req.CustomerID})
	}
	if req.Phone != nil {
		q.Where(sqrl.Eq{"Phone": req.Phone})
	}
	if req.Chasis != nil {
		q.Where(sqrl.Eq{"Chasis": req.Chasis})
	}
	if req.Engine != nil {
		q.Where(sqrl.Eq{"Engine": req.Engine})
	}

	if req.Limit > 0 {
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit))
	}
}
