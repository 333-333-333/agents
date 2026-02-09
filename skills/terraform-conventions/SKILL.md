---
name: terraform-conventions
description: >
  Terraform code conventions: file organization, naming, variables, outputs, formatting, and the plan/apply workflow.
  Trigger: When writing Terraform files, reviewing HCL code style, organizing .tf files, or adding new Terraform modules.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [infra]
  auto_invoke:
    - "Writing Terraform configuration files"
    - "Creating new Terraform modules"
    - "Organizing .tf file structure"
    - "Reviewing Terraform code style"
    - "Adding variables or outputs to Terraform"
    - "Running terraform plan or apply"
---

## When to Use

- Creating a new Terraform directory or module
- Adding resources, variables, or outputs to existing `.tf` files
- Reviewing Terraform code for style and organization
- Deciding how to split resources across files
- Running the plan/apply workflow

---

## Critical Patterns

### File Organization — One Concern Per File

Every Terraform directory MUST have these files. No exceptions.

| File | Contains | Rule |
|------|----------|------|
| `versions.tf` | `terraform {}` block — required providers, version constraints | ONE per directory, never duplicated |
| `provider.tf` | Provider configuration blocks | ONE provider per block, sensitive config via variables |
| `variables.tf` | ALL `variable` blocks | Grouped by section with comment dividers |
| `outputs.tf` | ALL `output` blocks | Mirror the resource structure |
| `main.tf` | Resources and data sources | Split into domain files when > 300 lines |

When `main.tf` grows beyond ~300 lines, split by domain:

```
main.tf          →  roles.tf, channels.tf, permissions.tf
                 →  network.tf, compute.tf, storage.tf
                 →  iam.tf, databases.tf, pubsub.tf
```

> See [assets/file-layout-example.tf](assets/file-layout-example.tf) for the standard layout.

### Naming Conventions

| Element | Convention | Good | Bad |
|---------|-----------|------|-----|
| Resources | `snake_case`, descriptive noun | `discord_text_channel.dev_backend` | `discord_text_channel.ch1` |
| Variables | `snake_case`, clear purpose | `server_name` | `name`, `sn` |
| Outputs | `snake_case`, match resource | `server_id` | `id`, `output1` |
| Locals | `snake_case` | `default_tags` | `dt` |
| Files | `lowercase.tf` | `variables.tf` | `Variables.tf`, `vars.tf` |
| Modules | `kebab-case` directory | `modules/vpc-network/` | `modules/VPCNetwork/` |

### Variable Rules

Every variable MUST have:
1. A `description` — no exceptions
2. A `type` — always explicit, never rely on inference
3. A `default` — ONLY if there's a sensible universal default
4. `sensitive = true` — for any secret (tokens, passwords, keys)
5. `validation` blocks — for values with known constraints

> See [assets/variables-example.tf](assets/variables-example.tf) for the full pattern.

### Output Rules

Every output MUST have:
1. A `description`
2. `sensitive = true` if it exposes secrets
3. Outputs should provide the IDs and attributes that OTHER modules or humans need

### Resource Organization Within Files

- Use **comment dividers** (`# ===`) to separate logical sections
- Group related resources together (e.g., a channel and its permissions)
- Data sources go BEFORE the resources that reference them
- Order: data sources → resources → permission overrides

> See [assets/resource-organization.tf](assets/resource-organization.tf) for the pattern.

### The Change Workflow

```
1. Edit .tf files
2. terraform fmt -recursive    # Auto-format all files
3. terraform validate          # Check syntax and types
4. terraform plan              # Review the diff — READ IT CAREFULLY
5. terraform apply             # Apply with confirmation
```

**NEVER** skip `plan`. Always review the diff before applying. Look for unexpected destroys or replacements.

### Formatting

- Run `terraform fmt` before every commit — non-negotiable
- Use 2-space indentation (Terraform default)
- Align `=` signs within a block for readability
- One blank line between resource blocks
- Comment dividers between logical sections

---

## Decision Tree

```
New Terraform directory?       → Create all 5 standard files (versions, provider, variables, main, outputs)
main.tf > 300 lines?           → Split by domain into separate .tf files
Variable has a secret?         → Add sensitive = true
Variable has known values?     → Add validation block
Output exposes a secret?       → Add sensitive = true
Multiple providers needed?     → One provider block each in provider.tf, alias if same type
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/file-layout-example.tf` | Standard file layout for a new Terraform directory |
| `assets/variables-example.tf` | Variable patterns with types, defaults, validation, and sensitivity |
| `assets/resource-organization.tf` | Resource grouping and comment divider conventions |

---

## Commands

```bash
terraform fmt -recursive   # Format all .tf files
terraform validate         # Check syntax
terraform plan             # Preview changes
```

---

## Anti-Patterns

| Don't | Do |
|-------|-----|
| Dump everything in `main.tf` | Split by domain when > 300 lines |
| Variables without descriptions | ALWAYS add `description` |
| Hardcode values in resources | Use variables with defaults |
| Skip `terraform plan` | ALWAYS review the plan before apply |
| Use short cryptic names | Use descriptive `snake_case` names |
| Rely on type inference | Always set explicit `type` on variables |
| Commit without `terraform fmt` | Format before every commit |

---

## Resources

- **Templates**: See [assets/](assets/) for file layout and variable patterns
