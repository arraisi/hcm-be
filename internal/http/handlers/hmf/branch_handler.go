package hmf

import (
	"context"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
	"github.com/arraisi/hcm-be/pkg/response"
)

type HmfService interface {
	GetBranches(ctx context.Context) ([]creditsimulation.BranchResponse, error)
}

// Handler handles HTTP requests for HMF operations
type Handler struct {
	svc HmfService
}

// New creates a new HMF Handler instance
func New(svc HmfService) Handler {
	return Handler{svc: svc}
}

// GetBranches handles the GET /api/v1/hcm/credit-simulation/branches endpoint
func (h *Handler) GetBranches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	branches, err := h.svc.GetBranches(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch branches")
		return
	}

	response.OK(w, branches, "Branches fetched successfully")
}
