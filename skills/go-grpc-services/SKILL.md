---
name: go-grpc-services
description: >
  Inter-service communication with gRPC and Protocol Buffers: contract definition, code generation, client/server patterns.
  Trigger: When creating gRPC services, defining protobuf contracts, or implementing service-to-service communication.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Defining gRPC protobuf contracts"
    - "Implementing service-to-service communication"
    - "Generating Go code from protobuf"
---

## When to Use

- Defining communication contracts between microservices
- Implementing gRPC server or client in a service
- Generating Go code from protobuf definitions
- Setting up service-to-service calls

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Proto = Contract** | Protobuf files are the single source of truth for inter-service APIs |
| **Proto lives with the server** | Each service owns its `.proto` files in `proto/` |
| **Client wraps gRPC** | Consumer services use a thin client wrapper, never raw gRPC stubs |
| **Domain doesn't know gRPC** | Domain ports define interfaces; gRPC is an infrastructure adapter |
| **Errors map to domain** | gRPC status codes translate to domain errors at the client boundary |

## Protobuf Definition

> **Reference:** [assets/caregiver.proto](assets/caregiver.proto)

## Code Generation

```bash
# Install tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate from service directory
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/*.proto
```

Add a Makefile target:

> **Reference:** [assets/Makefile](assets/Makefile)

## gRPC Server Implementation

> **Reference:** [assets/grpc_handler.go](assets/grpc_handler.go)

## gRPC Client Wrapper (Consumer Side)

> **Reference:** [assets/grpc_client.go](assets/grpc_client.go)

## Server Setup

> **Reference:** [assets/grpc_server.go](assets/grpc_server.go)

## Directory Layout

```
api/caregiver/
  proto/
    caregiver.proto                    # Source of truth
    caregiverv1/
      caregiver.pb.go                  # Generated — DO NOT EDIT
      caregiver_grpc.pb.go             # Generated — DO NOT EDIT
  internal/
    caregiver/
      infrastructure/
        handler/
          grpc.go                      # gRPC server implementation
```

## Commands

```bash
# Install protoc (macOS)
brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate code
make proto

# Test with grpcurl
brew install grpcurl
grpcurl -plaintext localhost:9090 list
grpcurl -plaintext -d '{"caregiver_id": "abc123"}' localhost:9090 caregiver.v1.CaregiverService/GetCaregiver
```

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Import proto stubs in domain layer | Domain defines its own interfaces; infra maps proto ↔ domain |
| Edit generated `.pb.go` files | Edit `.proto` and regenerate |
| Use raw gRPC stubs in application layer | Wrap in a client that implements a domain port |
| Share proto files via copy-paste | Use a shared proto repo or git submodule if needed |
| Skip error code mapping | Map gRPC codes to domain errors at client boundary |
