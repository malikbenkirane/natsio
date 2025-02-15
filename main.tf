
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "6.2.2"
    }
  }
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