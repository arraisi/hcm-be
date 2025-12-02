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

	// TODO: delete before insert new insert interested_part
	for _, interestedPart := range request.Data.Leads.InterestedPart {
		// Create interested part
		part := interestedPart.ToDomain(leadsData.LeadsID)
		if err := s.interestedPartRepo.CreateInterestedPart(ctx, tx, &part); err != nil {
			return errors.ErrInterestedPartCreateFailed
		}

		// TODO: delete before insert new package_parts/interested_part_item
		// Create package parts if it's a package type
		if interestedPart.InterestedPartType == "PACKAGE" && len(interestedPart.PackageParts) > 0 {
			for _, packagePart := range interestedPart.PackageParts {
				item := packagePart.ToDomain(leadsData.LeadsID, part.ID)
				if err := s.interestedPartRepo.CreateInterestedPartItem(ctx, tx, &item); err != nil {
					return errors.ErrInterestedPartItemCreateFailed
				}
			}
		}
	}

	// 3. Create finance simulation
	financeSimData := request.Data.FinanceSimulation
	financeSimulation := financeSimData.ToDomain(leadsData.FinanceSimulationID, leadsData.FinanceSimulationNumber, leadsData.LeadsID)

	if err := s.financeSimulationRepo.CreateFinanceSimulation(ctx, tx, &financeSimulation); err != nil {
		return errors.ErrFinanceSimulationCreateFailed
	}
	// TODO:delete before insert new finance_simulation_credit
	for _, creditResult := range financeSimData.CreditSimulationResults {
		credit := creditResult.ToDomain(leadsData.LeadsID, financeSimulation.ID)
		if err := s.financeSimulationRepo.CreateFinanceSimulationCredit(ctx, tx, &credit); err != nil {
			return errors.ErrFinanceSimulationCreditCreateFailed
		}
	}

	// 4. Create trade-in if flag is true
	tradeInData := request.Data.TradeIn
	tradeIn := tradeInData.ToDomain(leadsData.LeadsID)

	// TODO:get trade in before decide update or insert
	if err := s.tradeInRepo.CreateTradeIn(ctx, tx, &tradeIn); err != nil {
		return errors.ErrTradeInCreateFailed
	}
	// TODO: implement update trade in if already exists

	// Commit transaction
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
