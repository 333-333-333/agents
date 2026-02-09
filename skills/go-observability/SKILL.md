---
name: go-observability
description: >
  Vendor-neutral observability with OpenTelemetry: structured logging (slog), distributed tracing, metrics, and export to object storage.
  Trigger: When adding logging, tracing, metrics, or configuring observability pipelines in microservices.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Adding structured logging to Go services"
    - "Implementing distributed tracing with OpenTelemetry"
    - "Collecting and exporting metrics"
    - "Configuring observability pipelines"
---

## When to Use

- Adding structured logging to a service
- Implementing distributed tracing across microservices
- Collecting and exporting metrics
- Exporting observability data to object storage for long-term retention
- Setting up Gin/gRPC middleware for automatic instrumentation

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **OpenTelemetry standard** | OTel SDK for traces + metrics — vendor neutral, export anywhere |
| **slog for logging** | Go stdlib `log/slog` — no Zap, no Logrus, no external deps |
| **JSON in production** | Structured JSON logs in prod, text in dev |
| **Trace context propagation** | Every request gets a trace ID, propagated across services |
| **Export to object storage** | Periodic export of logs/metrics/traces to S3/GCS/MinIO for retention |
| **Cloud agnostic** | OpenTelemetry Collector handles export — backend is config, not code |

## Three Pillars

```
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│   Logging    │    │   Tracing    │    │   Metrics    │
│   (slog)     │    │   (OTel)     │    │   (OTel)     │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       └───────────┬───────┴───────────────────┘
                   │
         ┌─────────▼──────────┐
         │  OTel Collector    │     ← Deployed as sidecar or standalone
         └─────────┬──────────┘
                   │
       ┌───────────┼───────────┐
       ▼           ▼           ▼
   Object Store  Jaeger    Prometheus
   (long-term)  (traces)   (metrics)
```

## Structured Logging (slog)

> **Reference:** [`assets/logger.go`](assets/logger.go)

### Logging Convention

> **Reference:** [`assets/logging_convention.go`](assets/logging_convention.go)

## Distributed Tracing (OpenTelemetry)

> **Reference:** [`assets/tracing.go`](assets/tracing.go)

## Metrics (OpenTelemetry)

> **Reference:** [`assets/metrics.go`](assets/metrics.go)

### Custom Business Metrics

> **Reference:** [`assets/business_metrics.go`](assets/business_metrics.go)

## Gin Middleware

> **Reference:** [`assets/gin_middleware.go`](assets/gin_middleware.go)

## gRPC Interceptors

> **Reference:** [`assets/grpc_interceptors.go`](assets/grpc_interceptors.go)

## Export to Object Storage

Use the OpenTelemetry Collector to batch and export to object storage:

> **Reference:** [`assets/otel-collector-config.yaml`](assets/otel-collector-config.yaml)

## Local Dev Stack (docker-compose)

> **Reference:** [`assets/docker-compose-observability.yml`](assets/docker-compose-observability.yml)

## Commands

```bash
# OpenTelemetry SDK
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/sdk
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc

# Gin instrumentation
go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin

# gRPC instrumentation
go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc

# View traces locally
open http://localhost:16686  # Jaeger UI

# View metrics
open http://localhost:9090   # Prometheus UI
```

## Anti-Patterns

| Don't | Do |
|----------|-------|
| `log.Printf` / `fmt.Println` | `slog.InfoContext(ctx, ...)` with structured fields |
| Vendor-specific tracing SDKs | OpenTelemetry standard |
| Skip trace context propagation | Always pass `ctx`, use OTel propagators |
| Log sensitive data (passwords, tokens) | Log IDs and metadata only |
| Metrics in domain layer | Metrics in infrastructure, referenced via interface if needed |
| Custom log format per service | Same slog setup shared via `observability.SetupLogger()` |
