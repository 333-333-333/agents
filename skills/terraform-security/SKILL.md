---
name: terraform-security
description: >
  Terraform security practices: sensitive variables, secret management, state protection, .gitignore patterns, and CI/CD credential handling.
  Trigger: When handling secrets in Terraform, configuring state backends, reviewing .gitignore for Terraform, or setting up CI/CD pipelines for infrastructure.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [infra]
  auto_invoke:
    - "Adding secrets or tokens to Terraform variables"
    - "Configuring Terraform state backends"
    - "Setting up .gitignore for Terraform directories"
    - "Configuring CI/CD pipelines for Terraform"
    - "Reviewing Terraform for secret leaks"
    - "Protecting Terraform state files"
---

## When to Use

- Adding any secret, token, password, or API key to Terraform
- Configuring local or remote state backends
- Setting up `.gitignore` for a Terraform directory
- Creating CI/CD pipelines that run `terraform plan/apply`
- Reviewing Terraform code for security issues
- Deciding how to pass secrets to Terraform

---

## Critical Patterns

### The Three Things You NEVER Commit

| File/Pattern | Why | Consequence if Leaked |
|-------------|-----|----------------------|
| `*.tfvars` | Contains actual secret values | Full credential exposure |
| `*.tfstate` / `*.tfstate.backup` | Contains ALL resource attributes including secrets in plaintext | Infrastructure takeover |
| `.terraform/` | Contains provider binaries and cached state | Potential secret exposure |

These MUST be in `.gitignore`. No exceptions. Ever.

> See [assets/gitignore-terraform](assets/gitignore-terraform) for the complete `.gitignore` pattern.

### Sensitive Variables

ANY variable that holds a secret MUST be marked `sensitive = true`. This prevents Terraform from displaying the value in plan/apply output and logs.

```
variable "discord_token" {
  description = "Discord bot token"
  type        = string
  sensitive   = true
}
```

Terraform will show `(sensitive value)` instead of the actual value in all output.

### Passing Secrets — Priority Order

| Method | When to Use | Security Level |
|--------|-------------|----------------|
| `TF_VAR_` environment variable | CI/CD pipelines, automation | High — never touches disk |
| Secret manager (GCP Secret Manager, Vault) | Production infrastructure | Highest — encrypted, audited, rotated |
| `terraform.tfvars` file (local only) | Local development only | Medium — on disk but gitignored |
| `-var` CLI flag | One-off operations | Low — visible in shell history |
| Hardcoded in `.tf` files | **NEVER** | **Catastrophic** — in git forever |

> See [assets/secret-passing-examples.sh](assets/secret-passing-examples.sh) for the patterns.

### State File Protection

The Terraform state file contains the **full truth** about your infrastructure — including every attribute of every resource. For many providers, this includes secrets in plaintext.

**Local state** (default):
- Fine for single-operator, non-production setups
- MUST be gitignored
- Lives at `terraform.tfstate` in the working directory

**Remote state** (recommended for teams):
- Encrypted at rest (GCS, S3)
- State locking prevents concurrent modifications
- Access controlled via IAM

> See [assets/remote-backend-gcs.tf](assets/remote-backend-gcs.tf) for the GCS backend configuration.

### CI/CD Pipeline Security

When running Terraform in CI/CD (GitHub Actions, Cloud Build):

| Rule | Why |
|------|-----|
| Secrets via environment variables only | Never pass via CLI flags (visible in logs) |
| Use `TF_VAR_` prefix for Terraform variables | Standard Terraform convention |
| Store secrets in GitHub Secrets / GCP Secret Manager | Encrypted at rest, access-controlled |
| Run `plan` on PRs, `apply` only on merge to main | Prevent accidental infrastructure changes |
| Use `-input=false` flag | Prevent interactive prompts in CI |
| Use `-no-color` flag | Clean log output |
| NEVER log `terraform output` with sensitive values | Sensitive outputs leak in CI logs |

> See [assets/github-actions-terraform.yml](assets/github-actions-terraform.yml) for the GitHub Actions workflow.

### Provider Authentication

| Provider | Auth Method | Environment Variable |
|----------|------------|---------------------|
| GCP | Service account key or Workload Identity | `GOOGLE_APPLICATION_CREDENTIALS` or `GOOGLE_CLOUD_PROJECT` |
| Discord | Bot token | `TF_VAR_discord_token` |
| AWS | Access key or IAM role | `AWS_ACCESS_KEY_ID` + `AWS_SECRET_ACCESS_KEY` |

**Prefer Workload Identity Federation** over service account keys when possible — no long-lived credentials.

---

## Decision Tree

```
Adding a secret to Terraform?     → variable with sensitive = true, pass via TF_VAR_
Working alone on non-prod?        → Local state is fine (gitignored)
Working in a team?                → Remote backend with locking (GCS/S3)
Setting up CI/CD?                 → GitHub Secrets + TF_VAR_ + plan on PR, apply on merge
Reviewing a PR with .tf changes?  → Check for hardcoded secrets, missing sensitive flags
State file accidentally committed? → ROTATE ALL CREDENTIALS IMMEDIATELY, then remove from git history
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/gitignore-terraform` | Complete .gitignore pattern for Terraform directories |
| `assets/secret-passing-examples.sh` | How to pass secrets via environment variables and tfvars |
| `assets/remote-backend-gcs.tf` | GCS remote backend configuration with state locking |
| `assets/github-actions-terraform.yml` | GitHub Actions workflow for plan on PR, apply on merge |

---

## Commands

```bash
terraform plan -input=false   # Non-interactive plan (for CI)
terraform apply -input=false  # Non-interactive apply (for CI)
```

---

## Anti-Patterns

| Don't | Do |
|-------|-----|
| Hardcode secrets in `.tf` files | Use `sensitive` variables + `TF_VAR_` env vars |
| Commit `*.tfvars` with real values | Commit only `*.tfvars.example` with placeholder values |
| Commit `*.tfstate` | Add to `.gitignore`, use remote backend for teams |
| Pass secrets via `-var` flag | Use `TF_VAR_` environment variables |
| Run `apply` in CI on every push | Run `plan` on PRs, `apply` only on merge to main |
| Use service account keys in CI | Use Workload Identity Federation |
| Share state files via Slack/email | Use a remote backend with IAM access control |
| Ignore leaked state file | ROTATE ALL CREDENTIALS immediately |

---

## Resources

- **Templates**: See [assets/](assets/) for .gitignore, CI/CD, and backend patterns
