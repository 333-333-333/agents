package booking

import "context"

// Repository defines persistence operations for Booking aggregate.
type Repository interface {
	Save(ctx context.Context, booking *Booking) error
	FindByID(ctx context.Context, id BookingID) (*Booking, error)
	FindByOwner(ctx context.Context, ownerID string) ([]*Booking, error)
}

// EventPublisher defines async event emission for this domain.
type EventPublisher interface {
	PublishBookingCreated(ctx context.Context, event BookingCreatedEvent) error
	PublishBookingCancelled(ctx context.Context, event BookingCancelledEvent) error
}

// PaymentGateway defines interaction with the payment bounded context.
type PaymentGateway interface {
	ChargeOwner(ctx context.Context, ownerID string, amount Money) (TransactionID, error)
	RefundTransaction(ctx context.Context, txID TransactionID) error
}
