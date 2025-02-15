
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "6.2.2"
    }
    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "2.35.1"
    }
    helm = {
      source = "hashicorp/helm"
      version = "3.0.0-pre1"
    }
  }
  backend "kubernetes" {
    secret_suffix = "natsio-demo-state"
    namespace = "tofu-backend"
    config_path = "~/.kube/config"
    config_context = "kind-kind"
  }
}

provider "helm" {
  kubernetes = {
		config_path = "~/.kube/config"
		config_context = "kind-kind"
	}
}

provider "kubernetes" {
  config_path = "~/.kube/config"
  config_context = "kind-kind"
}

provider "github" {
  owner = "malikbenkirane"
}

module "repo" {
  source = "./modules/github"
}

output "repo_url" {
  value = module.repo.url
}
