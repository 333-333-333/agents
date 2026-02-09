// internal/booking/domain/port.go
package domain

import "context"

type EventPublisher interface {
	Publish(ctx context.Context, topic string, event Event) error
}

type EventSubscriber interface {
	Subscribe(ctx context.Context, topic string, handler EventHandler) error
	Close() error
}

type EventHandler func(ctx context.Context, event Event) error
