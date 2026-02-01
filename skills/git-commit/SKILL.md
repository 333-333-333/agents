---
name: git-commit
description: >
  Create well-structured git commits following conventional commits.
  Trigger: When committing changes, writing commit messages, or staging files.
license: Apache-2.0
metadata:
  author: template
  version: "1.1"
  type: generic
  scope: [root]
  auto_invoke:
    - "Creating git commits"
    - "Writing commit messages"
allowed-tools: Bash, Read, Glob, Grep
---

## When to Use

- Committing changes to the repository
- Writing or reviewing commit messages
- Staging files for commit
- Amending commits (with caution)

---

## Conventional Commits Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

| Type | Description | Example |
|------|-------------|---------|
| `feat` | New feature | `feat: add user authentication` |
| `fix` | Bug fix | `fix: resolve login timeout issue` |
| `docs` | Documentation only | `docs: update API reference` |
| `style` | Code style (formatting, semicolons) | `style: format with prettier` |
| `refactor` | Code change that neither fixes nor adds | `refactor: extract validation logic` |
| `perf` | Performance improvement | `perf: optimize database queries` |
| `test` | Adding or fixing tests | `test: add unit tests for auth` |
| `chore` | Maintenance tasks | `chore: update dependencies` |
| `ci` | CI/CD changes | `ci: add GitHub Actions workflow` |
| `build` | Build system changes | `build: upgrade webpack to v5` |

### Scope (Optional)

Indicates the section of the codebase:

```
feat(auth): add OAuth2 support
fix(api): handle null response
docs(readme): add installation steps
```

---

## Commit Message Guidelines

### Good Commit Messages

```
feat: add password strength indicator

- Show visual feedback for password strength
- Validate minimum requirements
- Display helpful suggestions

Closes #123
```

```
fix: prevent duplicate form submissions

Add debounce to submit button to prevent
multiple API calls when user clicks rapidly.
```

### Bad Commit Messages

```
# Too vague
fix: fix bug

# Not descriptive
update code

# Too long for subject
feat: add new feature that allows users to reset their password by clicking a link sent to their email
```

### Rules

| Rule | Example |
|------|---------|
| Subject line max 50 chars | `feat: add login page` |
| Use imperative mood | `add` not `added` or `adds` |
| No period at end of subject | `fix: resolve issue` not `fix: resolve issue.` |
| Blank line between subject and body | See examples above |
| Body wraps at 72 chars | For readability in terminals |
| Explain what and why, not how | Code shows how |

---

## Workflow

### Before Committing

```bash
# Check status
git status

# Review changes
git diff

# Review staged changes
git diff --staged
```

### Staging Files

```bash
# Stage specific files
git add path/to/file.js

# Stage all changes in directory
git add src/

# Stage all tracked files
git add -u

# Interactive staging (select hunks)
git add -p
```

### Creating the Commit

```bash
# Commit with message
git commit -m "feat: add user profile page"

# Commit with body (opens editor)
git commit

# Commit with multi-line message
git commit -m "feat: add user profile page" -m "Includes avatar upload and bio editing"
```

---

## Atomic Commits — Granularity Rules

Each commit MUST represent **one logical change**. When the user asks to commit, analyze `git diff` and split into multiple commits if the changes span different concerns.

### How to Decide

Ask: "If I needed to revert this, would I want to revert ALL of these changes together?" If the answer is no, split them.

| Signal | Action |
|--------|--------|
| Changes touch different domains/features | Split by domain |
| Mix of infra + feature code | Separate commits |
| New dependency + feature using it | Can be one commit (dependency serves the feature) |
| Tests + implementation they test | Same commit (they're one logical unit) |
| Formatting/linting + feature changes | Separate commits |
| Migration + repository + domain changes | Split if they're independently meaningful |

### Commit Order

When splitting, commits MUST be ordered so the repo **compiles and passes tests at every point**:

1. Infrastructure/config changes first (dependencies, migrations, configs)
2. Domain/core logic second
3. Application/service layer third
4. Transport/handler layer last
5. Documentation/chore changes independently

### Example — What NOT to Do

```bash
# BAD: One giant commit mixing everything
git add -A
git commit -m "feat(auth): add multi-role support, postgres, gateway proxy, docker-compose"
```

### Example — Correct Granularity

```bash
# 1. Domain change
git add api/auth/internal/auth/domain/
git commit -m "feat(auth): refactor user entity from single role to multi-role

- Change Role field to []Role with HasRole/AddRole/RemoveRole
- Add NewRoles validator with dedup and ErrNoRoles
- Update domain tests for multi-role scenarios"

# 2. Infrastructure adapters
git add api/auth/internal/auth/infrastructure/ api/auth/migrations/ api/auth/go.*
git commit -m "feat(auth): add PostgreSQL repository with pgx and sqlc

- Add user_roles migration and sqlc-generated queries
- Implement PostgresUserRepository and PostgresTokenBlacklist
- Add DB_PROVIDER switch in main.go (memory/postgres)"

# 3. Independent service change
git add api/gateway/
git commit -m "feat(gateway): add reverse proxy forwarding to auth service"

# 4. DevOps/tooling
git add docker-compose.yml Makefile skills/
git commit -m "chore: add docker-compose and local dev workflow"
```

### When a Single Commit Is Fine

- All changes serve ONE purpose (e.g., a bug fix touching 5 files)
- A small feature fully contained in one layer
- Pure refactoring with no behavior change

---

## Safety Rules

### Never Do

| Action | Risk |
|--------|------|
| `git commit --amend` after push | Rewrites public history |
| `git push --force` to shared branches | Overwrites others' work |
| Commit secrets/credentials | Security breach |
| Commit large binary files | Bloats repository |

### Amend Safely

Only amend if ALL conditions are met:
1. Commit has NOT been pushed
2. You are the author
3. No one else has based work on it

```bash
# Safe: amend unpushed commit
git commit --amend -m "feat: corrected message"

# Check if pushed
git status  # Shows "Your branch is ahead of..."
```

---

## Files to Never Commit

```gitignore
# Secrets
.env
.env.local
*.pem
credentials.json
secrets.yaml

# IDE
.idea/
.vscode/
*.swp

# Dependencies
node_modules/
vendor/
__pycache__/

# Build artifacts
dist/
build/
*.pyc
```

---

## Commands Reference

```bash
# View recent commits
git log --oneline -10

# View commit details
git show <commit-hash>

# Undo last commit (keep changes staged)
git reset --soft HEAD~1

# Undo last commit (keep changes unstaged)
git reset HEAD~1

# Discard last commit entirely (DANGEROUS)
git reset --hard HEAD~1

# Create commit with sign-off
git commit -s -m "feat: add feature"

# Create commit skipping hooks (use sparingly)
git commit --no-verify -m "wip: work in progress"
```

---

## Checklist

- [ ] Changes are related and atomic (one logical change)
- [ ] Commit message follows conventional format
- [ ] Subject line is under 50 characters
- [ ] No secrets or credentials included
- [ ] Tests pass (if applicable)
- [ ] Linting passes (if applicable)
