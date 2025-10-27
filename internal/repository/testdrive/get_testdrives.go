package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/elgris/sqrl"
)

func (r *repository) GetTestDrives(ctx context.Context, req testdrive.GetTestDriveRequest) ([]domain.TestDrive, error) {
	var testdrives []domain.TestDrive
	model := domain.TestDrive{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("i_id DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &testdrives, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return testdrives, nil
}
