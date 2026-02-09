---
name: go-api-docs
description: >
  API documentation with OpenAPI/Swagger generated from code annotations, plus service-level README conventions.
  Trigger: When documenting API endpoints, generating Swagger specs, or writing service documentation.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Documenting Go API endpoints"
    - "Generating OpenAPI/Swagger specs"
    - "Writing microservice README files"
---

## When to Use

- Documenting HTTP API endpoints
- Generating OpenAPI/Swagger specs
- Writing service README files
- Documenting gRPC services (protobuf is self-documenting)

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Swagger from annotations** | Use `swaggo/swag` to generate OpenAPI from Go comments |
| **Spec lives with service** | Each service generates its own `docs/swagger.json` |
| **Proto is self-documenting** | gRPC services use protobuf comments as docs |
| **README per service** | Each service has a README with setup, architecture, API summary |
| **Keep specs in sync** | Regenerate on every API change — add to CI |

## Swagger Annotations (swaggo/swag)

> See [assets/swagger_main.go](assets/swagger_main.go)

> See [assets/swagger_handler.go](assets/swagger_handler.go)

## Serve Swagger UI

> See [assets/swagger_ui.go](assets/swagger_ui.go)

## Protobuf Documentation

Protobuf files are self-documenting. Use comments:

> See [assets/caregiver.proto](assets/caregiver.proto)

## Service README Template

> See [assets/README_TEMPLATE.md](assets/README_TEMPLATE.md)

## Commands

```bash
# Install swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g cmd/server/main.go -o docs

# Dependencies
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
go get github.com/swaggo/swag

# View docs
open http://localhost:8080/swagger/index.html

# Makefile target
.PHONY: swagger
swagger:
	swag init -g cmd/server/main.go -o docs
```

## Anti-Patterns

| Don't | Do |
|----------|-------|
| Manually write OpenAPI YAML | Generate from code annotations |
| Swagger in production | Disable Swagger UI in production |
| Outdated README | Regenerate swagger on API changes, keep README current |
| No docs for events | Document published/consumed events in README |
| Undocumented protobuf | Add comments to every service, rpc, and message |
