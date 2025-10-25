# Webhook Handler Refactoring - Reusable Methods

## Overview

Webhook handler telah di-refactor untuk memisahkan logika validasi dan processing menjadi method-method yang reusable dan mudah di-test secara individual.

## Reusable Methods

### 1. Header Processing Methods

#### `extractHeaders(r *http.Request) (map[string]string, error)`

- **Purpose**: Mengekstrak semua required headers dari HTTP request
- **Returns**: Map berisi header values atau error jika ada header yang missing
- **Headers yang diekstrak**:
  - `Content-Type`
  - `X-API-Key`
  - `X-Signature`
  - `X-Event-Id`
  - `X-Event-Timestamp`

```go
headers, err := h.extractHeaders(r)
if err != nil {
    // Handle missing headers
}
```

#### `validateHeaders(headers map[string]string) error`

- **Purpose**: Memvalidasi format dan content dari headers
- **Validations**:
  - Content-Type harus `application/json`
  - X-Event-Id harus valid UUID v4
  - X-Event-Timestamp harus valid unix timestamp
  - X-Signature harus valid hex string (64 karakter)

```go
if err := h.validateHeaders(headers); err != nil {
    // Handle header validation errors
}
```

### 2. Authentication Methods

#### `validateAPIKey(headers map[string]string) error`

- **Purpose**: Memvalidasi API key dari header
- **Returns**: Error jika API key tidak cocok dengan config

```go
if err := h.validateAPIKey(headers); err != nil {
    // Handle invalid API key - return 401 Unauthorized
}
```

#### `validateSignature(body []byte, signature string) error`

- **Purpose**: Memvalidasi HMAC-SHA256 signature dari raw request body
- **Security**: Menggunakan constant-time comparison untuk mencegah timing attacks

```go
if err := h.validateSignature(body, headers["X-Signature"]); err != nil {
    // Handle signature verification failure - return 401 Unauthorized
}
```

#### `validateTimestamp(timestamp string) error`

- **Purpose**: Memvalidasi bahwa timestamp dalam rentang ±15 menit dari server time
- **Prevents**: Replay attacks dengan timestamp yang sudah expired

```go
if err := h.validateTimestamp(headers["X-Event-Timestamp"]); err != nil {
    // Handle timestamp validation failure
}
```

### 3. Payload Processing Methods

#### `validatePayload(bookingEvent *domain.BookingEvent) error`

- **Purpose**: Memvalidasi struktur dan business rules dari payload
- **Uses**: Struct validation tags dengan validator library

```go
if err := h.validatePayload(&bookingEvent); err != nil {
    // Handle payload validation errors
}
```

#### `validateEventIDMatch(headerEventID, bodyEventID string) error`

- **Purpose**: Memvalidasi bahwa event ID di header sama dengan di body
- **Security**: Mencegah event ID spoofing

```go
if err := h.validateEventIDMatch(headers["X-Event-Id"], bookingEvent.EventID); err != nil {
    // Handle event ID mismatch
}
```

### 4. Idempotency Methods

#### `checkIdempotency(eventID string) error`

- **Purpose**: Mengecek apakah event sudah pernah diproses
- **Returns**: Error jika duplicate event ID ditemukan

```go
if err := h.checkIdempotency(bookingEvent.EventID); err != nil {
    // Handle duplicate event - return 409 Conflict
}
```

#### `storeIdempotencyKey(eventID string) error`

- **Purpose**: Menyimpan event ID untuk future idempotency checks
- **TTL**: 24 jam (configurable)

```go
if err := h.storeIdempotencyKey(bookingEvent.EventID); err != nil {
    // Handle storage failure - return 500 Internal Server Error
}
```

### 5. Event Publishing Methods

#### `publishEvent(bookingEvent *domain.BookingEvent) error`

- **Purpose**: Mempublikasi normalized event ke message queue
- **Topic**: `test_drive.booking.received`
- **Key**: Event ID (untuk message ordering)

