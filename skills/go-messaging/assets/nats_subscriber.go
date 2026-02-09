// internal/booking/infrastructure/messaging/nats_subscriber.go
package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"api/booking/internal/booking/domain"
	"github.com/nats-io/nats.go"
)

type NATSSubscriber struct {
	conn *nats.Conn
	subs []*nats.Subscription
}

func NewNATSSubscriber(url string) (*NATSSubscriber, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}
	return &NATSSubscriber{conn: conn}, nil
}

func (s *NATSSubscriber) Subscribe(ctx context.Context, topic string, handler domain.EventHandler) error {
	sub, err := s.conn.Subscribe(topic, func(msg *nats.Msg) {
		var event domain.Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			slog.Error("failed to unmarshal event", "error", err, "topic", topic)
			return
		}
		if err := handler(ctx, event); err != nil {
			slog.Error("failed to handle event", "error", err, "topic", topic, "event_id", event.ID)
			// In production: implement retry / dead letter queue
		}
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to %s: %w", topic, err)
	}
	s.subs = append(s.subs, sub)
	return nil
}

func (s *NATSSubscriber) Close() error {
	for _, sub := range s.subs {
		_ = sub.Unsubscribe()
	}
	s.conn.Close()
	return nil
}
