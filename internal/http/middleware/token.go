package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/golang-jwt/jwt/v5"

	"github.com/arraisi/hcm-be/pkg/response"
)

type tokenClaimsKey struct{}

// TokenValidator validates incoming JWT service tokens.
type TokenValidator struct {
	secret []byte
	issuer string
}

// NewTokenValidator constructs a token validator using JWT config.
func NewTokenValidator(cfg config.JWTConfig) (*TokenValidator, error) {
	if cfg.Secret == "" {
		return nil, errors.New("jwt secret is required")
	}

	issuer := cfg.Issuer
	if issuer == "" {
		issuer = "hcm-be"
	}

	return &TokenValidator{
		secret: []byte(cfg.Secret),
		issuer: issuer,
	}, nil
}

// Middleware validates Authorization Bearer token and adds claims to context.
func (v *TokenValidator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" {
			response.Unauthorized(w, "authorization header is required")
			return
		}

		if !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			response.Unauthorized(w, "authorization header must use Bearer schema")
			return
		}

		rawToken := strings.TrimSpace(authHeader[len("bearer "):])
		if rawToken == "" {
			response.Unauthorized(w, "authorization token is empty")
			return
		}

		claims, err := v.validate(rawToken)
		if err != nil {
			log.Printf("token validation failed: %v", err)
			response.Unauthorized(w, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), tokenClaimsKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetTokenClaims retrieves JWT claims from context.
func GetTokenClaims(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(tokenClaimsKey{}).(jwt.MapClaims)
	return claims, ok
}

func (v *TokenValidator) validate(tokenStr string) (jwt.MapClaims, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	claims := jwt.MapClaims{}

	token, err := parser.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return v.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	if iss, ok := claims["iss"].(string); !ok || iss != v.issuer {
		return nil, errors.New("invalid issuer")
	}

	if typ, ok := claims["typ"].(string); !ok || typ != "service-token" {
		return nil, errors.New("invalid token type")
	}

	if apiKey, ok := claims["api_key"].(string); !ok || strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("api_key claim missing")
	}

	return claims, nil
}
