package usedcar

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/usedcar"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateUsedCar(ctx context.Context, tx *sqlx.Tx, req *domain.UsedCar) (int, error)
	UpdateUsedCar(ctx context.Context, tx *sqlx.Tx, req domain.UsedCar) error
	GetUsedCar(ctx context.Context, req usedcar.GetUsedCarRequest) (domain.UsedCar, error)
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
