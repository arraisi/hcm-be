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

	ErrTestDriveNotFound                        = errors.New("test drive booking not found")
	ErrTestDriveInvalidData                     = errors.New("invalid test drive data")
	ErrTestDriveCreateFailed                    = errors.New("failed to create test drive booking")
	ErrTestDriveUpdateFailed                    = errors.New("failed to update test drive booking")
	ErrTestDriveInvalidStatus                   = errors.New("invalid test drive status")
	ErrTestDriveStatusInvalid                   = errors.New("test drive status is invalid")
	ErrTestDriveLocationInvalid                 = errors.New("test drive location is invalid")
	ErrLeadsFollowUpStatusInvalid               = errors.New("leads follow-up status is invalid")
	ErrLeadsTypeInvalid                         = errors.New("leads type is invalid")
	ErrLeadsSourceInvalid                       = errors.New("leads source is invalid")
	ErrLeadsTAMLeadScoreInvalid                 = errors.New("leads TAM lead score is invalid")
	ErrLeadsOutletLeadScoreInvalid              = errors.New("leads outlet lead score is invalid")
	ErrTestDriveCancellationReasonInvalid       = errors.New("test drive cancellation reason is invalid")
	ErrTestDriveCancellationReasonRequired      = errors.New("cancellation reason is required for cancelled test drive")
	ErrTestDriveOtherCancellationReasonRequired = errors.New("other cancellation reason is required when cancellation reason is OTHERS")
	ErrTestDriveCustomerHasBooking              = errors.New("customer already has a test drive booking")
	// Test Drive-specific error list
	ErrListTestDrive = ErrList{
		{Error: ErrTestDriveNotFound, Code: http.StatusNotFound},
		{Error: ErrTestDriveInvalidData, Code: http.StatusBadRequest},
		{Error: ErrTestDriveCreateFailed, Code: http.StatusInternalServerError},
		{Error: ErrTestDriveUpdateFailed, Code: http.StatusInternalServerError},
		{Error: ErrTestDriveInvalidStatus, Code: http.StatusBadRequest},
		{Error: ErrTestDriveStatusInvalid, Code: http.StatusBadRequest},
		{Error: ErrTestDriveLocationInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsFollowUpStatusInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsTypeInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsSourceInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsTAMLeadScoreInvalid, Code: http.StatusBadRequest},
		{Error: ErrLeadsOutletLeadScoreInvalid, Code: http.StatusBadRequest},
		{Error: ErrTestDriveCancellationReasonInvalid, Code: http.StatusBadRequest},
		{Error: ErrTestDriveCancellationReasonRequired, Code: http.StatusBadRequest},
		{Error: ErrTestDriveOtherCancellationReasonRequired, Code: http.StatusBadRequest},
		{Error: ErrTestDriveCustomerHasBooking, Code: http.StatusConflict},
	}

	// Service Booking errors
	ErrServiceBookingCategoryInvalid   = errors.New("service booking category is invalid")
	ErrServiceBookingCustomerHasActive = errors.New("customer already has an active periodic maintenance booking")
	// Service Booking-specific error list
	ErrListServiceBooking = ErrList{
		{Error: ErrServiceBookingCategoryInvalid, Code: http.StatusBadRequest},
		{Error: ErrServiceBookingCustomerHasActive, Code: http.StatusConflict},
	}

	// Sales Order errors
	ErrSalesOrderNotFound     = errors.New("sales order not found")
	ErrSalesOrderInvalidData  = errors.New("invalid sales order data")
	ErrSalesOrderCreateFailed = errors.New("failed to create sales order")
	ErrSalesOrderUpdateFailed = errors.New("failed to update sales order")
	ErrSPKNotFound            = errors.New("SPK not found")
	ErrSPKInvalidStatus       = errors.New("invalid SPK status")
	ErrSPKCreateFailed        = errors.New("failed to create SPK")
	ErrSPKUpdateFailed        = errors.New("failed to update SPK")
	ErrPaymentInvalidStatus   = errors.New("invalid payment status")
	ErrDeliveryInvalidStatus  = errors.New("invalid delivery status")
	ErrInsuranceInvalidType   = errors.New("invalid insurance type")
	ErrAccessoryInvalidType   = errors.New("invalid accessory type")
	// Sales Order-specific error list
	ErrListSalesOrder = ErrList{
		{Error: ErrSalesOrderNotFound, Code: http.StatusNotFound},
		{Error: ErrSalesOrderInvalidData, Code: http.StatusBadRequest},
		{Error: ErrSalesOrderCreateFailed, Code: http.StatusInternalServerError},
		{Error: ErrSalesOrderUpdateFailed, Code: http.StatusInternalServerError},
		{Error: ErrSPKNotFound, Code: http.StatusNotFound},
		{Error: ErrSPKInvalidStatus, Code: http.StatusBadRequest},
		{Error: ErrSPKCreateFailed, Code: http.StatusInternalServerError},
		{Error: ErrSPKUpdateFailed, Code: http.StatusInternalServerError},
		{Error: ErrPaymentInvalidStatus, Code: http.StatusBadRequest},
		{Error: ErrDeliveryInvalidStatus, Code: http.StatusBadRequest},
		{Error: ErrInsuranceInvalidType, Code: http.StatusBadRequest},
		{Error: ErrAccessoryInvalidType, Code: http.StatusBadRequest},
	}

	ErrCustomerNotFound = errors.New("customer is not found")
	ErrListCustomer     = ErrList{
		{Error: ErrCustomerNotFound, Code: http.StatusNotFound},
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
