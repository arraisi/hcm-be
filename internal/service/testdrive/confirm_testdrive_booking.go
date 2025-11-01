package testdrive

import (
	"context"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (s *service) ConfirmTestDriveBooking(ctx context.Context, request testdrive.ConfirmTestDriveBookingRequest) error {
	testDrive, err := s.repo.GetTestDrive(ctx, testdrive.GetTestDriveRequest{
		TestDriveID: utils.ToPointer(request.TestDriveID),
	})
	if err != nil {
		return err
	}

	testDrive.Status = constants.TestDriveBookingStatusConfirmed

	customerData, err := s.customerRepo.GetCustomer(ctx, customer.GetCustomerRequest{
		CustomerID: testDrive.CustomerID,
	})
	if err != nil {
		return err
	}

	leadsData, err := s.leadRepo.GetLeads(ctx, leads.GetLeadsRequest{
		LeadsID: utils.ToPointer(testDrive.LeadsID),
	})
	if err != nil {
		return err
	}

	tdEventConfirmRequest := testdrive.TestDriveEvent{
		Process:   "test_drive_confirm",
		EventID:   testDrive.EventID,
		Timestamp: time.Now().Unix(),
		Data: testdrive.TestDriveEventData{
			OneAccount: customer.NewOneAccountRequest(customerData),
			TestDrive:  testdrive.NewTestDriveRequest(testDrive),
			Leads:      leads.NewLeadsRequest(leadsData),
			PICAssignment: utils.ToPointer(testdrive.PICAssignmentRequest{
				EmployeeID: request.EmployeeID,
			}),
		},
	}

	fmt.Printf("%+v", tdEventConfirmRequest)

	return nil
}
