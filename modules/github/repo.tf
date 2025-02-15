
resource "github_repository" "natsio" {
  name        = "natsio"
  description = ""
  visibility  = "public"
}

output "url" {
  value = github_repository.natsio.ssh_clone_url
}