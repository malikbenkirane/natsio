resource "kubernetes_deployment" "svctl" {
  metadata {
    name = "svctl"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "svctl"
      }
    }

    template {
      metadata {
        labels = {
          app = "svctl"
        }
      }

      spec {
        container {
          name  = "svctl"
          image = var.svctl_container_image
        }
      }
    }
  }
}
