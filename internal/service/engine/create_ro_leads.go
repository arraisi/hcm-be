package engine

//go:generate mockgen -package=engine -source=customer_segmentation.go -destination=customer_segmentation_mock_test.go
import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/arraisi/hcm-be/internal/domain/dto/roleads"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (s *service) CreateRoLeads(ctx context.Context, request engine.CreateRoLeadsRequest) error {
	req := customervehicle.GetCustomerVehiclePaginatedRequest{
		Limit:                   100,
		DecDateNotNull:          true,
		CarPaymentStatusNotNull: true,
		OutletCodeNotNull:       true,
		SalesNikNotNull:         true,
	}

	hasMore := true
	for hasMore {
		vehicles, more, err := s.customerVehicleSvc.GetCustomerVehiclePaginated(ctx, req)
		if err != nil {
			return err
		}

		// Create RO Leads for the fetched vehicles
		_, err = s.createRoLeads(ctx, request, vehicles)
		if err != nil {
			return err
		}

		hasMore = more

		if hasMore {
			req.Offset += req.Limit
		}
	}

	return nil
}

func (s *service) createRoLeads(ctx context.Context, request engine.CreateRoLeadsRequest, vehicles []domain.CustomerVehicle) ([]domain.RoLeads, error) {
	roLeads := make([]domain.RoLeads, 0, len(vehicles))
	roLeadsToBeDelete := make([]domain.RoLeads, 0)
	for _, vehicle := range vehicles {
		leads, err := s.getRoLeadsForVehicleThisMonth(ctx, vehicle)
		if err != nil {
			return nil, err
		}

		if leads.ID != "" && !request.ForceUpdate {
			// Skip if RO Leads already exists for this vehicle this month
			continue
		}
		if leads.ID != "" && request.ForceUpdate {
			// Mark existing leads for deletion if force update is true
			roLeadsToBeDelete = append(roLeadsToBeDelete, leads)
		}

		carAge := time.Now().Year() - vehicle.DecDate.Year()
		if carAge <= 0 {
			carAge = 1
		}

		createdAt := time.Now()

		roLead := domain.RoLeads{
			CustomerVehicleID:       vehicle.ID,
			CarAge:                  carAge,
			CarAgeScore:             s.getCategoryScore(carAge),
			CarPaymentStatusScore:   s.getPaymentStatusScore(utils.ToValue(vehicle.CarPaymentStatus)),
			CarServiceActivityScore: s.getServiceActivityScore(ctx, vehicle),
			CarServiceScore:         s.getServiceScore(ctx, vehicle),
			CustomerResponse:        engine.CustomerResponseNotSent,
			LeadsOutlet:             utils.ToValue(vehicle.OutletCode),
			LeadsSalesNik:           utils.ToValue(vehicle.SalesNik),
			CreatedAt:               createdAt,
			UpdatedAt:               &createdAt,
		}
		roLead.RoScore = s.calculateRoScore(&roLead)
		roLeads = append(roLeads, roLead)
	}

	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return roLeads, err
	}
	defer func() {
		_ = s.transactionRepo.RollbackTransaction(tx)
	}()

	for _, roLead := range roLeadsToBeDelete {
		err = s.roLeadsRepo.DeleteRoLeads(ctx, tx, roLead)
		if err != nil {
			return roLeads, err
		}
	}

	err = s.roLeadsRepo.CreateRoLeads(ctx, tx, roLeads)
	if err != nil {
		return roLeads, err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return roLeads, err
	}

	// TODO: send repeat order offers via WhatsApp

	return roLeads, nil
}

func (s *service) getRoLeadsForVehicleThisMonth(ctx context.Context, vehicle domain.CustomerVehicle) (domain.RoLeads, error) {
	leads, err := s.roLeadsRepo.GetRoLeads(ctx, roleads.GetRoLeadsRequest{
		CustomerVehicleID: utils.ToPointer(vehicle.ID),
		CreatedThisMonth:  true,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return domain.RoLeads{}, err
	}

	return leads, nil
}

func (s *service) getCategoryScore(q int) int {
	switch {
	case q < 4:
		return 0
	case q < 5:
		return 30
	case q < 6:
		return 60
	default:
		return 100
	}
}

func (s *service) getPaymentStatusScore(q string) int {
	switch q {
	case "Kredit lunas":
		return 50
	case "Cash":
		return 100
	default:
		return 0
	}
}

func (s *service) getServiceActivityScore(ctx context.Context, vehicle domain.CustomerVehicle) int {
	// TODO: get service booking with status confirmed in last 2 months by slot date
	// if len service bookings >= 2 return 100
	// if len service bookings == 1 return 50
	// else 0
	return 0
}

func (s *service) getServiceScore(ctx context.Context, vehicle domain.CustomerVehicle) int {
	// TODO: get service booking with status confirmed in last 12 months by slot date and is_major true
	// if len service bookings > 0 return 100
	// if len service bookings == 0 return 50
	return 0
}

func (s *service) calculateRoScore(rodata *domain.RoLeads) int {
	weight := domain.RoScoreWeight
	rodata.RoScore = int(float64(rodata.CarAgeScore)*weight["car_age"] +
		float64(rodata.CarPaymentStatusScore)*weight["car_payment_status"] +
		float64(rodata.CarServiceActivityScore)*weight["car_service_activity"] +
		float64(rodata.CarServiceScore)*weight["car_service"])
	return rodata.RoScore
}
