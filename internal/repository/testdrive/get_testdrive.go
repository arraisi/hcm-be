package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/elgris/sqrl"
)

func (r *repository) GetTestDrive(ctx context.Context, req testdrive.GetTestDriveRequest) (domain.TestDrive, error) {
	var model domain.TestDrive

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
