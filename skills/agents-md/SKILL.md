---
name: agents-md
description: >
  Create and maintain AGENTS.md files for services, components, and subagents with scoped skill references and hierarchical context.
  Trigger: When creating AGENTS.md for a new service or component, updating existing AGENTS.md, or defining subagent context boundaries.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [root]
  auto_invoke:
    - "Creating AGENTS.md for a new service or component"
    - "Updating an existing AGENTS.md file"
    - "Defining context boundaries for subagents"
    - "Adding scoped skill references to AGENTS.md"
---

## When to Use

- Creating an `AGENTS.md` for a new microservice, app, or infrastructure component
- Establishing context boundaries so agents/subagents understand their scope
- Referencing skills that are relevant to a specific component
- Defining what an agent SHOULD and SHOULD NOT know about

## Core Concept: Agent Hierarchy

```
AGENTS.md (root)                    ← Global norms, skill registry, project overview
├── api/AGENTS.md                   ← Backend architecture, shared API conventions
│   ├── api/auth/AGENTS.md          ← Auth service: domain model, endpoints, config
│   ├── api/booking/AGENTS.md       ← Booking service context
│   └── api/payment/AGENTS.md       ← Payment service context
├── mobile/AGENTS.md                ← Flutter mobile conventions
├── web/AGENTS.md                   ← Flutter web conventions
└── infra/AGENTS.md                 ← Terraform / IaC conventions
```

### Override Rule

When guidance conflicts, the **most specific** AGENTS.md wins:

```
root/AGENTS.md  <  api/AGENTS.md  <  api/auth/AGENTS.md
(general)          (backend)          (service-specific)
```

## AGENTS.md Levels

### Level 1 — Root (`./AGENTS.md`)

The root AGENTS.md is the **single source of truth** for:
- Project overview and tech stack map
- Skills registry (synced via `skill-sync`)
- Cross-cutting conventions (commits, PRs, code quality)
- Development setup

> Already exists in this repo. Do NOT recreate. Use `skill-sync` to update skill tables.

### Level 2 — Component (`{component}/AGENTS.md`)

Component-level files define **architectural context** for a major area:

| Component | Defines |
|-----------|---------|
| `api/` | Microservice architecture, communication patterns, shared conventions |
| `mobile/` | Flutter patterns, state management, navigation |
| `web/` | Web-specific conventions, SSR, routing |
| `infra/` | IaC conventions, environments, deployment |
| `docs/` | Documentation standards, LaTeX compilation |

### Level 3 — Service/Module (`{component}/{service}/AGENTS.md`)

The most specific level. Defines **everything an agent needs** to work on this service autonomously:

- What the service does (bounded context)
- Domain model (entities, value objects, ports)
- API surface (endpoints, request/response)
- Database schema and migrations
- Configuration and environment variables
- How to run and test locally
- What's NOT implemented (future work)
- Scoped skill references

## Critical Patterns

### 1. Language

ALL AGENTS.md content MUST be written in **English**. This includes section headers, descriptions, table content, and comments. Code examples naturally remain in English.

### 2. Scoped Skill References

Each AGENTS.md can reference skills that are relevant to its scope. This tells agents which skills to load when working in that directory.

```markdown
## Skills Reference

| Skill | Use For |
|-------|---------|
| `go-gin-handlers` | HTTP handlers, middleware, validation |
| `go-repository-pattern` | Persistence, PostgreSQL, migrations |
| `go-tdd` | Tests, mocks, table-driven tests |
```

**Rules for scoped references:**
- Only reference skills that are RELEVANT to this component
- A service AGENTS.md should reference implementation skills, not meta-skills
- Component AGENTS.md can reference broader architectural skills
- Root AGENTS.md contains the FULL registry (managed by `skill-sync`)

### 3. Required Sections by Level

#### Level 2 — Component AGENTS.md

| Section | Required | Purpose |
|---------|----------|---------|
| Overview | Yes | What this component area is |
| Architecture | Yes | Patterns, layer rules, diagrams |
| Service/Module Map | Yes | List of children with descriptions |
| Shared Conventions | Yes | Response formats, naming, error codes |
| Communication Patterns | If applicable | REST, gRPC, messaging between services |
| Development Workflow | Yes | How to run, test, add new services |
| Skills Reference | Yes | Scoped skills for this component |

#### Level 3 — Service/Module AGENTS.md

