package engine

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/arraisi/hcm-be/pkg/response"
)

// RunCustomerSegmentation handles the request to run customer segmentation
func (h Handler) RunCustomerSegmentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req engine.RunCustomerSegmentationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Engine] Failed to decode request: %v", err)
		response.BadRequest(w, "Invalid request payload")
		return
	}

	// Run the segmentation
	err := h.svc.CustomerSegmentation(ctx, req)
	if err != nil {
		log.Printf("[Engine] Failed to run customer segmentation: %v", err)
		response.Error(w, http.StatusInternalServerError, "Failed to run customer segmentation")
		return
	}

	response.OK(w, nil, "customer segmentation completed successfully")
}
