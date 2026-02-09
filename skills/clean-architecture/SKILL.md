---
name: clean-architecture
description: >
  Language-agnostic Clean Architecture principles: dependency rule, layer separation, ports & adapters, screaming architecture, and bounded contexts.
  Trigger: When designing service architecture, reviewing layer boundaries, onboarding to the codebase, or deciding where code belongs.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [root]
  auto_invoke:
    - "Designing service architecture"
    - "Reviewing layer boundaries and dependency direction"
    - "Deciding where a piece of code belongs"
    - "Onboarding to the codebase architecture"
    - "Evaluating if a dependency violates the architecture"
---

## When to Use

- Designing a new service or module from scratch
- Deciding which layer a piece of code belongs to
- Reviewing code for architecture violations
- Onboarding someone to the project's architecture
- Evaluating whether to add a dependency to a layer

This skill defines the **principles**. For language-specific implementation, see:
- Go: `go-microservice-scaffold` skill (directory structure, naming, scaffold steps)
- Flutter: `flutter-scaffold` skill (feature-based structure, layers)

---

## The Dependency Rule

This is the ONE rule that makes Clean Architecture work. Everything else follows from it.

```
┌─────────────────────────────────────┐
│          Infrastructure             │  Frameworks, drivers, UI, DB
│  ┌───────────────────────────────┐  │
│  │         Application           │  │  Use cases, orchestration
│  │  ┌─────────────────────────┐  │  │
│  │  │         Domain           │  │  │  Entities, value objects, ports
│  │  └─────────────────────────┘  │  │
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘

Dependencies ALWAYS point INWARD.
Inner layers NEVER know about outer layers.
```

**Concretely:**
- **Domain** imports NOTHING from application or infrastructure. Zero.
- **Application** imports from domain only. Never infrastructure.
- **Infrastructure** implements domain interfaces. It's the only layer that imports SDKs, drivers, frameworks.

### Why This Matters

The dependency rule exists so that:
1. **Domain logic is testable** without databases, HTTP, or any framework
2. **Infrastructure is swappable** — change your database, cloud provider, or framework without touching business logic
3. **The codebase communicates intent** — domain code reads like business rules, not framework boilerplate

---

## The Three Layers

### Domain (innermost)

The domain layer contains the **business rules** of your system. It has no dependencies on anything external.

| Contains | Description | Example |
|----------|-------------|---------|
| **Entities** | Objects with identity and lifecycle | `Booking`, `User`, `Pet` |
| **Value Objects** | Immutable objects defined by their attributes | `Money`, `Email`, `BookingStatus` |
| **Ports (interfaces)** | Contracts that infrastructure must fulfill | `BookingRepository`, `PaymentGateway` |
| **Domain Errors** | Business rule violations | `ErrBookingOverlap`, `ErrInsufficientFunds` |
| **Domain Services** | Logic that doesn't belong to a single entity | `PricingCalculator` |

**Rules:**
- Entities validate their own invariants (no invalid state)
- Value objects are immutable — create new ones instead of mutating
- Ports are defined by the domain, implemented by infrastructure
- Domain errors are specific and meaningful — never generic "error occurred"

### Application (middle)

The application layer contains **use cases** — the specific things your system can do.

| Contains | Description | Example |
|----------|-------------|---------|
| **Use Cases / Services** | Orchestrate domain objects to fulfill a request | `CreateBooking`, `CancelBooking` |
| **DTOs** | Input/output structures for use cases | `CreateBookingInput`, `BookingOutput` |
| **Application Errors** | Errors specific to use case orchestration | `ErrBookingNotFound` (repo returned nil) |

**Rules:**
- Use cases receive DTOs, not domain entities (prevent leaking domain internals)
- Use cases call ports (interfaces), not concrete implementations
- Use cases contain NO framework-specific code
- One use case = one reason to change

### Infrastructure (outermost)

The infrastructure layer contains **adapters** — concrete implementations of domain ports.

| Contains | Description | Example |
|----------|-------------|---------|
| **Handlers** | HTTP/gRPC entry points that call use cases | `BookingHandler` |
| **Repositories** | Database implementations of domain ports | `PostgresBookingRepository` |
| **Messaging** | Message broker adapters | `NATSPublisher`, `NATSSubscriber` |
| **External Services** | Third-party API clients | `FlowPaymentGateway`, `FCMNotificationSender` |
| **Configuration** | Environment loading, server setup | `Config`, `Server` |

