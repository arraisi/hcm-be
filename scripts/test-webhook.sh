#!/bin/bash

# Test script for webhook endpoint
# Usage: ./test-webhook.sh [endpoint_url]

WEBHOOK_URL="${1:-http://localhost:8080/webhook/test-drive-booking}"
API_KEY="your-webhook-api-key-here"
HMAC_SECRET="your-webhook-hmac-secret-here"

# Function to generate HMAC-SHA256 signature
generate_signature() {
    local payload="$1"
    echo -n "$payload" | openssl dgst -sha256 -hmac "$HMAC_SECRET" -hex | cut -d' ' -f2
}

# Function to generate UUID v4
generate_uuid() {
    if command -v uuidgen >/dev/null 2>&1; then
        uuidgen | tr '[:upper:]' '[:lower:]'
    else
        cat /proc/sys/kernel/random/uuid
    fi
}

# Generate test data
EVENT_ID=$(generate_uuid)
TIMESTAMP=$(date +%s)

# Create payload
PAYLOAD=$(cat <<EOF
{
  "process": "test drive request",
  "event_ID": "$EVENT_ID",
  "timestamp": $TIMESTAMP,
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
      "test_drive_ID": "$(generate_uuid)",
      "test_drive_number": "TUT010026-02-20241107959",
      "katashiki_code": "NSP170R-MWYXKD",
      "model": "Innova Zenix",
      "variant": "2.0 Q A/T",
      "created_datetime": $TIMESTAMP,
      "test_drive_datetime_start": $((TIMESTAMP + 3600)),
      "test_drive_datetime_end": $((TIMESTAMP + 7200)),
      "location": "DEALER",
      "outlet_ID": "AST01329",
      "outlet_name": "Astrido Toyota Bitung",
      "test_drive_status": "SUBMITTED",
      "cancellation_reason": null,
      "other_cancellation_reason": null,
      "customer_driving_consent": true
    },
    "leads": {
      "leads_ID": "$(generate_uuid)",
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
EOF
)

# Generate signature
SIGNATURE=$(generate_signature "$PAYLOAD")

echo "Testing webhook endpoint: $WEBHOOK_URL"
echo "Event ID: $EVENT_ID"
echo "Timestamp: $TIMESTAMP"
echo "Signature: $SIGNATURE"
echo ""

# Make the request
curl -X POST "$WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -H "X-Signature: $SIGNATURE" \
  -H "X-Event-Id: $EVENT_ID" \
  -H "X-Event-Timestamp: $TIMESTAMP" \
  -d "$PAYLOAD" \
  -w "\nHTTP Status: %{http_code}\nResponse Time: %{time_total}s\n" \
  -s

echo ""
echo "Test completed!"

# Test duplicate request
echo ""
echo "Testing duplicate request (should return 409)..."
curl -X POST "$WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -H "X-Signature: $SIGNATURE" \
  -H "X-Event-Id: $EVENT_ID" \
  -H "X-Event-Timestamp: $TIMESTAMP" \
  -d "$PAYLOAD" \
  -w "\nHTTP Status: %{http_code}\nResponse Time: %{time_total}s\n" \
  -s