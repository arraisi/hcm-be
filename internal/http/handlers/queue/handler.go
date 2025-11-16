package queue

import (
	"net/http"
	"strconv"

	"github.com/arraisi/hcm-be/internal/queue/inspector"
	"github.com/arraisi/hcm-be/pkg/response"
)

// Handler handles queue inspection HTTP requests
type Handler struct {
	inspector *inspector.Inspector
}

// NewHandler creates a new queue handler
func NewHandler(insp *inspector.Inspector) *Handler {
	return &Handler{
		inspector: insp,
	}
}

// GetQueueStats returns queue statistics
// GET /queue/stats?queue=default
func (h *Handler) GetQueueStats(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	stats, err := h.inspector.GetQueueStats(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, stats, "Queue statistics retrieved successfully")
}

// ListPendingTasks returns pending tasks in the queue
// GET /queue/pending?queue=default&limit=10
func (h *Handler) ListPendingTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	tasks, err := h.inspector.ListPendingTasks(queueName, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue": queueName,
		"count": len(tasks),
		"tasks": tasks,
	}, "Pending tasks retrieved successfully")
}

// ListActiveTasks returns active (processing) tasks in the queue
// GET /queue/active?queue=default
func (h *Handler) ListActiveTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	tasks, err := h.inspector.ListActiveTasks(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue": queueName,
		"count": len(tasks),
		"tasks": tasks,
	}, "Active tasks retrieved successfully")
}

// ListRetryTasks returns tasks waiting to be retried
// GET /queue/retry?queue=default&limit=10
func (h *Handler) ListRetryTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	tasks, err := h.inspector.ListRetryTasks(queueName, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue": queueName,
		"count": len(tasks),
		"tasks": tasks,
	}, "Retry tasks retrieved successfully")
}

// ListArchivedTasks returns archived (dead) tasks
// GET /queue/archived?queue=default&limit=10
func (h *Handler) ListArchivedTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	tasks, err := h.inspector.ListArchivedTasks(queueName, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue": queueName,
		"count": len(tasks),
		"tasks": tasks,
	}, "Archived tasks retrieved successfully")
}

// DeleteArchivedTasks deletes all archived tasks
// DELETE /queue/archived?queue=default
func (h *Handler) DeleteArchivedTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	count, err := h.inspector.DeleteAllArchivedTasks(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue":   queueName,
		"deleted": count,
	}, "Archived tasks deleted successfully")
}

// RunArchivedTasks runs all archived tasks immediately
// POST /queue/archived/run?queue=default
func (h *Handler) RunArchivedTasks(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	count, err := h.inspector.RunAllArchivedTasks(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue":    queueName,
		"enqueued": count,
	}, "Archived tasks enqueued for processing")
}

// PauseQueue pauses the queue
// POST /queue/pause?queue=default
func (h *Handler) PauseQueue(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	err := h.inspector.PauseQueue(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue":  queueName,
		"paused": true,
	}, "Queue paused successfully")
}

// UnpauseQueue unpauses the queue
// POST /queue/unpause?queue=default
func (h *Handler) UnpauseQueue(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	err := h.inspector.UnpauseQueue(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(w, map[string]interface{}{
		"queue":  queueName,
		"paused": false,
	}, "Queue unpaused successfully")
}

// GetAllStats returns comprehensive queue statistics
// GET /queue/all-stats?queue=default
func (h *Handler) GetAllStats(w http.ResponseWriter, r *http.Request) {
	queueName := r.URL.Query().Get("queue")
	if queueName == "" {
		queueName = "default"
	}

	stats, err := h.inspector.GetQueueStats(queueName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get sample tasks from each state
	pending, _ := h.inspector.ListPendingTasks(queueName, 5)
	active, _ := h.inspector.ListActiveTasks(queueName)
	retry, _ := h.inspector.ListRetryTasks(queueName, 5)
	archived, _ := h.inspector.ListArchivedTasks(queueName, 5)

	result := map[string]interface{}{
		"stats": stats,
		"samples": map[string]interface{}{
			"pending":  pending,
			"active":   active,
			"retry":    retry,
			"archived": archived,
		},
	}

	response.OK(w, result, "Comprehensive queue statistics retrieved successfully")
}
