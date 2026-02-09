//go:build integration

package repository_test

import (
	"context"
	"testing"
	// Adjust import paths for your service
	// "github.com/333-333-333/bastet/api/notification/internal/notification/domain"
	// "github.com/333-333-333/bastet/api/notification/internal/notification/infrastructure/repository"
)

func TestPostgresNotificationRepository_Save(t *testing.T) {
	// pool is set up by TestMain in postgres_setup_test.go
	// repo := repository.NewPostgresNotificationRepository(pool)
	ctx := context.Background()

	tests := []struct {
		name    string
		setup   func() // Pre-conditions
		verify  func() // Post-conditions to check in DB
		wantErr bool
	}{
		{
			name: "saves notification successfully",
			setup: func() {
				// Create notification via domain factory
				// n, _ := domain.NewNotification(domain.ChannelEmail, "user@test.com", "Subject", "Content")
			},
			verify: func() {
				// Query DB directly to verify persistence
				// var count int
				// err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM notifications").Scan(&count)
				// if err != nil { t.Fatal(err) }
				// if count != 1 { t.Errorf("expected 1 notification, got %d", count) }
			},
		},
		{
			name: "retrieves notification by ID",
			setup: func() {
				// Insert a notification
			},
			verify: func() {
				// repo.FindByID should return it with all fields
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean table before each test
			_, err := pool.Exec(ctx, "DELETE FROM notifications")
			if err != nil {
				t.Fatalf("failed to clean table: %v", err)
			}

			tt.setup()
			tt.verify()
		})
	}
}
