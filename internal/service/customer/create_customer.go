package customer

import (
	"context"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
)

func (s *service) CreateCustomer(ctx context.Context, request customer.CreateCustomerRequest) (customer.CreateCustomerResponse, error) {
	// Generate Hasjrat ID
	haID, err := s.hasjratIDSvc.GenerateHasjratID(ctx, hasjratid.GenerateRequest{
		SourceCode:       "H",
		CustomerType:     "personal",
		OutletCode:       request.OutletID,
		RegistrationDate: time.Now().Unix(),
	})
	if err != nil {
		return customer.CreateCustomerResponse{}, err
	}
	request.HasjratID = haID

	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return customer.CreateCustomerResponse{}, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = s.UpsertCustomerV2(ctx, tx, request.ToDomain())
	if err != nil {
		return customer.CreateCustomerResponse{}, err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return customer.CreateCustomerResponse{}, err
	}

	return customer.CreateCustomerResponse{
		HasjratID: haID,
	}, nil
}
