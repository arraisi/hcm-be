package inspector

import (
	"fmt"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/hibiken/asynq"
)

// Inspector provides methods to inspect and monitor Asynq queues
type Inspector struct {
	inspector *asynq.Inspector
}

// QueueStats contains statistics for a queue
type QueueStats struct {
	Queue       string `json:"queue"`
	Pending     int    `json:"pending"`
	Active      int    `json:"active"`
	Scheduled   int    `json:"scheduled"`
	Retry       int    `json:"retry"`
	Archived    int    `json:"archived"`
	Completed   int    `json:"completed"`
	Processed   int    `json:"processed"`
	Failed      int    `json:"failed"`
	Paused      bool   `json:"paused"`
	Size        int    `json:"size"`
	Latency     int64  `json:"latency_ms"`
	MemoryUsage int64  `json:"memory_usage_bytes"`
}

// TaskInfo contains information about a task
type TaskInfo struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Payload       string `json:"payload"`
	Queue         string `json:"queue"`
	MaxRetry      int    `json:"max_retry"`
	Retried       int    `json:"retried"`
	LastErr       string `json:"last_error,omitempty"`
	LastFailedAt  string `json:"last_failed_at,omitempty"`
	NextProcessAt string `json:"next_process_at,omitempty"`
}

// New creates a new Inspector instance
func New(cfg config.AsynqConfig) *Inspector {
	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr:     cfg.RedisAddr,
		DB:       cfg.RedisDB,
		Password: cfg.RedisPassword,
	})

	return &Inspector{
		inspector: inspector,
	}
}

// GetQueueStats returns statistics for the specified queue
func (i *Inspector) GetQueueStats(queueName string) (*QueueStats, error) {
	info, err := i.inspector.GetQueueInfo(queueName)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue info: %w", err)
	}

	return &QueueStats{
		Queue:       queueName,
		Pending:     info.Pending,
		Active:      info.Active,
		Scheduled:   info.Scheduled,
		Retry:       info.Retry,
		Archived:    info.Archived,
		Completed:   info.Completed,
		Processed:   info.Processed,
		Failed:      info.Failed,
		Paused:      info.Paused,
		Size:        info.Size,
		Latency:     info.Latency.Milliseconds(),
		MemoryUsage: info.MemoryUsage,
	}, nil
}

// ListPendingTasks returns all pending tasks in the queue
func (i *Inspector) ListPendingTasks(queueName string, limit int) ([]*TaskInfo, error) {
	tasks, err := i.inspector.ListPendingTasks(queueName, asynq.PageSize(limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list pending tasks: %w", err)
	}

	return i.convertToTaskInfo(tasks), nil
}

// ListActiveTasks returns all active (currently processing) tasks in the queue
func (i *Inspector) ListActiveTasks(queueName string) ([]*TaskInfo, error) {
	tasks, err := i.inspector.ListActiveTasks(queueName)
	if err != nil {
		return nil, fmt.Errorf("failed to list active tasks: %w", err)
	}

	return i.convertToTaskInfo(tasks), nil
}

// ListScheduledTasks returns all scheduled tasks in the queue
func (i *Inspector) ListScheduledTasks(queueName string, limit int) ([]*TaskInfo, error) {
	tasks, err := i.inspector.ListScheduledTasks(queueName, asynq.PageSize(limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list scheduled tasks: %w", err)
	}

	return i.convertToTaskInfo(tasks), nil
}

// ListRetryTasks returns all tasks that are waiting to be retried
func (i *Inspector) ListRetryTasks(queueName string, limit int) ([]*TaskInfo, error) {
	tasks, err := i.inspector.ListRetryTasks(queueName, asynq.PageSize(limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list retry tasks: %w", err)
	}

	return i.convertToTaskInfo(tasks), nil
}

// ListArchivedTasks returns all archived (dead) tasks in the queue
func (i *Inspector) ListArchivedTasks(queueName string, limit int) ([]*TaskInfo, error) {
	tasks, err := i.inspector.ListArchivedTasks(queueName, asynq.PageSize(limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list archived tasks: %w", err)
	}

	return i.convertToTaskInfo(tasks), nil
}

// DeleteTask deletes a task by ID
func (i *Inspector) DeleteTask(queueName, taskID string) error {
	return i.inspector.DeleteTask(queueName, taskID)
}

// RunTask runs a task immediately (useful for scheduled tasks)
func (i *Inspector) RunTask(queueName, taskID string) error {
	return i.inspector.RunTask(queueName, taskID)
}

// ArchiveTask archives a task (moves it to archived state)
func (i *Inspector) ArchiveTask(queueName, taskID string) error {
	return i.inspector.ArchiveTask(queueName, taskID)
}

// DeleteAllArchivedTasks deletes all archived tasks from the queue
func (i *Inspector) DeleteAllArchivedTasks(queueName string) (int, error) {
	return i.inspector.DeleteAllArchivedTasks(queueName)
}

// RunAllScheduledTasks runs all scheduled tasks immediately
func (i *Inspector) RunAllScheduledTasks(queueName string) (int, error) {
	return i.inspector.RunAllScheduledTasks(queueName)
}

// RunAllRetryTasks runs all retry tasks immediately
func (i *Inspector) RunAllRetryTasks(queueName string) (int, error) {
	return i.inspector.RunAllRetryTasks(queueName)
}

// RunAllArchivedTasks runs all archived tasks immediately
func (i *Inspector) RunAllArchivedTasks(queueName string) (int, error) {
	return i.inspector.RunAllArchivedTasks(queueName)
}

// PauseQueue pauses the queue (stops processing new tasks)
func (i *Inspector) PauseQueue(queueName string) error {
	return i.inspector.PauseQueue(queueName)
}

// UnpauseQueue unpauses the queue (resumes processing)
func (i *Inspector) UnpauseQueue(queueName string) error {
	return i.inspector.UnpauseQueue(queueName)
}

// Close closes the inspector connection
func (i *Inspector) Close() error {
	return i.inspector.Close()
}

// Helper function to convert asynq.TaskInfo to our TaskInfo
func (i *Inspector) convertToTaskInfo(tasks []*asynq.TaskInfo) []*TaskInfo {
	result := make([]*TaskInfo, len(tasks))
	for idx, task := range tasks {
		info := &TaskInfo{
			ID:       task.ID,
			Type:     task.Type,
			Payload:  string(task.Payload),
			Queue:    task.Queue,
			MaxRetry: task.MaxRetry,
			Retried:  task.Retried,
		}

		if task.LastErr != "" {
			info.LastErr = task.LastErr
		}

		if !task.LastFailedAt.IsZero() {
			info.LastFailedAt = task.LastFailedAt.Format("2006-01-02 15:04:05")
		}

		if !task.NextProcessAt.IsZero() {
			info.NextProcessAt = task.NextProcessAt.Format("2006-01-02 15:04:05")
		}

		result[idx] = info
	}
	return result
}
