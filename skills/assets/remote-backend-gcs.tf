# =============================================================================
# Remote Backend — Google Cloud Storage
# =============================================================================
# Stores Terraform state in a GCS bucket with:
# - Encryption at rest (Google-managed or CMEK)
# - State locking (prevents concurrent modifications)
# - Versioning (rollback to previous state)
# - IAM access control
#
# Prerequisites:
# 1. Create the bucket MANUALLY (do not manage it with this Terraform config)
# 2. Enable versioning on the bucket
# 3. Grant the service account roles/storage.objectAdmin on the bucket

# --- Add this to versions.tf ---
terraform {
  backend "gcs" {
    bucket = "bastet-terraform-state" # Create this bucket manually
    prefix = "discord"                # Each module gets its own prefix
  }
}

# =============================================================================
# Bucket creation (run this ONCE, separately, or create via gcloud)
# =============================================================================
# gcloud storage buckets create gs://bastet-terraform-state \
#   --project=bastet-prod \
#   --location=us-central1 \
#   --uniform-bucket-level-access \
#   --public-access-prevention=enforced
#
# gcloud storage buckets update gs://bastet-terraform-state \
#   --versioning

# =============================================================================
# Per-environment state isolation
# =============================================================================
# Use different prefixes for each environment:
#
#   prefix = "discord"         → gs://bastet-tf-state/discord/default.tfstate
#   prefix = "gcp/production"  → gs://bastet-tf-state/gcp/production/default.tfstate
#   prefix = "gcp/staging"     → gs://bastet-tf-state/gcp/staging/default.tfstate
