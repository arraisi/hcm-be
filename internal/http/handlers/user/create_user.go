package user

import (
	"encoding/json"
	"net/http"

	"tabeldata.com/hcm-be/internal/domain/dto/user"
	"tabeldata.com/hcm-be/pkg/errors"
	"tabeldata.com/hcm-be/pkg/response"
	"tabeldata.com/hcm-be/pkg/utils/validator"
)

// Create creates a new user
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid JSON payload")
		return
	}

	// Validate using the existing validator
	if err := validator.ValidateStruct(&req); err != nil {
		validationErrors := response.NormalizeValidationError(err)
		response.Validation(w, validationErrors)
		return
	}

	// Basic business logic validation
	if req.Email == "" || req.Name == "" {
		response.DomainError(w, errors.NewDomainError(errors.ErrValidation, "email and name are required"))
		return
	}

	if err := h.svc.Create(ctx, req); err != nil {
		// Check if it's a domain error first
		if errors.IsDomainError(err) {
			response.DomainError(w, err)
			return
		}

		// Handle specific error cases
		if err.Error() == "user already exists" {
			response.DomainError(w, errors.NewDomainError(errors.ErrConflict, "User with this email already exists"))
			return
		}

		// Generic server error
		response.DomainError(w, errors.Wrap(errors.ErrInternal, err))
		return
	}

	response.Created(w, map[string]interface{}{
		"message": "User created successfully",
	}, "User created successfully")
}
