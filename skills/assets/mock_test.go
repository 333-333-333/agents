// internal/booking/application/mock_test.go
package application_test

import (
	"api/booking/internal/booking/domain"
	"context"
)

// MockRepository implements domain.Repository for testing.
type MockRepository struct {
	SaveFunc     func(ctx context.Context, booking *domain.Booking) error
	FindByIDFunc func(ctx context.Context, id domain.BookingID) (*domain.Booking, error)
	// ... other methods
}

func (m *MockRepository) Save(ctx context.Context, booking *domain.Booking) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, booking)
	}
	return nil
}

func (m *MockRepository) FindByID(ctx context.Context, id domain.BookingID) (*domain.Booking, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(ctx, id)
	}
	return nil, domain.ErrBookingNotFound
}

// MockPublisher implements domain.EventPublisher for testing.
type MockPublisher struct {
	PublishFunc func(ctx context.Context, topic string, event domain.Event) error
	Published   []domain.Event
}

func (m *MockPublisher) Publish(ctx context.Context, topic string, event domain.Event) error {
	m.Published = append(m.Published, event)
	if m.PublishFunc != nil {
		return m.PublishFunc(ctx, topic, event)
	}
	return nil
}
