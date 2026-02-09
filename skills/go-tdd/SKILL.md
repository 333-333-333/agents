---
name: go-tdd
description: >
  Test-Driven Development workflow in Go: red-green-refactor cycle, table-driven tests, mocks via interfaces, integration tests with testcontainers, and contract tests between services.
  Trigger: When writing tests, implementing TDD, creating mocks, setting up integration tests, or creating contract tests.
metadata:
  author: 333-333-333
  version: "2.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Writing tests in Go"
    - "Implementing TDD workflow"
    - "Creating mocks for Go interfaces"
    - "Setting up integration tests with testcontainers"
    - "Writing integration tests with testcontainers"
    - "Creating contract tests between services"
    - "Testing repository implementations against real PostgreSQL"
    - "Verifying API contracts between gateway and services"
    - "Setting up integration test infrastructure"
---

## When to Use

- Starting implementation of any new feature (TDD = write test FIRST)
- Adding tests for existing code
- Creating mocks for domain ports
- Setting up integration tests with real databases
- Writing contract tests between services (producer/consumer)
- Reviewing test quality

---

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Test FIRST** | Write the failing test before writing production code. Always. |
| **Red → Green → Refactor** | Fail → pass → clean up. Never skip a step. |
| **Table-driven tests** | Go convention: use `[]struct` for test cases |
| **Mocks via interfaces** | No mock frameworks needed — domain ports ARE the mock boundary |
| **Test file colocation** | `service.go` → `service_test.go` in same package |
| **Integration tests tagged** | Use `//go:build integration` tag |
| **No test pollution** | Each test is independent, no shared mutable state |

---

## The TDD Cycle

```
1. RED    — Write a test that describes the desired behavior. Run it. It MUST fail.
2. GREEN  — Write the MINIMUM code to make the test pass. Nothing more.
3. REFACTOR — Clean up the code while keeping tests green.
4. REPEAT
```

### When Writing Code with AI (IMPORTANT)

```
1. DESCRIBE the behavior you want as a test
2. AI writes the failing test
3. Run it — confirm RED
4. AI writes the implementation
5. Run it — confirm GREEN
6. AI refactors if needed
7. Run it — confirm still GREEN
```

---

## Unit Tests

### Table-Driven Pattern

> See [assets/service_test.go](assets/service_test.go)

### Mock Pattern (No Framework Needed)

> See [assets/mock_test.go](assets/mock_test.go)

### Domain Entity Tests

> See [assets/entity_test.go](assets/entity_test.go)

### Handler Tests (httptest)

> See [assets/http_test.go](assets/http_test.go)

---

## Integration Tests (Testcontainers)

### Unit vs Integration

| Aspect | Unit Test | Integration Test |
|--------|-----------|-----------------|
| Database | In-memory implementation | Real PostgreSQL via testcontainers |
| External services | Mocked interfaces | Real or contract-verified |
| Speed | Milliseconds | Seconds (container startup) |
| Build tag | None | `//go:build integration` |
| File suffix | `_test.go` | `_integration_test.go` |
| When to run | Always (`go test ./...`) | CI + explicit (`go test -tags=integration`) |

### Testcontainers Pattern

Every integration test file follows this structure:

1. Build tag `//go:build integration` on line 1
2. `TestMain` function that starts the container and runs migrations
3. Shared `pool` variable for all tests in the package
4. Table-driven tests using the real pool
5. Cleanup in `TestMain` (container terminates automatically)

> See [assets/postgres_setup_test.go](assets/postgres_setup_test.go) for the TestMain pattern with PostgreSQL container setup.

> See [assets/repository_integration_test.go](assets/repository_integration_test.go) for repository integration test examples.

### Repository Test Strategy

Two implementations, two purposes:

| Layer | Implementation | Used By | What It Validates |
|-------|---------------|---------|-------------------|
| In-memory | `repository/memory.go` | Unit tests + local dev without Docker | Domain logic, application flow |
| PostgreSQL | `repository/postgres.go` | Integration tests (testcontainers) + deployed envs | Real SQL, constraints, indexes |

> See [assets/repository_test_strategy.go](assets/repository_test_strategy.go)

> See `go-repository-pattern` skill for in-memory implementation details.

---

## Contract Tests

