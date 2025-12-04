package leads

import (
	"strings"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/google/uuid"
)

// FinanceSimulationWebhookEvent represents the webhook event for finance simulation
type FinanceSimulationWebhookEvent struct {
	Process   string                       `json:"process" validate:"required"`
	EventID   string                       `json:"event_ID" validate:"required"`
	Timestamp int64                        `json:"timestamp" validate:"required"`
	Data      FinanceSimulationWebhookData `json:"data" validate:"required"`
}

// FinanceSimulationWebhookData represents the main data in the webhook
type FinanceSimulationWebhookData struct {
	OneAccount        customer.OneAccountRequest      `json:"one_account" validate:"required"`
	Leads             FinanceSimulationLeadsRequest   `json:"leads" validate:"required"`
	FinanceSimulation FinanceSimulationDetailsRequest `json:"finance_simulation" validate:"required"`
	TradeIn           TradeInRequest                  `json:"trade_in"`
}

// FinanceSimulationLeadsRequest represents the leads information in finance simulation webhook
type FinanceSimulationLeadsRequest struct {
	FinanceSimulationID             string           `json:"finance_simulation_ID" validate:"required"`
	FinanceSimulationNumber         string           `json:"finance_simulation_number" validate:"required"`
	LeadsID                         string           `json:"leads_ID" validate:"required"`
	CreatedDatetime                 int64            `json:"created_datetime"`
	KatashikiSuffix                 string           `json:"katashiki_suffix"`
	ColorCode                       string           `json:"color_code"`
	Model                           string           `json:"model"`
	Variant                         string           `json:"variant"`
	Color                           string           `json:"color"`
	VehicleOTRPrice                 float64          `json:"vehicle_otr_price"`
	OutletID                        string           `json:"outlet_ID"`
	OutletName                      string           `json:"outlet_name"`
	LeadsSource                     string           `json:"leads_source"`
	LeadsType                       string           `json:"leads_type"`
	LeadsFollowUpStatus             string           `json:"leads_follow_up_status"`
	LeadsPreferenceContactTimeStart string           `json:"leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string           `json:"leads_preference_contact_time_end"`
	AdditionalNotes                 *string          `json:"additional_notes"`
	Score                           Score            `json:"score"`
	InterestedPart                  []InterestedPart `json:"interested_part"`
}

// InterestedPart represents parts/accessories/merchandise interested by customer
type InterestedPart struct {
	InterestedPartType                 string        `json:"interested_part_type"`
	PackageID                          string        `json:"package_ID"`
	InterestedPartNumber               string        `json:"interested_part_number"`
	InterestedPartName                 string        `json:"interested_part_name"`
	InterestedPartQuantity             int           `json:"interested_part_quantity"`
	PackageParts                       []PackagePart `json:"package_parts"`
	InterestedPartSize                 string        `json:"interested_part_size"`
	InterestedPartColor                string        `json:"interested_part_color"`
	InterestedPartEstPrice             float64       `json:"interested_part_est_price"`
	InterestedPartInstallationEstPrice float64       `json:"interested_part_installation_est_price"`
	FlagInterestedPartNeedDownPayment  bool          `json:"flag_interested_part_need_down_payment"`
}

// PackagePart represents parts within a package
type PackagePart struct {
	InterestedPartNumber string `json:"interested_part_number"`
	InterestedPartName   string `json:"interested_part_name"`
}

// FinanceSimulationDetailsRequest represents the finance simulation details
type FinanceSimulationDetailsRequest struct {
	PaymentPreference       string                   `json:"payment_preference"`
	InsurancePeriod         string                   `json:"insurance_period"`
	CreditMethod            string                   `json:"credit_method"`
	PackageType             string                   `json:"package_type"`
	FirstPaymentType        string                   `json:"first_payment_type"`
	InsuranceType           string                   `json:"insurance_type"`
	InsurancePaymentType    string                   `json:"insurance_payment_type"`
	InsuranceCoverage       []string                 `json:"insurance_coverage"`
	CreditSimulationResults []CreditSimulationResult `json:"credit_simulation_results"`
}

// CreditSimulationResult represents credit simulation calculation results
type CreditSimulationResult struct {
	Tenor              int     `json:"tenor"`
	DownPayment        float64 `json:"down_payment"`
	TotalFirstPayment  float64 `json:"total_first_payment"`
	MonthlyInstallment float64 `json:"monthly_installment"`
}

// ToDomain converts FinanceSimulationLeadsRequest to domain.Leads
func (req *FinanceSimulationLeadsRequest) ToDomain(customerID string) domain.Leads {
	return domain.Leads{
		CustomerID:                      customerID,
		FinanceSimulationID:             &req.FinanceSimulationID,
		FinanceSimulationNumber:         &req.FinanceSimulationNumber,
		LeadsID:                         req.LeadsID,
		LeadsType:                       req.LeadsType,
		LeadsFollowUpStatus:             req.LeadsFollowUpStatus,
		LeadsPreferenceContactTimeStart: utils.ToPointer(req.LeadsPreferenceContactTimeStart),
		LeadsPreferenceContactTimeEnd:   utils.ToPointer(req.LeadsPreferenceContactTimeEnd),
		LeadSource:                      req.LeadsSource,
		AdditionalNotes:                 req.AdditionalNotes,
		TAMLeadScore:                    req.Score.TAMLeadScore,
		OutletLeadScore:                 req.Score.OutletLeadScore,
		KatashikiSuffix:                 utils.ToPointer(req.KatashikiSuffix),
		ColorCode:                       utils.ToPointer(req.ColorCode),
		Model:                           utils.ToPointer(req.Model),
		Variant:                         utils.ToPointer(req.Variant),
		Color:                           utils.ToPointer(req.Color),
		VehicleOTRPrice:                 utils.ToPointer(req.VehicleOTRPrice),
		OutletID:                        utils.ToPointer(req.OutletID),
		OutletName:                      utils.ToPointer(req.OutletName),
		CreatedDatetime:                 time.Unix(req.CreatedDatetime, 0),
		PurchasePlanCriteria:            utils.ToPointer(req.Score.Parameter.PurchasePlanCriteria),
		PaymentPreferCriteria:           utils.ToPointer(req.Score.Parameter.PaymentPreferCriteria),
		TestDriveCriteria:               utils.ToPointer(req.Score.Parameter.TestDriveCriteria),
		TradeInCriteria:                 utils.ToPointer(req.Score.Parameter.TradeInCriteria),
		BrowsingHistoryCriteria:         utils.ToPointer(req.Score.Parameter.BrowsingHistoryCriteria),
		VehicleAgeCriteria:              utils.ToPointer(req.Score.Parameter.VehicleAgeCriteria),
		NegotiationCriteria:             utils.ToPointer(req.Score.Parameter.NegotiationCriteria),
	}
}

// ToDomain converts FinanceSimulationDetailsRequest to domain.LeadsFinanceSimulation
func (req *FinanceSimulationDetailsRequest) ToDomain(simulationID, simulationNumber, leadsID string) domain.LeadsFinanceSimulation {
	return domain.LeadsFinanceSimulation{
		SimulationID:         simulationID,
		SimulationNumber:     simulationNumber,
		LeadsID:              leadsID,
		PaymentPreference:    req.PaymentPreference,
		InsurancePeriod:      req.InsurancePeriod,
		CreditMethod:         req.CreditMethod,
		PackageType:          req.PackageType,
		FirstPaymentType:     req.FirstPaymentType,
		InsuranceType:        req.InsuranceType,
		InsurancePaymentType: req.InsurancePaymentType,
		InsuranceCoverage:    strings.Join(req.InsuranceCoverage, ","),
	}
}

// ToDomain converts InterestedPart to domain.LeadsInterestedPart
func (ip *InterestedPart) ToDomain(leadsID string) domain.LeadsInterestedPart {
	var packageID *string
	if ip.PackageID != "" {
		packageID = &ip.PackageID
	}

	var partSize *string
	if ip.InterestedPartSize != "" {
		partSize = &ip.InterestedPartSize
	}

	var partColor *string
	if ip.InterestedPartColor != "" {
		partColor = &ip.InterestedPartColor
	}

	return domain.LeadsInterestedPart{
		ID:                       uuid.New().String(),
		LeadsID:                  leadsID,
		PartType:                 ip.InterestedPartType,
		PackageID:                packageID,
		PartNumber:               ip.InterestedPartNumber,
		PartName:                 ip.InterestedPartName,
		PartQuantity:             ip.InterestedPartQuantity,
		PartSize:                 partSize,
		PartColor:                partColor,
		PartEstPrice:             &ip.InterestedPartEstPrice,
		PartInstallationEstPrice: &ip.InterestedPartInstallationEstPrice,
		FlagPartNeedDownPayment:  ip.FlagInterestedPartNeedDownPayment,
		CreatedAt:                time.Now(),
	}
}

// ToDomain converts PackagePart to domain.LeadsInterestedPartItem
func (pp *PackagePart) ToDomain(leadsID, interestedPartID string) domain.LeadsInterestedPartItem {
	return domain.LeadsInterestedPartItem{
		ID:                    uuid.New().String(),
		LeadsID:               leadsID,
		LeadsInterestedPartID: interestedPartID,
		PartNumber:            pp.InterestedPartNumber,
		PartName:              pp.InterestedPartName,
		CreatedAt:             time.Now(),
	}
}

// ToDomain converts CreditSimulationResult to domain.LeadsFinanceSimulationCredit
func (csr *CreditSimulationResult) ToDomain(leadsID, financeSimulationID string) domain.LeadsFinanceSimulationCredit {
	return domain.LeadsFinanceSimulationCredit{
		ID:                       uuid.New().String(),
		LeadsID:                  leadsID,
		LeadsFinanceSimulationID: financeSimulationID,
		Tenor:                    csr.Tenor,
		DownPayment:              csr.DownPayment,
		TotalFirstPayment:        csr.TotalFirstPayment,
		MonthlyInstallment:       csr.MonthlyInstallment,
		IsActive:                 true,
	}
}
