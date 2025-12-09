package creditsimulation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)

func (r *branchRepository) GetCreditSimulation(branchCode, assetGroupCode, assetTypeCode, calculationType, price, installment, downPayment, dpPersen string) ([]creditsimulation.CreditSimulationDetailResponse, error) {
	apiURL := "http://apk.hmf.co.id:7070/mobile/marketing/credit-simulation.html"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	validAttrib := generateAttrib()
	
	// Menyiapkan Query Parameters
	q := req.URL.Query()
	q.Add("action", "selectCreditSimulationExt") //
	q.Add("attrib", validAttrib)
	q.Add("branchCode", branchCode)
	q.Add("assetGroupCode", assetGroupCode)
	q.Add("assetTypeCode", assetTypeCode)
	q.Add("calculationType", calculationType) 
	q.Add("price", price)
	q.Add("installment", installment) 
	q.Add("downPayment", downPayment) 
	q.Add("dpPersen", dpPersen) 
	
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body) 
		return nil, fmt.Errorf("External API returned status %d. Body: %s", resp.StatusCode, string(raw))
	}
	
	raw, err := io.ReadAll(resp.Body)

	var result []creditsimulation.CreditSimulationDetailResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal credit simulation detail response: %w", err)
	}

	return result, nil
}