variable "svctl_container_image" {
  default = "localhost:5501/natsio/svctl"
}
variable "namespace" {
  default = "natsio-demo"
}
variable "nats_namespace" {
  default = "nats"
}
variable "nats_url" {
  default = "nats://nats.nats.svc.cluster.local:4222"
}
