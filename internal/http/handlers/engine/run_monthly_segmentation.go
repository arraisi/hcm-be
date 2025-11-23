package engine

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/arraisi/hcm-be/pkg/response"
)

type RunMonthlySegmentationRequest struct {
	ForceUpdate bool `json:"force_update"`
}

// RunMonthlySegmentation handles the request to run monthly customer segmentation
func (h Handler) RunMonthlySegmentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RunMonthlySegmentationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Engine] Failed to decode request: %v", err)
		response.BadRequest(w, "Invalid request payload")
		return
	}

	// Run the segmentation
	err := h.svc.RunMonthlySegmentation(ctx, engine.RunMonthlySegmentationRequest{
		ForceUpdate: req.ForceUpdate,
	})
	if err != nil {
		log.Printf("[Engine] Failed to run monthly segmentation: %v", err)
		response.Error(w, http.StatusInternalServerError, "Failed to run monthly segmentation")
		return
	}

	response.OK(w, nil, "Monthly segmentation completed successfully")
}
