resource "kubernetes_namespace" "namespace" {
  metadata {
    name = var.namespace
  }
}

# fixme see also nats.tf
#resource "kubernetes_namespace" "nats_namespace" {
#  metadata {
#    name = var.nats_namespace
#  }
#}