| Section | Required | Purpose |
|---------|----------|---------|
| Overview | Yes | What this service does, bounded context |
| Tech Stack | Yes | Technologies used (table format) |
| Architecture | Yes | Directory structure with annotations |
| Domain Model | Yes | Entities, value objects, ports |
| API Endpoints | If HTTP | Method, path, description, auth |
| Error Catalog | If HTTP | Domain errors mapped to HTTP status codes |
| Database | If persistent | Provider, schema, migrations |
| Configuration | Yes | Env vars with defaults and descriptions |
| Running | Yes | Commands to run, test, build |
| What's NOT Implemented | Recommended | Future work, missing features |
| Skills Reference | Optional | Service-specific skill references |

### Error Catalog Convention

Every service with HTTP endpoints MUST document its error catalog. This enables agents to:
- Implement correct error handling in handlers
- Return consistent error responses across services
- Map domain errors to the right HTTP status codes

The catalog has two parts:

**1. Domain Errors** — internal sentinel errors defined in `domain/error.go`:

```markdown
| Error | Description |
|-------|-------------|
| `ErrUserNotFound` | User does not exist |
| `ErrEmailTaken` | Email already registered |
```

**2. HTTP Error Mapping** — how domain errors translate to HTTP responses:

```markdown
| Domain Error | HTTP Status | Error Code | Message |
|-------------|-------------|------------|---------|
| `ErrInvalidEmail` | 400 | `VALIDATION_ERROR` | Invalid email format |
| `ErrEmailTaken` | 409 | `CONFLICT` | Email already registered |
| `ErrInvalidCredentials` | 401 | `UNAUTHORIZED` | Invalid credentials |
| _(unhandled)_ | 500 | `INTERNAL_ERROR` | An unexpected error occurred |
```

Error codes follow the shared conventions in the parent `api/AGENTS.md`. Domain errors are service-specific.

### 4. Do NOT Duplicate Parent Context

Each level provides INCREMENTAL context. Don't repeat what the parent already says.

```
❌ api/auth/AGENTS.md repeating the full clean architecture explanation
✅ api/auth/AGENTS.md saying "Follows clean architecture (see api/AGENTS.md)"
   then showing the SPECIFIC directory structure for auth
```

### 5. Keep It Scannable

Agents parse AGENTS.md for quick context. Optimize for scanning:
- Use **tables** for structured data (endpoints, env vars, migrations)
- Use **code blocks** for directory trees and commands
- Use **bold** for key terms on first mention
- Keep descriptions to ONE line where possible
- Use `---` dividers between major sections

## Templates

Use the templates in `assets/` as starting points:

- **[component-agents.md](assets/component-agents.md)** — Level 2 template for `{component}/AGENTS.md`
- **[service-agents.md](assets/service-agents.md)** — Level 3 template for `{component}/{service}/AGENTS.md`

## Decision Tree: Do I Need an AGENTS.md?

```
Is this a new microservice?                          → YES, Level 3
Is this a new major component (mobile/, web/)?       → YES, Level 2
Is this a subdirectory within a service?             → NO, parent covers it
Is this a shared library or package?                 → MAYBE, if complex enough
Does the directory have unique conventions?           → YES
Does it just follow parent conventions?              → NO
```

## Commands

```bash
# Create AGENTS.md for a new service
cp skills/agents-md/assets/service-agents.md api/{service}/AGENTS.md
# Then fill in service-specific details

# Create AGENTS.md for a new component
cp skills/agents-md/assets/component-agents.md {component}/AGENTS.md
# Then fill in component-specific details

# After creating, sync skills if you added skill references
./skills/skill-sync/assets/sync.sh
```

## Anti-Patterns

| Don't | Do |
|-------|-----|
| Write AGENTS.md in Spanish or mixed languages | Write ALL content in English |
| Duplicate parent context verbatim | Reference parent, add incremental details |
| List every skill in every AGENTS.md | Only reference skills relevant to scope |
| Write paragraphs of explanation | Use tables, code blocks, bullet points |
| Include implementation code | Show directory structures and interfaces |
| Describe how the code works line by line | Describe WHAT the service does and its boundaries |
| Leave AGENTS.md stale after major changes | Update AGENTS.md when architecture changes |

## Resources

- **Templates**: See [assets/](assets/) for component and service AGENTS.md templates
