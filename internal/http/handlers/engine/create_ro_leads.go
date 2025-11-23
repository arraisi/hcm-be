package engine

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/arraisi/hcm-be/pkg/response"
)

// CreateRoLeads handles the request to create ro leads
func (h Handler) CreateRoLeads(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req engine.CreateRoLeadsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Engine] Failed to decode request: %v", err)
		response.BadRequest(w, "Invalid request payload")
		return
	}

	// Run the segmentation
	err := h.svc.CreateRoLeads(ctx, req)
	if err != nil {
		log.Printf("[Engine] Failed to create ro leads: %v", err)
		response.Error(w, http.StatusInternalServerError, "Failed to create ro leads")
		return
	}

	response.OK(w, nil, "create ro leads completed successfully")
}
