---
name: project-docs
description: >
  Conventions for writing project documents: technical language, medium complexity,
  Chilean peso economics, and structured output under docs/.
  Trigger: When generating project documents, proposals, reports, or any formal deliverable.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [docs/latex]
  auto_invoke:
    - "Writing project proposals or reports"
    - "Generating formal project deliverables"
    - "Creating financial or social impact documents"
    - "Structuring documents under docs/"
---

## When to Use

- Writing project proposals, reports, or formal deliverables
- Creating documents that include financial or social impact data
- Any document that will live under `docs/`
- Establishing tone, language, and formatting for project output

---

## File Organization

The `docs/` directory is split into two top-level areas:

| Directory | Purpose | Organized by |
|-----------|---------|--------------|
| `docs/latex/` | LaTeX documents (reports, proposals, analyses) | Business/project area |
| `docs/diagrams/` | Mermaid diagrams (`.mmd`) | Software domain |

### LaTeX Documents — `docs/latex/`

```
docs/latex/
├── finance/          # Financial analysis, budgets, cost projections
│   ├── budget-2025.tex
│   └── cost-analysis.tex
├── project/          # Project plans, proposals, status reports
│   ├── proposal.tex
│   └── status-report-q1.tex
├── processes/        # Process documentation, workflows, SOPs
│   ├── onboarding.tex
│   └── ci-cd-pipeline.tex
├── legal/            # Contracts, terms, compliance
├── marketing/        # Campaigns, brand guidelines
├── engineering/      # Technical design docs, architecture
└── research/         # Market research, benchmarks, studies
```

#### Choosing the LaTeX Subdirectory

```
Financial analysis, budgets, ROI, costs    → latex/finance/
Project plans, proposals, status updates   → latex/project/
Workflows, SOPs, pipelines, methodologies  → latex/processes/
Contracts, compliance, legal terms         → latex/legal/
Campaigns, positioning, brand materials    → latex/marketing/
Architecture, design docs, technical specs → latex/engineering/
Studies, benchmarks, market analysis       → latex/research/
```

If a document spans multiple domains, place it in the **primary** domain and add a reference note in the secondary.

### Mermaid Diagrams — `docs/diagrams/`

Organized by **software domain** (not diagram type). The diagram type is encoded in the filename prefix.

```
docs/diagrams/
├── auth/
│   ├── flow-login.mmd
│   ├── seq-oauth-google.mmd
│   └── er-users.mmd
├── payments/
│   ├── flow-checkout.mmd
│   ├── seq-flow-cl-webhook.mmd
│   └── state-payment-lifecycle.mmd
├── booking/
│   └── flow-reservation.mmd
└── architecture/          # Cross-cutting, system-level diagrams
    ├── c4-system-context.mmd
    └── deployment-gcp.mmd
```

#### Filename Prefixes for Diagram Type

| Prefix | Diagram type | Example |
|--------|-------------|---------|
| `flow-` | Flowchart / process flow | `flow-login.mmd` |
| `seq-` | Sequence diagram | `seq-oauth-google.mmd` |
| `er-` | Entity-relationship | `er-users.mmd` |
| `class-` | Class diagram | `class-booking-model.mmd` |
| `state-` | State diagram | `state-payment-lifecycle.mmd` |
| `c4-` | C4 model diagram | `c4-system-context.mmd` |
| `gantt-` | Gantt chart | `gantt-mvp-timeline.mmd` |

#### Choosing the Diagram Subdirectory

Create subdirectories matching software domains as they emerge (e.g., `auth/`, `payments/`, `booking/`, `notifications/`). Use `architecture/` for system-wide diagrams that don't belong to a single domain.

### Naming Convention (both LaTeX and Mermaid)

| Rule | Example |
|------|---------|
| Subdirectory in **English** (lowercase) | `finance/`, `auth/`, `payments/` |
| File name in **English** (lowercase, hyphens) | `cost-analysis.tex`, `flow-login.mmd` |
| No spaces, no uppercase, no special characters | `budget-2025.tex` (not `Budget 2025.tex`) |

---

## Writing Conventions

### Language and Tone

| Aspect | Convention |
|--------|-----------|
| **Language** | Spanish (Chile), except technical terms that are standard in English |
| **Register** | Technical-professional, direct, impersonal where possible |
| **Complexity** | Medium — accessible to stakeholders with general technical literacy, not overly academic |
| **Voice** | Prefer impersonal constructions: "Se estima...", "El análisis indica..." over "Yo creo..." |
| **Technical terms** | Use original English terms for established concepts (e.g., _sprint_, _pipeline_, _deploy_, _stakeholder_) with brief explanation on first use if the audience may not know them |

