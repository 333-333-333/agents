---
name: github-actions-nx
description: >
  GitHub Actions CI/CD conventions for Nx monorepos: affected-based pipelines, changed-file detection,
  matrix jobs for renderable artifacts, Go service testing, and artifact bundling.
  Trigger: When creating or modifying GitHub Actions workflows, adding CI for new Nx projects,
  or optimizing pipeline execution.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [.github]
  auto_invoke:
    - "Creating GitHub Actions workflows"
    - "Adding CI/CD for a new Nx project"
    - "Adding a new service or library to CI"
    - "Optimizing GitHub Actions pipeline performance"
    - "Adding renderable artifact pipelines (LaTeX, Mermaid, etc.)"
    - "Modifying .github/workflows/ files"
---

## When to Use

Use this skill when:
- Creating or modifying any file under `.github/workflows/`
- Adding a new Nx project that needs CI (service, library, docs)
- Adding a pipeline for files that need rendering/compilation (LaTeX, Mermaid, Protobuf, etc.)
- Reviewing or optimizing CI execution time

---

## Core Principles

1. **Nx drives project detection** — Use `nrwl/nx-set-shas` + `nx show projects --affected` to know WHICH projects changed. Never hardcode path checks with `git diff` alone.
2. **Only run what changed** — Every job MUST have an `if:` condition gated on the affected project output.
3. **Only render what changed** — For renderable artifacts (LaTeX, Mermaid, Protobuf, etc.), use `git diff` to detect the specific FILES that changed, then compile/render only those files via a matrix strategy.
4. **Parallel by default** — Independent projects run in parallel. Use `needs:` only for real dependencies (e.g., bundle depends on compile).
5. **Fail fast is off for artifacts** — Use `fail-fast: false` in matrix strategies so one broken file doesn't block others.

---

## Pipeline Architecture

Every CI workflow follows a **3-layer** structure:

```
Layer 1: Detection        → "affected" job (Nx + git diff)
Layer 2: Execution        → Per-project jobs (test, build, render)
Layer 3: Aggregation      → Bundle jobs (merge artifacts, cleanup)
```

The `affected` job is always the first job and the SINGLE source of truth. All other jobs depend on it via `needs: affected`.

---

## Critical Patterns

### Affected Detection Job

The `affected` job does two things:
1. **Project-level**: Which Nx projects changed? (boolean outputs)
2. **File-level**: Which specific files changed? (JSON array outputs for matrix strategies)

Project outputs follow a naming convention: the Nx project name with the `docs-` prefix stripped. So `docs-latex` becomes `latex`, `docs-diagrams` becomes `diagrams`, and `api-gateway` stays `api-gateway`.

File-level outputs are JSON arrays (e.g., `["docs/latex/project/file.tex"]`) consumed by `fromJson()` in matrix strategies. An empty array `[]` means nothing changed.

> See [assets/affected-job.yml](assets/affected-job.yml) for the full affected detection pattern.

### Adding a New Nx Project to CI

When a new Nx project is created (e.g., `api/booking` with name `api-booking`):

1. Add the project name to the loop in the `affected` job's "Check affected projects" step
2. Add its key to the job's `outputs:` map
3. Create the corresponding execution job with `needs: affected` and an `if:` condition
4. If it produces renderable artifacts, also add file-level detection in the "Detect changed files" step

### Renderable Artifacts Pattern

Files that need compilation or rendering (LaTeX, Mermaid, Protobuf, SVG, etc.) follow the **detect-render-bundle** pattern:

1. **Detect**: In the `affected` job, use `git diff` to find changed source files, resolve to root documents if needed, output as JSON array
2. **Render**: Matrix job that compiles/renders each file individually
3. **Bundle**: Aggregation job that downloads individual artifacts, merges them, and cleans up temporary artifacts

Key rules:
- Matrix source is the JSON array from the `affected` job
- Guard with BOTH project-level AND file-level conditions: `if: needs.affected.outputs.project == 'true' && needs.affected.outputs.changed-files != '[]'`
- Individual artifacts use `retention-days: 1` (temporary)
- Bundle artifact uses `retention-days: 30`
- Bundle job uses `if: always() && needs.render-job.result != 'skipped'`

For files with dependencies (e.g., a LaTeX module included by a root document), the detection step must resolve the changed file to its root compilable document.

> See [assets/renderable-job.yml](assets/renderable-job.yml) for the full render + bundle pattern.

### Go Service Pattern

Go microservices follow a standard 3-step job: test, build binary, build Docker image. Use `actions/setup-go@v5` with `go-version-file` to pin the Go version from the service's `go.mod`.

> See [assets/go-service-job.yml](assets/go-service-job.yml) for the full Go service CI pattern.

---

## Decision Tree

```
New Nx project added?
  ├─ It produces renderable artifacts (PDF, SVG, etc.)
  │    → Use renderable pattern: detect changed files + matrix + bundle
  │    → See assets/renderable-job.yml
  ├─ It's a Go microservice
  │    → Use Go service pattern: test + build + docker
  │    → See assets/go-service-job.yml
  ├─ It's a Flutter app
  │    → Use Flutter pattern: analyze + test + build
  │    → (extend when mobile/ is added)
  └─ It's a generic library
       → Use simple pattern: lint + test
       → Gate with if: needs.affected.outputs.{key} == 'true'
```

---

## Workflow File Conventions

| Convention | Rule |
|------------|------|
| Location | `.github/workflows/ci.yml` (single CI file) |
| Triggers | `push` to `main` + all `pull_request` |
| Permissions | Minimal: `actions: read`, `contents: read` |
| Job naming | Descriptive `name:` field, kebab-case `id` |
| Sections | Separated by `# ====` comment banners with numbered sections |
| Runner | `ubuntu-latest` unless specific container needed |
| Node version | Pin to `20` via `actions/setup-node@v4` |
| Go version | Pin via `go-version-file` pointing to service's `go.mod` |
| Artifacts | Named `{type}-{hash}` for individuals, `bastet-{category}` for bundles |

---

## Assets

| File | Description |
|------|-------------|
| `assets/affected-job.yml` | Affected detection job with project-level and file-level outputs |
| `assets/renderable-job.yml` | Matrix render + bundle pattern for compilable artifacts |
| `assets/go-service-job.yml` | Go microservice test + build + Docker pattern |

---

## Resources

- **Current CI**: See [.github/workflows/ci.yml](../../../.github/workflows/ci.yml) for the live implementation
- **Nx docs**: `nrwl/nx-set-shas` action sets `NX_BASE` and `NX_HEAD` environment variables
