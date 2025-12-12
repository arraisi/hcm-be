package customerreminder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/customerreminder"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (s *service) CreateCustomerReminder(ctx context.Context, request customerreminder.Request) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if !outletAllowed(s.cfg.Condition.OutletIDs, request.Data.OutletID) {
		// outlet not allowed → nothing to do, just commit empty tx
		return s.transactionRepo.CommitTransaction(tx)
	}

	// ---- Process each reminder ----
	for _, reminder := range request.Data.Reminders {
		// 1) Upsert customer
		c := reminder.OneAccount.ToCustomerModel()
		customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, c)
		if err != nil {
			return err
		}

		// 2) Upsert customer vehicle
		cv := reminder.CustomerVehicle.ToCustomerVehicleModel()
		cv.CustomerID = customerID
		cv.OneAccountID = utils.ToValue(c.OneAccountID)
		customerVehicleID, err := s.customerVehicleSvc.UpsertCustomerVehicleV2(ctx, tx, cv)
		if err != nil {
			return err
		}

		// 3) Upsert customer reminder
		cr := reminder.ReminderDetail.ToDomainCustomerReminder()
		cr.CustomerID = customerID
		cr.CustomerVehicleID = customerVehicleID
		cr.OutletID = request.Data.OutletID
		if _, err := s.UpsertCustomerReminder(ctx, tx, cr); err != nil {
			return err
		}
	}

	return s.transactionRepo.CommitTransaction(tx)
}

func outletAllowed(allowed []string, outletID string) bool {
	for _, v := range allowed {
		if v == "ALL" || v == outletID {
			return true
		}
	}
	return false
}
