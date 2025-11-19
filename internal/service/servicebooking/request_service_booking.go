package servicebooking

import (
	"context"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/pkg/constants"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) RequestServiceBooking(ctx context.Context, event servicebooking.ServiceBookingEvent) error {
	// Validate service category
	if _, ok := constants.ServiceCategoryMap[event.Data.ServiceBookingRequest.ServiceCategory]; !ok {
		return errorx.ErrServiceBookingCategoryInvalid
	}

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

	// Check if customer already has an active periodic maintenance booking
	if event.Data.ServiceBookingRequest.ServiceCategory == constants.ServiceCategoryPeriodicMaintenance {
		err := s.checkActivePeriodicMaintenance(ctx, customerID)
		if err != nil {
			return err
		}
	}

	customerVehicleID, err := s.customerVehicleSvc.UpsertCustomerVehicle(ctx, tx, customerID, event.Data.OneAccount.OneAccountID, event.Data.CustomerVehicle)
	if err != nil {
		return err
	}

	serviceBookingID, err := s.UpsertServiceBooking(ctx, tx, customerID, customerVehicleID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingWarranty(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingRecalls(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingJobs(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingParts(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingVehicleInsurance(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.handleServiceBookingDamageImages(ctx, tx, serviceBookingID, event)
	if err != nil {
		return err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	// TODO: call store procedure to sync to external dms after sales system

	return nil
}

func (s *service) handleServiceBookingWarranty(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	err := s.repo.DeleteServiceBookingWarranty(ctx, tx, servicebooking.DeleteServiceBookingWarranty{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}
	for _, warranty := range event.Data.ServiceBookingRequest.Warranty {
		serviceWarranty := warranty.ToModel(serviceBookingID)
		err := s.repo.CreateServiceBookingWarranty(ctx, tx, &serviceWarranty)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) handleServiceBookingRecalls(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	err := s.repo.DeleteServiceBookingRecall(ctx, tx, servicebooking.DeleteServiceBookingRecall{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}
	for _, recall := range event.Data.ServiceBookingRequest.Recalls {
		recallPart := recall.ToModel(serviceBookingID)
		err := s.repo.CreateServiceBookingRecall(ctx, tx, &recallPart)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) handleServiceBookingJobs(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	err := s.repo.DeleteServiceBookingJob(ctx, tx, servicebooking.DeleteServiceBookingJob{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}
	for _, job := range event.Data.Job {
		serviceJob := job.ToDomain(serviceBookingID)
		err := s.repo.CreateServiceBookingJob(ctx, tx, &serviceJob)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) handleServiceBookingParts(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	err := s.repo.DeleteServiceBookingPart(ctx, tx, servicebooking.DeleteServiceBookingPart{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}
	err = s.repo.DeleteServiceBookingPartItem(ctx, tx, servicebooking.DeleteServiceBookingPartItem{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}
	for _, part := range event.Data.Part {
		part, partItem := part.ToDomain(serviceBookingID)
		err := s.repo.CreateServiceBookingPart(ctx, tx, &part)
		if err != nil {
			return err
		}

		if len(partItem) > 0 {
			for _, item := range partItem {
				err := s.repo.CreateServiceBookingPartItem(ctx, tx, serviceBookingID, part.ID, &item)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// checkActivePeriodicMaintenance checks if the customer has any active periodic maintenance bookings
func (s *service) checkActivePeriodicMaintenance(ctx context.Context, customerID string) error {
	// Get all service bookings for the customer with periodic maintenance category
	// and active statuses
	bookings, err := s.repo.GetServiceBookings(ctx, servicebooking.GetServiceBooking{
		CustomerID:      utils.ToPointer(customerID),
		ServiceCategory: utils.ToPointer(constants.ServiceCategoryPeriodicMaintenance),
	})
	if err != nil {
		return err
	}

	// Check if any booking has an active status
	for _, booking := range bookings {
		for _, activeStatus := range constants.ActiveServiceBookingStatuses {
			if booking.ServiceBookingStatus == activeStatus {
				return errorx.ErrServiceBookingCustomerHasActive
			}
		}
	}

	return nil
}

func (s *service) handleServiceBookingVehicleInsurance(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	// If no vehicle insurance data is provided, skip processing
	if event.Data.VehicleInsurance == nil {
		return nil
	}

	// Delete existing vehicle insurance and policies
	err := s.repo.DeleteServiceBookingVehicleInsurancePolicy(ctx, tx, servicebooking.DeleteServiceBookingVehicleInsurancePolicy{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}

	err = s.repo.DeleteServiceBookingVehicleInsurance(ctx, tx, servicebooking.DeleteServiceBookingVehicleInsurance{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}

	vehicleInsurance := event.Data.VehicleInsurance.ToModel(serviceBookingID)
	err = s.repo.CreateServiceBookingVehicleInsurance(ctx, tx, &vehicleInsurance)
	if err != nil {
		return err
	}

	// Create insurance policies
	for _, policy := range event.Data.VehicleInsurance.Policies {
		policyModel := policy.ToModel(vehicleInsurance.ID, serviceBookingID)
		err := s.repo.CreateServiceBookingVehicleInsurancePolicy(ctx, tx, &policyModel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) handleServiceBookingDamageImages(ctx context.Context, tx *sqlx.Tx, serviceBookingID string, event servicebooking.ServiceBookingEvent) error {
	if len(event.Data.ServiceBookingRequest.DamageImage) == 0 {
		return nil
	}

	// Delete existing damage images
	err := s.repo.DeleteServiceBookingImage(ctx, tx, servicebooking.DeleteServiceBookingDamageImage{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return err
	}

	// Create new damage image records
	// TODO: Integrate with object storage service to upload base64 images and get URLs
	// For now, we store the base64 strings directly (as per requirement to expose a hook for future implementation)
	for _, imageData := range event.Data.ServiceBookingRequest.DamageImage {
		damageImage := domain.ServiceBookingImage{
			ServiceBookingID: serviceBookingID,
			ImageURL:         imageData, // TODO: Replace with uploaded URL from object storage
			CreatedAt:        time.Now().UTC(),
			CreatedBy:        constants.System,
			UpdatedAt:        time.Now().UTC(),
			UpdatedBy:        constants.System,
		}
		err := s.repo.CreateServiceBookingImage(ctx, tx, &damageImage)
		if err != nil {
			return err
		}
	}

	return nil
}
