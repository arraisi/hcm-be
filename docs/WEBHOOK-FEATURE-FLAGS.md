# Webhook Feature Flags

The webhook endpoint supports configurable validation through feature flags in the `config.yaml` file. This allows you to enable or disable specific validation steps based on your deployment requirements.

## Configuration

Add the following to your `config.yaml`:

```yaml
featureFlag:
  webhook:
    enableSignatureValidation: true
    enableTimestampValidation: true
```

## Feature Flags

### `enableSignatureValidation`

**Default:** `true`

Controls whether HMAC-SHA256 signature verification is performed on incoming webhook requests.

- **When `true`:** The `X-Signature` header is verified against the request body using the configured HMAC secret
- **When `false`:** Signature verification is skipped

**Use cases for disabling:**

- Development/testing environments
- Internal networks where signature verification is not required
- Debugging webhook integration issues

### `enableTimestampValidation`

**Default:** `true`

Controls whether timestamp validation is performed on incoming webhook requests.

- **When `true`:** The `X-Event-Timestamp` header is validated to be within 5 minutes of the current time
- **When `false`:** Timestamp validation is skipped

**Use cases for disabling:**

- Testing with old or static test data
- Environments with clock synchronization issues
- Replay of historical webhook events

## Security Considerations

### Always Enabled Validations

The following validations are **always performed** regardless of feature flag settings:

1. **API Key Validation**: The `X-API-Key` header must match the configured API key
2. **Header Format Validation**: All required headers must be present and properly formatted
3. **Content Type Validation**: Request must have `Content-Type: application/json`
4. **Payload Validation**: JSON structure and data types are validated
5. **Idempotency**: Duplicate event IDs are rejected

### Security Recommendations

- **Production environments** should have both feature flags set to `true`
- **Signature validation** provides the strongest security guarantee
- **Timestamp validation** prevents replay attacks
- Only disable these validations in trusted environments or for specific testing scenarios

## Configuration Examples

### Maximum Security (Production)

```yaml
featureFlag:
  webhook:
    enableSignatureValidation: true
    enableTimestampValidation: true
```

### Development/Testing

```yaml
featureFlag:
  webhook:
    enableSignatureValidation: false # Skip signature for easier testing
    enableTimestampValidation: true # Keep timestamp validation
```

### Debugging Mode

```yaml
featureFlag:
  webhook:
    enableSignatureValidation: false # Skip signature validation
    enableTimestampValidation: false # Skip timestamp validation
```

## Testing Feature Flags

Use the provided test script to verify feature flag behavior:

```bash
./scripts/test-feature-flags.sh
```

This script will:

1. Test valid requests with all validations
2. Test invalid signatures (should fail when `enableSignatureValidation: true`)
3. Test old timestamps (should fail when `enableTimestampValidation: true`)

## Environment Variables

Feature flags can also be set via environment variables:

```bash
export APP_FEATUREFLAG_WEBHOOK_ENABLESIGNATUREVALIDATION=false
export APP_FEATUREFLAG_WEBHOOK_ENABLETIMESTAMPVALIDATION=false
```

## Implementation Details

The feature flags are checked in the `validateWebhookRequest` method:

1. **Header validation** and **API key check** are always performed
2. **Signature validation** is performed only if `enableSignatureValidation` is `true`
3. **Timestamp validation** is performed only if `enableTimestampValidation` is `true`

This ensures that critical security checks (API key) are never bypassed, while allowing flexibility for signature and timestamp validation based on deployment needs.
