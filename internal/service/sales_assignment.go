package service

import (
	"context"
	"log"
)

type SalesAssignmentService interface {
	RunDailySalesAssignment(ctx context.Context) error
}

type salesAssignService struct {
}

func NewSalesAssignmentService() SalesAssignmentService {
	return &salesAssignService{}
}

func (s *salesAssignService) RunDailySalesAssignment(ctx context.Context) error {
	// TODO: actual logic
	log.Println("[Scheduler] Running daily sales assignment")
	return nil
}
