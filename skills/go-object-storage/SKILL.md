---
name: go-object-storage
description: >
  Cloud-agnostic object storage with abstract interfaces and implementations for S3, GCS, and MinIO.
  Trigger: When storing files, logs, metrics exports, backups, or any binary objects in cloud storage.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Storing files or objects in cloud storage"
    - "Exporting observability data to object storage"
    - "Configuring S3/GCS/MinIO integration"
---

## When to Use

- Storing user-uploaded files (pet photos, caregiver documents)
- Exporting observability data (logs, metrics, traces) to long-term storage
- Storing backups or reports
- Any file/blob persistence need

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Interface first** | `ObjectStore` interface in shared package, never import SDK in domain |
| **Cloud agnostic** | Same interface for S3, GCS, MinIO — swap via config |
| **Bucket per concern** | Separate buckets: `uploads`, `exports`, `backups` |
| **Signed URLs for clients** | Never expose raw credentials; generate pre-signed URLs |
| **Content-addressed when possible** | Use hash-based keys for immutable objects |

## Interface

> **Reference:** [`assets/storage.go`](assets/storage.go) — `internal/shared/storage/storage.go`

Defines `ObjectStore` interface with `Put`, `Get`, `Delete`, `Exists`, `SignedURL`, and `List` methods, plus `PutOptions` and `ObjectInfo` types.

## S3-Compatible Implementation (S3 / MinIO)

> **Reference:** [`assets/s3.go`](assets/s3.go) — `internal/shared/storage/s3.go`

Uses AWS SDK v2. Supports custom endpoints and path-style access for MinIO compatibility.

## Key Naming Conventions

```
{concern}/{service}/{date}/{filename}

Examples:
  uploads/caregiver/2025-01-15/photo-abc123.jpg
  exports/observability/2025-01-15/traces-14h.jsonl
  exports/observability/2025-01-15/metrics-14h.jsonl
  backups/booking/2025-01-15/booking-db-full.sql.gz
```

## Configuration

> **Reference:** [`assets/config.go`](assets/config.go) — `StorageConfig` + `NewObjectStore` factory

Factory function switches on `Provider` to create the appropriate implementation (`s3`/`minio` or `gcs`).

## Local Development with MinIO

> **Reference:** [`assets/docker-compose-minio.yml`](assets/docker-compose-minio.yml)

```bash
# Config for local dev
STORAGE_PROVIDER=minio
STORAGE_ENDPOINT=http://localhost:9000
STORAGE_REGION=us-east-1
STORAGE_FORCE_PATH_STYLE=true
```

## Commands

```bash
# AWS SDK v2
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3

# MinIO for local dev
docker-compose up minio

# Create buckets via MinIO client
brew install minio/stable/mc
mc alias set local http://localhost:9000 minioadmin minioadmin
mc mb local/bastet-uploads
mc mb local/bastet-exports
mc mb local/bastet-backups
```

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| Import AWS SDK in domain layer | Use `ObjectStore` interface |
| Hardcode bucket names | Configure via environment variables |
| Send files through your API | Generate pre-signed URLs, client uploads directly |
| Store files without organized key structure | Use `{concern}/{service}/{date}/{file}` pattern |
| Single bucket for everything | Separate buckets per concern (uploads, exports, backups) |
