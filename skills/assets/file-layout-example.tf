# =============================================================================
# Standard Terraform Directory Layout
# =============================================================================
#
# Every Terraform directory follows this structure:
#
#   {module}/
#     versions.tf              # Required providers and Terraform version
#     provider.tf              # Provider configuration
#     variables.tf             # All input variables
#     main.tf                  # Resources and data sources
#     outputs.tf               # All outputs
#     terraform.tfvars.example # Safe example (NEVER commit real .tfvars)
#
# When main.tf grows beyond ~300 lines, split by domain:
#
#   {module}/
#     versions.tf
#     provider.tf
#     variables.tf
#     iam.tf                   # IAM roles, service accounts, bindings
#     network.tf               # VPC, subnets, firewall rules
#     compute.tf               # Cloud Run, GCE, GKE
#     storage.tf               # GCS buckets, Cloud SQL
#     outputs.tf
#
# =============================================================================

# --- versions.tf ---
# terraform {
#   required_version = ">= 1.5"
#
#   required_providers {
#     google = {
#       source  = "hashicorp/google"
#       version = "~> 5.0"
#     }
#   }
# }

# --- provider.tf ---
# provider "google" {
#   project = var.project_id
#   region  = var.region
# }

# --- variables.tf ---
# variable "project_id" {
#   description = "GCP project ID"
#   type        = string
# }

# --- main.tf ---
# resource "google_cloud_run_service" "api" {
#   name     = "bastet-api"
#   location = var.region
#   # ...
# }

# --- outputs.tf ---
# output "api_url" {
#   description = "Cloud Run service URL"
#   value       = google_cloud_run_service.api.status[0].url
# }
