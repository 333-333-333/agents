# =============================================================================
# Variable Patterns
# =============================================================================
# Every variable MUST have: description, type.
# Add default ONLY when there's a sensible universal default.
# Add sensitive = true for ANY secret.
# Add validation for values with known constraints.

# --- Required variable (no default) ---
variable "project_id" {
  description = "GCP project ID where resources will be created"
  type        = string
}

# --- Variable with sensible default ---
variable "region" {
  description = "GCP region for resource deployment"
  type        = string
  default     = "us-central1"
}

# --- Sensitive variable ---
variable "api_token" {
  description = "API token for external service authentication"
  type        = string
  sensitive   = true
}

# --- Variable with validation ---
variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "development"

  validation {
    condition     = contains(["development", "staging", "production"], var.environment)
    error_message = "Environment must be one of: development, staging, production."
  }
}

# --- Numeric with validation ---
variable "instance_count" {
  description = "Number of instances to deploy"
  type        = number
  default     = 1

  validation {
    condition     = var.instance_count >= 1 && var.instance_count <= 10
    error_message = "Instance count must be between 1 and 10."
  }
}

# --- Complex type (map) ---
variable "tags" {
  description = "Resource tags applied to all created resources"
  type        = map(string)
  default = {
    project = "bastet"
    managed = "terraform"
  }
}

# --- Complex type (list of objects) ---
variable "notification_channels" {
  description = "List of notification channels for alerting"
  type = list(object({
    name  = string
    type  = string
    email = optional(string)
  }))
  default = []
}

# --- Boolean with clear default ---
variable "enable_cdn" {
  description = "Enable Cloud CDN for the load balancer"
  type        = bool
  default     = false
}
