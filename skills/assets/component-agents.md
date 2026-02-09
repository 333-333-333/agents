# {Component Name} — AGENTS.md

## Overview

{One paragraph: what this component area is and its role in the platform.}

---

## Architecture

{Architecture pattern, layer rules, dependency direction. Include ASCII diagram if helpful.}

```
{component}/
  {service-1}/          # {Brief description}
  {service-2}/          # {Brief description}
  shared/               # {Cross-cutting concerns if applicable}
```

---

## Service Map

| Service | Description | Port |
|---------|-------------|------|
| `{service-1}` | {What it does} | {Local port} |
| `{service-2}` | {What it does} | {Local port} |

---

## Shared Conventions

{Conventions that apply to ALL services/modules within this component.}

### Response Format

{Standard response envelope, error codes, etc.}

### Naming

| Element | Convention | Example |
|---------|-----------|---------|
| {element} | {convention} | {example} |

---

## Communication Patterns

| Direction | Protocol | When |
|-----------|----------|------|
| {source} → {target} | {protocol} | {when} |

---

## Development Workflow

```bash
# Run all services
{command}

# Run single service
{command}

# Run tests
{command}
```

---

## Skills Reference

> Skills relevant to working within this component.

| Skill | Use For |
|-------|---------|
| `{skill-name}` | {When to use} |

### Auto-invoke Skills

| Action | Skill |
|--------|-------|
| {Action description} | `{skill-name}` |
