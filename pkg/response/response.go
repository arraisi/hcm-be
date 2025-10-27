package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"tabeldata.com/hcm-be/pkg/errors"
)

// Response represents the unified JSON response envelope
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
}

// ErrorResponse represents an error response with additional error details
type ErrorResponse struct {
	Data    interface{}   `json:"data"`
	Message string        `json:"message"`
	Meta    interface{}   `json:"meta,omitempty"`
	Error   *ErrorDetails `json:"error,omitempty"`
}

// ErrorDetails provides structured error information
type ErrorDetails struct {
	Code    string                 `json:"code"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// FieldValidationError represents a field validation error
type FieldValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// JSON sends a JSON response with the given HTTP status code and value.
func JSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// OK sends a 200 OK response with data and optional message
func OK(w http.ResponseWriter, data interface{}, message ...string) {
	msg := "Success"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	resp := Response{
		Data:    data,
		Message: msg,
	}
	JSON(w, http.StatusOK, resp)
}

// Created sends a 201 Created response with data and optional message
func Created(w http.ResponseWriter, data interface{}, message ...string) {
	msg := "Resource created successfully"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	resp := Response{
		Data:    data,
		Message: msg,
	}
	JSON(w, http.StatusCreated, resp)
}

// NoContent sends a 204 No Content response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Error sends a generic error response
func Error(w http.ResponseWriter, code int, msg string) {
	resp := ErrorResponse{
		Data:    nil,
		Message: msg,
	}
	JSON(w, code, resp)
}

// DomainError sends an error response based on domain error
func DomainError(w http.ResponseWriter, err error) {
	if domainErr, ok := err.(*errors.DomainError); ok {
		resp := ErrorResponse{
			Data:    nil,
			Message: domainErr.Message,
			Error: &ErrorDetails{
				Code:    domainErr.Code,
				Details: domainErr.Details,
			},
		}
		JSON(w, domainErr.StatusCode, resp)
		return
	}

	// Fallback for non-domain errors
	InternalServerError(w, "Internal server error")
}

// Validation sends a validation error response
func Validation(w http.ResponseWriter, validationErrors []FieldValidationError) {
	details := make(map[string]interface{})
	details["validation_errors"] = validationErrors

	resp := ErrorResponse{
		Data:    nil,
		Message: "Validation failed",
		Error: &ErrorDetails{
			Code:    "VALIDATION_FAILED",
			Details: details,
		},
	}
	JSON(w, http.StatusBadRequest, resp)
}

// BadRequest sends a 400 Bad Request response
func BadRequest(w http.ResponseWriter, message string) {
	Error(w, http.StatusBadRequest, message)
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Authentication required"
	}
	Error(w, http.StatusUnauthorized, message)
}

// Forbidden sends a 403 Forbidden response
func Forbidden(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Access denied"
	}
	Error(w, http.StatusForbidden, message)
}

// NotFound sends a 404 Not Found response
func NotFound(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Resource not found"
	}
	Error(w, http.StatusNotFound, message)
}

// Conflict sends a 409 Conflict response
func Conflict(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Resource conflict"
	}
	Error(w, http.StatusConflict, message)
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Internal server error"
	}
	Error(w, http.StatusInternalServerError, message)
}

// ErrorResponseJSON sends an error response using ErrorResponse
func ErrorResponseJSON(w http.ResponseWriter, errorResponse *errors.ErrorResponse) {
	// Check if it's a domain error
	if errorResponse.IsDomainError() {
		DomainError(w, errorResponse.Err)
		return
	}

	// Check for validation errors
	if strings.Contains(errorResponse.Error(), "validation") {
		BadRequest(w, errorResponse.Error())
		return
	}

	// Default to the error response's status code
	message := errorResponse.Error()
	response := ErrorResponse{
		Data:    nil,
		Message: message,
		Error: &ErrorDetails{
			Code: "ERROR",
		},
	}

	JSON(w, errorResponse.HTTPStatus(), response)
}

// NormalizeValidationError converts validation errors to a standard format
func NormalizeValidationError(err error) []FieldValidationError {
	var validationErrors []FieldValidationError

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			validationErrors = append(validationErrors, FieldValidationError{
				Field:   fieldErr.Field(),
				Message: getValidationMessage(fieldErr),
				Value:   fieldErr.Value().(string),
			})
		}
	} else {
		// Handle custom validation errors (like the ones from your existing validator)
		parts := strings.Split(err.Error(), ";")
		for _, part := range parts {
			validationErrors = append(validationErrors, FieldValidationError{
				Field:   "unknown",
				Message: strings.TrimSpace(part),
			})
		}
	}

	return validationErrors
}

// getValidationMessage returns a human-readable validation error message
func getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters long"
	case "len":
		return fe.Field() + " must be exactly " + fe.Param() + " characters long"
	case "numeric":
		return fe.Field() + " must be a number"
	case "alpha":
		return fe.Field() + " must contain only letters"
	case "alphanum":
		return fe.Field() + " must contain only letters and numbers"
	default:
		return fe.Field() + " is invalid"
	}
}
