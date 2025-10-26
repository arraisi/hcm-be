# HTTP Response Layer Documentation

## Overview

This document describes the unified HTTP response layer implemented across the codebase. All HTTP endpoints now return consistent JSON responses with proper error handling, typed domain errors, and panic recovery.

## Features

- **Unified JSON envelope**: All responses follow the same structure
- **Centralized response helpers**: Functions for common HTTP status codes
- **Typed domain errors**: Structured error handling with HTTP status mapping
- **Panic-safe recovery**: Middleware that catches panics and returns JSON 500
- **Request ID tracking**: Every request gets a unique identifier
- **Validation error normalization**: Consistent format for validation errors

## Response Structure

All HTTP responses follow this unified structure:

```json
{
  "data": <any>,           // The actual response payload
  "message": "<string>",   // Human-readable message
  "meta": <optional>,      // Optional metadata (pagination, etc.)
  "error": {               // Only present in error responses
    "code": "<string>",    // Error code for programmatic handling
    "details": <object>    // Additional error details
  }
}
```

## Success Responses

### Basic Success (200 OK)

```go
response.OK(w, userData, "User retrieved successfully")
```

### Created (201 Created)

```go
response.Created(w, newUser, "User created successfully")
```

### No Content (204 No Content)

```go
response.NoContent(w)
```

### With Metadata (Pagination)

```go
meta := map[string]interface{}{
    "pagination": map[string]interface{}{
        "total": 100,
        "limit": 10,
        "offset": 0,
    },
}

resp := response.Response{
    Data:    users,
    Message: "Users retrieved successfully",
    Meta:    meta,
}
response.JSON(w, http.StatusOK, resp)
```

## Error Responses

### Domain Errors

Domain errors provide structured error handling with automatic HTTP status mapping:

```go
// In service layer
func (s *UserService) GetUser(id string) error {
    if id == "" {
        return errors.NewDomainError(errors.ErrBadRequest, "User ID is required")
    }

    if userNotFound {
        return errors.NewDomainError(errors.ErrNotFound, "User not found")
    }

    return nil
}

// In handler
if err := service.GetUser(id); err != nil {
    response.DomainError(w, err) // Automatically maps to correct HTTP status
    return
}
```

### Validation Errors

Validation errors are normalized to a consistent format:

```go
// Using go-playground/validator
if err := validate.Struct(req); err != nil {
    validationErrors := response.NormalizeValidationError(err)
    response.Validation(w, validationErrors)
    return
}

// Custom validation errors
customErr := fmt.Errorf("email is required;name must be at least 2 characters")
validationErrors := response.NormalizeValidationError(customErr)
response.Validation(w, validationErrors)
```

### Quick Error Responses

```go
response.BadRequest(w, "Invalid parameters")
response.NotFound(w, "Resource not found")
response.Unauthorized(w, "Authentication required")
response.Forbidden(w, "Access denied")
response.Conflict(w, "Resource already exists")
response.InternalServerError(w, "Something went wrong")
```

## Domain Error Types

Pre-defined domain errors with appropriate HTTP status codes:

| Error Type              | HTTP Status | Usage                      |
| ----------------------- | ----------- | -------------------------- |
| `ErrBadRequest`         | 400         | Invalid request parameters |
| `ErrUnauthorized`       | 401         | Authentication required    |
| `ErrForbidden`          | 403         | Access denied              |
| `ErrNotFound`           | 404         | Resource not found         |
| `ErrConflict`           | 409         | Resource conflict          |
| `ErrValidation`         | 400         | Validation failure         |
| `ErrInternal`           | 500         | Internal server error      |
| `ErrNotImplemented`     | 501         | Feature not implemented    |
| `ErrServiceUnavailable` | 503         | Service unavailable        |

### Creating Custom Domain Errors

```go
// Simple error with custom message
err := errors.NewDomainError(errors.ErrConflict, "Email already exists")

// Error with additional details
err := errors.NewDomainErrorWithDetails(
    errors.ErrConflict,
    "User with this email already exists",
    map[string]interface{}{
        "conflicting_field": "email",
        "existing_id": "123",
    },
)

// Wrapping existing errors
err := errors.Wrap(errors.ErrInternal, dbError)
```

## Middleware

### Request ID Middleware

Automatically generates or extracts request IDs:

```go
handler = middleware.RequestID(handler)

// Access request ID in handlers
requestID := middleware.GetRequestID(ctx)
```

### Recovery Middleware

Catches panics and returns JSON 500 responses:

```go
handler = middleware.Recovery(handler)
```

### Middleware Stack Example

```go
var handler http.Handler = yourRouter

// Apply middleware in order
handler = middleware.RequestID(handler)
handler = middleware.Recovery(handler)
```

## Migration Guide

### Before (Old System)

```go
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(id)
    if err != nil {
        if err.Error() == "user not found" {
            response.Error(w, http.StatusNotFound, "user not found")
            return
        }
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }
    response.JSON(w, http.StatusOK, map[string]any{"data": user, "message": ""})
}
```

### After (New System)

```go
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(id)
    if err != nil {
        response.DomainError(w, err) // Automatic status mapping
        return
    }
    response.OK(w, user, "User retrieved successfully")
}
```

## Response Examples

### Success Response

```json
{
	"data": {
		"id": "123",
		"name": "John Doe",
		"email": "john@example.com"
	},
	"message": "User retrieved successfully"
}
```

### Validation Error Response

```json
{
	"data": null,
	"message": "Validation failed",
	"error": {
		"code": "VALIDATION_FAILED",
		"details": {
			"validation_errors": [
				{
					"field": "email",
					"message": "email is required"
				},
				{
					"field": "name",
					"message": "name must be at least 2 characters long"
				}
			]
		}
	}
}
```

### Domain Error Response

```json
{
	"data": null,
	"message": "User not found",
	"error": {
		"code": "NOT_FOUND",
		"details": {}
	}
}
```

### Paginated Response

```json
{
	"data": [
		{ "id": "123", "name": "John Doe" },
		{ "id": "124", "name": "Jane Doe" }
	],
	"message": "Users retrieved successfully",
	"meta": {
		"pagination": {
			"total": 100,
			"limit": 10,
			"offset": 0,
			"pages": 10
		}
	}
}
```

## Best Practices

1. **Use domain errors in service layer**: Always return structured domain errors instead of generic errors
2. **Consistent messaging**: Provide meaningful, user-friendly messages
3. **Include metadata**: Use the meta field for pagination, versioning, or other contextual information
4. **Validate early**: Use validation middleware and normalize errors consistently
5. **Handle specific cases**: Map specific error conditions to appropriate domain errors
6. **Log with request IDs**: Include request IDs in logs for traceability

## Testing

The response system can be easily tested:

```go
func TestHandler(t *testing.T) {
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    var resp response.Response
    json.Unmarshal(w.Body.Bytes(), &resp)

    assert.Equal(t, "User retrieved successfully", resp.Message)
    assert.NotNil(t, resp.Data)
}
```

## Backward Compatibility

The new system is designed to coexist with existing code. Old handlers can be migrated incrementally without breaking the API contract.
