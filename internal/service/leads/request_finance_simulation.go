package leads

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	dtoLeads "github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) RequestFinanceSimulation(ctx context.Context, request dtoLeads.FinanceSimulationWebhookEvent) error {
	outletData, err := s.outletRepo.GetOutletCodeByTAMOutletID(ctx, request.Data.Leads.OutletID)
	if err != nil {
		return err
	}

	// Start transaction
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// 1. Upsert customer
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, request.Data.OneAccount, hasjratid.GenerateRequest{
		SourceCode:       "H",
		CustomerType:     "personal",
		TamOutletID:      request.Data.Leads.OutletID,
		OutletCode:       outletData.OutletCode,
		RegistrationDate: time.Now().Unix(),
	})
	if err != nil {
		return fmt.Errorf("failed to upsert customer: %w", err)
	}

	// 2. Upsert leads
	leadsData := request.Data.Leads
	leads, err := s.processLeadsFinanceSim(ctx, tx, leadsData, customerID)
	if err != nil {
		return err
	}

	// 3. Process interested parts
	if err := s.processInterestedParts(ctx, tx, leads.ID, request.Data.Leads.InterestedPart); err != nil {
		return err
	}

	// 4. Process finance simulation
	if err := s.processLeadsFinanceSimulation(ctx, tx, leads.ID, leadsData.FinanceSimulationID, leadsData.FinanceSimulationNumber, request.Data.FinanceSimulation); err != nil {
		return err
	}

	// 5. Upsert trade-in
	if err := s.processTradeIn(ctx, tx, leads.ID, request.Data.TradeIn); err != nil {
		return err
	}

	// Commit transaction
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *service) processLeadsFinanceSim(ctx context.Context, tx *sqlx.Tx, leadsData dtoLeads.FinanceSimulationLeadsRequest, customerID string) (domain.Leads, error) {
	leads := leadsData.ToDomain(customerID)
	leads.CreatedBy = "SYSTEM"
	leads.UpdatedAt = time.Now()
	leads.UpdatedBy = utils.ToPointer("SYSTEM")

	existingLeads, err := s.leadsRepo.GetLeads(ctx, dtoLeads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsData.LeadsID),
	})
	if err != nil && err != sql.ErrNoRows {
		return leads, fmt.Errorf("failed to check existing leads: %w", err)
	}

	if existingLeads.ID != "" {
		leads.ID = existingLeads.ID
		if err := s.leadsRepo.UpdateLeads(ctx, tx, leads); err != nil {
			return leads, errors.ErrLeadsUpdateFailed
		}
	} else {
		leads.CreatedAt = time.Now()
		if err := s.leadsRepo.CreateLeads(ctx, tx, &leads); err != nil {
			return leads, errors.ErrLeadsCreateFailed
		}
	}
	return leads, nil
}

// processLeadsFinanceSimulation handles upsert logic for finance simulation and credits
func (s *service) processLeadsFinanceSimulation(ctx context.Context, tx *sqlx.Tx, leadsID, financeSimulationID, financeSimulationNumber string, financeSimData dtoLeads.FinanceSimulationDetailsRequest) error {
	financeSimulation := financeSimData.ToDomain(financeSimulationID, financeSimulationNumber, leadsID)

	// Check if finance simulation already exists
	existingFinanceSim, err := s.financeSimulationRepo.GetFinanceSimulation(ctx, dtoLeads.GetFinanceSimulationRequest{
		SimulationID: utils.ToPointer(financeSimulationID),
		LeadsID:      utils.ToPointer(leadsID),
	})
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing finance simulation: %w", err)
	}

	if existingFinanceSim.ID != "" {
		// Update existing finance simulation
		financeSimulation.ID = existingFinanceSim.ID
		if err := s.financeSimulationRepo.UpdateFinanceSimulation(ctx, tx, financeSimulation); err != nil {
			return errors.ErrFinanceSimulationUpdateFailed
		}
	} else {
		// Create new finance simulation
		if err := s.financeSimulationRepo.CreateFinanceSimulation(ctx, tx, &financeSimulation); err != nil {
			return errors.ErrFinanceSimulationCreateFailed
		}
	}

	// Delete existing credits before inserting new ones
	if err := s.financeSimulationRepo.DeleteCreditsByLeadsID(ctx, tx, leadsID); err != nil {
		return fmt.Errorf("failed to delete finance simulation credits: %w", err)
	}

	// Insert new credits
	for _, creditResult := range financeSimData.CreditSimulationResults {
		credit := creditResult.ToDomain(leadsID, financeSimulation.ID)
		if err := s.financeSimulationRepo.CreateFinanceSimulationCredit(ctx, tx, &credit); err != nil {
			return errors.ErrFinanceSimulationCreditCreateFailed
		}
	}

	return nil
}
