package appraisal

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// GetAppraisal gets a single appraisal
func (r *repository) GetAppraisal(ctx context.Context, req appraisal.GetAppraisalRequest) (domain.Appraisal, error) {
	var model domain.Appraisal

	query := sqrl.
		Select(model.SelectColumns()...).
		From(model.TableName())
	req.Apply(query)

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
