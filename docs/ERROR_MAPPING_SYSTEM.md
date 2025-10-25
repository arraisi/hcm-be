# Custom Error Mapping System (WIMS-like Pattern)

This document describes the custom error mapping system implemented to provide consistent HTTP status code mapping from service-level errors to HTTP responses.

## Overview

The error mapping system follows a pattern similar to WIMS where:
1. **Service Layer**: Returns specific error constants
2. **Handler Layer**: Uses `NewErrorResponseFromList` to map service errors to HTTP status codes
3. **Response Layer**: Automatically sends appropriate HTTP responses

## Architecture

### 1. Error Constants (`pkg/errors/constants.go`)

Define domain-specific errors with HTTP status code mappings:

```go
// User domain errors
var (
    ErrUserNotFound       = errors.New("user not found")
    ErrUserAlreadyExists  = errors.New("user already exists")
    ErrUserInvalidEmail   = errors.New("invalid email format")
    // ... more errors
)

// Error list with status code mappings
var ErrListUser = ErrList{
    {Error: ErrUserNotFound, Code: http.StatusNotFound},
    {Error: ErrUserAlreadyExists, Code: http.StatusConflict},
    {Error: ErrUserInvalidEmail, Code: http.StatusBadRequest},
    // ... more mappings
}
```

### 2. Error Response Type (`pkg/errors/response.go`)

The `ErrorResponse` type handles error-to-HTTP mapping:

```go
type ErrorResponse struct {
    Err  error
    Code int
}

// Map service error to HTTP status using error list
func NewErrorResponseFromList(err error, errList ErrList) *ErrorResponse
```

### 3. Response Helpers (`pkg/response/json.go`)

Helper function for sending error responses:

```go
func ErrorResponseJSON(w http.ResponseWriter, errorResponse *errors.ErrorResponse)
```

## Usage Pattern

### Service Layer

Return specific error constants:

```go
func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.ErrUserNotFound
        }
        return nil, errors.ErrDatabaseConnection
    }
    return user, nil
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) error {
    // Check if user exists
    if s.userExists(req.Email) {
        return errors.ErrUserAlreadyExists
    }
    
    // Validate email
    if !isValidEmail(req.Email) {
        return errors.ErrUserInvalidEmail
    }
    
    // Create user...
    return nil
}
```

### Handler Layer

Use `NewErrorResponseFromList` to map errors:

```go
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    // ... request parsing and validation ...
    
    if err := h.svc.Create(ctx, req); err != nil {
        // Map service error to HTTP status code
        errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
        response.ErrorResponseJSON(w, errorResponse)
        return
    }
    
    response.Created(w, data, "User created successfully")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    user, err := h.svc.Get(ctx, req)
    if err != nil {
        // Automatic error mapping
        errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
        response.ErrorResponseJSON(w, errorResponse)
        return
    }
    
    response.OK(w, user, "User retrieved successfully")
}
```

## Error Mapping Examples

| Service Error | HTTP Status | Description |
|---------------|-------------|-------------|
| `ErrUserNotFound` | 404 | User not found |
| `ErrUserAlreadyExists` | 409 | User already exists (conflict) |
| `ErrUserInvalidEmail` | 400 | Invalid email format |
| `ErrUserInvalidData` | 400 | Invalid user data |
| `ErrUserUnauthorized` | 401 | Unauthorized access |
| `ErrUserForbidden` | 403 | Forbidden access |
| `ErrDatabaseConnection` | 500 | Database connection failed |
| `ErrExternalService` | 502 | External service error |
| `ErrTimeout` | 408 | Request timeout |

## Benefits

1. **Consistency**: Standardized error handling across all handlers
2. **Maintainability**: Centralized error-to-status mapping
3. **Flexibility**: Easy to add new error types and modify mappings
4. **Separation of Concerns**: Service layer focuses on business logic, handlers focus on HTTP concerns
5. **Type Safety**: Compile-time error checking for error constants

## Migration Guide

### Before (Manual Error Handling)
```go
if err := h.svc.Create(ctx, req); err != nil {
    if errors.IsDomainError(err) {
        response.DomainError(w, err)
        return
    }
    
    if err.Error() == "user already exists" {
        response.DomainError(w, errors.NewDomainError(errors.ErrConflict, "User exists"))
        return
    }
    
    response.DomainError(w, errors.Wrap(errors.ErrInternal, err))
    return
}
```

### After (Error List Mapping)
```go
if err := h.svc.Create(ctx, req); err != nil {
    errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
    response.ErrorResponseJSON(w, errorResponse)
    return
}
```

## Custom Error Lists

Create domain-specific error lists:

```go
// Order domain errors
var ErrListOrder = ErrList{
    {Error: ErrOrderNotFound, Code: http.StatusNotFound},
    {Error: ErrOrderCancelled, Code: http.StatusGone},
    {Error: ErrOrderPaymentFailed, Code: http.StatusPaymentRequired},
}

// Extend with common errors
func GetOrderErrorList() ErrList {
    return ErrListOrder.Extend(ErrListCommon)
}
```

## Best Practices

1. **Service Layer**: Always return predefined error constants
2. **Error Lists**: Group related errors by domain
3. **Handler Layer**: Use `NewErrorResponseFromList` consistently
4. **Documentation**: Document all error codes and their meanings
5. **Testing**: Test error mappings with unit tests

## Testing Error Mappings

```go
func TestErrorMapping(t *testing.T) {
    testCases := []struct {
        serviceError   error
        expectedStatus int
    }{
        {errors.ErrUserNotFound, http.StatusNotFound},
        {errors.ErrUserAlreadyExists, http.StatusConflict},
        {errors.ErrDatabaseConnection, http.StatusInternalServerError},
    }
    
    for _, tc := range testCases {
        errorResponse := errors.NewErrorResponseFromList(tc.serviceError, errors.ErrListUser)
        assert.Equal(t, tc.expectedStatus, errorResponse.HTTPStatus())
    }
}
```