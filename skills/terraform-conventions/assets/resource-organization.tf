# =============================================================================
# Resource Organization Conventions
# =============================================================================
# 1. Comment dividers (===) separate logical sections
# 2. Data sources BEFORE resources that reference them
# 3. Group related resources (e.g., a service + its IAM binding)
# 4. Order: data → resources → permissions/bindings
# 5. One blank line between resource blocks

# =============================================================================
# Data Sources
# =============================================================================
# Place ALL data sources at the top of the file (or section).
# Data sources represent READ-ONLY lookups.

# data "google_project" "current" {}
#
# data "google_compute_default_service_account" "default" {
#   project = var.project_id
# }

# =============================================================================
# Cloud Run — API Service
# =============================================================================
# Group the service with its IAM and domain mapping.

# resource "google_cloud_run_service" "api" {
#   name     = "${var.project_name}-api"
#   location = var.region
#
#   template {
#     spec {
#       containers {
#         image = var.api_image
#       }
#     }
#   }
# }
#
# resource "google_cloud_run_service_iam_member" "api_public" {
#   service  = google_cloud_run_service.api.name
#   location = google_cloud_run_service.api.location
#   role     = "roles/run.invoker"
#   member   = "allUsers"
# }

# =============================================================================
# Cloud SQL — Database
# =============================================================================
# Group the instance with its database and user.

# resource "google_sql_database_instance" "main" {
#   name             = "${var.project_name}-db"
#   database_version = "POSTGRES_16"
#   region           = var.region
#
#   settings {
#     tier = var.db_tier
#   }
# }
#
# resource "google_sql_database" "app" {
#   name     = "bastet"
#   instance = google_sql_database_instance.main.name
# }
#
# resource "google_sql_user" "app" {
#   name     = "bastet"
#   instance = google_sql_database_instance.main.name
#   password = var.db_password
# }
