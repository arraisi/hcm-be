package servicebooking

import (
	"context"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// ConfirmServiceBookingGR processes General Repair service booking confirmation from webhook event
func (s *service) ConfirmServiceBookingGR(ctx context.Context, request servicebooking.ServiceBookingEvent) error {
	return s.confirmServiceBookingFromEvent(ctx, request)
}

// ConfirmServiceBookingBP processes Body and Paint service booking confirmation from webhook event
func (s *service) ConfirmServiceBookingBP(ctx context.Context, request servicebooking.ServiceBookingEvent) error {
	return s.confirmServiceBookingFromEvent(ctx, request)
}

// confirmServiceBookingFromEvent contains the shared logic for confirming service bookings from webhook events
// This method handles both GR (General Repair) and BP (Body and Paint) bookings as they share the same process
func (s *service) confirmServiceBookingFromEvent(ctx context.Context, request servicebooking.ServiceBookingEvent) error {
	// Start transaction
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Upsert customer
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, request.Data.OneAccount)
	if err != nil {
		return err
	}

	// Upsert customer vehicle with one_account_ID
	customerVehicleID, err := s.customerVehicleSvc.UpsertCustomerVehicle(ctx, tx, customerID, request.Data.OneAccount.OneAccountID, request.Data.CustomerVehicle)
	if err != nil {
		return err
	}

	// Get existing service booking
	existingServiceBooking, err := s.repo.GetServiceBooking(ctx, servicebooking.GetServiceBooking{
		ServiceBookingID: utils.ToPointer(request.Data.ServiceBookingRequest.BookingId),
	})
	if err != nil {
		return err
	}

	// Convert to service booking model and update with new data
	serviceBookingModel := request.ToServiceBookingModel(customerID, customerVehicleID)
	serviceBookingModel.ID = existingServiceBooking.ID // Preserve existing ID
	serviceBookingModel.ServiceBookingStatus = request.Data.ServiceBookingRequest.BookingStatus

	// Update service booking in database
	err = s.repo.UpdateServiceBooking(ctx, tx, serviceBookingModel)
	if err != nil {
		return err
	}

	serviceBookingID := existingServiceBooking.ID

	// Update jobs using existing handler
	err = s.handleServiceBookingJobs(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Update parts using existing handler
	err = s.handleServiceBookingParts(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Update warranties using existing handler
	err = s.handleServiceBookingWarranty(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Update recalls using existing handler
	err = s.handleServiceBookingRecalls(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Update vehicle insurance using existing handler
	err = s.handleServiceBookingVehicleInsurance(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Update damage images using existing handler
	err = s.handleServiceBookingDamageImages(ctx, tx, serviceBookingID, request)
	if err != nil {
		return err
	}

	// Commit transaction
	if err = s.transactionRepo.CommitTransaction(tx); err != nil {
		return err
	}

	// Enqueue the task to Asynq for external API call
	payload := queue.DIDXServiceBookingConfirmPayload{
		ServiceBookingEvent: request,
	}

	err = s.queueClient.EnqueueDIDXServiceBookingConfirm(context.Background(), payload)
	if err != nil {
		return fmt.Errorf("failed to enqueue DIDX confirm task: %w", err)
	}

	return nil
}

// ConfirmServiceBooking deprecated
func (s *service) ConfirmServiceBooking(ctx context.Context, request servicebooking.ConfirmServiceBookingRequest) error {
	serviceBooking, err := s.repo.GetServiceBooking(ctx, servicebooking.GetServiceBooking{
		ServiceBookingID: utils.ToPointer(request.ServiceBookingID),
	})
	if err != nil {
		return err
	}

	serviceBooking.ServiceBookingStatus = constants.ServiceBookingStatusManuallyConfirmed

	customerData, err := s.customerRepo.GetCustomer(ctx, customer.GetCustomerRequest{
		CustomerID: serviceBooking.CustomerID,
	})
	if err != nil {
		return err
	}

	customerVehicle, err := s.customerVehicleSvc.GetCustomerVehicle(ctx, customervehicle.GetCustomerVehicleRequest{
		ID: utils.ToPointer(serviceBooking.CustomerVehicleID),
	})
	if err != nil {
		return err
	}

	jobs, err := s.getJobs(ctx, serviceBooking.ID)
	if err != nil {
		return err
	}

	part, partItems, err := s.getPartsAndItems(ctx, serviceBooking.ID)
	if err != nil {
		return err
	}
	parts := servicebooking.NewPartsRequest(part, partItems)

	warranties, err := s.getWarranties(ctx, serviceBooking.ID)
	if err != nil {
		return err
	}

	serviceBookingRecalls, err := s.getServiceBookingRecalls(ctx, serviceBooking.ID)
	if err != nil {
		return err
	}
	recalls := servicebooking.NewRecallsRequest(serviceBookingRecalls)

	vehicleInsurance, err := s.getVehicleInsurance(ctx, serviceBooking.ID)
	if err != nil {
		return err
	}

	sbEventConfirmRequest := servicebooking.ServiceBookingEvent{
		Process:   serviceBooking.BookingType,
		EventID:   serviceBooking.EventID,
		Timestamp: int(time.Now().Unix()),
		Data: servicebooking.DataRequest{
			OneAccount:            customer.NewOneAccountRequest(customerData),
			ServiceBookingRequest: servicebooking.NewServiceBookingRequest(serviceBooking, warranties, recalls),
			CustomerVehicle:       customervehicle.NewCustomerVehicleRequest(customerVehicle),
			Job:                   jobs,
			Part:                  parts,
			VehicleInsurance:      vehicleInsurance,
		},
	}

	// Enqueue the task to Asynq instead of calling DIDX directly
	// Use context.Background() to ensure the enqueue operation completes
	// even if the parent request context is cancelled
	payload := queue.DIDXServiceBookingConfirmPayload{
		ServiceBookingEvent: sbEventConfirmRequest,
	}

	err = s.queueClient.EnqueueDIDXServiceBookingConfirm(context.Background(), payload)
	if err != nil {
		return fmt.Errorf("failed to enqueue DIDX confirm task: %w", err)
	}

	return nil
}

func (s *service) getServiceBookingRecalls(ctx context.Context, serviceBookingID string) ([]domain.ServiceBookingRecall, error) {
	serviceBookingRecalls, err := s.repo.GetServiceBookingRecalls(ctx, servicebooking.GetServiceBookingRecall{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return nil, err
	}
	return serviceBookingRecalls, nil
}

func (s *service) getWarranties(ctx context.Context, serviceBookingID string) ([]servicebooking.WarrantyRequest, error) {
	serviceBookingWarranties, err := s.repo.GetServiceBookingWarranties(ctx, servicebooking.GetServiceBookingWarranty{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return nil, err
	}
	warranties := servicebooking.NewWarrantiesRequest(serviceBookingWarranties)
	return warranties, nil
}

func (s *service) getJobs(ctx context.Context, serviceBookingID string) ([]servicebooking.JobRequest, error) {
	job, err := s.repo.GetServiceBookingJobs(ctx, servicebooking.GetServiceBookingJob{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return nil, err
	}
	jobs := servicebooking.NewJobsRequest(job)
	return jobs, nil
}

func (s *service) getPartsAndItems(ctx context.Context, serviceBookingID string) ([]domain.ServiceBookingPart, []domain.ServiceBookingPartItem, error) {
	part, err := s.repo.GetServiceBookingParts(ctx, servicebooking.GetServiceBookingPart{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return nil, nil, err
	}

	var partItems []domain.ServiceBookingPartItem
	for _, p := range part {
		if p.PartType == "PACKAGE" {
			items, err := s.repo.GetServiceBookingPartItems(ctx, servicebooking.GetServiceBookingPartItem{
				ServiceBookingPartID: utils.ToPointer(p.ID),
			})
			if err != nil {
				return nil, nil, err
			}
			partItems = append(partItems, items...)
		}
	}
	return part, partItems, nil
}

func (s *service) getVehicleInsurance(ctx context.Context, serviceBookingID string) (*servicebooking.VehicleInsuranceRequest, error) {
	// Get vehicle insurance
	vehicleInsurance, err := s.repo.GetServiceBookingVehicleInsurance(ctx, servicebooking.GetServiceBookingVehicleInsurance{
		ServiceBookingID: utils.ToPointer(serviceBookingID),
	})
	if err != nil {
		return nil, err
	}

	// If no insurance data found, return nil (not an error)
	if vehicleInsurance.ID == "" {
		return nil, nil
	}

	// Get vehicle insurance policies
	policies, err := s.repo.GetServiceBookingVehicleInsurancePolicies(ctx, servicebooking.GetServiceBookingVehicleInsurancePolicy{
		VehicleInsuranceID: utils.ToPointer(vehicleInsurance.ID),
	})
	if err != nil {
		return nil, err
	}

	// Convert to request DTO
	return servicebooking.NewVehicleInsuranceRequest(vehicleInsurance, policies), nil
}
