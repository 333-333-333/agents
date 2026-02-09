// cmd/server/main.go
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"api/booking/internal/booking/application"
	bookingHandler "api/booking/internal/booking/infrastructure/handler"
	bookingMessaging "api/booking/internal/booking/infrastructure/messaging"
	bookingRepo "api/booking/internal/booking/infrastructure/repository"
	"api/booking/internal/shared/config"
	"api/booking/internal/shared/server"
	sharedStorage "api/booking/internal/shared/storage"
)

func main() {
	// 1. Load config
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// 2. Setup logger
	setupLogger(cfg.Env)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// 3. Infrastructure — databases
	db, err := config.NewPostgresDB(cfg.Database)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// 4. Infrastructure — messaging
	publisher, err := newPublisher(cfg.Messaging)
	if err != nil {
		slog.Error("failed to create publisher", "error", err)
		os.Exit(1)
	}
	defer publisher.Close()

	// 5. Infrastructure — storage
	store, err := sharedStorage.NewObjectStore(ctx, cfg.Storage)
	if err != nil {
		slog.Error("failed to create object store", "error", err)
		os.Exit(1)
	}

	// 6. Repositories (adapters)
	bookingRepository := bookingRepo.NewPostgresBookingRepository(db)

	// 7. Application services (use cases)
	bookingService := application.NewBookingService(bookingRepository, publisher, nil)

	// 8. HTTP handlers
	bookingHTTP := bookingHandler.NewBookingHandler(bookingService)

	// 9. Router
	router := server.NewRouter(cfg.Env)
	v1 := router.Group("/api/v1")
	bookingHTTP.RegisterRoutes(v1)

	// 10. Start server
	slog.Info("starting server", "http_port", cfg.HTTP.Port, "env", cfg.Env)
	if err := server.ListenAndServe(ctx, router, cfg.HTTP.Port); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}

func setupLogger(env string) {
	var handler slog.Handler
	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(handler))
}
