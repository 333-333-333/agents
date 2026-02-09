#!/usr/bin/env bash
# =============================================================================
# How to Pass Secrets to Terraform
# =============================================================================
# Priority order (most secure first):
# 1. Secret manager (GCP Secret Manager, HashiCorp Vault)
# 2. TF_VAR_ environment variables
# 3. terraform.tfvars file (local dev only, gitignored)
# 4. -var flag (avoid — visible in shell history)

# =============================================================================
# Method 1: TF_VAR_ Environment Variables (RECOMMENDED for CI/CD)
# =============================================================================
# Terraform automatically reads TF_VAR_{variable_name} from the environment.
# The variable name must match EXACTLY (case-sensitive).

export TF_VAR_discord_token="your-bot-token-here"
export TF_VAR_db_password="your-db-password-here"

terraform plan # Variables are automatically picked up
terraform apply

# =============================================================================
# Method 2: terraform.tfvars (LOCAL DEV ONLY)
# =============================================================================
# Create from the example file, fill in real values.
# This file is gitignored and NEVER committed.

# cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars with your real values
# terraform plan   # Automatically reads terraform.tfvars

# =============================================================================
# Method 3: GCP Secret Manager (PRODUCTION)
# =============================================================================
# Store secrets in GCP Secret Manager and read them with a data source.
# This keeps secrets out of Terraform variables entirely.
#
# In your .tf file:
#
#   data "google_secret_manager_secret_version" "discord_token" {
#     secret  = "discord-bot-token"
#     project = var.project_id
#   }
#
#   # Use: data.google_secret_manager_secret_version.discord_token.secret_data

# =============================================================================
# Method 4: -var flag (AVOID — visible in shell history)
# =============================================================================
# terraform plan -var="discord_token=your-token-here"
# ^^^ This is visible in: shell history, process list, CI logs
# Only use for non-sensitive values or one-off debugging.
