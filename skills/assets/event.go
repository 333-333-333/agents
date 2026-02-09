// internal/booking/domain/event.go
package domain

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Data      any       `json:"data"`
}

type BookingCreatedEvent struct {
	BookingID   string `json:"booking_id"`
	OwnerID     string `json:"owner_id"`
	CaregiverID string `json:"caregiver_id"`
	ServiceType string `json:"service_type"`
	TotalCLP    int64  `json:"total_clp"`
}

type BookingCancelledEvent struct {
	BookingID string `json:"booking_id"`
	Reason    string `json:"reason"`
}
