package usedcar

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

func (r *repository) GetLeadsScore(ctx context.Context, leadsID string) (domain.LeadsScore, error) {
	var model domain.LeadsScore

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		Where(sqrl.Eq{"i_leads_id": leadsID})

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
