---
name: go-service-bootstrap
description: >
  Service bootstrap in Go: manual dependency injection, environment-based configuration, composition root wiring, graceful shutdown, and multi-environment conventions.
  Trigger: When wiring dependencies in main.go, setting up configuration, managing secrets, or configuring environment-specific behavior.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [api]
  auto_invoke:
    - "Wiring dependencies in Go main.go"
    - "Setting up environment-based configuration"
    - "Managing secrets in Go microservices"
    - "Configuring environment-specific behavior in Go"
    - "Setting up local, dev, staging, or production environments"
    - "Managing environment variables per deployment stage"
---

## When to Use

- Setting up `main.go` to wire all dependencies
- Adding environment-based configuration
- Managing secrets (DB passwords, API keys)
- Deciding what changes between local/dev/staging/prod
- Adding a new adapter/implementation swap
- Setting up docker-compose profiles

---

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Manual DI** | No DI frameworks (no Wire, no Dig) — Go's simplicity is the feature |
| **Constructor injection** | All dependencies via `NewXxx(deps...)` constructors |
| **Config from environment** | All configuration from env vars, no config files in production |
| **Secrets via env** | Secrets injected via environment, never hardcoded or in repo |
| **Wire in main.go** | `cmd/server/main.go` is the ONLY place that knows all concrete types |
| **Same code, different config** | Never `if env == "production"` in domain — only in config and infra |

---

## Environments

| Environment | `APP_ENV` | Purpose |
|-------------|-----------|---------|
| **local** | `local` | Developer machine, docker-compose, fast feedback |
| **dev** | `development` | Shared cloud environment for team integration |
| **staging** | `staging` | Pre-production replica, QA and smoke tests |
| **production** | `production` | Live, real users, real money |

### What Changes Per Environment

#### Logging

| Setting | local | dev | staging | production |
|---------|-------|-----|---------|------------|
| Format | text | JSON | JSON | JSON |
| Level | debug | debug | info | info |
| Source location | yes | yes | no | no |

#### Database

| Setting | local (no Docker) | local (Docker) | dev | staging | production |
|---------|-------------------|----------------|-----|---------|------------|
| Provider | `memory` | `postgres` | `postgres` | `postgres` | `postgres` |
| Host | N/A | `localhost` | Cloud SQL / RDS | Cloud SQL / RDS | Cloud SQL / RDS |
| SSL | N/A | `disable` | `require` | `require` | `verify-full` |
| Max connections | N/A | 5 | 10 | 20 | 25 |
| Password source | N/A | `.env` file | Secret manager | Secret manager | Secret manager |

The in-memory provider uses a pure Go `map`-based repository — zero external dependencies. See `go-repository-pattern` skill for implementation details.

#### Observability

| Setting | local | dev | staging | production |
|---------|-------|-----|---------|------------|
| Tracing | disabled or Jaeger local | OTel Collector → Jaeger | OTel Collector → Jaeger + S3 | OTel Collector → Jaeger + S3 |
| Trace sampling | 100% | 100% | 50% | 10% |
| Metrics | disabled | Prometheus | Prometheus | Prometheus + S3 export |

#### Messaging

| Setting | local | dev | staging | production |
|---------|-------|-----|---------|------------|
| Provider | NATS (docker-compose) | NATS cluster | NATS cluster | NATS cluster |
| URL | `nats://localhost:4222` | Cloud endpoint | Cloud endpoint | Cloud endpoint |

#### Object Storage

| Setting | local | dev | staging | production |
|---------|-------|-----|---------|------------|
| Provider | MinIO (docker-compose) | S3 / GCS | S3 / GCS | S3 / GCS |
| Endpoint | `http://localhost:9000` | (default SDK) | (default SDK) | (default SDK) |
| Force path style | `true` | `false` | `false` | `false` |

#### Security

