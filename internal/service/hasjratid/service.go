package hasjratid

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
	GetNextSequence(ctx context.Context, sourceCode string, customerTypeCode string, outletCode string, year string) (uint64, error)
	GetNextSequenceV2(ctx context.Context, sourceCode string, customerTypeCode string, outletCode string, year string) (uint64, error)
}

type OutletRepository interface {
	GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	OutletRepo      OutletRepository
}

type service struct {
	transactionRepo transactionRepository
	repo            Repository
	outletRepo      OutletRepository
}

func New(container ServiceContainer) *service {
	return &service{
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		outletRepo:      container.OutletRepo,
	}
}
