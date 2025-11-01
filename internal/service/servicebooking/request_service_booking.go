package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
)

func (s *service) RequestServiceBooking(ctx context.Context, event servicebooking.ServiceBookingEvent) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.transactionRepo.RollbackTransaction(tx)
	}()

	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, event.Data.OneAccount)
	if err != nil {
		return err
	}

	customerVehicleID, err := s.customerVehicleSvc.UpsertCustomerVehicle(ctx, tx, customerID, event.Data.OneAccount.OneAccountID, event.Data.CustomerVehicle)
	if err != nil {
		return err
	}

	serviceBookingID, err := s.UpsertServiceBooking(ctx, tx, customerID, customerVehicleID, event)
	if err != nil {
		return err
	}

	for _, warranty := range event.Data.Warranty {
		serviceWarranty := warranty.ToModel(serviceBookingID)
		err := s.repo.CreateServiceBookingWarranty(ctx, tx, &serviceWarranty)
		if err != nil {
			return err
		}
	}

	for _, recall := range event.Data.Recalls {
		for _, part := range recall.AffectedParts {
			recallPart := recall.ToModel(serviceBookingID, part)
			err := s.repo.CreateServiceBookingRecall(ctx, tx, &recallPart)
			if err != nil {
				return err
			}
		}
	}

	for _, job := range event.Job {
		serviceJob := job.ToDomain(serviceBookingID)
		err := s.repo.CreateServiceBookingJob(ctx, tx, &serviceJob)
		if err != nil {
			return err
		}
	}

	for _, part := range event.Part {
		part, partItem := part.ToDomain(serviceBookingID)
		err := s.repo.CreateServiceBookingPart(ctx, tx, &part)
		if err != nil {
			return err
		}

		if len(partItem) > 0 {
			for _, item := range partItem {
				err := s.repo.CreateServiceBookingPartItem(ctx, tx, part.ID, &item)
				if err != nil {
					return err
				}
			}
		}
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	return nil
}
