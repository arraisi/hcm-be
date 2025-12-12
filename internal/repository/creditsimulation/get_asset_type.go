package creditsimulation


import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)


func (r *branchRepository) GetAssetTypes(branchCode string, assetGroupCode string) ([]creditsimulation.AssetTypeResponse, error) {
	apiURL := "http://apk.hmf.co.id:7070/mobile/marketing/credit-simulation.html"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	validAttrib := generateAttrib()
	
	// Prepare Query Parameters
	q := req.URL.Query()
	q.Add("action", "selectAssetTypeExt") 
	q.Add("attrib", validAttrib)
	q.Add("branchCode", branchCode)    
	q.Add("assetGroupCode", assetGroupCode) 
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body) 
		errorMessage := fmt.Sprintf("External API returned status %d. Body: %s", resp.StatusCode, string(raw))
		return nil, fmt.Errorf(errorMessage)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result []creditsimulation.AssetTypeResponse 
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal asset type response: %w. Raw body: %s", err, string(raw))
	}

	return result, nil
}


func (r *branchRepository) GetMinMaxInstallments(branchCode, assetGroupCode, assetTypeCode, calculationType, price string) ([]creditsimulation.InstallmentResponse, error) {
	apiURL := "http://apk.hmf.co.id:7070/mobile/marketing/credit-simulation.html"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	validAttrib := generateAttrib() 
	
	// Menyiapkan Query Parameters
	q := req.URL.Query()
	q.Add("action", "selectAssetPriceListDetailSingleExt")
	q.Add("attrib", validAttrib)
	q.Add("branchCode", branchCode)
	q.Add("assetGroupCode", assetGroupCode)
	q.Add("assetTypeCode", assetTypeCode)
	q.Add("calculationType", calculationType)
	q.Add("price", price)
	
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body) 
		errorMessage := fmt.Sprintf("External API returned status %d. Body: %s", resp.StatusCode, string(raw))
		return nil, fmt.Errorf(errorMessage)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result []creditsimulation.InstallmentResponse 
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal installment response: %w. Raw body: %s", err, string(raw))
	}

	return result, nil
}