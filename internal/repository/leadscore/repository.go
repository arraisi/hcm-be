package leadscore

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"tabeldata.com/hcm-be/internal/config"
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
	cfg *config.Config
	db  iDB
}

// New creates a new customer repository instance
func New(cfg *config.Config, db iDB) *repository {
	return &repository{
		db:  db,
		cfg: cfg,
	}
}
