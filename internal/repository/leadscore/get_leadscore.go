package leadscore

import (
	"context"

	"github.com/elgris/sqrl"
	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/leads"
)

func (r *repository) GetLeadScore(ctx context.Context, req leads.GetLeadScoreRequest) (domain.LeadScore, error) {
	var model domain.LeadScore

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	err = r.db.GetContext(ctx, &model, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
