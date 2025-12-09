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
	cfg         *config.Config
}

// New creates a new Asynq client instance
func New(cfg *config.Config) *client {
	c := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     cfg.Asynq.RedisAddr,
		DB:       cfg.Asynq.RedisDB,
		Password: cfg.Asynq.RedisPassword,
	})
	return &client{
		asynqClient: c,
		cfg:         cfg,
	}
}

// EnqueueDIDXServiceBookingConfirm enqueues a DIDX confirm task with custom retry configuration
func (c *client) EnqueueDIDXServiceBookingConfirm(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXServiceBookingConfirmPayload
	body, ok := payload.(queue.DIDXServiceBookingConfirmPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DIDXServiceBookingConfirmPayload")
	}

	task, err := queue.NewDIDXServiceBookingConfirmTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDIDXServiceBookingConfirmTask taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDIDXTestDriveConfirm enqueues a DIDX test drive confirm task with custom retry configuration
func (c *client) EnqueueDIDXTestDriveConfirm(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXTestDriveConfirmPayload
	body, ok := payload.(queue.DIDXTestDriveConfirmPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DIDXTestDriveConfirmPayload")
	}

	task, err := queue.NewDIDXTestDriveConfirmTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDIDXTestDriveConfirmTask taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDIDXAppraisalConfirm enqueues a DIDX appraisal confirm task with custom retry configuration
func (c *client) EnqueueDIDXAppraisalConfirm(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXAppraisalConfirmPayload
	body, ok := payload.(queue.DIDXAppraisalConfirmPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DIDXAppraisalConfirmPayload")
	}

	task, err := queue.NewDIDXAppraisalConfirmTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDIDXAppraisalConfirmTask taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDMSTestDriveRequest enqueues a DMS test drive request task with custom retry configuration
func (c *client) EnqueueDMSTestDriveRequest(ctx context.Context, payload interface{}) error {
	// Type assert to DIDXServiceBookingConfirmPayload
	body, ok := payload.(queue.DMSTestDriveRequestPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DMSTestDriveRequestPayload")
	}

	task, err := queue.NewDMSTestDriveRequestTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDMSTestDriveRequest taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDMSCreateOneAccess enqueues a DMS create one access task with custom retry configuration
func (c *client) EnqueueDMSCreateOneAccess(ctx context.Context, payload interface{}) error {
	// Type asserts to DMSCreateOneAccessPayload
	body, ok := payload.(queue.DMSCreateOneAccessPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.EnqueueDMSCreateOneAccess")
	}

	task, err := queue.NewDMSCreateOneAccessTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDMSCreateOneAccess taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDMSCreateToyotaID enqueues a DMS create toyota id task with custom retry configuration
func (c *client) EnqueueDMSCreateToyotaID(ctx context.Context, payload interface{}) error {
	// Type asserts to DMSCreateToyotaIDPayload
	body, ok := payload.(queue.DMSCreateToyotaIDPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.EnqueueDMSCreateToyotaID")
	}

	task, err := queue.NewDMSCreateToyotaIDTask(body)
	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDMSCreateToyotaID taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDMSAppraisalBookingRequest enqueues a DMS Appraisal Booking request task with custom retry configuration
func (c *client) EnqueueDMSAppraisalBookingRequest(ctx context.Context, payload interface{}) error {
	// Type asserts to DMSAppraisalBookingRequestPayload
	body, ok := payload.(queue.DMSAppraisalBookingRequestPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.EnqueueDMSAppraisalBookingRequest")
	}

	task, err := queue.NewDMSAppraisalBookingRequestTask(body)

	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDMSAppraisalBookingRequest taskInfo: %+v\n", taskInfo)

	return err
}

// EnqueueDMSCreateGetOffer enqueues a DMS create get offer task with custom retry configuration
func (c *client) EnqueueDMSCreateGetOffer(ctx context.Context, payload interface{}) error {
	// Type asserts to DMSCreateGetOfferPayload
	body, ok := payload.(queue.DMSCreateGetOfferPayload)
	if !ok {
		return fmt.Errorf("invalid payload type: expected queue.DMSCreateGetOfferPayload")
	}

	task, err := queue.NewDMSCreateGetOfferTask(body)

	if err != nil {
		return err
	}

	// Retry 3 times with custom backoff: 1m, 5m, 10m
	taskInfo, err := c.asynqClient.EnqueueContext(
		ctx,
		task,
		asynq.Queue(c.cfg.Asynq.Queue),
		asynq.MaxRetry(3),
		asynq.Retention(24*time.Hour), // Keep task info for 24 hours after completion
		asynq.Timeout(30*time.Second), // Task timeout
		asynq.Unique(5*time.Minute),   // Prevent duplicate tasks within 5 minutes
	)

	fmt.Printf("EnqueueDMSCreateGetOffer taskInfo: %+v\n", taskInfo)

	return err
}

// Close closes the Asynq client connection
func (c *client) Close() error {
	return c.asynqClient.Close()
}
