package leads

import (
	"context"
	"log"
)

type OutletAssignmentService interface {
	RunDailyOutletAssignment(ctx context.Context) error
}

type outletAssignService struct {
}

func NewOutletAssignmentService() OutletAssignmentService {
	return &outletAssignService{}
}

func (s *outletAssignService) RunDailyOutletAssignment(ctx context.Context) error {
	// TODO: actual logic
	log.Println("[Scheduler] Running daily outlet assignment")
	return nil
}
