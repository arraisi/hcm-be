package transaction

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type iDB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type repository struct {
	db iDB
}

// New creates a new transaction repository instance
func New(db iDB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) BeginTransaction(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return r.db.BeginTxx(ctx, opts)
}

func (r *repository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (r *repository) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}