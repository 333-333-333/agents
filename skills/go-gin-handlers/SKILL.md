---
name: go-gin-handlers
description: >
  HTTP handler patterns with Gin framework: routing, middleware, validation, error handling, and API versioning.
  Trigger: When creating HTTP endpoints, adding middleware, handling request validation, or standardizing error responses.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Creating HTTP endpoints with Gin"
    - "Adding middleware to Go services"
    - "Standardizing API error responses"
---

## When to Use

- Creating new HTTP endpoints in a microservice
- Adding or modifying middleware
- Standardizing error responses
- Implementing request validation
- Setting up API versioning

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Handlers are thin** | Handlers ONLY parse input, call application service, format output |
| **No business logic** | Zero domain logic in handlers — delegate to application layer |
| **Standardized responses** | ALL endpoints return the same JSON envelope |
| **Validation at the edge** | Validate request DTOs in the handler layer using binding tags |
| **Context propagation** | Always pass `ctx` from Gin to application service |

## Response Envelope

ALL API responses follow this structure:

> **Reference:** [assets/response.go](assets/response.go)

## Handler Pattern

> **Reference:** [assets/handler.go](assets/handler.go)

## Error Mapping

> **Reference:** [assets/errors.go](assets/errors.go)

## Router Setup with Versioning

> **Reference:** [assets/router.go](assets/router.go)

## Middleware Pattern

> **Reference:** [assets/auth_middleware.go](assets/auth_middleware.go)

## Commands

```bash
# Install Gin
go get -u github.com/gin-gonic/gin

# Run server locally
go run cmd/server/main.go

# Test endpoint
curl -X POST http://localhost:8080/api/v1/bookings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"caregiver_id": "uuid", "service_type": "walk", ...}'
```

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Business logic in handlers | Call application service, return result |
| Raw `c.JSON(500, ...)` scattered | Use `server.Fail()` / `server.OK()` |
| Validate in application layer | Use Gin binding tags + `ShouldBindJSON` |
| Return different JSON shapes | Always use Response envelope |
| Parse JWT in every handler | Use auth middleware, read from `c.GetString("user_id")` |
