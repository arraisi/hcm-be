package user

//go:generate mockgen -source=repository.go -package=user -destination=repository_mock_test.go
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
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type repository struct {
	db iDB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db iDB) *repository {
	return &repository{
		db: db,
	}
}
