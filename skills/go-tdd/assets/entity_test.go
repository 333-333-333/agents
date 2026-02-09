// internal/booking/domain/entity_test.go
package domain_test

import (
	"testing"
	"time"

	"api/booking/internal/booking/domain"
)

func TestNewBooking(t *testing.T) {
	tests := []struct {
		name        string
		ownerID     string
		caregiverID string
		serviceType domain.ServiceType
		startAt     time.Time
		endAt       time.Time
		wantErr     bool
	}{
		{
			name:        "valid booking",
			ownerID:     "owner-1",
			caregiverID: "caregiver-1",
			serviceType: domain.Walk,
			startAt:     time.Now().Add(24 * time.Hour),
			endAt:       time.Now().Add(25 * time.Hour),
			wantErr:     false,
		},
		{
			name:        "end before start",
			ownerID:     "owner-1",
			caregiverID: "caregiver-1",
			serviceType: domain.Walk,
			startAt:     time.Now().Add(25 * time.Hour),
			endAt:       time.Now().Add(24 * time.Hour),
			wantErr:     true,
		},
		{
			name:        "empty owner",
			ownerID:     "",
			caregiverID: "caregiver-1",
			serviceType: domain.Walk,
			startAt:     time.Now().Add(24 * time.Hour),
			endAt:       time.Now().Add(25 * time.Hour),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewBooking(tt.ownerID, tt.caregiverID, tt.serviceType,
				domain.Schedule{StartAt: tt.startAt, EndAt: tt.endAt})
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
