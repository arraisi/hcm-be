package engine

//go:generate mockgen -package=engine -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
)

type Service interface {
	RunMonthlySegmentation(ctx context.Context, request engine.RunMonthlySegmentationRequest) error
}

// Handler handles HTTP requests for engine operations
type Handler struct {
	cfg *config.Config
	svc Service
}

// New creates a new Handler instance
func New(cfg *config.Config, svc Service) Handler {
	return Handler{
		cfg: cfg,
		svc: svc,
	}
}
