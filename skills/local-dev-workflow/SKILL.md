---
name: local-dev-workflow
description: >
  Local development workflow: Makefiles, port conventions, running services, and environment setup.
  Trigger: When adding new services, updating Makefiles, configuring local ports, or setting up local development.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.1"
  type: project
  scope: [api]
  auto_invoke:
    - Adding a new microservice to the project
    - Updating Makefiles for services
    - Configuring local development ports
    - Setting up local development environment
    - Adding service to docker-compose
---

## When to Use

- Adding a new microservice to `api/`
- Updating or creating Makefiles
- Changing port assignments
- Modifying docker-compose services
- Setting up a developer's local environment

## Port Convention

Every service gets a unique local port. In Docker, all services use `:8080` internally.

| Service | Local port | Docker internal | `pg-{service}` host port |
|---------|-----------|-----------------|--------------------------|
| gateway | 8080 | 8080 | — |
| auth | 8081 | 8080 | 5433 |
| booking | 8082 | 8080 | 5434 |
| caregiver | 8083 | 8080 | 5435 |
| payment | 8084 | 8080 | 5436 |
| notification | 8085 | 8080 | 5437 |
| pet | 8086 | 8080 | 5438 |

**Rule**: next service = next sequential port. Never reuse ports.

## Environment Modes

| Mode | `APP_ENV` | `DB_PROVIDER` | How to run |
|------|-----------|---------------|------------|
| Local (no Docker) | `local` | `memory` | `make run` or `make run-{service}` |
| Local + PostgreSQL | `development` | `postgres` | `make run-{service}-pg` (needs Docker for DB) |
| Docker full stack | `development` | `postgres` | `make up` |
| Staging | `staging` | `postgres` | Cloud deployment |
| Production | `production` | `postgres` | Cloud deployment |

## Root Makefile Commands

```bash
make run            # Start ALL services in background (in-memory)
make stop           # Stop all background services
make logs           # Tail all service logs
make run-{service}  # Start single service
make test           # Run all tests
make test-{service} # Test single service
make build          # Build all binaries
make up             # docker-compose up -d
make down           # docker-compose down -v
make sqlc           # Regenerate sqlc code for all services
```

## Per-Service Makefile Template

When creating a new service, its `Makefile` MUST follow the template in [`assets/service.Makefile`](assets/service.Makefile).

Replace `{name}`, `{assigned port from table}`, `{pg host port from table}`, `{service_name}`, and `{domain}` with actual values.

## Checklist: Adding a New Service

When adding a new service to the project:

1. [ ] Assign next sequential port in the port table above
2. [ ] Create `api/{service}/Makefile` following the template
3. [ ] Add `run-{service}` and `test-{service}` targets to root `Makefile`
4. [ ] Add `pg-{service}` container to `docker-compose.yml`
5. [ ] Add service container to `docker-compose.yml`
6. [ ] Add service URL env var to gateway in `docker-compose.yml`
7. [ ] Add proxy route in gateway's `main.go`
8. [ ] Add `{SERVICE}_SERVICE_URL` to gateway's `config.go`
9. [ ] Update this skill's port table if needed

## Docker-Compose: Database per Service Pattern

Each service gets its own PostgreSQL container following the template in [`assets/docker-compose.service.yml`](assets/docker-compose.service.yml).

Replace `{service}` and `{host port}` with actual values from the port table.

**Critical**: dev passwords are service name (e.g. `auth`/`auth`). Production uses secrets management.

## File Locations

```
api/Makefile                      # API orchestrator — runs all services
api/gateway/Makefile              # Gateway service
api/auth/Makefile                 # Auth service
api/{service}/Makefile            # Future services
api/docker-compose.yml            # Docker full stack
api/logs/                         # Local dev log files (gitignored)
api/.pids/                        # Background process PIDs (gitignored)
```
