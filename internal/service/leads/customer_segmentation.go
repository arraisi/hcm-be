package leads

import (
	"context"
	"log"
)

type CustomerSegmentationService interface {
	RunMonthlySegmentation(ctx context.Context) error
}

type customerSegService struct {
}

func NewCustomerSegmentationService() CustomerSegmentationService {
	return &customerSegService{}
}

func (s *customerSegService) RunMonthlySegmentation(ctx context.Context) error {
	// TODO: actual logic
	log.Println("[Scheduler] Running monthly customer segmentation")
	return nil
}
