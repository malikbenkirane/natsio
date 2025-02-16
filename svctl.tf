resource "kubernetes_deployment" "svctl" {
  metadata {
    name = "svctl"
    namespace = kubernetes_namespace.namespace.metadata[0].name
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
	  command = ["/svctl", "serve"]
          name  = "svctl"
          image = var.svctl_container_image
	  env {
	    name = "NATS_URL"
	    value = var.nats_url
	  }
        }
      }
    }
  }
}
