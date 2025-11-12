package servicebooking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

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
		Process:   "service_booking_confirm",
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

	marshal, err := json.Marshal(sbEventConfirmRequest)
	if err != nil {
		return err
	}
	fmt.Printf("Service Booking Confirm Request: %s\n", string(marshal))

	err = s.apimDIDXSvc.ConfirmServiceBooking(ctx, sbEventConfirmRequest)
	if err != nil {
		return err
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
