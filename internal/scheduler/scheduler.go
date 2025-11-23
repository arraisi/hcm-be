package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/go-co-op/gocron"
)

type EngineService interface {
	RunMonthlySegmentation(ctx context.Context) error
	RunDailyOutletAssignment(ctx context.Context) error
	RunDailySalesAssignment(ctx context.Context) error
}

type Scheduler struct {
	scheduler *gocron.Scheduler
	cfg       config.SchedulerConfig

	engineSvc EngineService
}

func New(cfg config.SchedulerConfig,
	engineSvc EngineService,
) (*Scheduler, error) {
	log.Printf("[Scheduler] Config: %+v", cfg)
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	s := gocron.NewScheduler(loc)

	return &Scheduler{
		scheduler: s,
		cfg:       cfg,
		engineSvc: engineSvc,
	}, nil
}

func (s *Scheduler) Start() error {
	// Register jobs
	if _, err := s.scheduler.Cron(s.cfg.CustomerSegCron).Do(func() {
		ctx := context.Background()
		if err := s.engineSvc.RunMonthlySegmentation(ctx); err != nil {
			log.Printf("[Scheduler] Error running monthly segmentation: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.OutletAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.engineSvc.RunDailyOutletAssignment(ctx); err != nil {
			log.Printf("[Scheduler] Error running daily outlet assignment: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.SalesAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.engineSvc.RunDailySalesAssignment(ctx); err != nil {
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
