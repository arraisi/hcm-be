package hasjratid

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
)

// GenerateHasjratID builds an ID with the format:
//
//	HA + SourceCode(1) + CustomerTypeCode(1) + Outlet(5) + Year(2) + Seq(7)
//
// Example:
//
//	HAHR1010125000001
//
// Parameters:
//   - ctx:            request context
//   - c:              customer entity (for customer type)
//   - sourceCode:     "H" / "C" (customer source, e.g. H = Hasjrat, C = Campaign, etc.)
//   - tamOutletID:    optional TAM outlet ID; if provided, it will be resolved via OutletRepo
//   - outletCode:     optional outlet code; used when tamOutletID is empty
//   - registrationDate: UNIX timestamp (seconds) used to derive the year (YY)
func (s *service) GenerateHasjratID(
	ctx context.Context,
	c domain.Customer,
	sourceCode string,
	tamOutletID string,
	outletCode string,
	registrationDate int64,
) (string, error) {
	const prefix = "HA"

	// --- Normalize & validate source code (H / C) ---
	sourceCode = strings.ToUpper(strings.TrimSpace(sourceCode))
	if len(sourceCode) != 1 {
		return "", fmt.Errorf("source code must be exactly 1 character, got %q", sourceCode)
	}

	// --- Map customer type text to 1-letter code (R/G/C) ---
	customerTypeCode, err := MapCustomerTypeTextToCode(c.CustomerType)
	if err != nil {
		return "", err
	}

	// --- Resolve outlet code ---
	var resolvedOutletCode string

	tamOutletID = strings.TrimSpace(tamOutletID)
	outletCode = strings.TrimSpace(outletCode)

	switch {
	case tamOutletID != "":
		// Use TAM outlet ID → resolve to outlet via repository
		outlet, err := s.outletRepo.GetOutletCodeByTAMOutletID(ctx, tamOutletID)
		if err != nil {
			return "", fmt.Errorf("failed to get outlet from TAM outlet ID %q: %w", tamOutletID, err)
		}
		if outlet == nil || strings.TrimSpace(outlet.OutletCode) == "" {
			return "", fmt.Errorf("outlet not found or invalid for TAM outlet ID %q", tamOutletID)
		}
		resolvedOutletCode = outlet.OutletCode

	case outletCode != "":
		// Use outlet code directly
		resolvedOutletCode = outletCode

	default:
		return "", fmt.Errorf("either tamOutletID or outletCode must be provided")
	}

	// --- Normalize outlet code to 5 digits ---
	outletCodePadded, err := padOutletCode(resolvedOutletCode)
	if err != nil {
		return "", err
	}

	// --- Derive 2-digit year from registration UNIX timestamp ---
	if registrationDate <= 0 {
		return "", fmt.Errorf("registrationDate must be a valid UNIX timestamp, got %d", registrationDate)
	}

	regTime := time.Unix(registrationDate, 0)
	yearStr := fmt.Sprintf("%02d", regTime.Year()%100)

	// --- Get next sequence value from DB (running number 1..n) ---
	seq, err := s.repo.GetNextSequence(
		ctx,
		sourceCode,
		customerTypeCode,
		outletCodePadded,
		yearStr,
	)
	if err != nil {
		return "", fmt.Errorf("failed to get next sequence: %w", err)
	}

	// Format running number as 7 digits
	running := fmt.Sprintf("%07d", seq)

	// --- Build final Hasjrat ID ---
	hasjratID := fmt.Sprintf(
		"%s%s%s%s%s%s",
		prefix,           // HA
		sourceCode,       // H / C
		customerTypeCode, // R / G / C
		outletCodePadded, // 5-digit outlet
		yearStr,          // 2-digit year
		running,          // 7-digit sequence
	)

	return hasjratID, nil
}

// MapCustomerTypeTextToCode converts free-text customer type into a 1-letter code.
// Rules:
//   - "Personal" or "INDIVIDUAL" → "R"
//   - "Government"/"Goverment"   → "G"
//   - "corporate" or "COMPANY"   → "C"
func MapCustomerTypeTextToCode(input string) (string, error) {
	normalized := strings.ToLower(strings.TrimSpace(input))

	switch normalized {
	case "personal", "individual":
		return "R", nil
	case "government", "goverment": // include a common typo
		return "G", nil
	case "corporate", "company":
		return "C", nil
	default:
		return "", fmt.Errorf("unsupported customer type: %q", input)
	}
}

// padOutletCode ensures the outlet code is 5 characters by left-padding with zeros.
func padOutletCode(outletCode string) (string, error) {
	outletCode = strings.TrimSpace(outletCode)
	if len(outletCode) == 0 {
		return "", fmt.Errorf("outlet code is empty")
	}
	if len(outletCode) > 5 {
		return "", fmt.Errorf("outlet code %q is longer than 5 characters", outletCode)
	}
	return fmt.Sprintf("%05s", outletCode), nil
}
