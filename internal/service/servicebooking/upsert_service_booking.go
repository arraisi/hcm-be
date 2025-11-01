package servicebooking

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertServiceBooking(ctx context.Context, tx *sqlx.Tx, customerID, customerVehicleID string, sb servicebooking.ServiceBookingEvent) (string, error) {
	serviceBooking, err := s.repo.GetServiceBooking(ctx, servicebooking.GetServiceBooking{
		ServiceBookingID:     utils.ToPointer(sb.Data.BookingId),
		ServiceBookingNumber: utils.ToPointer(sb.Data.BookingNumber),
	})
	if err == nil {
		// Found → update
		sb := sb.ToServiceBookingModel(customerID, customerVehicleID)
		sb.ID = serviceBooking.ID

		err = s.repo.UpdateServiceBooking(ctx, tx, sb)
		if err != nil {
			return serviceBooking.ID, err
		}
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		serviceBookingModel := sb.ToServiceBookingModel(customerID, customerVehicleID)
		err := s.repo.CreateServiceBooking(ctx, tx, &serviceBookingModel)
		if err != nil {
			return serviceBookingModel.ID, err
		}
		return serviceBookingModel.ID, nil
	}

	return serviceBooking.ID, err
}
