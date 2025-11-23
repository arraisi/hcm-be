package engine

//go:generate mockgen -package=engine -source=service.go -destination=service_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/roleads"
	"github.com/jmoiron/sqlx"
)

type CustomerVehicleService interface {
	GetCustomerVehiclePaginated(ctx context.Context, req customervehicle.GetCustomerVehiclePaginatedRequest) ([]domain.CustomerVehicle, bool, error)
}

type TransactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type RoLeadsRepository interface {
	CreateRoLeads(ctx context.Context, tx *sqlx.Tx, req []domain.RoLeads) error
	GetRoLeads(ctx context.Context, req roleads.GetRoLeadsRequest) (domain.RoLeads, error)
	DeleteRoLeads(ctx context.Context, tx *sqlx.Tx, req domain.RoLeads) error
}

type service struct {
	transactionRepo    TransactionRepository
	roLeadsRepo        RoLeadsRepository
	customerVehicleSvc CustomerVehicleService
}

func New(transactionRepo TransactionRepository, roLeadsRepo RoLeadsRepository, customerVehicleSvc CustomerVehicleService) *service {
	return &service{
		customerVehicleSvc: customerVehicleSvc,
		roLeadsRepo:        roLeadsRepo,
		transactionRepo:    transactionRepo,
	}
}
