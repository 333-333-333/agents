// internal/booking/infrastructure/messaging/memory_publisher.go
package messaging

import (
	"api/booking/internal/booking/domain"
	"context"
	"sync"
)

type InMemoryPublisher struct {
	mu     sync.Mutex
	events map[string][]domain.Event
}

func NewInMemoryPublisher() *InMemoryPublisher {
	return &InMemoryPublisher{events: make(map[string][]domain.Event)}
}

func (p *InMemoryPublisher) Publish(ctx context.Context, topic string, event domain.Event) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.events[topic] = append(p.events[topic], event)
	return nil
}

func (p *InMemoryPublisher) Events(topic string) []domain.Event {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.events[topic]
}
