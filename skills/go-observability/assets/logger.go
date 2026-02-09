// internal/shared/observability/logger.go
package observability

import (
	"context"
	"log/slog"
	"os"
)

func SetupLogger(env string) {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		AddSource: true,
	}

	if env == "production" {
		opts.Level = slog.LevelInfo
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}

// ContextLogger returns a logger enriched with trace context.
func ContextLogger(ctx context.Context) *slog.Logger {
	logger := slog.Default()

	// Add trace ID if present
	traceID := TraceIDFromContext(ctx)
	if traceID != "" {
		logger = logger.With("trace_id", traceID)
	}

	return logger
}
