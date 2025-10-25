package user

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// Create creates a new user
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrInvalidJSON, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate using the existing validator
	if err := validator.ValidateStruct(&req); err != nil {
		validationErrors := response.NormalizeValidationError(err)
		response.Validation(w, validationErrors)
		return
	}

	if err := h.svc.Create(ctx, req); err != nil {
		// Use NewErrorResponseFromList to determine HTTP status code
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.Created(w, map[string]interface{}{
		"message": "User created successfully",
	}, "User created successfully")
}
