package appraisal

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// GetAppraisalByBooking gets a single appraisal by booking ID and booking number
func (r *repository) GetAppraisalByBooking(
	ctx context.Context,
	appraisalBookingID string,
	appraisalBookingNumber string,
) (domain.Appraisal, error) {
	var model domain.Appraisal

	query := sqrl.
		Select(model.SelectColumns()...).
		From(model.TableName()).
		Where(sqrl.Eq{
			"i_appraisal_booking_id":     appraisalBookingID,
			"c_appraisal_booking_number": appraisalBookingNumber,
		})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &model, sqlQuery, args...)
	if err != nil {
		return model, err
	}

	return model, nil
}

// GetAppraisalByID gets a single appraisal by primary key
func (r *repository) GetAppraisalByID(ctx context.Context, id string) (domain.Appraisal, error) {
	var model domain.Appraisal

	query := sqrl.
		Select(model.SelectColumns()...).
		From(model.TableName()).
		Where(sqrl.Eq{"i_id": id})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &model, sqlQuery, args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
