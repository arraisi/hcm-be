package appraisal

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

// UpdateAppraisal updates an appraisal (for confirm / update events)
func (r *repository) UpdateAppraisal(ctx context.Context, tx *sqlx.Tx, req domain.Appraisal) error {
	model := domain.Appraisal{}

	setMap := req.ToUpdateMap()
	if len(setMap) == 0 {
		// nothing to update, just return
		return nil
	}

	query, args, err := sqrl.
		Update(model.TableName()).
		SetMap(setMap).
		Where(sqrl.Eq{"i_id": req.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}

// Optional: UpdateAppraisalByBooking (if kamu mau update pakai bookingID + bookingNumber)
func (r *repository) UpdateAppraisalByBooking(
	ctx context.Context,
	tx *sqlx.Tx,
	req domain.Appraisal,
) error {
	model := domain.Appraisal{}

	setMap := req.ToUpdateMap()
	if len(setMap) == 0 {
		return nil
	}

	query, args, err := sqrl.
		Update(model.TableName()).
		SetMap(setMap).
		Where(sqrl.Eq{
			"i_appraisal_booking_id":     req.AppraisalBookingID,
			"c_appraisal_booking_number": req.AppraisalBookingNumber,
		}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
