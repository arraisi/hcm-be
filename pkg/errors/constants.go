package errors

import (
	"errors"
	"net/http"
)

// ErrItem represents an error with its associated HTTP status code
type ErrItem struct {
	Error error
	Code  int
}

// ErrList represents a list of error mappings
type ErrList []ErrItem

var (
	// User domain errors
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserInvalidEmail  = errors.New("invalid email format")
	ErrUserInvalidData   = errors.New("invalid user data")
	ErrUserUnauthorized  = errors.New("unauthorized access")
	ErrUserForbidden     = errors.New("forbidden access")
	ErrUserIDRequired    = errors.New("user ID is required")
	// User-specific error list
	ErrListUser = ErrList{
		{Error: ErrUserNotFound, Code: http.StatusNotFound},
		{Error: ErrUserAlreadyExists, Code: http.StatusConflict},
		{Error: ErrUserInvalidEmail, Code: http.StatusBadRequest},
		{Error: ErrUserInvalidData, Code: http.StatusBadRequest},
		{Error: ErrUserUnauthorized, Code: http.StatusUnauthorized},
		{Error: ErrUserForbidden, Code: http.StatusForbidden},
		{Error: ErrUserIDRequired, Code: http.StatusBadRequest},
	}

	// Webhook domain errors
	ErrWebhookInvalidPayload    = errors.New("invalid JSON payload")
	ErrWebhookInvalidHeaders    = errors.New("header extraction failed")
	ErrWebhookInvalidAPIKey     = errors.New("invalid API key")
	ErrWebhookInvalidSignature  = errors.New("signature verification failed")
	ErrWebhookDuplicateEvent    = errors.New("duplicate event ID")
	ErrWebhookIdempotencyFailed = errors.New("failed to store idempotency key")
	ErrWebhookPublishFailed     = errors.New("failed to publish event")
	ErrWebhookUnsupportedStatus = errors.New("unsupported test drive status")
	ErrWebhookValidationFailed  = errors.New("payload validation failed")
	ErrWebhookReadBodyFailed    = errors.New("failed to read request body")
	ErrWebhookInvalidTimestamp  = errors.New("invalid timestamp format")
	ErrWebhookInvalidEventID    = errors.New("invalid event ID format")
	// Webhook-specific error list
	ErrListWebhook = ErrList{
		{Error: ErrWebhookInvalidPayload, Code: http.StatusBadRequest},
		{Error: ErrWebhookInvalidHeaders, Code: http.StatusBadRequest},
		{Error: ErrWebhookInvalidAPIKey, Code: http.StatusUnauthorized},
		{Error: ErrWebhookInvalidSignature, Code: http.StatusUnauthorized},
		{Error: ErrWebhookDuplicateEvent, Code: http.StatusConflict},
		{Error: ErrWebhookIdempotencyFailed, Code: http.StatusInternalServerError},
		{Error: ErrWebhookPublishFailed, Code: http.StatusInternalServerError},
		{Error: ErrWebhookUnsupportedStatus, Code: http.StatusBadRequest},
		{Error: ErrWebhookValidationFailed, Code: http.StatusBadRequest},
		{Error: ErrWebhookReadBodyFailed, Code: http.StatusBadRequest},
		{Error: ErrWebhookInvalidTimestamp, Code: http.StatusBadRequest},
		{Error: ErrWebhookInvalidEventID, Code: http.StatusBadRequest},
	}

	ErrTestDriveNotFound          = errors.New("test drive booking not found")
	ErrTestDriveInvalidData       = errors.New("invalid test drive data")
	ErrTestDriveCreateFailed      = errors.New("failed to create test drive booking")
	ErrTestDriveUpdateFailed      = errors.New("failed to update test drive booking")
	ErrTestDriveInvalidStatus     = errors.New("invalid test drive status")
	ErrTestDriveStatusInvalid     = errors.New("test drive status is invalid")
	ErrLeadsFollowUpStatusInvalid = errors.New("leads follow-up status is invalid")
	// Test Drive-specific error list
	ErrListTestDrive = ErrList{
		{Error: ErrTestDriveNotFound, Code: http.StatusNotFound},
		{Error: ErrTestDriveInvalidData, Code: http.StatusBadRequest},
		{Error: ErrTestDriveCreateFailed, Code: http.StatusInternalServerError},
		{Error: ErrTestDriveUpdateFailed, Code: http.StatusInternalServerError},
		{Error: ErrTestDriveInvalidStatus, Code: http.StatusBadRequest},
		{Error: ErrTestDriveStatusInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsFollowUpStatusInvalid, Code: http.StatusBadRequest},
	}

	// Service errors
	ErrDatabaseConnection = errors.New("database connection failed")
	ErrExternalService    = errors.New("external service error")
	ErrTimeout            = errors.New("request timeout")
	// Database-specific error list
	ErrListDatabase = ErrList{
		{Error: ErrDatabaseConnection, Code: http.StatusInternalServerError},
		{Error: ErrExternalService, Code: http.StatusBadGateway},
	}

	// Validation errors
	ErrValidationFailed = errors.New("validation failed")
	ErrInvalidJSON      = errors.New("invalid JSON payload")
	ErrMissingRequired  = errors.New("missing required fields")
	// Common error list that can be extended by specific domains
	ErrListCommon = ErrList{
		{Error: ErrValidationFailed, Code: http.StatusBadRequest},
		{Error: ErrInvalidJSON, Code: http.StatusBadRequest},
		{Error: ErrMissingRequired, Code: http.StatusBadRequest},
		{Error: ErrTimeout, Code: http.StatusRequestTimeout},
		{Error: ErrDatabaseConnection, Code: http.StatusInternalServerError},
		{Error: ErrExternalService, Code: http.StatusBadGateway},
	}
)

// Extend allows combining error lists
func (el ErrList) Extend(other ErrList) ErrList {
	return append(el, other...)
}
