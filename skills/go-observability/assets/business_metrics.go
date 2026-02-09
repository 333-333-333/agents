// internal/booking/infrastructure/metrics.go
package infrastructure

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("bastet/booking")

var (
	BookingsCreated, _ = meter.Int64Counter("bookings.created",
		metric.WithDescription("Total bookings created"),
	)
	BookingDuration, _ = meter.Float64Histogram("bookings.duration_seconds",
		metric.WithDescription("Time to process booking creation"),
	)
	ActiveBookings, _ = meter.Int64UpDownCounter("bookings.active",
		metric.WithDescription("Current active bookings"),
	)
)
