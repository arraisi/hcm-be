package leads

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/elgris/sqrl"
)

func (r *repository) GetLeadsTestDrive(ctx context.Context, req leads.GetLeadsTestDriveRequest) ([]domain.Leads, error) {
	var model domain.Leads
	var results []domain.Leads

	query := sqrl.Select(model.LeadsTestDriveColumns()...).
		From(model.TableNameAlias()).
		LeftJoin("dbo.tm_testdrive td on td.i_leads_id = l.i_leads_id")

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &results, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}
