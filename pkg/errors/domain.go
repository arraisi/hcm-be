package errors

import (
	"fmt"
	"net/http"
)

// DomainError represents application-specific errors with associated HTTP status codes
type DomainError struct {
	Code       string
	Message    string
	StatusCode int
	Details    map[string]interface{}
}

func (e *DomainError) Error() string {
	return e.Message
}

// HTTPStatus returns the HTTP status code for this error
func (e *DomainError) HTTPStatus() int {
	return e.StatusCode
}

// Common domain errors
var (
	// Client errors (4xx)
	ErrNotFound = &DomainError{
		Code:       "NOT_FOUND",
		Message:    "Resource not found",
		StatusCode: http.StatusNotFound,
	}

	ErrBadRequest = &DomainError{
		Code:       "BAD_REQUEST",
		Message:    "Invalid request",
		StatusCode: http.StatusBadRequest,
	}

	ErrUnauthorized = &DomainError{
		Code:       "UNAUTHORIZED",
		Message:    "Authentication required",
		StatusCode: http.StatusUnauthorized,
	}

	ErrForbidden = &DomainError{
		Code:       "FORBIDDEN",
		Message:    "Access denied",
		StatusCode: http.StatusForbidden,
	}

	ErrConflict = &DomainError{
		Code:       "CONFLICT",
		Message:    "Resource conflict",
		StatusCode: http.StatusConflict,
	}

	ErrValidation = &DomainError{
		Code:       "VALIDATION_FAILED",
		Message:    "Validation failed",
		StatusCode: http.StatusBadRequest,
	}

	// Server errors (5xx)
	ErrInternal = &DomainError{
		Code:       "INTERNAL_ERROR",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}

	ErrNotImplemented = &DomainError{
		Code:       "NOT_IMPLEMENTED",
		Message:    "Feature not implemented",
		StatusCode: http.StatusNotImplemented,
	}

	ErrServiceUnavailable = &DomainError{
		Code:       "SERVICE_UNAVAILABLE",
		Message:    "Service temporarily unavailable",
		StatusCode: http.StatusServiceUnavailable,
	}

	ErrNoEligibleSalesAssignment = &DomainError{
		Code:       "NO_ELIGIBLE_SALES",
		Message:    "No eligible sales found for assignment",
		StatusCode: http.StatusNotFound,
	}
)

// NewDomainError creates a new domain error with custom message
func NewDomainError(base *DomainError, message string) *DomainError {
	return &DomainError{
		Code:       base.Code,
		Message:    message,
		StatusCode: base.StatusCode,
		Details:    make(map[string]interface{}),
	}
}

// NewDomainErrorWithDetails creates a new domain error with custom message and details
func NewDomainErrorWithDetails(base *DomainError, message string, details map[string]interface{}) *DomainError {
	return &DomainError{
		Code:       base.Code,
		Message:    message,
		StatusCode: base.StatusCode,
		Details:    details,
	}
}

// Wrap creates a new domain error that wraps an existing error
func Wrap(base *DomainError, err error) *DomainError {
	message := base.Message
	if err != nil {
		message = fmt.Sprintf("%s: %v", base.Message, err)
	}

	return &DomainError{
		Code:       base.Code,
		Message:    message,
		StatusCode: base.StatusCode,
		Details:    make(map[string]interface{}),
	}
}

// GetHTTPStatus extracts HTTP status code from error
func GetHTTPStatus(err error) int {
	if domainErr, ok := err.(*DomainError); ok {
		return domainErr.HTTPStatus()
	}
	return http.StatusInternalServerError
}

// IsDomainError checks if an error is a domain error
func IsDomainError(err error) bool {
	_, ok := err.(*DomainError)
	return ok
}
