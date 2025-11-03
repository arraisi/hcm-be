package servicebooking

import (
	"strings"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/elgris/sqrl"
)

type GetServiceBookingRecall struct {
	ID               *string
	RecallID         *string
	ServiceBookingID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingRecall) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.RecallID != nil {
		q.Where(sqrl.Eq{"i_recall_id": req.RecallID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
}

type DeleteServiceBookingRecall struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingRecall) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}

type RecallRequest struct {
	RecallID          string   `json:"recall_ID"`
	RecallDate        string   `json:"recall_date"`
	RecallDescription string   `json:"recall_description"`
	AffectedParts     []string `json:"affected_parts"`
}

func (r *RecallRequest) ToModel(bookingID string) domain.ServiceBookingRecall {
	now := time.Now()
	return domain.ServiceBookingRecall{
		ServiceBookingID:  bookingID,
		RecallID:          r.RecallID,
		RecallDate:        r.RecallDate,
		RecallDescription: r.RecallDescription,
		AffectedParts:     strings.Join(r.AffectedParts, ","),
		CreatedAt:         now.UTC(),
		CreatedBy:         constants.System,
		UpdatedAt:         now.UTC(),
		UpdatedBy:         constants.System,
	}
}

func NewRecallsRequest(recalls []domain.ServiceBookingRecall) []RecallRequest {
	var recallsRequest []RecallRequest
	seen := make(map[string]bool)
	for _, recall := range recalls {
		if !seen[recall.RecallID] {
			recallsRequest = append(recallsRequest, RecallRequest{
				RecallID:          recall.RecallID,
				RecallDate:        recall.RecallDate,
				RecallDescription: recall.RecallDescription,
				AffectedParts:     strings.Split(recall.AffectedParts, ","),
			})
			seen[recall.RecallID] = true
		}
	}
	return recallsRequest
}
