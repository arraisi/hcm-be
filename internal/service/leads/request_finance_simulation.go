package leads

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	dtoLeads "github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
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
	leads, err := s.processLeads(ctx, tx, leadsData, customerID)
	if err != nil {
		return err
	}

	// Delete existing interested parts before inserting new ones
	if err := s.interestedPartRepo.DeleteInterestedPartItemsByLeadsID(ctx, tx, leads.ID); err != nil {
		return fmt.Errorf("failed to delete interested part items: %w", err)
	}
	if err := s.interestedPartRepo.DeleteInterestedPartByLeadsID(ctx, tx, leads.ID); err != nil {
		return fmt.Errorf("failed to delete interested parts: %w", err)
	}

	for _, interestedPart := range request.Data.Leads.InterestedPart {
		// Create interested part
		part := interestedPart.ToDomain(leads.ID)
		if err := s.interestedPartRepo.CreateInterestedPart(ctx, tx, &part); err != nil {
			return errors.ErrInterestedPartCreateFailed
		}

		// Create package parts if it's a package type
		if interestedPart.InterestedPartType == "PACKAGE" && len(interestedPart.PackageParts) > 0 {
			for _, packagePart := range interestedPart.PackageParts {
				item := packagePart.ToDomain(leads.ID, part.ID)
				if err := s.interestedPartRepo.CreateInterestedPartItem(ctx, tx, &item); err != nil {
					return errors.ErrInterestedPartItemCreateFailed
				}
			}
		}
	}

	// 3. Upsert finance simulation
	financeSimData := request.Data.FinanceSimulation
	financeSimulation := financeSimData.ToDomain(leadsData.FinanceSimulationID, leadsData.FinanceSimulationNumber, leads.ID)

	// Check if finance simulation already exists
	existingFinanceSim, err := s.financeSimulationRepo.GetFinanceSimulation(ctx, dtoLeads.GetFinanceSimulationRequest{
		SimulationID: utils.ToPointer(leadsData.FinanceSimulationID),
		LeadsID:      utils.ToPointer(leads.ID),
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
	if err := s.financeSimulationRepo.DeleteCreditsByLeadsID(ctx, tx, leads.ID); err != nil {
		return fmt.Errorf("failed to delete finance simulation credits: %w", err)
	}

	for _, creditResult := range financeSimData.CreditSimulationResults {
		credit := creditResult.ToDomain(leads.ID, financeSimulation.ID)
		if err := s.financeSimulationRepo.CreateFinanceSimulationCredit(ctx, tx, &credit); err != nil {
			return errors.ErrFinanceSimulationCreditCreateFailed
		}
	}

	// 4. Upsert trade-in
	tradeInData := request.Data.TradeIn
	tradeIn := tradeInData.ToDomain(leads.ID)

	// Check if trade-in already exists
	existingTradeIn, err := s.tradeInRepo.GetTradeIn(ctx, dtoLeads.GetTradeInRequest{
		LeadsID: utils.ToPointer(leads.ID),
	})
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing trade-in: %w", err)
	}

	if existingTradeIn.ID != "" {
		// Update existing trade-in
		tradeIn.ID = existingTradeIn.ID
		if err := s.tradeInRepo.UpdateTradeIn(ctx, tx, tradeIn); err != nil {
			return errors.ErrTradeInUpdateFailed
		}
	} else {
		// Create new trade-in
		if err := s.tradeInRepo.CreateTradeIn(ctx, tx, &tradeIn); err != nil {
			return errors.ErrTradeInCreateFailed
		}
	}

	// Commit transaction
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *service) processLeads(ctx context.Context, tx *sqlx.Tx, leadsData dtoLeads.FinanceSimulationLeadsRequest, customerID string) (domain.Leads, error) {
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
