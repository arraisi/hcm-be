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

func (s *service) RequestGetOffer(ctx context.Context, request dtoLeads.GetOfferWebhookEvent) error {
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
	leads, err := s.processLeadsGetOffer(ctx, tx, leadsData, customerID)
	if err != nil {
		return err
	}

	// 3. Process interested parts (delete old and insert new)
	if err := s.processInterestedParts(ctx, tx, leads.ID, request.Data.InterestedPart); err != nil {
		return err
	}

	// 4. Upsert trade-in
	if err := s.processTradeIn(ctx, tx, leads.ID, request.Data.TradeIn); err != nil {
		return err
	}

	// Commit transaction
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// processLeadsGetOffer handles upsert logic for leads from get offer request
func (s *service) processLeadsGetOffer(ctx context.Context, tx *sqlx.Tx, leadsData dtoLeads.GetOfferLeadsRequest, customerID string) (domain.Leads, error) {
	leads := leadsData.ToDomain(customerID)
	leads.CreatedBy = "SYSTEM"
	leads.UpdatedAt = time.Now()
	leads.UpdatedBy = utils.ToPointer("SYSTEM")

	// Check if leads already exists
	existingLeads, err := s.leadsRepo.GetLeads(ctx, dtoLeads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsData.LeadsID),
	})
	if err != nil && err != sql.ErrNoRows {
		return domain.Leads{}, fmt.Errorf("failed to check existing leads: %w", err)
	}

	if existingLeads.ID != "" {
		// Update existing leads
		leads.ID = existingLeads.ID
		if err := s.leadsRepo.UpdateLeads(ctx, tx, leads); err != nil {
			return domain.Leads{}, errors.ErrLeadsUpdateFailed
		}
	} else {
		// Create new leads
		leads.CreatedAt = time.Now()
		if err := s.leadsRepo.CreateLeads(ctx, tx, &leads); err != nil {
			return domain.Leads{}, errors.ErrLeadsCreateFailed
		}
	}

	return leads, nil
}

// processInterestedParts handles delete and insert for interested parts
func (s *service) processInterestedParts(ctx context.Context, tx *sqlx.Tx, leadsID string, interestedParts []dtoLeads.InterestedPart) error {
	// Delete existing interested parts before inserting new ones
	if err := s.interestedPartRepo.DeleteInterestedPartItemsByLeadsID(ctx, tx, leadsID); err != nil {
		return fmt.Errorf("failed to delete interested part items: %w", err)
	}
	if err := s.interestedPartRepo.DeleteInterestedPartByLeadsID(ctx, tx, leadsID); err != nil {
		return fmt.Errorf("failed to delete interested parts: %w", err)
	}

	// Insert new interested parts
	for _, interestedPart := range interestedParts {
		// Create interested part
		part := interestedPart.ToDomain(leadsID)
		if err := s.interestedPartRepo.CreateInterestedPart(ctx, tx, &part); err != nil {
			return errors.ErrInterestedPartCreateFailed
		}

		// Create package parts if it's a package type
		if interestedPart.InterestedPartType == "PACKAGE" && len(interestedPart.PackageParts) > 0 {
			for _, packagePart := range interestedPart.PackageParts {
				item := packagePart.ToDomain(leadsID, part.ID)
				if err := s.interestedPartRepo.CreateInterestedPartItem(ctx, tx, &item); err != nil {
					return errors.ErrInterestedPartItemCreateFailed
				}
			}
		}
	}

	return nil
}

// processTradeIn handles upsert logic for trade-in
func (s *service) processTradeIn(ctx context.Context, tx *sqlx.Tx, leadsID string, tradeInData dtoLeads.TradeInRequest) error {
	tradeIn := tradeInData.ToDomain(leadsID)

	// Check if trade-in already exists
	existingTradeIn, err := s.tradeInRepo.GetTradeIn(ctx, dtoLeads.GetTradeInRequest{
		LeadsID: utils.ToPointer(leadsID),
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

	return nil
}
