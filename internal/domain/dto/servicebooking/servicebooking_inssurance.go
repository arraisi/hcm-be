package servicebooking

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

// VehicleInsuranceRequest represents vehicle insurance information
type VehicleInsuranceRequest struct {
	InsuranceProvider      string          `json:"insurance_provider"`
	InsuranceProviderOther string          `json:"insurance_provider_other"`
	InsurancePolicyNumber  string          `json:"insurance_policy_number"`
	Policies               []PolicyRequest `json:"policies"`
}

// PolicyRequest represents individual insurance policy
type PolicyRequest struct {
	InsuranceType          string   `json:"insurance_type"`
	InsuranceCoverage      []string `json:"insurance_coverage"`
	InsuranceStartDatetime int64    `json:"insurance_start_datetime"`
	InsuranceEndDatetime   int64    `json:"insurance_end_datetime"`
	InsuranceStatus        string   `json:"insurance_status"`
}

// ToModel converts VehicleInsuranceRequest to domain model
func (v *VehicleInsuranceRequest) ToModel(serviceBookingID string) domain.ServiceBookingVehicleInsurance {
	now := time.Now()
	return domain.ServiceBookingVehicleInsurance{
		ServiceBookingID:       serviceBookingID,
		InsuranceProvider:      v.InsuranceProvider,
		InsuranceProviderOther: v.InsuranceProviderOther,
		InsurancePolicyNumber:  v.InsurancePolicyNumber,
		CreatedAt:              now.UTC(),
		CreatedBy:              constants.System,
		UpdatedAt:              now.UTC(),
		UpdatedBy:              constants.System,
	}
}

// ToModel converts PolicyRequest to domain model
func (p *PolicyRequest) ToModel(vehicleInsuranceID, serviceBookingID string) domain.ServiceBookingVehicleInsurancePolicy {
	now := time.Now()
	return domain.ServiceBookingVehicleInsurancePolicy{
		ServiceBookingID:       serviceBookingID,
		VehicleInsuranceID:     vehicleInsuranceID,
		InsuranceType:          p.InsuranceType,
		InsuranceCoverage:      p.InsuranceCoverage,
		InsuranceStartDatetime: utils.GetTimeUnix(p.InsuranceStartDatetime).UTC(),
		InsuranceEndDatetime:   utils.GetTimeUnix(p.InsuranceEndDatetime).UTC(),
		InsuranceStatus:        p.InsuranceStatus,
		CreatedAt:              now.UTC(),
		CreatedBy:              constants.System,
		UpdatedAt:              now.UTC(),
		UpdatedBy:              constants.System,
	}
}

// GetServiceBookingVehicleInsurance represents query filters for vehicle insurance
type GetServiceBookingVehicleInsurance struct {
	ID               *string
	ServiceBookingID *string
}

// Apply applies the filters to the query builder
func (req *GetServiceBookingVehicleInsurance) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
}

// DeleteServiceBookingVehicleInsurance represents delete filters
type DeleteServiceBookingVehicleInsurance struct {
	ServiceBookingID *string
}

// Apply applies the filters to the delete builder
func (d *DeleteServiceBookingVehicleInsurance) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}

// GetServiceBookingVehicleInsurancePolicy represents query filters for insurance policy
type GetServiceBookingVehicleInsurancePolicy struct {
	ID                 *string
	ServiceBookingID   *string
	VehicleInsuranceID *string
}

// Apply applies the filters to the query builder
func (req *GetServiceBookingVehicleInsurancePolicy) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
	if req.VehicleInsuranceID != nil {
		q.Where(sqrl.Eq{"i_vehicle_insurance_id": req.VehicleInsuranceID})
	}
}

// DeleteServiceBookingVehicleInsurancePolicy represents delete filters
type DeleteServiceBookingVehicleInsurancePolicy struct {
	ServiceBookingID *string
}

// Apply applies the filters to the delete builder
func (d *DeleteServiceBookingVehicleInsurancePolicy) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}
