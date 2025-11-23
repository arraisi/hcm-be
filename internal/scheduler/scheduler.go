package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/service"
	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
	cfg       config.SchedulerConfig

	customerSegSvc  service.CustomerSegmentationService
	outletAssignSvc service.OutletAssignmentService
	salesAssignSvc  service.SalesAssignmentService
}

func New(cfg config.SchedulerConfig,
	customerSegSvc service.CustomerSegmentationService,
	outletAssignSvc service.OutletAssignmentService,
	salesAssignSvc service.SalesAssignmentService,
) (*Scheduler, error) {
	log.Printf("[Scheduler] Config: %+v", cfg)
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	s := gocron.NewScheduler(loc)

	return &Scheduler{
		scheduler:       s,
		cfg:             cfg,
		customerSegSvc:  customerSegSvc,
		outletAssignSvc: outletAssignSvc,
		salesAssignSvc:  salesAssignSvc,
	}, nil
}

func (s *Scheduler) Start() error {
	// Register jobs
	if _, err := s.scheduler.Cron(s.cfg.CustomerSegCron).Do(func() {
		ctx := context.Background()
		if err := s.customerSegSvc.RunMonthlySegmentation(ctx); err != nil {
			log.Printf("[Scheduler] Error running monthly segmentation: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.OutletAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.outletAssignSvc.RunDailyOutletAssignment(ctx); err != nil {
			log.Printf("[Scheduler] Error running daily outlet assignment: %v", err)
		}
	}); err != nil {
		return err
	}

	if _, err := s.scheduler.Cron(s.cfg.SalesAssignCron).Do(func() {
		ctx := context.Background()
		if err := s.salesAssignSvc.RunDailySalesAssignment(ctx); err != nil {
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
