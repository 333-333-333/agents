// internal/booking/application/service_test.go
package application_test

import (
	"context"
	"testing"
	"time"

	"api/booking/internal/booking/application"
	"api/booking/internal/booking/domain"
)

func TestBookingService_CreateBooking(t *testing.T) {
	tests := []struct {
		name    string
		input   application.CreateBookingInput
		setup   func(repo *MockRepository, pub *MockPublisher)
		wantErr bool
		errType error
	}{
		{
			name: "successful booking creation",
			input: application.CreateBookingInput{
				OwnerID:     "owner-1",
				CaregiverID: "caregiver-1",
				ServiceType: "walk",
				StartAt:     time.Now().Add(24 * time.Hour),
				EndAt:       time.Now().Add(25 * time.Hour),
				PetIDs:      []string{"pet-1"},
			},
			setup: func(repo *MockRepository, pub *MockPublisher) {
				repo.SaveFunc = func(ctx context.Context, b *domain.Booking) error {
					return nil
				}
				pub.PublishFunc = func(ctx context.Context, topic string, event domain.Event) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "fails when start time is in the past",
			input: application.CreateBookingInput{
				OwnerID:     "owner-1",
				CaregiverID: "caregiver-1",
				ServiceType: "walk",
				StartAt:     time.Now().Add(-1 * time.Hour),
				EndAt:       time.Now(),
				PetIDs:      []string{"pet-1"},
			},
			setup:   func(repo *MockRepository, pub *MockPublisher) {},
			wantErr: true,
		},
		{
			name: "fails when no pets provided",
			input: application.CreateBookingInput{
				OwnerID:     "owner-1",
				CaregiverID: "caregiver-1",
				ServiceType: "walk",
				StartAt:     time.Now().Add(24 * time.Hour),
				EndAt:       time.Now().Add(25 * time.Hour),
				PetIDs:      []string{},
			},
			setup:   func(repo *MockRepository, pub *MockPublisher) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepository{}
			pub := &MockPublisher{}
			tt.setup(repo, pub)

			svc := application.NewBookingService(repo, pub, nil)
			_, err := svc.CreateBooking(context.Background(), tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