### Examples

```
✅ "El costo estimado del módulo de autenticación es de $2.500.000 CLP,
    considerando 3 sprints de desarrollo."

❌ "Yo pienso que el módulo de autenticación podría llegar a costar
    alrededor de dos millones y medio de pesos más o menos."

✅ "Se implementará un pipeline de CI/CD (integración y despliegue continuo)
    utilizando GitHub Actions."

❌ "Vamos a hacer un sistema automático para subir el código."
```

### Complexity Guidelines

- **DO**: Use precise terminology, quantify claims, cite sources when available
- **DO**: Define acronyms and English terms on first use
- **DO**: Structure with clear sections, tables, and bullet points
- **DON'T**: Use jargon without context or assume deep domain expertise
- **DON'T**: Write walls of text — prefer structured, scannable content
- **DON'T**: Oversimplify to the point of losing technical accuracy

---

## Economics and Currency

### Currency: Chilean Peso (CLP)

ALL monetary values MUST be expressed in **pesos chilenos (CLP)**.

| Rule | Example |
|------|---------|
| Symbol | `$` followed by amount, then `CLP` suffix for clarity |
| Thousands separator | Period (`.`) |
| No decimal places | CLP has no centavos in practice |
| Full format | `$2.500.000 CLP` |
| In tables/shorthand | `$2.500.000` (CLP implied by context) |

### Conversion Requirements

When referencing values originally in foreign currency:

1. **Search for the current exchange rate** using web tools if needed
2. **Show both values**: original and CLP equivalent
3. **State the exchange rate used** and the date

```
El servicio de hosting tiene un costo de USD $49/mes
(~$45.000 CLP/mes al tipo de cambio de $920 CLP/USD, enero 2026).
```

### Social and Economic Metrics

When quantifying social or economic impact:

- Express in CLP, even if the source uses another currency
- Use Chilean economic references where possible (UF, UTM, IPC, sueldo mínimo)
- Reference publicly available data from:
  - Banco Central de Chile
  - INE (Instituto Nacional de Estadísticas)
  - SII (Servicio de Impuestos Internos)
  - CMF (Comisión para el Mercado Financiero)

```
El costo operativo mensual equivale a 2,3 UTM
($143.000 CLP aprox., valor UTM enero 2026: $62.171).
```

### Web Search for Data

When economic data is required and not available locally:

1. **ALWAYS search** for the most current values (UTM, UF, exchange rates, minimum wage)
2. **State the source** and date of the data
3. **Flag stale data** if the document may be read months later

```
Nota: Valores económicos calculados con datos de enero 2026.
Verificar vigencia antes de utilizar para decisiones financieras.
```

---

## Document Structure Standards

Every formal document SHOULD follow this general structure:

1. **Título y metadata** — Title, author, date, version
2. **Resumen ejecutivo** — 1 paragraph summary of the document
3. **Contexto** — Why this document exists, background
4. **Contenido principal** — Organized in numbered sections
5. **Impacto económico** (if applicable) — Costs, benefits, ROI in CLP
6. **Conclusiones** — Key takeaways
7. **Anexos** (if applicable) — Supporting data, references

### Section Numbering

Use hierarchical numbering: `1.`, `1.1.`, `1.1.1.`

### Tables and Figures

- Every table MUST have a caption
- Every figure MUST have a caption and source
- Use consistent formatting across the document

---

## Dependencies

| Skill | Relationship |
|-------|-------------|
| [`corporate-colors`](../corporate-colors/SKILL.md) | Use corporate palette for styled/branded documents |
| `latex-corporate-docs` _(future)_ | LaTeX-specific implementation of these conventions |

---

## Commands

```bash
# Create a new LaTeX document directory
mkdir -p docs/latex/{domain}

# Create a new diagram domain directory
mkdir -p docs/diagrams/{domain}

# Verify document structure
tree docs/ 2>/dev/null || find docs/ -type f | sort
```

---

## Checklist for New Documents

- [ ] LaTeX placed under `docs/latex/{domain}/`, diagrams under `docs/diagrams/{domain}/`
- [ ] Written in Spanish (Chile) with technical-professional register
- [ ] Medium complexity, accessible to general technical audience
- [ ] All monetary values in CLP with proper formatting
- [ ] Foreign currency values converted with exchange rate and date
- [ ] Economic references sourced and dated
- [ ] Document follows standard structure (título, resumen, contexto, contenido, conclusiones)
- [ ] Tables and figures captioned
