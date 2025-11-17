package asynqclient

import (
	"context"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

type client struct {
	asynqClient *asynq.Client
	cfg         config.AsynqConfig
}

// New creates a new Asynq client instance
func New(cfg config.AsynqConfig) *client {
	c := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     cfg.RedisAddr,
		DB:       cfg.RedisDB,
		Password: cfg.RedisPassword,
	})
	return &client{
		asynqClient: c,
		cfg:         cfg,
	}
}

// EnqueueDIDXConfirm enqueues a DIDX confirm task with custom retry configuration
func (c *client) EnqueueDIDXConfirm(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXConfirmPayload
	body, ok := payload.(queue.DIDXConfirmPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DIDXConfirmPayload")
	}

	task, err := queue.NewDIDXConfirmTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	_, err = c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)
	return err
}

// EnqueueDMSTestDriveRequest enqueues a DMS test drive request task with custom retry configuration
func (c *client) EnqueueDMSTestDriveRequest(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXConfirmPayload
	body, ok := payload.(queue.DMSTestDriveRequestPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DMSTestDriveRequestPayload")
	}

	task, err := queue.NewDMSTestDriveRequestTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	_, err = c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)
	return err
}

// Close closes the Asynq client connection
func (c *client) Close() error {
	return c.asynqClient.Close()
}
