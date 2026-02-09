---
name: latex-version-control
description: >
  Add and maintain version control tables in LaTeX corporate documents.
  Trigger: When creating new LaTeX documents, updating document versions, or reviewing document history.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [docs/latex]
  auto_invoke:
    - "Creating new LaTeX documents"
    - "Updating version of an existing LaTeX document"
    - "Reviewing document change history"
---

## When to Use

- Creating any new LaTeX document (add initial version entry)
- Updating an existing document with meaningful changes
- Reviewing or auditing document history

---

## Dependencies

This skill works alongside:

| Skill | Purpose |
|-------|---------|
| **`latex-corporate-docs`** | Corporate template and document structure |

---

## Version Control Table

Every LaTeX document MUST include a version control table immediately after `\maketitle` and `\thispagestyle{fancy}` (if present), BEFORE any content (confidentiality notices, resumen ejecutivo, etc.).

### LaTeX Macro

The `\versiontable` command is defined in the document preamble. It uses the corporate color palette for consistent styling.

> Source: [`assets/version_macros.tex`](assets/version_macros.tex)

### Usage

Place the macro definition in the preamble (after custom macros section), then call it right after `\maketitle`:

> Example: [`assets/version_usage.tex`](assets/version_usage.tex)

### Critical Rule: Always Ask for Author

**Before adding or updating a version entry, the agent MUST ask the user who made the changes.** Never assume "Equipo Bastet" or any generic author. The version table tracks individual accountability, so the author field must reflect the real person.

Example prompt:
> "¿Quién es el autor de estos cambios para la tabla de versiones?"

### Rules

| Rule | Detail |
|------|--------|
| **Placement** | Immediately after `\maketitle` (and `\thispagestyle{fancy}` if present), before any other content |
| **Ordering** | Newest version LAST (chronological, top-to-bottom) |
| **Version format** | `MAJOR.MINOR` — increment MAJOR for structural changes, MINOR for content updates |
| **Date format** | `Mes Año` (e.g., `Enero 2026`) — matches `\date{}` convention |
| **Author** | The specific person who made the changes — **ALWAYS ask the user**, never assume |
| **Changes** | Brief description of what changed (1 sentence) |
| **Initial version** | Always `1.0` with description `Versión inicial del documento` |

### When to Increment Version

| Change Type | Version Bump | Example |
|-------------|-------------|---------|
| New document | 1.0 | Initial creation |
| Content updates, corrections | +0.1 | Fix typos, update figures |
| New sections added | +0.1 | Add appendix, new analysis |
| Major restructuring | +1.0 | Rewrite, change scope |
| Data refresh (economics) | +0.1 | Update CLP values, exchange rates |

---

## Checklist

- [ ] `\versiontable` and `\versionrow` macros defined in preamble
- [ ] Version table placed immediately after `\maketitle`
- [ ] At least one version entry (1.0 — initial)
- [ ] Versions in chronological order (oldest first)
- [ ] Each entry has version, date, author, and change description
