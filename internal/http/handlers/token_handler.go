package handlers

import (
	"log"
	"net/http"
	"strings"

	authService "github.com/arraisi/hcm-be/internal/service/auth"
	"github.com/arraisi/hcm-be/pkg/response"
)

type TokenHandler struct {
	tokenService authService.TokenService
}

func NewTokenHandler(tokenService authService.TokenService) TokenHandler {
	return TokenHandler{tokenService: tokenService}
}

// Generate creates a service JWT using the provided API key header.
func (h TokenHandler) Generate(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.TrimSpace(r.Header.Get("X-Api-Key"))
	if apiKey == "" {
		response.BadRequest(w, "x-api-key header is required")
		return
	}

	token, err := h.tokenService.GenerateToken(r.Context(), apiKey)
	if err != nil {
		log.Printf("token generation failed: %v", err)
		response.InternalServerError(w, "failed to generate token")
		return
	}

	resp := response.Response{
		Data: map[string]string{
			"token": token,
		},
		Message: "ok",
	}

	response.JSON(w, http.StatusOK, resp)
}
