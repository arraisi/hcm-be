package creditsimulation

import (
	"encoding/json"
	"net/http"

	service "github.com/arraisi/hcm-be/internal/service/creditsimulation"
)


type CreditSimulationHandler struct {
	branchService service.BranchService
}

func NewCreditSimulationHandler(branchService service.BranchService) *CreditSimulationHandler {
	return &CreditSimulationHandler{
		branchService: branchService,
	}
}


func (h *CreditSimulationHandler) GetBranches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    
	result, err := h.branchService.GetBranches()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}


func (h *CreditSimulationHandler) GetOutlets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil branchCode dari Query Parameter
	branchCode := r.URL.Query().Get("branchCode")
	if branchCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "branchCode is required as a query parameter",
		})
		return
	}

	// 2. Panggil Service
	result, err := h.branchService.GetOutlets(branchCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 3. Respon Sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}

func (h *CreditSimulationHandler) GetAssetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil branchCode dari Query Parameter
	branchCode := r.URL.Query().Get("branchCode")
	if branchCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "branchCode is required as a query parameter",
		})
		return
	}

	// 2. Panggil Service
	result, err := h.branchService.GetAssetGroups(branchCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 3. Respon Sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}

func (h *CreditSimulationHandler) GetAssetTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil branchCode dan assetGroupCode dari Query Parameter
	branchCode := r.URL.Query().Get("branchCode")
	assetGroupCode := r.URL.Query().Get("assetGroupCode")

	if branchCode == "" || assetGroupCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "branchCode and assetGroupCode are required as query parameters",
		})
		return
	}

	// 2. Panggil Service
	result, err := h.branchService.GetAssetTypes(branchCode, assetGroupCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 3. Respon Sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}



func (h *CreditSimulationHandler) GetMinMaxInstallments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil semua 5 parameter wajib dari Query Parameter 
	branchCode := r.URL.Query().Get("branchCode")
	assetGroupCode := r.URL.Query().Get("assetGroupCode")
	assetTypeCode := r.URL.Query().Get("assetTypeCode")
	price := r.URL.Query().Get("price")
	
    // Defaulting calculationType
	calculationType := r.URL.Query().Get("calculationType")
    if calculationType == "" {
        calculationType = "INSTALLMENT" 
    }

    // Cek semua parameter wajib
	if branchCode == "" || assetGroupCode == "" || assetTypeCode == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "branchCode, assetGroupCode, assetTypeCode, and price are required as query parameters",
		})
		return
	}

	// 2. Panggil Service
	result, err := h.branchService.GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 3. Respon Sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}

// GetCreditSimulation By Installment
func (h *CreditSimulationHandler) GetCreditSimulationByInstallment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil Parameter dari Query Parameter
	branchCode := r.URL.Query().Get("branchCode")
	assetGroupCode := r.URL.Query().Get("assetGroupCode")
	assetTypeCode := r.URL.Query().Get("assetTypeCode")
	price := r.URL.Query().Get("price")
    installment := r.URL.Query().Get("installment") 
    
	// 2. Validasi Parameter Wajib
	if branchCode == "" || assetGroupCode == "" || assetTypeCode == "" || price == "" || installment == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "Parameter wajib (branchCode, assetGroupCode, assetTypeCode, price, installment) tidak lengkap.",
		})
		return
	}
    
	// 3. Panggil Service
	// Catatan: Parameter downPayment dan dpPersen dihandle di Service (default 0)
	result, err := h.branchService.GetCreditSimulationByInstallment(branchCode, assetGroupCode, assetTypeCode, price, installment)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 4. Respon Sukses (Mengembalikan array pricelist yang difilter)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}


// GetCreditSimulation By DOWN_PAYMENT
func (h *CreditSimulationHandler) GetCreditSimulationByDownPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 1. Ambil Parameter dari Query Parameter
	branchCode := r.URL.Query().Get("branchCode")
	assetGroupCode := r.URL.Query().Get("assetGroupCode")
	assetTypeCode := r.URL.Query().Get("assetTypeCode")
	price := r.URL.Query().Get("price")
    downPayment := r.URL.Query().Get("downPayment") 
    
	// 2. Validasi Parameter Wajib
	if branchCode == "" || assetGroupCode == "" || assetTypeCode == "" || price == "" || downPayment == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": "Parameter wajib (branchCode, assetGroupCode, assetTypeCode, price, downPayment) tidak lengkap.",
		})
		return
	}
    
	// 3. Panggil Service
	// Catatan: Parameter installment dan dpPersen dihandle di Service (default 0 dan dihitung)
	result, err := h.branchService.GetCreditSimulationByDownPayment(branchCode, assetGroupCode, assetTypeCode, price, downPayment)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	// 4. Respon Sukses (Mengembalikan array pricelist yang difilter)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":  result,
	})
}

