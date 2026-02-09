# {Service Name} Service — AGENTS.md

## Overview

{One paragraph: what this service does, its bounded context, key responsibilities.}

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | {language} |
| Framework | {framework} |
| Database | {database} |
| {other} | {technology} |

---

## Architecture

{Brief note on architecture pattern, reference parent if applicable.}

```
internal/
  {domain}/
    domain/          # {What's in domain layer}
    application/     # {What's in application layer}
    infrastructure/
      handler/       # {HTTP/gRPC handlers}
      repository/    # {Persistence implementations}
  shared/            # {Cross-domain concerns}
```

---

## Domain Model

### Entities

- **{Entity}**: `{field1}`, `{field2}`, `{field3}`

### Value Objects

- **{ValueObject}**: {Description}

### Ports (Interfaces)

| Port | Responsibility |
|------|---------------|
| `{PortName}` | {What it does} |

---

## API Endpoints

Base path: `{base-path}`

| Method | Path | Description | Auth Required |
|--------|------|-------------|---------------|
| {METHOD} | `{path}` | {description} | {Yes/No} |

---

## Error Catalog

### Domain Errors

> Sentinel errors defined in `domain/error.go`.

| Error | Description |
|-------|-------------|
| `{ErrName}` | {description} |

### HTTP Error Mapping

> How domain errors map to HTTP responses in the handler layer.

| Domain Error | HTTP Status | Error Code | Message |
|-------------|-------------|------------|---------|
| `{ErrName}` | {status} | `{CODE}` | {message} |
| _(unhandled)_ | 500 | `INTERNAL_ERROR` | An unexpected error occurred |

---

## Database

**Provider**: {database provider}

### Migrations

| Migration | Table | Purpose |
|-----------|-------|---------|
| {number} | `{table}` | {purpose} |

---

## Configuration (Environment Variables)

| Variable | Default | Description |
|----------|---------|-------------|
| `{VAR_NAME}` | `{default}` | {description} |

---

## Running

```bash
{command to run}          # {description}
{command to test}         # {description}
```

---

## What's NOT Implemented (Future Work)

- {Feature 1}
- {Feature 2}

---

## Skills Reference

> Skills relevant to working on this service.

| Skill | Use For |
|-------|---------|
| `{skill-name}` | {When to use} |
