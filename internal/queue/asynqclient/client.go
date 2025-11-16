package asynqclient

import (
	"context"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// Client defines the interface for Asynq client operations
type Client interface {
	EnqueueDIDXConfirm(ctx context.Context, payload interface{}) error
	Close() error
}

type client struct {
	asynqClient *asynq.Client
	cfg         config.AsynqConfig
}

// New creates a new Asynq client instance
func New(cfg config.AsynqConfig) Client {
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
	didxPayload, ok := payload.(queue.DIDXConfirmPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DIDXConfirmPayload")
	}

	task, err := queue.NewDIDXConfirmTask(didxPayload)
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
