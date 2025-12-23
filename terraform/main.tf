# Generate SSH keys for Flux to access GitHub
module "tls_keys" {
  source = "github.com/den-vasyliev/tf-hashicorp-tls-keys"
}

# Create local Kind cluster
module "kind_cluster" {
  source = "github.com/den-vasyliev/tf-kind-cluster"
}

# Configure GitHub repository with deploy key
module "github_repository" {
  source = "github.com/den-vasyliev/tf-github-repository"

  github_owner             = var.github_owner
  github_token             = var.github_token
  repository_name          = var.flux_github_repo
  public_key_openssh       = module.tls_keys.public_key_openssh
  public_key_openssh_title = "flux-deploy-key"
}

# Bootstrap FluxCD into the cluster
module "flux_bootstrap" {
  source = "github.com/den-vasyliev/tf-fluxcd-flux-bootstrap"

  github_token      = var.github_token
  github_repository = "${var.github_owner}/${var.flux_github_repo}"
  private_key       = module.tls_keys.private_key_pem
  target_path       = var.flux_target_path
}
