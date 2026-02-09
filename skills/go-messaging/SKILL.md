---
name: go-messaging
description: >
  Cloud-agnostic asynchronous messaging with abstract interfaces and implementations for NATS, RabbitMQ, and Google Pub/Sub.
  Trigger: When implementing event-driven communication, publishing domain events, or creating event subscribers.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Publishing domain events in Go"
    - "Implementing async messaging between services"
    - "Creating event subscribers"
---

## When to Use

- Publishing domain events after state changes
- Subscribing to events from other services
- Choosing a messaging backend for the project
- Implementing async workflows (notifications, emails, analytics)

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Interface in domain** | `EventPublisher` is a domain port |
| **Implementation swappable** | NATS, RabbitMQ, or Pub/Sub behind same interface |
| **Events are immutable facts** | Past tense: `BookingCreated`, `PaymentCompleted` |
| **Events carry data** | Include enough data so consumers don't need to call back |
| **Idempotent consumers** | Subscribers MUST handle duplicate messages |
| **Cloud agnostic** | Application layer never imports messaging SDK |

## Event Definition (Domain)

> **Reference:** [assets/event.go](assets/event.go)

## Domain Port

> **Reference:** [assets/port.go](assets/port.go)

## NATS Implementation

> **Reference:** [assets/nats_publisher.go](assets/nats_publisher.go)

> **Reference:** [assets/nats_subscriber.go](assets/nats_subscriber.go)

## Topic Naming Convention

```
{service}.{domain}.{event}

Examples:
  booking.booking.created
  booking.booking.cancelled
  payment.transaction.completed
  caregiver.profile.verified
  notification.email.requested
```

## Publishing from Application Layer

> **Reference:** [assets/publish_usage.go](assets/publish_usage.go)

## In-Memory Implementation (Testing)

> **Reference:** [assets/memory_publisher.go](assets/memory_publisher.go)

## Commands

```bash
# NATS
brew install nats-server
nats-server  # Start locally

# Dependencies
go get github.com/nats-io/nats.go

# RabbitMQ alternative
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
go get github.com/rabbitmq/amqp091-go
```

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Import NATS/RabbitMQ in domain | Define `EventPublisher` interface in domain |
| Events named as commands (`CreateBooking`) | Events are past tense facts (`BookingCreated`) |
| Consumer calls back to producer service | Include enough data in the event payload |
| Assume exactly-once delivery | Design idempotent consumers |
| Synchronous event publishing blocking the request | Use goroutines or accept best-effort for non-critical events |
