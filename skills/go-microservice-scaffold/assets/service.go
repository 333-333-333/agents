package application

import (
	"api/booking/internal/booking/domain"
	"context"
)

type BookingService struct {
	repo     domain.Repository
	events   domain.EventPublisher
	payments domain.PaymentGateway
}

func NewBookingService(
	repo domain.Repository,
	events domain.EventPublisher,
	payments domain.PaymentGateway,
) *BookingService {
	return &BookingService{repo: repo, events: events, payments: payments}
}

func (s *BookingService) CreateBooking(ctx context.Context, input CreateBookingInput) (*CreateBookingOutput, error) {
	// 1. Domain logic — create entity
	booking, err := domain.NewBooking(input.OwnerID, input.CaregiverID, input.ServiceType, input.Schedule)
	if err != nil {
		return nil, err
	}

	// 2. Persist
	if err := s.repo.Save(ctx, booking); err != nil {
		return nil, err
	}

	// 3. Side effects
	_ = s.events.PublishBookingCreated(ctx, domain.BookingCreatedEvent{BookingID: booking.ID})

	return &CreateBookingOutput{ID: booking.ID.String()}, nil
}
