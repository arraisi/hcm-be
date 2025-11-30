package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/go-co-op/gocron"
)

type RoAutomationService interface {
	RunRoAutomation(ctx context.Context) error
	RoAutomationOutletAssignment(ctx context.Context) error
	RoAutomationSalesAssignment(ctx context.Context) error
}

type Scheduler struct {
	scheduler       *gocron.Scheduler
	cfg             config.SchedulerConfig
	roAutomationSvc RoAutomationService
}

func New(cfg config.SchedulerConfig, roAutomationSvc RoAutomationService) (*Scheduler, error) {
	log.Printf("[Scheduler] Config: %+v", cfg)
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	s := gocron.NewScheduler(loc)

	return &Scheduler{
		scheduler:       s,
		cfg:             cfg,
		roAutomationSvc: roAutomationSvc,
	}, nil
}

func (s *Scheduler) Start() error {
	// Register jobs
	if _, err := s.scheduler.Cron(s.cfg.CustomerSegCron).Do(func() {
		ctx := context.Background()
		if err := s.roAutomationSvc.RunRoAutomation(ctx); err != nil {
			log.Printf("[Scheduler] Error running monthly segmentation: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.OutletAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.roAutomationSvc.RoAutomationOutletAssignment(ctx); err != nil {
			log.Printf("[Scheduler] Error running daily outlet assignment: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.SalesAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.roAutomationSvc.RoAutomationSalesAssignment(ctx); err != nil {
			log.Printf("[Scheduler] Error running daily sales assignment: %v", err)
		}
	}); err != nil {
		return err
	}

	s.scheduler.StartAsync()
	log.Println("[Scheduler] Started")
	return nil
}

func (s *Scheduler) Stop(ctx context.Context) {
	s.scheduler.Stop()
	log.Println("[Scheduler] Stopped")
}