| Setting | local | dev | staging | production |
|---------|-------|-----|---------|------------|
| CORS origins | `*` | `*.bastet.dev` | `*.bastet.dev` | `bastet.cl` |
| Rate limiting | disabled | soft limits | production limits | production limits |
| JWT validation | mock / disabled | real | real | real |
| Swagger UI | enabled | enabled | enabled | **disabled** |

---

## Configuration

> See [assets/config.go](assets/config.go) — centralized config struct with typed helpers.

All configuration loaded once via `config.Load()`. Typed helpers: `getEnv`, `requireEnv`, `getEnvInt`, `getEnvBool`, `getEnvDuration`.

---

## main.go — The Composition Root

> See [assets/main.go](assets/main.go)

The composition root follows a numbered sequence:
1. Load config
2. Setup logger
3. Infrastructure — databases
4. Infrastructure — messaging
5. Infrastructure — storage
6. Repositories (adapters)
7. Application services (use cases)
8. HTTP handlers
9. Router
10. Start server

---

## Logger Setup

> See [assets/logger.go](assets/logger.go) — environment-aware slog configuration (text for local, JSON for others).

---

## Health Check Convention

Every service exposes two endpoints:

| Endpoint | Purpose | When it fails |
|----------|---------|---------------|
| `GET /health` | Liveness — is the process alive? | Restart the container |
| `GET /ready` | Readiness — can it serve traffic? | Stop sending traffic |

> See [assets/health_check.go](assets/health_check.go)

---

## Graceful Shutdown

> See [assets/server.go](assets/server.go) — listens for context cancellation, 10-second shutdown timeout.

---

## Environment Variable Files

```
api/{service}/
  .env.example        # Committed — all keys, safe defaults for local
  .env                 # NOT committed — developer's local overrides
```

> See [assets/env.example](assets/env.example)

**NEVER commit `.env`** — only `.env.example`.

### Secrets (dev/staging/prod)

Secrets are NEVER in env files for non-local environments. Use:

| Provider | Tool |
|----------|------|
| GCP | Google Secret Manager |
| AWS | AWS Secrets Manager / SSM Parameter Store |
| Kubernetes | Sealed Secrets / External Secrets Operator |

Secrets are injected as environment variables at runtime by the orchestrator.

---

## Docker-Compose Profiles

> See [assets/docker-compose.example.yml](assets/docker-compose.example.yml)

```bash
# Start everything
docker-compose --profile local up -d

# Start only infra (DB, NATS, MinIO) — run services natively
docker-compose --profile infra up -d

# Start only observability stack
docker-compose --profile observability up -d
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/config.go` | Centralized config struct with typed env helpers |
| `assets/main.go` | Composition root with numbered wiring sequence |
| `assets/logger.go` | Environment-aware slog setup |
| `assets/server.go` | Graceful HTTP server with shutdown |
| `assets/health_check.go` | Liveness and readiness endpoints |
| `assets/env.example` | Environment variable template for local dev |
| `assets/docker-compose.example.yml` | Docker-compose with profiles for local dev |

---

## Commands

```bash
# Local dev with .env
cp .env.example .env
docker-compose --profile infra up -d
go run ./cmd/server

# Or use direnv
echo 'dotenv' > .envrc && direnv allow

# Run with specific env
APP_ENV=development go run ./cmd/server

# Verify environment
curl http://localhost:8080/health
```

---

## Anti-Patterns

| Don't | Do |
|----------|-------|
| DI frameworks (Wire, Dig, fx) | Manual constructor injection in main.go |
| Global variables for DB/config | Pass dependencies through constructors |
| Config files (YAML/TOML) in production | Environment variables only |
| Secrets in repo or config files | Inject via env vars, use secret manager in CI/CD |
| Init functions with side effects | Explicit initialization in main.go |
| Scattered `os.Getenv` calls | Centralized `config.Load()` |
| `if env == "production"` in domain | Environment logic in config and infrastructure only |
| Different code paths per environment | Same code, different configuration |
| Skip SSL in staging | Staging mirrors production security |
| Disable rate limiting in staging | Staging mirrors production limits |
| Hardcode environment values | Always read from `os.Getenv` via `config.Load()` |
