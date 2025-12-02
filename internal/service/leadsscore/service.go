package usedcar

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateLeadsScore(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsScore) (string, error)
	UpdateLeadsScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadsScore) error
	GetLeadsScore(ctx context.Context, leadsID string) (domain.LeadsScore, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
}

type service struct {
	transactionRepo transactionRepository
	repo            Repository
}

func New(container ServiceContainer) *service {
	return &service{
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
	}
}
