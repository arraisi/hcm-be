package leads

import (
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// FinanceSimulationRequest represents the finance simulation request
type FinanceSimulationRequest struct {
	SimulationID         string `json:"simulation_id" validate:"required"`
	SimulationNumber     string `json:"simulation_number" validate:"required"`
	LeadsID              string `json:"leads_id" validate:"required"`
	PaymentPreference    string `json:"payment_preference"`
	InsurancePeriod      string `json:"insurance_period"`
	CreditMethod         string `json:"credit_method"`
	PackageType          string `json:"package_type"`
	FirstPaymentType     string `json:"first_payment_type"`
	InsuranceType        string `json:"insurance_type"`
	InsurancePaymentType string `json:"insurance_payment_type"`
	InsuranceCoverage    string `json:"insurance_coverage"`
}

// GetFinanceSimulationRequest represents the request parameters for getting finance simulations
type GetFinanceSimulationRequest struct {
	ID               *string
	SimulationID     *string
	SimulationNumber *string
	LeadsID          *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetFinanceSimulationRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": *req.ID})
	}
	if req.SimulationID != nil {
		q.Where(sqrl.Eq{"i_simulation_id": *req.SimulationID})
	}
	if req.SimulationNumber != nil {
		q.Where(sqrl.Eq{"c_simulation_number": *req.SimulationNumber})
	}
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"i_leads_id": *req.LeadsID})
	}
}

// GetFinanceSimulationsRequest represents the request parameters for getting multiple finance simulations
type GetFinanceSimulationsRequest struct {
	LeadsID  *string
	Page     int
	PageSize int
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetFinanceSimulationsRequest) Apply(q *sqrl.SelectBuilder) {
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"i_leads_id": *req.LeadsID})
	}

	if req.PageSize > 0 {
		// Calculate offset: (page - 1) * pageSize
		offset := 0
		if req.Page > 1 {
			offset = (req.Page - 1) * req.PageSize
		}
		// Use pageSize + 1 to detect if there's a next page
		limit := req.PageSize + 1
		q.Suffix("OFFSET ? ROWS FETCH NEXT ? ROWS ONLY", offset, limit)
	}
}

// GetFinanceSimulationsResponse represents the response for getting multiple finance simulations
type GetFinanceSimulationsResponse struct {
	Data       []domain.LeadsFinanceSimulation `json:"data"`
	Pagination Pagination                      `json:"pagination"`
}

// Pagination represents pagination information
type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

// ToDomain converts the request to the internal LeadsFinanceSimulation model
func (req *FinanceSimulationRequest) ToDomain() domain.LeadsFinanceSimulation {
	return domain.LeadsFinanceSimulation{
		SimulationID:         req.SimulationID,
		SimulationNumber:     req.SimulationNumber,
		LeadsID:              req.LeadsID,
		PaymentPreference:    req.PaymentPreference,
		InsurancePeriod:      req.InsurancePeriod,
		CreditMethod:         req.CreditMethod,
		PackageType:          req.PackageType,
		FirstPaymentType:     req.FirstPaymentType,
		InsuranceType:        req.InsuranceType,
		InsurancePaymentType: req.InsurancePaymentType,
		InsuranceCoverage:    req.InsuranceCoverage,
	}
}
