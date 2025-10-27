package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// SignatureVerifier handles HMAC SHA-256 signature verification for webhooks
type SignatureVerifier struct {
	secret []byte
}

// NewSignatureVerifier creates a new signature verifier with the given secret
func NewSignatureVerifier(secret string) *SignatureVerifier {
	return &SignatureVerifier{
		secret: []byte(secret),
	}
}

// GenerateSignature generates HMAC SHA-256 signature for the given payload
func (sv *SignatureVerifier) GenerateSignature(payload []byte) string {
	h := hmac.New(sha256.New, sv.secret)
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifySignature verifies if the provided signature matches the expected signature for the payload
func (sv *SignatureVerifier) VerifySignature(payload []byte, providedSignature string) error {
	expectedSignature := sv.GenerateSignature(payload)

	// Use hmac.Equal to prevent timing attacks
	providedBytes, err := hex.DecodeString(providedSignature)
	if err != nil {
		return fmt.Errorf("invalid signature format: %w", err)
	}

	expectedBytes, err := hex.DecodeString(expectedSignature)
	if err != nil {
		return fmt.Errorf("failed to decode expected signature: %w", err)
	}

	if !hmac.Equal(providedBytes, expectedBytes) {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}
