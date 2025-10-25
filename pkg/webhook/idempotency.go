package webhook

import (
	"sync"
	"time"
)

// IdempotencyStore defines the interface for storing and checking idempotency keys
type IdempotencyStore interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
	// Cleanup removes expired entries (for implementations that support TTL)
	Cleanup()
}

// InMemoryIdempotencyStore implements IdempotencyStore using in-memory storage
type InMemoryIdempotencyStore struct {
	mu      sync.RWMutex
	entries map[string]time.Time
	ttl     time.Duration
}

// NewInMemoryIdempotencyStore creates a new in-memory idempotency store
func NewInMemoryIdempotencyStore(ttl time.Duration) *InMemoryIdempotencyStore {
	store := &InMemoryIdempotencyStore{
		entries: make(map[string]time.Time),
		ttl:     ttl,
	}

	// Start cleanup goroutine
	go store.cleanupRoutine()

	return store
}

// Exists checks if the given event ID already exists
func (s *InMemoryIdempotencyStore) Exists(eventID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	timestamp, exists := s.entries[eventID]
	if !exists {
		return false
	}

	// Check if entry has expired
	if time.Since(timestamp) > s.ttl {
		// Entry has expired, remove it and return false
		delete(s.entries, eventID)
		return false
	}

	return true
}

// Store stores the event ID to prevent duplicate processing
func (s *InMemoryIdempotencyStore) Store(eventID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entries[eventID] = time.Now()
	return nil
}

// Cleanup removes expired entries
func (s *InMemoryIdempotencyStore) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for eventID, timestamp := range s.entries {
		if now.Sub(timestamp) > s.ttl {
			delete(s.entries, eventID)
		}
	}
}

// cleanupRoutine runs cleanup every minute
func (s *InMemoryIdempotencyStore) cleanupRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		s.Cleanup()
	}
}