```go
if err := h.publishEvent(&bookingEvent); err != nil {
    // Handle publish failure - return 500 Internal Server Error
}
```

### 6. Composite Methods

#### `validateWebhookRequest(headers map[string]string, body []byte) error`

- **Purpose**: Menjalankan semua validasi header dan authentication
- **Sequence**:
  1. Header format validation
  2. API key validation
  3. Signature validation
  4. Timestamp validation

```go
if err := h.validateWebhookRequest(headers, body); err != nil {
    // Determine appropriate HTTP status code based on error type
    statusCode := h.getStatusCodeForError(err)
    h.sendErrorResponse(w, statusCode, err.Error())
    return
}
```

#### `validateAndProcessBooking(headers map[string]string, bookingEvent *domain.BookingEvent) error`

- **Purpose**: Menjalankan semua validasi payload dan processing
- **Sequence**:
  1. Payload structure validation
  2. Event ID match validation
  3. Idempotency check
  4. Store idempotency key
  5. Publish to message queue

```go
if err := h.validateAndProcessBooking(headers, &bookingEvent); err != nil {
    // Handle processing errors with appropriate status codes
}
```

## Error Handling Strategy

### HTTP Status Code Mapping

- **400 Bad Request**: Validation errors, malformed JSON, timestamp issues
- **401 Unauthorized**: Invalid API key, signature verification failure
- **409 Conflict**: Duplicate event ID (idempotency violation)
- **500 Internal Server Error**: Storage failure, message queue publish failure

### Error Message Standards

- Semua error messages mengikuti Go convention (lowercase, no capitalization)
- Error messages bersifat descriptive untuk debugging
- Sensitive information tidak di-expose di error response

## Benefits of Refactoring

### 1. **Maintainability**

- Setiap validation logic terpisah dan focused
- Easy to modify individual validation rules
- Clear separation of concerns

### 2. **Testability**

- Setiap method dapat di-unit test secara individual
- Mock dependencies dapat di-inject dengan mudah
- Test coverage lebih comprehensive

### 3. **Reusability**

- Methods dapat digunakan untuk webhook endpoints lain
- Validation logic dapat di-share antar handlers
- Consistent error handling patterns

### 4. **Readability**

- Main handler method lebih clean dan readable
- Business logic flow yang jelas
- Self-documenting method names

### 5. **Error Handling**

- Consistent error response format
- Appropriate HTTP status codes
- Proper error logging and debugging info

## Usage Example

```go
func (h *WebhookHandler) AnotherWebhookEndpoint(w http.ResponseWriter, r *http.Request) {
    // Read body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
        return
    }

    // Extract and validate headers (reusable!)
    headers, err := h.extractHeaders(r)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    // Validate webhook requirements (reusable!)
    if err := h.validateWebhookRequest(headers, body); err != nil {
        statusCode := h.getStatusCodeForError(err)
        h.sendErrorResponse(w, statusCode, err.Error())
        return
    }

    // Parse and process specific payload for this endpoint
    // ... endpoint-specific logic
}
```

## Testing Strategy

Dengan refactoring ini, testing menjadi lebih modular:

```go
func TestWebhookHandler_validateAPIKey(t *testing.T) {
    handler := setupTestHandler()

    tests := []struct {
        name        string
        apiKey      string
        expectError bool
    }{
        {"valid key", "correct-api-key", false},
        {"invalid key", "wrong-key", true},
        {"empty key", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            headers := map[string]string{"X-API-Key": tt.apiKey}
            err := handler.validateAPIKey(headers)

            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

## Future Enhancements

1. **Metrics**: Tambahkan metrics untuk setiap validation step
2. **Logging**: Enhanced logging untuk debugging dan monitoring
3. **Circuit Breaker**: Untuk external dependencies (MQ, storage)
4. **Rate Limiting**: Per API key atau per IP
5. **Configuration**: Make validation rules configurable
6. **Middleware**: Convert some validations to reusable middleware
