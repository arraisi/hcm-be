package mq

import (
	"encoding/json"
	"log"
)

// Publisher defines the interface for message queue publishing
type Publisher interface {
	// Publish publishes a message to the specified topic/queue with a key
	Publish(topic string, key string, payload interface{}) error
	// Close closes the publisher connection
	Close() error
}

// TestDriveEvent represents the normalized event payload for MQ publishing
type TestDriveEvent struct {
	EventID   string      `json:"eventId"`
	EventType string      `json:"eventType"`
	Timestamp int64       `json:"timestamp"`
	Source    string      `json:"source"`
	Data      interface{} `json:"data"`
}

// InMemoryPublisher implements Publisher interface for testing/development
// In production, this should be replaced with actual MQ implementation (RabbitMQ, Kafka, etc.)
type InMemoryPublisher struct {
	messages []PublishedMessage
}

// PublishedMessage represents a message that was published
type PublishedMessage struct {
	Topic   string      `json:"topic"`
	Key     string      `json:"key"`
	Payload interface{} `json:"payload"`
}

// NewInMemoryPublisher creates a new in-memory publisher for testing
func NewInMemoryPublisher() *InMemoryPublisher {
	return &InMemoryPublisher{
		messages: make([]PublishedMessage, 0),
	}
}

// Publish publishes a message to the in-memory store
func (p *InMemoryPublisher) Publish(topic string, key string, payload interface{}) error {
	message := PublishedMessage{
		Topic:   topic,
		Key:     key,
		Payload: payload,
	}

	p.messages = append(p.messages, message)

	// Log for debugging
	payloadJSON, _ := json.Marshal(payload)
	log.Printf("Published message to topic '%s' with key '%s': %s", topic, key, string(payloadJSON))

	return nil
}

// Close closes the publisher (no-op for in-memory implementation)
func (p *InMemoryPublisher) Close() error {
	return nil
}

// GetMessages returns all published messages (for testing)
func (p *InMemoryPublisher) GetMessages() []PublishedMessage {
	return p.messages
}

// Clear clears all published messages (for testing)
func (p *InMemoryPublisher) Clear() {
	p.messages = make([]PublishedMessage, 0)
}
