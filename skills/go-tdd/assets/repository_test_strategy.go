// --- Unit tests use in-memory repos (fast, zero deps) ---

func TestBookingService_CreateBooking(t *testing.T) {
	repo := repository.NewInMemoryBookingRepository()
	pub := messaging.NewInMemoryPublisher()
	svc := application.NewBookingService(repo, pub, nil)
	// ...test use case logic...
}

// --- Integration tests use testcontainers (real DB, validates SQL) ---

//go:build integration

func TestPostgresBookingRepository_SaveAndFind(t *testing.T) {
	db, cleanup := setupPostgres(t) // testcontainers
	defer cleanup()
	repo := repository.NewPostgresBookingRepository(db)
	// ...test actual SQL...
}