Contract tests verify that service boundaries (API contracts) don't break silently. Two perspectives:

| Role | What it tests | Who owns it |
|------|--------------|-------------|
| **Producer** | "My API responses match the documented contract" | Service team |
| **Consumer** | "I can parse the response the other service returns" | Client team |

Contract tests run against a real HTTP server (httptest) with in-memory repos — they validate the **shape** of requests/responses, NOT business logic.

> See [assets/contract_producer_test.go](assets/contract_producer_test.go) for producer contract test pattern.

> See [assets/contract_consumer_test.go](assets/contract_consumer_test.go) for consumer contract test pattern.

### Handler Integration Tests

Handler integration tests use `httptest.NewServer` with the real Gin router, real service, and real (or in-memory) repos. They test the full HTTP stack including:

- JSON serialization/deserialization
- HTTP status codes for all error paths
- Request validation (binding tags)
- Response envelope format (`{"data": ...}` / `{"error": ...}`)

> See [assets/handler_integration_test.go](assets/handler_integration_test.go) for the full handler integration test pattern.

---

## Decision Tree

```
Testing domain logic (entity, value object)?
  → Unit test with table-driven pattern

Testing a use case / application service?
  → Unit test with mocked ports (in-memory repo)

Testing a repository method against real SQL?
  → Integration test with testcontainers

Testing a handler's HTTP contract?
  → Contract test with httptest (in-memory repo)

Testing cross-service communication?
  → Consumer contract test (parse expected responses)

Testing migrations work?
  → Integration test: run migrations, verify schema

Testing full HTTP stack (handler → service → repo)?
  → Handler integration test with httptest
```

---

## Test Organization

```
internal/booking/
  domain/
    entity.go
    entity_test.go           # Unit tests — pure domain logic
    port.go
  application/
    service.go
    service_test.go          # Unit tests — use cases with in-memory repos
    mock_test.go             # Mock implementations for non-repo ports
  infrastructure/
    repository/
      memory.go              # In-memory implementation
      memory_test.go         # Unit tests — verify in-memory behavior
      postgres.go            # PostgreSQL implementation
      postgres_integration_test.go  # Integration — real DB via testcontainers
    handler/
      http.go
      http_test.go           # HTTP handler tests — httptest
      http_contract_test.go  # Contract tests — response shape validation
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/service_test.go` | Table-driven unit test for application service |
| `assets/mock_test.go` | Manual mock implementation via interfaces |
| `assets/entity_test.go` | Domain entity unit tests |
| `assets/http_test.go` | Handler tests with httptest |
| `assets/postgres_setup_test.go` | TestMain: start PostgreSQL container, run migrations, expose pool |
| `assets/repository_integration_test.go` | Repository integration tests with real PostgreSQL |
| `assets/repository_test_strategy.go` | In-memory vs PostgreSQL test strategy |
| `assets/contract_producer_test.go` | Producer contract test: verify API response shape |
| `assets/contract_consumer_test.go` | Consumer contract test: parse responses from other services |
| `assets/handler_integration_test.go` | Full HTTP handler integration test with httptest |
| `assets/Makefile` | Test-related Makefile targets |

---

## Commands

```bash
# Run all unit tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package
go test -v ./internal/booking/application/...

# Run integration tests only
go test -v -tags=integration ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test
go test -v -run TestBookingService_CreateBooking ./internal/booking/application/

# Race detection
go test -race ./...

# Testcontainers dependency
go get github.com/testcontainers/testcontainers-go
go get github.com/testcontainers/testcontainers-go/modules/postgres
```

## Makefile Targets

> See [assets/Makefile](assets/Makefile)

---

## Anti-Patterns

| Don't | Do |
|----------|-------|
| Write code first, tests later | Write test FIRST (red), then code (green) |
| Use testify/mock or mockgen | Use manual mocks via interfaces — simpler, explicit |
| Test implementation details | Test behavior — inputs and outputs |
| Shared state between tests | Each test creates its own fixtures |
| Skip integration tests | Use testcontainers — real DB, real behavior |
| One giant test function | Table-driven tests with descriptive names |
| `t.Log` for assertions | Use `t.Errorf` / `t.Fatalf` with clear messages |
| Skip contract tests for APIs | Validate response shapes to catch drift early |
