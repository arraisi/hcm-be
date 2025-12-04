package sales

import (
	"context"
	"fmt"
	"sort"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/constants"
	pkgErrors "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type TestDriveRepository interface {
	GetTestDrives(ctx context.Context, req testdrive.GetTestDriveRequest) ([]domain.TestDrive, error)
}

type Repository interface {
	GetSalesScoring(ctx context.Context, req sales.GetSalesAssignmentRequest) (sales.GetSalesScoringResponse, error)
}

type LeadsRepository interface {
	GetLeadsTestDrive(ctx context.Context, req leads.GetLeadsTestDriveRequest) (domain.LeadsList, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	TestDriveRepo   TestDriveRepository
	LeadsRepo       LeadsRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
	testDriveRepo   TestDriveRepository
	leadsRepo       LeadsRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		testDriveRepo:   container.TestDriveRepo,
		leadsRepo:       container.LeadsRepo,
	}
}

// fetchAllSalesScoring retrieves all sales scoring data handling pagination automatically
func (s *service) fetchAllSalesScoring(ctx context.Context, request sales.GetSalesAssignmentRequest) (domain.SalesScorings, error) {
	var allSalesData domain.SalesScorings
	page := 1
	const pageSize = 100  // Use larger page size for efficiency
	const maxPages = 1000 // Safety limit to prevent infinite loops

	for page <= maxPages {
		result, err := s.repo.GetSalesScoring(ctx, sales.GetSalesAssignmentRequest{
			TAMOutletCode: request.TAMOutletCode,
			OutletCode:    request.OutletCode,
			Periode:       request.Periode,
			NIK:           request.NIK,
			BranchCode:    request.BranchCode,
			Page:          page,
			PageSize:      pageSize,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get sales scoring (page %d): %w", page, err)
		}

		// Merge results
		allSalesData = append(allSalesData, result.Data...)

		// Check if we've reached the last page
		if !result.Pagination.HasNext {
			break
		}

		page++
	}

	if page > maxPages {
		return nil, fmt.Errorf("reached maximum page limit (%d), possible infinite loop", maxPages)
	}

	return allSalesData, nil
}

// getActiveTestDriveCounts retrieves the count of active test drives for each sales person
func (s *service) getActiveTestDriveCounts(ctx context.Context, salesNIKs []string) (map[string]int, error) {
	if len(salesNIKs) == 0 {
		return make(map[string]int), nil
	}

	// Fetch all test drive leads for the given sales IDs with ongoing statuses
	testDriveLeads, err := s.leadsRepo.GetLeadsTestDrive(ctx, leads.GetLeadsTestDriveRequest{
		SalesIDs:        salesNIKs,
		TestDriveStatus: constants.TestDriveOnGoingStatus,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get test drive leads: %w", err)
	}

	// Count test drives per sales person
	counts := make(map[string]int)
	for _, lead := range testDriveLeads {
		if lead.SalesID != nil && *lead.SalesID != "" {
			counts[*lead.SalesID]++
		}
	}

	return counts, nil
}

// pickBestSalesCandidate selects the best sales person from candidates based on business rules
//
// Selection criteria (in order of priority):
//  1. Filter: activeTestDriveCount < maxActive
//  2. Sort by: highest performance score
//  3. Tie-breaker 1: lowest active test drive count
//  4. Tie-breaker 2: smallest NIK lexicographically
//
// Returns ErrNoEligibleSalesAssignment if no candidate qualifies
func pickBestSalesCandidate(candidates domain.SalesScorings, maxActive int) (*domain.SalesScoring, error) {
	// Filter eligible candidates
	eligible := make(domain.SalesScorings, 0, len(candidates))
	for i := range candidates {
		if candidates[i].ActiveTestDriveCount < maxActive {
			eligible = append(eligible, candidates[i])
		}
	}

	// Check if any candidate is eligible
	if len(eligible) == 0 {
		return nil, pkgErrors.NewDomainErrorWithDetails(
			pkgErrors.ErrNoEligibleSalesAssignment,
			fmt.Sprintf("All sales have reached maximum active test drives (>=%d)", maxActive),
			map[string]interface{}{
				"max_active_test_drives": maxActive,
				"total_candidates":       len(candidates),
			},
		)
	}

	// Sort eligible candidates by business rules
	sort.Slice(eligible, func(i, j int) bool {
		scoreI := eligible[i].GetPerformanceScore()
		scoreJ := eligible[j].GetPerformanceScore()

		// Primary: Higher performance score wins
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}

		// Tie-breaker 1: Lower active test drive count wins
		if eligible[i].ActiveTestDriveCount != eligible[j].ActiveTestDriveCount {
			return eligible[i].ActiveTestDriveCount < eligible[j].ActiveTestDriveCount
		}

		// Tie-breaker 2: Smaller NIK lexicographically (deterministic)
		return eligible[i].NIK < eligible[j].NIK
	})

	// Return the best candidate (first after sorting)
	return &eligible[0], nil
}
