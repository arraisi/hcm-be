package sales

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
	pkgErrors "github.com/arraisi/hcm-be/pkg/errors"
)

const (
	// maxActiveTestDrives defines the maximum number of active test drives a sales can handle
	maxActiveTestDrives = 15
)

// GetSalesAssignment performs load-balanced test drive assignment based on performance and current workload.
//
// Business Rules:
//  1. Fetches all sales scoring data for the specified outlet (handles pagination automatically)
//  2. Queries test drive leads for all sales to count active test drives per sales person
//  3. Only considers sales with active test drive count < 15 as eligible
//  4. From eligible candidates, selects the best match based on:
//     - Highest performance score (Performa_nilai)
//     - If tied: lowest active test drive count
//     - If still tied: smallest NIK lexicographically (for deterministic results)
//  5. Returns business error if no sales person is eligible
//
// Returns:
//   - The selected sales person with their active test drive count populated
//   - ErrNoEligibleSalesAssignment if no one qualifies
//   - Repository errors wrapped with context
func (s *service) GetSalesAssignment(ctx context.Context, request sales.GetSalesAssignmentRequest) (*domain.SalesScoring, error) {
	// Step 1: Fetch all sales scoring data for the outlet (with pagination)
	salesData, err := s.fetchAllSalesScoring(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sales scoring data: %w", err)
	}

	// Handle empty result early
	if len(salesData) == 0 {
		return nil, pkgErrors.NewDomainError(
			pkgErrors.ErrNoEligibleSalesAssignment,
			"No sales data found for the specified outlet",
		)
	}

	// Step 2: Get all unique NIKs
	niks := salesData.GetUniqueNIKs()

	// Step 3: Count active test drives per sales person
	activeTestDriveCounts, err := s.getActiveTestDriveCounts(ctx, niks)
	if err != nil {
		return nil, fmt.Errorf("failed to get test drive counts: %w", err)
	}

	// Step 4: Enrich sales data with active test drive counts
	for i := range salesData {
		salesData[i].ActiveTestDriveCount = activeTestDriveCounts[salesData[i].NIK]
	}

	// Step 5: Select the best eligible sales candidate
	selected, err := pickBestSalesCandidate(salesData, maxActiveTestDrives)
	if err != nil {
		return nil, err
	}

	return selected, nil
}
