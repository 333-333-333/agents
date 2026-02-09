---
name: go-docker-deploy
description: >
  Multi-stage Dockerfiles, docker-compose for local dev, and cloud-agnostic deployment patterns for Cloud Run, ECS, or Kubernetes.
  Trigger: When containerizing a microservice, setting up local dev environment, or configuring deployment.
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [api]
  auto_invoke:
    - "Containerizing a Go microservice"
    - "Setting up docker-compose for local development"
    - "Deploying microservices to Cloud Run, ECS, or Kubernetes"
---

## When to Use

- Containerizing a Go microservice
- Setting up local development with docker-compose
- Deploying to Cloud Run, ECS, or Kubernetes
- Optimizing Docker image size

## Critical Patterns

| Pattern | Rule |
|---------|------|
| **Multi-stage builds** | Build in Go image, run in distroless/scratch |
| **One service, one Dockerfile** | Each service has its own Dockerfile |
| **docker-compose for dev** | Full local stack: services + Postgres + NATS + MinIO |
| **Cloud agnostic deploy** | Same image deploys to Cloud Run, ECS, or k8s |
| **Health checks** | Every service exposes `/health` |

## Dockerfile (Multi-Stage)

> **Reference:** [`assets/Dockerfile.example`](assets/Dockerfile.example)

## .dockerignore

> **Reference:** [`assets/dockerignore`](assets/dockerignore)

## docker-compose (Full Local Stack)

> **Reference:** [`assets/docker-compose.example.yml`](assets/docker-compose.example.yml)

## Init Script (Multiple Databases)

> **Reference:** [`assets/init-databases.sql`](assets/init-databases.sql)

## Makefile (Service Level)

> **Reference:** [`assets/service.Makefile`](assets/service.Makefile)

## Root Makefile

> **Reference:** [`assets/root.Makefile`](assets/root.Makefile)

## Cloud-Agnostic Deployment

The SAME Docker image works on any platform:

### Cloud Run (GCP)

> **Reference:** [`assets/deploy-cloud-run.sh`](assets/deploy-cloud-run.sh)

### ECS (AWS)

> **Reference:** [`assets/ecs-task.json`](assets/ecs-task.json)

### Kubernetes

> **Reference:** [`assets/k8s-deployment.yaml`](assets/k8s-deployment.yaml)

## Commands

```bash
# Full stack up
docker-compose up -d

# Rebuild single service
docker-compose build booking && docker-compose up -d booking

# View logs
docker-compose logs -f booking

# Run migrations inside container
docker-compose exec booking /migrations/run.sh

# Shell into container (distroless has no shell — use debug variant if needed)
docker-compose exec postgres psql -U postgres -d bastet_booking
```

## Anti-Patterns

| Don't | Do |
|----------|-------|
| Full Go image in production | Distroless or scratch for minimal attack surface |
| Cloud-specific entrypoint scripts | Same Dockerfile everywhere, config via env vars |
| `docker-compose` in production | docker-compose for dev only, use orchestrator in prod |
| Shared database containers between services | Each service gets its own database |
| Hardcoded ports | Configure via environment variables |
