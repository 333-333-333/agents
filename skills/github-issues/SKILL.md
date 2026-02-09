---
name: github-issues
description: >
  Create GitHub issues in Spanish following project conventions: structured sections, conventional title prefixes, and consistent labeling.
  Trigger: When creating GitHub issues, reporting bugs, or proposing features via gh CLI.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [root]
  auto_invoke:
    - "Creating GitHub issues"
    - "Reporting bugs via gh CLI"
    - "Proposing features as issues"
---

## When to Use

- Creating a new GitHub issue (feature, bug, chore, etc.)
- Reporting a bug found during development or CI
- Proposing a feature or enhancement for discussion

---

## Critical Patterns

### Language

All issue titles and bodies MUST be written in **Spanish (Chile)**. Code references, file paths, and technical terms may remain in English.

### Title Format

Issue titles follow conventional-commit style prefixes:

| Type | When | Example |
|------|------|---------|
| `feat:` | New functionality | `feat: publicar artefactos de contenedor desde main` |
| `feat(scope):` | Scoped feature | `feat(ci): pipeline de despliegue a Cloud Run` |
| `fix:` | Bug report | `fix: webhook de Discord falla con error 'sender is required'` |
| `chore:` | Maintenance | `chore: actualizar dependencias de Node` |
| `docs:` | Documentation | `docs: agregar guía de contribución` |
| `refactor:` | Code improvement | `refactor: extraer lógica de validación` |

### Body Structure

Use the appropriate template based on issue type:

- **Features / enhancements**: Contexto, Propuesta, Alcance, Criterios de aceptación
- **Bugs**: Descripción, Causa raíz, Pasos para reproducir, Solución propuesta, Referencias

> See [assets/issue-template-feat.md](assets/issue-template-feat.md) for the feature template.

> See [assets/issue-template-bug.md](assets/issue-template-bug.md) for the bug template.

### Labels

Apply the most appropriate label from the repository's existing labels:

| Label | When |
|-------|------|
| `enhancement` | `feat:` issues |
| `bug` | `fix:` issues (bugs) |
| `documentation` | `docs:` issues |

### Checklist Before Creating

- Title uses conventional-commit prefix
- Body is in Spanish
- Sections follow the appropriate template
- Label is applied
- Acceptance criteria are verifiable (checkboxes)

---

## Commands

```bash
gh issue create --title "feat: ..." --label enhancement --body "$(cat <<'EOF' ... EOF)"
gh issue create --title "fix: ..." --label bug --body "$(cat <<'EOF' ... EOF)"
gh issue list --limit 10  # review existing issues for style
```

---

## Resources

- **Templates**: See [assets/](assets/) for issue body templates (feat, bug)
