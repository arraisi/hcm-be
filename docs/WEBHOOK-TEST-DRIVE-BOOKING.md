# Test Drive Booking Webhook Documentation

## Overview
Endpoint webhook untuk menerima Test Drive Booking Request dari DI/DX Platform menggunakan arsitektur Clean Architecture.

## Endpoint Details
- **Method**: `POST`
- **Path**: `/webhook/test-drive-booking`
- **Content-Type**: `application/json`

## Authentication & Security

### Required Headers
| Header | Type | Description |
|--------|------|-------------|
| `Content-Type` | string | Must be `application/json` |
| `X-API-Key` | string | API key for authentication (configured in config) |
| `X-Signature` | string | HMAC-SHA256 hex signature of raw request body |
| `X-Event-Id` | string | UUID v4 for idempotency |
| `X-Event-Timestamp` | string | Unix timestamp in seconds |

### Signature Generation
The signature is generated using HMAC-SHA256:
```
signature = hex(HMAC_SHA256(hmacSecret, rawRequestBody))
```

## Request Body Schema

```json
{
  "process": "test drive request",
  "event_ID": "05dbe854-74a4-4e0d-be00-da098d3569d6",
  "timestamp": 1708726960,
  "data": {
    "one_account": {
      "one_account_ID": "GMA04GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
      "first_name": "John",
      "last_name": "Doe",
      "gender": "MALE",
      "phone_number": "1234567890",
      "email": "john.doe@example.com"
    },
    "test_drive": {
      "test_drive_ID": "0d5be854-74a4-4e0d-be00-da098d3529d5",
      "test_drive_number": "TUT010026-02-20241107959",
      "katashiki_code": "NSP170R-MWYXKD",
      "model": "Innova Zenix",
      "variant": "2.0 Q A/T",
      "created_datetime": 1709096400,
      "test_drive_datetime_start": 1709085600,
      "test_drive_datetime_end": 1709388000,
      "location": "DEALER",
      "outlet_ID": "AST01329",
      "outlet_name": "Astrido Toyota Bitung",
      "test_drive_status": "SUBMITTED",
      "cancellation_reason": null,
      "other_cancellation_reason": null,
      "customer_driving_consent": true
    },
    "leads": {
      "leads_ID": "44ae2529-98e4-41f4-bae8-f305f609932d",
      "leads_type": "TEST_DRIVE_REQUEST",
      "leads_follow_up_status": "ON_CONSIDERATION",
      "leads_preference_contact_time_start": "09:30",
      "leads_preference_contact_time_end": "10:30",
      "leads_source": "OFFLINE_WALK_IN_OR_CALL_IN",
      "additional_notes": "I am loyal customer can you give me a discount?"
    },
    "score": {
      "iam_lead_score": "HOT",
      "outlet_lead_score": "MEDIUM",
      "parameter": {
        "purchase_plan_criteria": "31_DAYS_TO_INFINITE",
        "payment_prefer_criteria": "CASH",
        "negotiation_criteria": "HAVE_STARTED_NEGOTIATIONS",
        "test_drive_criteria": "COMPLETED",
        "trade_in_criteria": "DELIVERY",
        "browsing_history_criteria": "MORE_THAN_5_PAGES",
        "vehicle_age_criteria": "MORE_THAN_2.5_YEARS"
      }
    }
  }
}
```

## Validation Rules

### Header Validation
- `X-API-Key` must match the configured API key
- `X-Signature` must be valid HMAC-SHA256 hex string
- `X-Event-Id` must be valid UUID v4
- `X-Event-Timestamp` must be within ±15 minutes of server time

### Body Validation
- `process` must be "test drive request"
- `event_ID` must be UUID v4 and match `X-Event-Id` header
- `data.leads.leads_type` must be "TEST_DRIVE_REQUEST"
- `data.test_drive.test_drive_status` must be one of: "SUBMITTED", "CHANGE_REQUEST", "CANCEL_SUBMITTED"
- `data.one_account.gender` must be "MALE" or "FEMALE"
- `data.one_account.email` is optional but must be valid email if provided
- All timestamp fields must be valid Unix timestamps

### Idempotency
- Duplicate `event_ID` within 24 hours will be rejected with `409 Conflict`
- Uses in-memory storage (can be extended to Redis for production)

## Response Format

### Success Response (202 Accepted)
```json
{
  "data": {
    "eventId": "05dbe854-74a4-4e0d-be00-da098d3569d6",
    "status": "RECEIVED"
  },
  "message": "accepted"
}
```

### Error Response (4xx/5xx)
```json
{
  "data": {},
  "message": "human readable error message"
}
```

## Error Codes

| Status Code | Description |
|-------------|-------------|
| 202 | Request accepted and published to MQ |
| 400 | Bad request (validation error, invalid JSON, etc.) |
| 401 | Unauthorized (invalid API key or signature) |
| 409 | Conflict (duplicate event ID) |
| 500 | Internal server error |

## Internal Processing

1. **Header Validation**: Validates all required headers and formats
2. **Authentication**: Verifies API key and HMAC signature
3. **Timestamp Check**: Ensures request is within acceptable time window
4. **Body Validation**: Validates JSON structure and business rules
5. **Idempotency Check**: Prevents duplicate processing
6. **Message Publishing**: Publishes normalized event to MQ topic `test_drive.booking.received`

## Configuration

Update your `config.yaml`:

```yaml
webhook:
  apiKey: 'your-webhook-api-key-here'
  hmacSecret: 'your-webhook-hmac-secret-here'
```

Or set environment variables:
- `APP_WEBHOOK_APIKEY`
- `APP_WEBHOOK_HMACSECRET`

## Testing

Use the provided test files to verify webhook functionality:
- `internal/http/handlers/webhook_test.go`

Run tests:
```bash
go test ./internal/http/handlers/webhook_test.go
```

## Message Queue Integration

Successfully processed webhooks are published to:
- **Topic**: `test_drive.booking.received`
- **Key**: `event_ID` (for message ordering)
- **Payload**: Normalized event structure

### MQ Event Format
```json
{
  "eventId": "05dbe854-74a4-4e0d-be00-da098d3569d6",
  "eventType": "test_drive_booking_received",
  "timestamp": 1708726960,
  "source": "webhook",
  "data": {
    // Original booking event data
  }
}
```

## Architecture Components

### Domain Models
- `BookingEvent`: Root aggregate for webhook payload
- `OneAccount`: Customer account information
- `TestDrive`: Test drive details
- `Leads`: Sales lead information
- `Score`: Lead scoring data

### Utilities
- `SignatureVerifier`: HMAC-SHA256 signature verification
- `Validator`: Comprehensive request validation
- `IdempotencyStore`: Duplicate event prevention
- `Publisher`: Message queue publishing interface

### Clean Architecture Layers
- **Domain**: Business entities and rules
- **Handlers**: HTTP request handling
- **Utilities**: Cross-cutting concerns (validation, security)
- **Infrastructure**: MQ publishing (with in-memory fallback)

## Development Notes

- Current implementation uses in-memory MQ publisher for testing
- For production, replace with actual MQ implementation (RabbitMQ, Kafka, etc.)
- Idempotency store currently uses in-memory storage with TTL
- Consider Redis-based idempotency store for production
- All validations follow fail-fast principle
- Comprehensive error messages for debugging