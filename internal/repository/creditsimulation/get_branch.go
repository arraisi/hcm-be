package creditsimulation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)

func (r *branchRepository) GetBranches() ([]creditsimulation.BranchResponse, error) {
	apiURL := "http://apk.hmf.co.id:7070/mobile/marketing/credit-simulation.html"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	validAttrib := generateAttrib()
    
	q := req.URL.Query()
	q.Add("action", "selectAvailableBranch")
	q.Add("attrib", validAttrib) 
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
        errorMessage := fmt.Sprintf("External API returned status %d.", resp.StatusCode)
        return nil, fmt.Errorf(errorMessage)
    }

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result []creditsimulation.BranchResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w. Raw body: %s", err, string(raw))
	}

	return result, nil
}