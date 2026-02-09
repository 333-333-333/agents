// GOOD: structured, contextual, with trace correlation
slog.InfoContext(ctx, "booking created",
    "booking_id", booking.ID,
    "owner_id", booking.OwnerID,
    "service_type", booking.ServiceType,
)

slog.ErrorContext(ctx, "failed to save booking",
    "error", err,
    "booking_id", booking.ID,
)

// BAD: unstructured, no context
log.Printf("booking created: %s", booking.ID)
