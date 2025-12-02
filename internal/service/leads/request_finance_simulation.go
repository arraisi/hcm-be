package leads

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	dtoLeads "github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (s *service) RequestFinanceSimulation(ctx context.Context, request dtoLeads.FinanceSimulationWebhookEvent) error {
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
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, request.Data.OneAccount)
	if err != nil {
		return fmt.Errorf("failed to upsert customer: %w", err)
	}

	// 2. Upsert leads
	leadsData := request.Data.Leads
	leads := leadsData.ToDomain(customerID)
	leads.CreatedBy = "SYSTEM"
	leads.UpdatedAt = time.Now()
	leads.UpdatedBy = utils.ToPointer("SYSTEM")

	// Check if leads already exists
	existingLeads, err := s.leadsRepo.GetLeads(ctx, dtoLeads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsData.LeadsID),
	})
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing leads: %w", err)
	}

	if existingLeads.ID != "" {
		// Update existing leads
		leads.ID = existingLeads.ID
		if err := s.leadsRepo.UpdateLeads(ctx, tx, leads); err != nil {
			return errors.ErrLeadsUpdateFailed
		}
	} else {
		// Create new leads
		leads.CreatedAt = time.Now()
		if err := s.leadsRepo.CreateLeads(ctx, tx, &leads); err != nil {
			return errors.ErrLeadsCreateFailed
		}
	}

	// TODO: delete insert interested_part
	// TODO: delete insert package_parts/interested_part_item

	// 3. Create finance simulation
	financeSimData := request.Data.FinanceSimulation
	financeSimulation := financeSimData.ToDomain(leadsData.FinanceSimulationID, leadsData.FinanceSimulationNumber, leadsData.LeadsID)

	if err := s.financeSimulationRepo.CreateFinanceSimulation(ctx, tx, &financeSimulation); err != nil {
		return errors.ErrFinanceSimulationCreateFailed
	}
	// TODO:delete insert finance_simulation_credit

	// 4. Create trade-in if flag is true
	tradeInData := request.Data.TradeIn
	tradeIn := tradeInData.ToDomain(leadsData.LeadsID)

	if err := s.tradeInRepo.CreateTradeIn(ctx, tx, &tradeIn); err != nil {
		return errors.ErrTradeInCreateFailed
	}

	// Commit transaction
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
