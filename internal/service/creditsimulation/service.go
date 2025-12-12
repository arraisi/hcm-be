package creditsimulation

import (
	"fmt"
	"strconv"
	"strings"
	"math"


	dto "github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
	repo "github.com/arraisi/hcm-be/internal/repository/creditsimulation"
)

//  DEKLARASI INTERFACE 
type BranchService interface {
	GetBranches() ([]dto.BranchResponse, error)
	GetOutlets(branchCode string) ([]dto.OutletResponse, error)
	GetAssetGroups(branchCode string) ([]dto.AssetGroupResponse, error)
	GetAssetTypes(branchCode string, assetGroupCode string) ([]dto.AssetTypeResponse, error)
	GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price string) (dto.AggregatedInstallmentResponse, error)
	GetCreditSimulationByInstallment(branchCode, assetGroupCode, assetTypeCode, price, installment string) ([]dto.CreditSimulationDetailResponse, error)
	GetCreditSimulationByDownPayment(branchCode, assetGroupCode, assetTypeCode, price, downPayment string) ([]dto.CreditSimulationDetailResponse, error)
}


type branchService struct {
	repo repo.BranchRepository
}

func NewBranchService(r repo.BranchRepository) BranchService {
	return &branchService{repo: r}
}

// --- IMPLEMENTASI METHOD ---

func (s *branchService) GetBranches() ([]dto.BranchResponse, error) {
	return s.repo.GetBranches()
}

func (s *branchService) GetOutlets(branchCode string) ([]dto.OutletResponse, error) {
	return s.repo.GetOutlets(branchCode)
}

func (s *branchService) GetAssetGroups(branchCode string) ([]dto.AssetGroupResponse, error) {
	return s.repo.GetAssetGroups(branchCode)
}

func (s *branchService) GetAssetTypes(branchCode string, assetGroupCode string) ([]dto.AssetTypeResponse, error) {
	return s.repo.GetAssetTypes(branchCode, assetGroupCode)
}


// Helper function untuk membersihkan dan mengonversi string mata uang
func cleanAndConvertCurrency(s string) (float64, error) {
	cleaned := strings.ReplaceAll(s, ".", "")
	f, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0, fmt.Errorf("gagal mengonversi string angsuran: %w", err)
	}
	return f, nil
}

// Implementasi GetMinMaxInstallments 
func (s *branchService) GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price string) (dto.AggregatedInstallmentResponse, error) {
	rawResults, err := s.repo.GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price)
	if err != nil {
		return dto.AggregatedInstallmentResponse{}, err
	}

	if len(rawResults) == 0 {
		return dto.AggregatedInstallmentResponse{}, fmt.Errorf("tidak ada data angsuran yang ditemukan")
	}

	var currentMin float64
	var currentMax float64 
	var finalMinStr string
	var finalMaxStr string
	
	min1, err := cleanAndConvertCurrency(rawResults[0].InstallmentMin)
	if err != nil {
		return dto.AggregatedInstallmentResponse{}, err
	}
	max1, err := cleanAndConvertCurrency(rawResults[0].InstallmentMax)
	if err != nil {
		return dto.AggregatedInstallmentResponse{}, err
	}
	
	currentMin = min1
	currentMax = max1
	finalMinStr = rawResults[0].InstallmentMin
	finalMaxStr = rawResults[0].InstallmentMax

	// Iterasi untuk mencari min terkecil dan max terbesar
	for i := 1; i < len(rawResults); i++ {
		item := rawResults[i]
		
		minVal, err := cleanAndConvertCurrency(item.InstallmentMin)
		if err != nil {
			return dto.AggregatedInstallmentResponse{}, err 
		}
		maxVal, err := cleanAndConvertCurrency(item.InstallmentMax)
		if err != nil {
			return dto.AggregatedInstallmentResponse{}, err 
		}

		// Ambil InstallmentMin terkecil
		if minVal < currentMin {
			currentMin = minVal
			finalMinStr = item.InstallmentMin
		}
		
		// Ambil InstallmentMax terbesar
		if maxVal > currentMax {
			currentMax = maxVal
			finalMaxStr = item.InstallmentMax
		}
	}

	return dto.AggregatedInstallmentResponse{
		InstallmentMin: finalMinStr,
		InstallmentMax: finalMaxStr,
	}, nil
}

// Helper function untuk menghitung dpPersen
func calculateDpPersen(downPayment, price string) (string, error) {

    dp, err := strconv.ParseFloat(downPayment, 64)
    if err != nil { return "0", fmt.Errorf("invalid downPayment format: %w", err) }
    
    p, err := strconv.ParseFloat(price, 64)
    if err != nil { return "0", fmt.Errorf("invalid price format: %w", err) }

    if p == 0 { return "0", nil }

    // Rumus: (downPayment/ price) x 100 dan dibulatkan Math.round
    dpPersenFloat := math.Round((dp / p) * 100)
    
    return strconv.Itoa(int(dpPersenFloat)), nil 
}


// GetCreditSimulation By INSTALLMENT
func (s *branchService) GetCreditSimulationByInstallment(branchCode, assetGroupCode, assetTypeCode, price, installment string) ([]dto.CreditSimulationDetailResponse, error) {
	downPayment := "0"
	dpPersen := "0" 
    calculationType := "INSTALLMENT"
    
	return s.repo.GetCreditSimulation(branchCode, assetGroupCode, assetTypeCode, calculationType, price, installment, downPayment, dpPersen)
}

// GetCreditSimulation By DOWN_PAYMENT
func (s *branchService) GetCreditSimulationByDownPayment(branchCode, assetGroupCode, assetTypeCode, price, downPayment string) ([]dto.CreditSimulationDetailResponse, error) {
	installment := "0" 
    calculationType := "DOWN_PAYMENT"

    // WAJIB: Hitung dpPersen
    dpPersen, err := calculateDpPersen(downPayment, price)
    if err != nil {
        return nil, fmt.Errorf("gagal menghitung dpPersen: %w", err)
    }
    
	return s.repo.GetCreditSimulation(branchCode, assetGroupCode, assetTypeCode, calculationType, price, installment, downPayment, dpPersen)
}
