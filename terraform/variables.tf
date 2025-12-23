variable "github_owner" {
  description = "GitHub username"
  type        = string
}

variable "github_token" {
  description = "GitHub PAT with repo permissions"
  type        = string
  sensitive   = true
}

variable "flux_github_repo" {
  description = "Repository name for Flux configuration"
  type        = string
  default     = "flux-infra"
}

variable "flux_target_path" {
  description = "Path in repo for Flux manifests"
  type        = string
  default     = "clusters/kind"
}
