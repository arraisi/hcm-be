package sales

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
	"github.com/elgris/sqrl"
)

type iDB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type repository struct {
	db iDB
}

// New creates a new sales repository instance
func New(db iDB) *repository {
	return &repository{
		db: db,
	}
}

// GetSalesScoring retrieves sales scoring data based on request filters
func (r *repository) GetSalesScoring(ctx context.Context, req sales.GetSalesScoringRequest) (sales.GetSalesScoringResponse, error) {
	var salesScorings []domain.SalesScoring
	model := domain.SalesScoring{}

	// Base query with LEFT JOIN
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName() + " ss").
		LeftJoin("dbo.tr_outlet o ON ss.Outlet_code = o.c_outlet")

	// Apply filters from request
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return sales.GetSalesScoringResponse{}, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &salesScorings, sqlQuery, args...)
	if err != nil {
		return sales.GetSalesScoringResponse{}, err
	}

	// Determine if there's a next page using the pageSize + 1 technique
	hasNext := false
	if req.PageSize > 0 && len(salesScorings) > req.PageSize {
		hasNext = true
		salesScorings = salesScorings[:req.PageSize] // Trim the extra record
	}

	// Set default page to 1 if not specified
	page := req.Page
	if page < 1 {
		page = 1
	}

	// Set default pageSize if not specified
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = len(salesScorings)
	}

	return sales.GetSalesScoringResponse{
		Data: salesScorings,
		Pagination: sales.Pagination{
			Page:     page,
			PageSize: pageSize,
			HasNext:  hasNext,
		},
	}, nil
}

// GetSalesScoringByTAMOutlet retrieves sales scoring data by TAM outlet code
func (r *repository) GetSalesScoringByTAMOutlet(ctx context.Context, tamOutletCode string) ([]domain.SalesScoring, error) {
	var salesScoring []domain.SalesScoring
	model := domain.SalesScoring{}

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName() + " ss").
		LeftJoin("dbo.tr_outlet o ON ss.Outlet_code = o.c_outlet").
		Where(sqrl.Eq{"o.c_tamoutlet": tamOutletCode})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &salesScoring, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return salesScoring, nil
}

// GetAllSalesScoring retrieves all sales scoring data with outlet information
func (r *repository) GetAllSalesScoring(ctx context.Context) ([]domain.SalesScoring, error) {
	var salesScoring []domain.SalesScoring
	model := domain.SalesScoring{}

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName() + " ss").
		LeftJoin("dbo.tr_outlet o ON ss.Outlet_code = o.c_outlet")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &salesScoring, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return salesScoring, nil
}

// GetSalesScoringByOutletCode retrieves sales scoring data by outlet code
func (r *repository) GetSalesScoringByOutletCode(ctx context.Context, outletCode string) ([]domain.SalesScoring, error) {
	var salesScoring []domain.SalesScoring
	model := domain.SalesScoring{}

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName() + " ss").
		LeftJoin("dbo.tr_outlet o ON ss.Outlet_code = o.c_outlet").
		Where(sqrl.Eq{"ss.Outlet_code": outletCode})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &salesScoring, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return salesScoring, nil
}
