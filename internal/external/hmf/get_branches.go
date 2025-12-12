package hmf

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)

func (c *client) GetBranches(ctx context.Context) ([]creditsimulation.BranchResponse, error) {
	header := map[string]string{
		"Accept":     "application/json",
		"User-Agent": "Go-http-client",
	}
	token := c.cfg.Http.HMFApi.APIKey

	// Build URL with query parameters
	baseURL := c.cfg.Http.HMFApi.BaseUrl
	validAttrib := generateAttrib()

	url := fmt.Sprintf("%s?action=selectAvailableBranch&attrib=%s", baseURL, validAttrib)

	result, err := c.httpUtil.Get(ctx, url, token, header)
	if err != nil {
		return nil, err
	}

	var resp []creditsimulation.BranchResponse
	if err := json.Unmarshal(result, &resp); err != nil {
		return resp, fmt.Errorf("parse response: %w", err)
	}

	return resp, nil
}

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
