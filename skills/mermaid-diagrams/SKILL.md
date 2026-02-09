---
name: mermaid-diagrams
description: >
  Create Mermaid diagrams following project conventions: domain-based organization,
  filename prefixes by diagram type, and software engineering best practices.
  Trigger: When creating diagrams, flowcharts, sequence diagrams, ER diagrams, or any visual documentation in Mermaid.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [docs/diagrams]
  auto_invoke:
    - "Creating Mermaid diagrams or flowcharts"
    - "Documenting software architecture visually"
    - "Creating sequence, ER, state, or class diagrams"
    - "Adding diagrams under docs/diagrams/"
---

## When to Use

- Creating any visual diagram for the project
- Documenting flows, sequences, data models, or architecture
- Adding `.mmd` files under `docs/diagrams/`

---

## Dependencies

| Skill | Purpose | Path |
|-------|---------|------|
| **`project-docs`** | File organization, `docs/diagrams/` structure, naming conventions | [`../project-docs/SKILL.md`](../project-docs/SKILL.md) |
| **`corporate-colors`** | Catppuccin palette for optional color coding | [`../corporate-colors/SKILL.md`](../corporate-colors/SKILL.md) |

---

## File Placement

Diagrams live in `docs/diagrams/`, organized by **software domain** (not diagram type):

```
docs/diagrams/
├── auth/
│   ├── flow-login.mmd
│   ├── seq-oauth-google.mmd
│   └── er-users.mmd
├── payments/
│   ├── flow-checkout.mmd
│   └── state-payment-lifecycle.mmd
├── booking/
│   └── flow-reservation.mmd
└── architecture/          # Cross-cutting, system-level diagrams
    ├── c4-system-context.mmd
    └── deployment-gcp.mmd
```

**Why by domain?** When working on a feature, you want everything related in one place. The diagram type is encoded in the filename prefix so you can still filter across domains.

### Choosing the Domain Subdirectory

```
User authentication, login, OAuth       → auth/
Payment processing, Flow.cl, billing    → payments/
Reservations, scheduling, availability  → booking/
Push notifications, emails, alerts      → notifications/
Pet profiles, caretaker profiles        → profiles/
System-level, deployment, infra, C4     → architecture/
```

---

## Filename Convention

Format: `{type-prefix}-{descriptive-name}.mmd`

| Prefix | Diagram type |
|--------|-------------|
| `flow-` | Flowchart / process flow |
| `seq-` | Sequence diagram |
| `er-` | Entity-relationship |
| `class-` | Class diagram |
| `state-` | State diagram |
| `c4-` | C4 model diagram |
| `gantt-` | Gantt chart |

---

## Critical Patterns

### 1. Title Comments (mandatory)

Every diagram MUST start with two comment lines:

```
%% Short description in Spanish
%% Dominio: {domain} | Tipo: {type}
```

### 2. Node IDs: Descriptive English camelCase

```
✅  userOpen[Usuario abre la app]
❌  A[Usuario abre la app]
```

### 3. Labels in Spanish, IDs in English

```
validatePayment[Validar pago] --> processCharge[Procesar cobro]
```

### 4. Arrow Styles

| Arrow | Meaning |
|-------|---------|
| `-->` | Normal flow |
| `-.->` | Optional / async |
| `==>` | Critical path |
| `--x` | Error / failure |

### 5. Subgraphs for System Boundaries

Use subgraphs to separate: mobile, API, external services, databases. See [assets/example-subgraphs.mmd](assets/example-subgraphs.mmd).

### 6. Keep Diagrams Focused

- **One diagram = one concept**
- Flowcharts: max ~15 nodes, split if larger
- Sequence diagrams: max ~6 participants, split by interaction boundary

### 7. Color Coding (optional)

Use Catppuccin Latte colors from `corporate-colors` skill. See [assets/example-styles.mmd](assets/example-styles.mmd).

---

## Examples

Reference examples for each diagram type in [assets/](assets/):

| File | Type | Shows |
|------|------|-------|
| [`example-flow.mmd`](assets/example-flow.mmd) | Flowchart | User journey with decisions |
| [`example-seq.mmd`](assets/example-seq.mmd) | Sequence | API interaction with Flow.cl |
| [`example-er.mmd`](assets/example-er.mmd) | ER | Core data model |
| [`example-state.mmd`](assets/example-state.mmd) | State | Booking lifecycle |
| [`example-subgraphs.mmd`](assets/example-subgraphs.mmd) | Flowchart | Bounded contexts |
| [`example-styles.mmd`](assets/example-styles.mmd) | Flowchart | Catppuccin color coding |

---

## Commands

```bash
# Create a new diagram domain directory
mkdir -p docs/diagrams/{domain}

# Preview (requires: npm install -g @mermaid-js/mermaid-cli)
mmdc -i docs/diagrams/{domain}/{diagram}.mmd -o docs/diagrams/{domain}/{diagram}.svg

# Batch render all diagrams to SVG
find docs/diagrams -name '*.mmd' -exec sh -c 'mmdc -i "$1" -o "${1%.mmd}.svg"' _ {} \;

# Validate syntax (dry run)
mmdc -i docs/diagrams/{domain}/{diagram}.mmd -o /dev/null
```

---

## Checklist

- [ ] File at `docs/diagrams/{domain}/{prefix}-{name}.mmd`
- [ ] Title comment on lines 1-2
- [ ] Descriptive English camelCase node IDs
- [ ] Spanish labels
- [ ] Consistent arrow styles
- [ ] Subgraphs for system boundaries
- [ ] Focused (not overloaded)
- [ ] Renders correctly with `mmdc` or GitHub/IDE preview
