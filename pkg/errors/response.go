package errors

import (
	"database/sql"
	"errors"
	"net/http"
)

// ErrorResponse represents an error response with HTTP status code mapping
type ErrorResponse struct {
	Err  error
	Code int
}

// NewErrorResponse creates a new ErrorResponse with the given status code and error
func NewErrorResponse(code int, err error) *ErrorResponse {
	// Handle common database errors
	if errors.Is(err, sql.ErrNoRows) {
		code = http.StatusNotFound
		err = ErrUserNotFound
	}

	return &ErrorResponse{
		Code: code,
		Err:  err,
	}
}

// NewErrorResponseFromList creates an ErrorResponse by matching the error against a list of known errors
func NewErrorResponseFromList(err error, errList ErrList) *ErrorResponse {
	var (
		code    = http.StatusInternalServerError
		respErr = err
	)

	// Combine with common error list
	errList = errList.Extend(ErrListCommon)

	// Check for exact error matches
	for _, v := range errList {
		if errors.Is(err, v.Error) {
			code = v.Code
			respErr = err

			// Special handling for database no rows error
			if errors.Is(err, sql.ErrNoRows) {
				respErr = ErrUserNotFound
			}
			break
		}
	}

	// If no match found and it's a domain error, use its status code
	if code == http.StatusInternalServerError {
		if domainErr, ok := err.(*DomainError); ok {
			code = domainErr.HTTPStatus()
			respErr = err
		}
	}

	return &ErrorResponse{
		Code: code,
		Err:  respErr,
	}
}

// Error implements the error interface
func (er *ErrorResponse) Error() string {
	if er.Err != nil {
		return er.Err.Error()
	}
	return "unknown error"
}

// HTTPStatus returns the HTTP status code
func (er *ErrorResponse) HTTPStatus() int {
	return er.Code
}

// IsDomainError checks if the underlying error is a domain error
func (er *ErrorResponse) IsDomainError() bool {
	_, ok := er.Err.(*DomainError)
	return ok
}

// GetDomainError returns the underlying domain error if it exists
func (er *ErrorResponse) GetDomainError() *DomainError {
	if domainErr, ok := er.Err.(*DomainError); ok {
		return domainErr
	}
	return nil
}
