package appraisal

import (
	"github.com/elgris/sqrl"
)

// GetAppraisalRequest represents the request parameters for getting users
type GetAppraisalRequest struct {
	AppraisalBookingID     string
	AppraisalBookingNumber string
	LeadsID                string
	OneAccountID           string
	VIN                    string
	OutletID               string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetAppraisalRequest) Apply(q *sqrl.SelectBuilder) {
	if req.AppraisalBookingID != "" {
		q.Where(sqrl.Eq{"i_appraisal_booking_id": req.AppraisalBookingID})
	}

	if req.AppraisalBookingNumber != "" {
		q.Where(sqrl.Eq{"c_appraisal_booking_number": req.AppraisalBookingNumber})
	}

	if req.LeadsID != "" {
		q.Where(sqrl.Eq{"i_leads_id": req.LeadsID})
	}

	if req.OneAccountID != "" {
		q.Where(sqrl.Eq{"i_one_account_id": req.OneAccountID})
	}

	if req.VIN != "" {
		q.Where(sqrl.Eq{"c_vin": req.VIN})
	}

	if req.OutletID != "" {
		q.Where(sqrl.Eq{"i_outlet_id": req.OutletID})
	}
}
