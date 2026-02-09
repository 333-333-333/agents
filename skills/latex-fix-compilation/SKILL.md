---
name: latex-fix-compilation
description: >
  Diagnose and fix LaTeX compilation errors using CI error logs committed to the repo.
  Trigger: When a LaTeX document fails to compile in CI, when .errlog files exist, or when user asks to fix a broken LaTeX build.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [docs/latex]
  auto_invoke:
    - "Fixing LaTeX compilation errors from CI"
    - "Diagnosing .errlog files in docs/latex/.error-logs/"
    - "Resolving broken LaTeX builds"
---

## When to Use

- A LaTeX document failed to compile in the CI pipeline
- An `.errlog` file exists in `docs/latex/.error-logs/`
- User asks to fix a LaTeX compilation error specifying a document name

---

## How It Works

The CI pipeline (`compile-docs.yml`) compiles each LaTeX document in a separate parallel job. When compilation **fails**, the pipeline:

1. Saves the LaTeX `.log` as `docs/latex/.error-logs/{document-name}.errlog`
2. Uploads it as a GitHub Actions artifact (`errlog-{name}`)
3. Commits the `.errlog` to the repo so it's available locally

---

## Error Log Location

```
docs/latex/.error-logs/
├── business-case.errlog        # failed compilation log for business-case.tex
├── interview-owners.errlog     # failed compilation log for interview-owners.tex
└── ...
```

File naming convention: `{document-name}.errlog` maps to `docs/latex/**/{document-name}.tex`.

---

## Resolution Workflow

Follow these steps **in order**. Do NOT skip steps.

### Step 1 — Pull latest changes

```bash
git pull origin main
```

The error logs are committed by CI, so you need the latest state.

### Step 2 — Identify the failing document

```bash
ls docs/latex/.error-logs/
```

Each `.errlog` file corresponds to a document that failed. The filename (without `.errlog`) is the document name.

If the user specified a document, look for `{document-name}.errlog`.

### Step 3 — Read and analyze the error log

```bash
# Read the full error log
cat docs/latex/.error-logs/{document-name}.errlog
```

Focus on these patterns in the log:

| Pattern | Meaning | Common Fix |
|---------|---------|------------|
| `! LaTeX Error:` | LaTeX-level error | Read the message, usually missing package or bad command |
| `! Undefined control sequence` | Unknown command | Check spelling, add missing `\usepackage{}` |
| `! Missing $ inserted` | Math mode issue | Wrap in `$...$` or escape special chars (`\_`, `\&`, `\%`) |
| `! File ... not found` | Missing input file | Check `\input{}` / `\include{}` paths are relative to root |
| `! Emergency stop` | Fatal — look at lines ABOVE this | The real error is earlier in the log |
| `Runaway argument` | Unclosed brace or environment | Find the mismatched `{` or `\begin{}` without `\end{}` |
| `! Package ... Error:` | Package-specific error | Read the package docs or check options |
| `! I can't find file` | Wrong path in `\input` | Paths must be relative to the compilation working directory (repo root) |

### Step 4 — Locate and fix the source `.tex` file

Find the root document:

```bash
# Find the .tex file for this document
find docs/latex/ -name "{document-name}.tex" -type f
```

Read the file, find the line referenced in the error log, and fix the issue.

**Common fixes:**

1. **Missing package** → Add `\usepackage{package-name}` to preamble AND add it to `extra_packages` in `.github/workflows/compile-docs.yml` if it's not a standard package.

2. **Bad \input path** → Paths in `\input{}` must be relative to the repo root (the `working_directory` in the CI action), e.g. `\input{docs/latex/research/owners/preamble}` — NOT `\input{owners/preamble}`.

3. **Special characters** → Escape: `\_` `\&` `\%` `\#` `\$` `\{` `\}`

4. **Encoding issues** → Ensure `\usepackage[utf8]{inputenc}` is present.

5. **Missing module file** → If `\input{path/to/module}` references a file that doesn't exist, create it or remove the `\input`.

### Step 5 — Delete the error log

After fixing, remove the `.errlog` file so it doesn't linger:

```bash
rm docs/latex/.error-logs/{document-name}.errlog
```

If the directory is now empty, remove it too:

```bash
rmdir docs/latex/.error-logs/ 2>/dev/null || true
```

### Step 6 — Commit and push to trigger recompilation

```bash
git add docs/latex/
git commit -m "fix(docs): resolve compilation error in {document-name}"
git push
```

The push will trigger the CI pipeline again, which will attempt to recompile the fixed document. The pipeline only triggers on changes to `docs/latex/**/*.tex`.

### Step 7 — Verify

After pushing, check the CI pipeline status:

```bash
gh run list --workflow=compile-docs.yml --limit=1
```

If it fails again, repeat from Step 2.

---

## Critical Rules

1. **ALWAYS read the `.errlog` first** — Don't guess. The log tells you exactly what's wrong.
2. **Fix the `.tex` source, not the log** — The log is read-only diagnostic output.
3. **Delete the `.errlog` after fixing** — Stale logs cause confusion.
4. **One fix per commit** — If multiple documents failed, fix and commit each separately so CI can recompile them independently.
5. **Check `extra_packages` in the workflow** — If you add a new `\usepackage{}`, make sure the package is listed in `.github/workflows/compile-docs.yml` under `extra_packages`.
6. **Paths are relative to repo root** — The CI compiles with `working_directory: .` (repo root). All `\input{}` paths must work from there.

---

## Commands Reference

```bash
# Check for error logs
ls docs/latex/.error-logs/

# Read a specific error log
cat docs/latex/.error-logs/{name}.errlog

# Find the source .tex file
find docs/latex/ -name "{name}.tex"

# Clean up after fix
rm docs/latex/.error-logs/{name}.errlog

# Check CI status
gh run list --workflow=compile-docs.yml --limit=3

# Watch a running CI job
gh run watch
```
