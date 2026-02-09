---
name: latex-corporate-docs
description: >
  Generate LaTeX corporate documents following project writing conventions and corporate color palette.
  Trigger: When creating LaTeX documents, reports, proposals, or any formal PDF deliverable.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [docs/latex]
  auto_invoke:
    - "Creating LaTeX documents or reports"
    - "Generating PDF deliverables for the project"
    - "Writing formal proposals or analyses in LaTeX"
---

## When to Use

- Creating any LaTeX document for the project
- Generating PDF reports, proposals, budgets, or analyses
- Formatting formal deliverables with corporate branding

---

## Dependencies — MUST Read First

Before writing any LaTeX document, the agent MUST load these skills:

| Skill | Purpose | Path |
|-------|---------|------|
| **`project-docs`** | Writing conventions, language, currency (CLP), doc structure | [`../project-docs/SKILL.md`](../project-docs/SKILL.md) |
| **`corporate-colors`** | Catppuccin color palette, light/dark mode definitions | [`../corporate-colors/SKILL.md`](../corporate-colors/SKILL.md) |
| **`latex-version-control`** | Version control table in every document | [`../latex-version-control/SKILL.md`](../latex-version-control/SKILL.md) |

All rules from `project-docs` apply: Spanish (Chile), technical register, medium complexity, CLP currency, web search for economic data, and `docs/latex/{domain}/` file placement.

---

## File Placement

LaTeX files MUST follow the `project-docs` directory convention:

```
docs/latex/
├── finance/
│   └── budget-2025.tex
├── project/
│   └── proposal.tex
├── processes/
│   └── ci-cd-pipeline.tex
└── ...
```

Output PDFs go alongside their `.tex` source. Build artifacts (`.aux`, `.log`, `.out`) should be gitignored.

---

## Critical Pattern: Paragraph Annotation Comments

**Every paragraph** in the document body MUST have a short LaTeX comment on the line immediately above it. This comment describes in very few words what the paragraph will discuss.

Purpose: Guarantee **global coherence and cohesion** — the comment acts as an outline node, making it possible to review the document's logical flow by reading only comments.

### Rules

| Rule | Detail |
|------|--------|
| **Placement** | Line immediately above the paragraph, no blank line between comment and text |
| **Format** | `% <topic in 3-8 words>` |
| **Language** | Same as document (Spanish), unless a technical English term is clearer |
| **Coverage** | Every `\par`, every paragraph after a blank line, every `\item` body |
| **Exceptions** | Titles, captions, and single-line table cells do not need annotation |

### Example

```latex
\section{Análisis de Costos}

% Contexto general del análisis de costos
El presente análisis detalla los costos asociados al desarrollo del módulo
de autenticación, considerando recursos humanos, infraestructura y licencias
de software.

% Desglose de costos por categoría
Los costos se dividen en tres categorías principales: desarrollo (\$1.800.000 CLP),
infraestructura cloud (\$450.000 CLP/mes) y licencias de terceros (\$250.000 CLP/año).
```

### Anti-patterns

```latex
% BAD: Comentario demasiado vago
% Costos
El presente análisis detalla los costos...

% BAD: Comentario separado del párrafo por línea en blanco
% Desglose de costos por categoría

Los costos se dividen en tres categorías...
```

---

## Corporate Colors in LaTeX

Use the Catppuccin palette from `corporate-colors`. Define colors in the preamble using `xcolor`:

### Light Mode (Latte) — Default for Print

> See [assets/colors-latte.tex](assets/colors-latte.tex)

### Dark Mode (Mocha) — For Screen/Digital PDFs

> See [assets/colors-mocha.tex](assets/colors-mocha.tex)

### Color Usage Convention

| Element | Color | Example |
|---------|-------|---------|
| Section titles | `primary` | `\color{primary}\section{...}` |
| Hyperlinks | `primary` | `\hypersetup{colorlinks, linkcolor=primary, urlcolor=primary}` |
| Highlighted boxes | `info` bg, `textPrimary` fg | `\colorbox{info!10}{...}` |
| Warning callouts | `warning` | `\textcolor{warning}{Advertencia:}` |
| Error/critical | `error` | `\textcolor{error}{Crítico:}` |
| Success indicators | `success` | `\textcolor{success}{\checkmark}` |
| Table headers | `bgMantle` bg | `\rowcolor{bgMantle}` |
| Table rules | `border` | `\arrayrulecolor{border}` |
| Captions | `textSecondary` | via caption package styling |

---

## Document Template Structure

Every LaTeX document MUST follow the corporate template skeleton.

> See [assets/corporate-template.tex](assets/corporate-template.tex)

---

## Paragraph Comment Workflow

When writing a document, follow this process:

1. **Outline first** — Draft the full section structure with only comments (no prose)
2. **Review outline** — Read comments top to bottom to verify logical flow
3. **Write prose** — Fill in each paragraph below its comment
4. **Verify coherence** — Re-read only comments to confirm the document still flows

```latex
% Example: outline-first approach

\section{Plan de Implementación}

% Resumen del enfoque de implementación en 3 fases
...

% Fase 1: diseño de arquitectura y prototipos (semanas 1-4)
...

% Fase 2: desarrollo iterativo con sprints de 2 semanas (semanas 5-12)
...
```

---

## Currency and Economics in LaTeX

Follow `project-docs` rules. CLP currency macros:

> See [assets/currency-macros.tex](assets/currency-macros.tex)

Usage example:

```latex
% Costo mensual del hosting
El hosting tiene un costo de \clp{45.000}/mes.

% Conversión desde USD con tipo de cambio
El servicio SaaS cuesta \usdtoclp{49}{45.000}{920}, enero 2026.
```

---

## Commands

```bash
# Compile LaTeX document
pdflatex -output-directory=docs/latex/{domain} docs/latex/{domain}/{document}.tex

# Full build with references
cd docs/latex/{domain} && pdflatex {document}.tex && bibtex {document} && pdflatex {document}.tex && pdflatex {document}.tex

# Clean build artifacts
find docs/latex/ -name "*.aux" -o -name "*.log" -o -name "*.out" -o -name "*.toc" -o -name "*.bbl" -o -name "*.blg" | xargs rm -f

# Watch and rebuild (requires latexmk)
latexmk -pdf -pvc -outdir=docs/latex/{domain} docs/latex/{domain}/{document}.tex
```

---

## Checklist for New LaTeX Documents

- [ ] `project-docs` and `corporate-colors` skills loaded
- [ ] File placed at `docs/latex/{domain}/{name}.tex`
- [ ] Catppuccin color definitions in preamble
- [ ] Every body paragraph has a short comment above it (3-8 words, no blank line between)
- [ ] Comments alone read as a coherent outline
- [ ] Spanish (Chile), technical register, medium complexity
- [ ] All monetary values in CLP, foreign currency converted with rate and date
- [ ] Standard sections: resumen ejecutivo, contexto, contenido, conclusiones
- [ ] Tables and figures captioned
- [ ] Hyperlinks use corporate `primary` color
- [ ] Version control table present (see `latex-version-control` skill)
- [ ] Build artifacts in `.gitignore`

---

## Resources

- **Templates**: See [assets/](assets/) for `corporate-template.tex`, color definitions, and currency macros
- **Color palette**: See [`corporate-colors`](../corporate-colors/SKILL.md) for full Catppuccin mappings
- **Writing conventions**: See [`project-docs`](../project-docs/SKILL.md) for language, currency, and structure rules
