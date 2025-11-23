package auth

import (
	"context"
	"errors"

	"github.com/arraisi/hcm-be/internal/auth"
)

// TokenService defines operations for generating service tokens.
type TokenService interface {
	GenerateToken(ctx context.Context, apiKey string) (string, error)
}

type tokenService struct {
	generator *auth.ServiceTokenGenerator
}

// NewTokenService creates a token service using the provided JWT generator.
func NewTokenService(generator *auth.ServiceTokenGenerator) TokenService {
	return &tokenService{generator: generator}
}

// GenerateToken creates a signed JWT containing the provided API key.
func (s *tokenService) GenerateToken(_ context.Context, apiKey string) (string, error) {
	if s.generator == nil {
		return "", errors.New("token generator is not configured")
	}
	if apiKey == "" {
		return "", errors.New("api key is required")
	}

	return s.generator.GenerateServiceToken(apiKey)
}
