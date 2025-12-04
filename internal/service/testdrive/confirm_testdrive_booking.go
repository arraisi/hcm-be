package testdrive

import (
	"context"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/employee"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// ConfirmTestDrive processes test drive confirmation from webhook event
func (s *service) ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error {
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

	// Get existing test drive
	existingTestDrive, err := s.repo.GetTestDrive(ctx, testdrive.GetTestDriveRequest{
		TestDriveID: utils.ToPointer(request.Data.TestDrive.TestDriveID),
	})
	if err != nil {
		return err
	}

	// Convert to test drive model and update with new data
	testDriveModel := request.ToTestDriveModel(customerID)
	testDriveModel.ID = existingTestDrive.ID // Preserve existing ID
	testDriveModel.Status = request.Data.TestDrive.TestDriveStatus

	// Update test drive in database
	err = s.repo.UpdateTestDrive(ctx, tx, testDriveModel)
	if err != nil {
		return err
	}

	// Update leads
	leadsModel := request.Data.Leads.ToDomain(customerID)

	// Get existing leads to preserve ID
	existingLeads, err := s.leadRepo.GetLeads(ctx, leads.GetLeadsRequest{
		LeadsID: utils.ToPointer(request.Data.Leads.LeadsID),
	})
	if err != nil {
		return err
	}
	leadsModel.ID = existingLeads.ID

	err = s.leadRepo.UpdateLeads(ctx, tx, leadsModel)
	if err != nil {
		return err
	}

	// Commit transaction
	if err = s.transactionRepo.CommitTransaction(tx); err != nil {
		return err
	}

	// If there's PIC assignment, enqueue task to send confirmation to external API
	// Enqueue the task to Asynq for external API call
	err = s.queueClient.EnqueueDIDXTestDriveConfirm(context.Background(), queue.DIDXTestDriveConfirmPayload{
		TestDriveEvent: request,
	})
	if err != nil {
		return err
	}

	return nil
}

// ConfirmTestDriveBooking deprecated
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
		LeadsID: testDrive.LeadsID,
	})
	if err != nil {
		return err
	}

	employeeData, err := s.employeeRepo.GetEmployee(ctx, employee.GetEmployeeRequest{
		EmployeeID: utils.ToPointer(request.EmployeeID),
	})
	if err != nil {
		return err
	}

	tdEventConfirmRequest := testdrive.TestDriveEvent{
		Process:   constants.ProcessTestDriveConfirmed,
		EventID:   testDrive.EventID,
		Timestamp: time.Now().Unix(),
		Data: testdrive.TestDriveEventData{
			OneAccount: customer.NewOneAccountRequest(customerData),
			TestDrive:  testdrive.NewTestDriveRequest(testDrive),
			Leads:      leads.NewLeadsRequest(leadsData),
			PICAssignment: utils.ToPointer(testdrive.PICAssignmentRequest{
				EmployeeID: employeeData.EmployeeID,
				FirstName:  employeeData.EmployeeName,
			}),
		},
	}

	err = s.apimDIDXSvc.Confirm(ctx, tdEventConfirmRequest)
	if err != nil {
		return err
	}

	return nil
}
