func (s *BookingService) CreateBooking(ctx context.Context, input CreateBookingInput) (*CreateBookingOutput, error) {
    booking, err := domain.NewBooking(/* ... */)
    if err != nil {
        return nil, err
    }

    if err := s.repo.Save(ctx, booking); err != nil {
        return nil, err
    }

    // Publish event — fire and forget or with error handling based on criticality
    event := domain.Event{
        ID:        uuid.NewString(),
        Type:      "booking.booking.created",
        Timestamp: time.Now(),
        Data: domain.BookingCreatedEvent{
            BookingID:   booking.ID.String(),
            OwnerID:     booking.OwnerID,
            CaregiverID: booking.CaregiverID,
            ServiceType: string(booking.ServiceType),
            TotalCLP:    booking.Total.Amount,
        },
    }

    if err := s.events.Publish(ctx, "booking.booking.created", event); err != nil {
        slog.Error("failed to publish booking created event", "error", err, "booking_id", booking.ID)
        // Don't fail the operation — event publishing is best-effort here
    }

    return &CreateBookingOutput{ID: booking.ID.String()}, nil
}
