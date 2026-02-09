// internal/booking/infrastructure/messaging/nats_publisher.go
package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"api/booking/internal/booking/domain"
	"github.com/nats-io/nats.go"
)

type NATSPublisher struct {
	conn *nats.Conn
}

func NewNATSPublisher(url string) (*NATSPublisher, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}
	return &NATSPublisher{conn: conn}, nil
}

func (p *NATSPublisher) Publish(ctx context.Context, topic string, event domain.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	return p.conn.Publish(topic, data)
}

func (p *NATSPublisher) Close() error {
	p.conn.Close()
	return nil
}
