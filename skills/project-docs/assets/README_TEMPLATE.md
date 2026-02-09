# {Service Name} Service

{One-line description}

## Architecture

This service owns the **{domain}** bounded context.

| Layer | Location |
|-------|----------|
| Domain | `internal/{domain}/domain/` |
| Application | `internal/{domain}/application/` |
| Infrastructure | `internal/{domain}/infrastructure/` |

## API

### HTTP (REST)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/bookings` | Create a booking |
| GET | `/api/v1/bookings/:id` | Get booking by ID |
| GET | `/api/v1/bookings` | List bookings for owner |
| PATCH | `/api/v1/bookings/:id/cancel` | Cancel a booking |

Swagger UI: `http://localhost:8080/swagger/index.html`

### gRPC

See `proto/*.proto` for service definitions.

### Events Published

| Topic | Event | Description |
|-------|-------|-------------|
| `booking.booking.created` | BookingCreated | Emitted when a new booking is created |
| `booking.booking.cancelled` | BookingCancelled | Emitted when a booking is cancelled |

### Events Consumed

| Topic | Handler | Description |
|-------|---------|-------------|
| `payment.transaction.completed` | `handlePaymentCompleted` | Confirms booking after payment |

## Setup

```bash
# Run locally
go run ./cmd/server

# Run tests
make test

# Run with infrastructure
docker-compose up -d postgres nats minio
make migrate-up
make run

# Generate swagger
make swagger

# Generate protobuf
make proto
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_ENV` | `development` | Environment (development/staging/production) |
| `HTTP_PORT` | `8080` | HTTP server port |
| `GRPC_PORT` | `9090` | gRPC server port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PASSWORD` | — | PostgreSQL password (required) |