**Rules:**
- Handlers are thin — parse request, call use case, format response
- Repositories implement domain port interfaces exactly
- External service clients implement domain port interfaces
- ALL framework imports live here and only here

---

## Ports & Adapters

Ports & Adapters (Hexagonal Architecture) is the mechanism that enforces the dependency rule.

```
                    ┌──────────────────────┐
   HTTP Handler ──▶ │                      │ ──▶ PostgreSQL Repository
                    │   Domain + Use Cases │
  gRPC Handler ──▶  │                      │ ──▶ NATS Publisher
                    │   (defines ports)    │
   CLI Command ──▶  │                      │ ──▶ S3 Object Store
                    └──────────────────────┘
        ▲ Driving adapters           Driven adapters ▶
        (call the system)            (called by the system)
```

- **Ports** = interfaces defined in the domain layer
- **Driving adapters** = entry points that call use cases (HTTP handlers, gRPC, CLI)
- **Driven adapters** = implementations called by use cases (repositories, message publishers, external APIs)

The domain defines WHAT it needs (ports). Infrastructure decides HOW to provide it (adapters).

---

## Screaming Architecture

> "Your architecture should scream the business domain, not the framework."
> — Robert C. Martin

The directory structure communicates WHAT the system does, not how it's built.

| Bad (screams framework) | Good (screams domain) |
|---|---|
| `controllers/` | `booking/` |
| `models/` | `caregiver/` |
| `services/` | `payment/` |
| `repositories/` | `notification/` |
| `utils/` | `pet/` |

A new developer should understand the business domain by looking at the directory structure, without reading a single line of code.

---

## Bounded Contexts

Each service (or module) owns ONE bounded context — a clear boundary around a business domain.

| Principle | Rule |
|-----------|------|
| **Own your data** | Each context has its own database/schema. Never share tables. |
| **Explicit boundaries** | Communication between contexts goes through defined interfaces (APIs, events). |
| **Ubiquitous language** | Each context uses its own terminology. A "User" in auth ≠ "User" in booking. |
| **No circular dependencies** | If A depends on B and B depends on A, merge them or extract a third context. |

---

## Where Does This Code Belong?

Use this decision tree when you're unsure:

```
Does it represent a business concept (entity, rule, constraint)?
  → Domain layer

Does it orchestrate multiple domain objects to fulfill a use case?
  → Application layer

Does it talk to a database, API, message broker, or framework?
  → Infrastructure layer

Does it define WHAT is needed without specifying HOW?
  → Domain layer (as a port/interface)

Does it implement a domain interface with a specific technology?
  → Infrastructure layer (as an adapter)

Is it a DTO that enters or leaves a use case?
  → Application layer

Is it configuration, logging setup, or server bootstrap?
  → Infrastructure layer (shared/)
```

---

## Common Violations

| Violation | Why it's bad | Fix |
|-----------|-------------|-----|
| Domain imports `database/sql` | Domain now depends on a specific database | Define a repository interface in domain |
| Use case imports `gin.Context` | Application tied to HTTP framework | Handler extracts data, passes plain DTOs |
| Handler contains business logic | Logic untestable without HTTP setup | Move logic to use case, handler just delegates |
| Domain entity has JSON tags | Domain coupled to serialization format | Use separate DTOs in infrastructure |
| Shared database between services | Tight coupling, can't evolve independently | Each service owns its data |
| `utils/` package | Grab bag with no cohesion | Put helpers where they belong by domain |

---

## Testing Implications

Clean architecture makes testing natural:

| Layer | Test type | Dependencies |
|-------|-----------|-------------|
| Domain | Unit tests | None — pure logic |
| Application | Unit tests | In-memory implementations of ports |
| Infrastructure | Integration tests | Real databases (testcontainers), real APIs |
| Full stack | E2E tests | All layers wired together |

The dependency rule means domain and application tests are **fast, isolated, and reliable** — they don't need databases, network, or frameworks.

---

## Resources

- Robert C. Martin — *Clean Architecture* (2017)
- Alistair Cockburn — *Hexagonal Architecture* (Ports & Adapters)
- Vaughn Vernon — *Implementing Domain-Driven Design* (Bounded Contexts)
