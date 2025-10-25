package webhook

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Validator handles validation for webhook requests
type Validator struct {
	validator *validator.Validate
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	return fmt.Sprintf("validation failed: %s", ve[0].Message)
}

// NewValidator creates a new webhook validator
func NewValidator() *Validator {
	v := validator.New()

	// Register custom validators
	_ = v.RegisterValidation("uuid4", validateUUID4)

	return &Validator{
		validator: v,
	}
}

// ValidateHeaders validates the required webhook headers
func (v *Validator) ValidateHeaders(headers map[string]string) error {
	var errors ValidationErrors

	// Check required headers
	requiredHeaders := []string{
		"Content-Type",
		"X-API-Key",
		"X-Signature",
		"X-Event-Id",
		"X-Event-Timestamp",
	}

	for _, header := range requiredHeaders {
		if value, exists := headers[header]; !exists || value == "" {
			errors = append(errors, ValidationError{
				Field:   header,
				Message: fmt.Sprintf("header %s is required", header),
			})
		}
	}

	// Validate Content-Type
	if contentType, exists := headers["Content-Type"]; exists {
		if contentType != "application/json" {
			errors = append(errors, ValidationError{
				Field:   "Content-Type",
				Message: "Content-Type must be application/json",
			})
		}
	}

	// Validate X-Event-Id format (UUID v4)
	if eventID, exists := headers["X-Event-Id"]; exists && eventID != "" {
		if !isValidUUID4(eventID) {
			errors = append(errors, ValidationError{
				Field:   "X-Event-Id",
				Message: "X-Event-Id must be a valid UUID v4",
			})
		}
	}

	// Validate X-Event-Timestamp format (unix timestamp)
	if timestamp, exists := headers["X-Event-Timestamp"]; exists && timestamp != "" {
		if !isValidUnixTimestamp(timestamp) {
			errors = append(errors, ValidationError{
				Field:   "X-Event-Timestamp",
				Message: "X-Event-Timestamp must be a valid unix timestamp",
			})
		}
	}

	// Validate X-Signature format (hex)
	if signature, exists := headers["X-Signature"]; exists && signature != "" {
		if !isValidHexSignature(signature) {
			errors = append(errors, ValidationError{
				Field:   "X-Signature",
				Message: "X-Signature must be a valid hex string",
			})
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// ValidateTimestamp validates the timestamp is within acceptable range (±15 minutes)
func (v *Validator) ValidateTimestamp(timestampStr string) error {
	// Parse the timestamp
	var timestamp int64
	if _, err := fmt.Sscanf(timestampStr, "%d", &timestamp); err != nil {
		return fmt.Errorf("invalid timestamp format")
	}

	requestTime := time.Unix(timestamp, 0)
	now := time.Now()

	// Check if timestamp is within ±15 minutes
	diff := now.Sub(requestTime)
	if diff < -15*time.Minute || diff > 15*time.Minute {
		return fmt.Errorf("timestamp is outside acceptable range (±15 minutes)")
	}

	return nil
}

// ValidateEventIDMatch validates that header event ID matches body event ID
func (v *Validator) ValidateEventIDMatch(headerEventID, bodyEventID string) error {
	if headerEventID != bodyEventID {
		return fmt.Errorf("header X-Event-Id does not match body event_ID")
	}
	return nil
}

// ValidateStruct validates a struct using the validator tags
func (v *Validator) ValidateStruct(s interface{}) error {
	if err := v.validator.Struct(s); err != nil {
		var errors ValidationErrors

		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				message = fmt.Sprintf("%s must be a valid email", err.Field())
			case "oneof":
				message = fmt.Sprintf("%s must be one of: %s", err.Field(), err.Param())
			case "eq":
				message = fmt.Sprintf("%s must equal %s", err.Field(), err.Param())
			case "uuid4":
				message = fmt.Sprintf("%s must be a valid UUID v4", err.Field())
			default:
				message = fmt.Sprintf("%s validation failed", err.Field())
			}

			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Message: message,
			})
		}

		return errors
	}

	return nil
}

// Helper functions

func isValidUUID4(u string) bool {
	parsed, err := uuid.Parse(u)
	if err != nil {
		return false
	}
	return parsed.Version() == 4
}

func isValidUnixTimestamp(timestamp string) bool {
	var ts int64
	_, err := fmt.Sscanf(timestamp, "%d", &ts)
	return err == nil && ts > 0
}

func isValidHexSignature(signature string) bool {
	// Check if it's a valid hex string (64 characters for SHA-256)
	matched, _ := regexp.MatchString("^[a-fA-F0-9]{64}$", signature)
	return matched
}

// Custom validator for UUID v4
func validateUUID4(fl validator.FieldLevel) bool {
	return isValidUUID4(fl.Field().String())
}
