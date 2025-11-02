package servicebooking

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

type GetServiceBookingPart struct {
	ID               *string
	ServiceBookingID *string
	PartType         *string
	PackageID        *string
	PartNumber       *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingPart) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
	if req.PartType != nil {
		q.Where(sqrl.Eq{"c_part_type": req.PartType})
	}
	if req.PackageID != nil {
		q.Where(sqrl.Eq{"i_package_id": req.PackageID})
	}
	if req.PartNumber != nil {
		q.Where(sqrl.Eq{"c_part_number": req.PartNumber})
	}
}

type DeleteServiceBookingPart struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingPart) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}

type PartRequest struct {
	PartType                 string            `json:"part_type"`
	PackageID                string            `json:"package_ID"`
	PartNumber               string            `json:"part_number"`
	PartName                 string            `json:"part_name"`
	PartQuantity             int32             `json:"part_quantity"`
	PackageParts             []PartItemRequest `json:"package_parts"`
	PartSize                 string            `json:"part_size"`
	PartColor                string            `json:"part_color"`
	PartEstPrice             float32           `json:"part_est_price"`
	PartInstallationEstPrice float32           `json:"part_installation_est_price"`
	FlagPartNeedDownPayment  bool              `json:"flag_part_need_down_payment"`
}

func (p *PartRequest) ToDomain(serviceBookingID string) (domain.ServiceBookingPart, []domain.ServiceBookingPartItem) {
	now := time.Now()

	partItems := make([]domain.ServiceBookingPartItem, 0, len(p.PackageParts))
	for _, item := range p.PackageParts {
		partItems = append(partItems, domain.ServiceBookingPartItem{
			PartNumber: item.PartNumber,
			PartName:   item.PartName,
			CreatedAt:  now.UTC(),
			CreatedBy:  constants.System,
			UpdatedAt:  now.UTC(),
			UpdatedBy:  constants.System,
		})
	}

	return domain.ServiceBookingPart{
		ServiceBookingID:         serviceBookingID,
		PartType:                 p.PartType,
		PackageID:                utils.ToPointer(p.PackageID),
		PartNumber:               utils.ToPointer(p.PartNumber),
		PartName:                 p.PartName,
		PartQuantity:             p.PartQuantity,
		PartSize:                 utils.ToPointer(p.PartSize),
		PartColor:                utils.ToPointer(p.PartColor),
		PartEstPrice:             p.PartEstPrice,
		PartInstallationEstPrice: p.PartInstallationEstPrice,
		FlagPartNeedDownPayment:  p.FlagPartNeedDownPayment,
		CreatedAt:                now.UTC(),
		CreatedBy:                constants.System,
		UpdatedAt:                now.UTC(),
		UpdatedBy:                constants.System,
	}, partItems
}

func NewPartsRequest(parts []domain.ServiceBookingPart, partItems []domain.ServiceBookingPartItem) []PartRequest {
	partItemsMap := make(map[string][]PartItemRequest)
	for _, item := range partItems {
		partItemsMap[item.PartNumber] = append(partItemsMap[item.PartNumber], PartItemRequest{
			PartNumber: item.PartNumber,
			PartName:   item.PartName,
		})
	}

	var partsRequest []PartRequest
	for _, part := range parts {
		partsRequest = append(partsRequest, PartRequest{
			PartType:                 part.PartType,
			PackageID:                utils.ToValue(part.PackageID),
			PartNumber:               utils.ToValue(part.PartNumber),
			PartName:                 part.PartName,
			PartQuantity:             part.PartQuantity,
			PackageParts:             partItemsMap[utils.ToValue(part.PartNumber)],
			PartSize:                 utils.ToValue(part.PartSize),
			PartColor:                utils.ToValue(part.PartColor),
			PartEstPrice:             part.PartEstPrice,
			PartInstallationEstPrice: part.PartInstallationEstPrice,
			FlagPartNeedDownPayment:  part.FlagPartNeedDownPayment,
		})
	}
	return partsRequest
}
