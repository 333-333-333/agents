---
name: go-microservice-scaffold
description: >
  Scaffold a new Go microservice: directory structure, naming conventions, file layout, and step-by-step creation guide following clean architecture.
  Trigger: When creating a new microservice, adding a new domain, or reviewing service structure.
metadata:
  author: 333-333-333
  version: "2.0"
  type: project
  scope: [api]
  auto_invoke:
    - "Creating a new Go microservice"
    - "Adding a new domain to a microservice"
    - "Reviewing microservice directory structure"
---

## When to Use

- Creating a new microservice from scratch
- Adding a new domain/module to an existing service
- Reviewing or refactoring service structure
- Onboarding someone to the Go service layout

> **Architecture principles** (dependency rule, ports & adapters, layer responsibilities) live in the `clean-architecture` skill. Load it first if you need the WHY. This skill covers the HOW for Go.

---

## Service Structure

Each microservice lives under `api/` and follows this layout:

### Single-Domain Service (Most Common)

When a service has one primary domain (e.g., `notification`, `booking`):

```
api/
  {service-name}/              # e.g., booking, payment, notification
    cmd/
      server/
        main.go                # Entry point — composition root
    internal/
      domain/                  # Domain layer: entities, value objects, ports
        entity.go
        value_object.go
        port.go
        error.go
      application/             # Application layer: use cases, DTOs
        service.go
        dto.go
      infrastructure/          # Infrastructure layer: implementations
        config/                # Environment configuration
          config.go
        database/              # Database connection setup
          postgres.go
        handler/               # HTTP/gRPC handlers
          http.go
          grpc.go
        middleware/            # HTTP middleware
          auth.go
          logging.go
        repository/            # Persistence implementations
          postgres.go
          memory.go
        messaging/             # Message pub/sub implementations
          publisher.go
          subscriber.go
        server/                # Server setup
          http.go
          grpc.go
    proto/                     # Protobuf definitions for this service
      {domain}.proto
    migrations/                # SQL migrations
      000001_initial.up.sql
      000001_initial.down.sql
    tests/
      e2e/                     # E2E tests (Bruno collections)
        bruno.json
        environments/
          local.bru
        {domain}/
          *.bru
      load/                    # Load tests (k6 scripts)
        smoke.js
        load.js
    Dockerfile
    go.mod
    go.sum
    Makefile
    README.md
```

> **Test location convention**: See `testing-strategy` skill for where each test level goes (unit, integration, contract, E2E, load).

### Multi-Domain Service (Rare)

Only use this when a service has multiple distinct domains (e.g., `booking` + `availability`):

```
api/
  {service-name}/
    cmd/server/main.go
    internal/
      {domain-1}/              # e.g., booking
        domain/
        application/
        infrastructure/
      {domain-2}/              # e.g., availability
        domain/
        application/
        infrastructure/
      shared/                  # ONLY for truly cross-domain concerns
        middleware/
        config/
```

> **Rule of thumb**: If you have `internal/shared/`, reconsider. Most "shared" code belongs in `infrastructure/`.

---

## Naming Conventions

| Element | Convention | Example |
|---------|-----------|---------|
| Service directory | lowercase singular noun | `api/booking/` |
| Domain directory | `domain/` at internal root | `internal/domain/` |
| Entity file | `entity.go` | `internal/domain/entity.go` |
| Port file | `port.go` | `internal/domain/port.go` |
| Use case file | `service.go` | `internal/application/service.go` |
| Package names | short, lowercase, singular | `package booking` |
| Interfaces | verb-er or descriptive | `BookingRepository`, `PaymentGateway` |
| Structs | noun | `Booking`, `CreateBookingInput` |
| Domain errors | `Err` prefix | `ErrBookingNotFound`, `ErrBookingOverlap` |
| Environment variables | `SCREAMING_SNAKE` | `DB_PASSWORD` |

---

## Domain Layer — port.go Pattern

> See [assets/port.go](assets/port.go)

## Application Layer — service.go Pattern

> See [assets/service.go](assets/service.go)

---

## Adding a New Domain to a Service

For single-domain services (most common):

1. Add to existing structure:
   - Entities → `internal/domain/entity.go`
   - Use cases → `internal/application/service.go`
   - Adapters → `internal/infrastructure/`

For multi-domain services (rare):

1. Create directory: `internal/{domain-name}/domain/`, `application/`, `infrastructure/`
2. Define entities in `internal/{domain-name}/domain/entity.go`
3. Define ports in `internal/{domain-name}/domain/port.go`
4. Implement use cases in `internal/{domain-name}/application/service.go`
5. Implement adapters in `internal/{domain-name}/infrastructure/`
6. Wire in `cmd/server/main.go` (see `go-service-bootstrap` skill)

## Adding a New Microservice

1. Create `api/{service-name}/` with full structure above
2. Initialize Go module: `go mod init github.com/{org}/bastet/api/{service-name}`
3. Add Dockerfile, Makefile, README.md (see `go-docker-deploy` skill)
4. Add protobuf definitions if service exposes gRPC (see `go-grpc-services` skill)
5. Add migration files if service has persistence (see `go-repository-pattern` skill)
6. Wire dependencies in main.go (see `go-service-bootstrap` skill)
7. Register in docker-compose and gateway routing (see `local-dev-workflow` skill)

---

## Commands

```bash
# Create a new SINGLE-DOMAIN service scaffold (most common)
mkdir -p api/{service}/cmd/server
mkdir -p api/{service}/internal/{domain,application,infrastructure/{config,database,handler,middleware,repository,messaging,server}}
mkdir -p api/{service}/proto
mkdir -p api/{service}/migrations

# Initialize Go module
cd api/{service} && go mod init github.com/{org}/bastet/api/{service}

# Verify structure
tree api/{service}
```

---

## Anti-Patterns

| Don't | Do |
|----------|-------|
| `api/booking/controllers/` | `api/booking/internal/infrastructure/handler/` |
| `api/booking/models/` | `api/booking/internal/domain/entity.go` |
| `internal/shared/` in single-domain services | Put in `internal/infrastructure/` |
| `internal/{domain}/` when service has one domain | Use `internal/domain/` directly |
| Import cloud SDK in domain | Define interface in domain, implement in infrastructure |
| Shared database between services | Each service owns its data |
| `utils/` package | Put helpers in the domain or infrastructure where they belong |
| Business logic in handlers | Handlers call application service, that's it |

---

## Related Skills

| Skill | What it covers |
|-------|----------------|
| `clean-architecture` | **Principles**: dependency rule, layers, ports & adapters (load first!) |
| `go-service-bootstrap` | Configuration, DI wiring, environments, graceful shutdown |
| `go-gin-handlers` | HTTP handler patterns, middleware, error responses |
| `go-repository-pattern` | Persistence, migrations, sqlc |
| `go-docker-deploy` | Dockerfile, docker-compose, deployment |
