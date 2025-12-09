package creditsimulation

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "time"

    "github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)

//  Helper Function 
func generateAttrib() string {
	securityValue := "Toyota_a3n9H1HMF"
	secretKey := "Emf_prod@2022"
	now := time.Now()
	dynamicYearMonth := now.Format("200601") 
	dataToHash := securityValue + dynamicYearMonth
    
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(dataToHash))
	hashBytes := h.Sum(nil) 
	attribValue := base64.StdEncoding.EncodeToString(hashBytes)
	
	return attribValue
}

// --- Interface Definition (Shared) ---
type BranchRepository interface {
	GetBranches() ([]creditsimulation.BranchResponse, error)
	GetOutlets(branchCode string) ([]creditsimulation.OutletResponse, error)
	GetAssetGroups(branchCode string) ([]creditsimulation.AssetGroupResponse, error)
	GetAssetTypes(branchCode string, assetGroupCode string) ([]creditsimulation.AssetTypeResponse, error)
	GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price string) ([]creditsimulation.InstallmentResponse, error)
	GetCreditSimulation(branchCode, assetGroupCode, assetTypeCode, calculationType, price, installment, downPayment, dpPersen string) ([]creditsimulation.CreditSimulationDetailResponse, error)
}


type branchRepository struct{}

func NewBranchRepository() BranchRepository {
	return &branchRepository{}
}