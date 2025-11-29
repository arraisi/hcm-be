package auth

import (
	"errors"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

const defaultIssuer = "hcm-be"

// ServiceTokenGenerator handles creation of signed service tokens.
type ServiceTokenGenerator struct {
	secret []byte
	issuer string
}

// New constructs a generator using the provided JWT config.
func New(cfg config.JWTConfig) (*ServiceTokenGenerator, error) {
	if cfg.Secret == "" {
		return nil, errors.New("jwt secret is required")
	}

	issuer := cfg.Issuer
	if issuer == "" {
		issuer = defaultIssuer
	}

	return &ServiceTokenGenerator{
		secret: []byte(cfg.Secret),
		issuer: issuer,
	}, nil
}

// GenerateServiceToken creates a signed JWT with the provided API key.
func (g *ServiceTokenGenerator) GenerateServiceToken(apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("api key is required")
	}

	claims := jwt.MapClaims{
		"api_key": apiKey,
		"iss":     g.issuer,
		"typ":     "service-token",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(g.secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
