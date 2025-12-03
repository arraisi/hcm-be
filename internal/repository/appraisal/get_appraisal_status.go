package appraisal

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// GetStatusUpdatesByAppraisalID returns all status history rows for an appraisal,
// ordered by datetime ascending.
func (r *repository) GetStatusUpdatesByAppraisalID(
	ctx context.Context,
	appraisalID string,
) ([]domain.AppraisalStatusUpdate, error) {
	var models []domain.AppraisalStatusUpdate

	model := domain.AppraisalStatusUpdate{}

	query := sqrl.
		Select(model.SelectColumns()...).
		From(model.TableName()).
		Where(sqrl.Eq{"i_appraisal_id": appraisalID}).
		OrderBy("d_trade_in_status_datetime ASC")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return models, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &models, sqlQuery, args...)
	if err != nil {
		return models, err
	}

	return models, nil
}
