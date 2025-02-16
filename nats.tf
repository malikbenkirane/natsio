# Fixme namespace has no effect and resources are in default namespace
#resource "helm_release" "nats" {
#  name = "nats"
#  namespace = kubernetes_namespace.nats_namespace.metadata[0].name
#  repository = "https://nats-io.github.io/k8s/helm/charts"
#  chart = "nats"
#}
