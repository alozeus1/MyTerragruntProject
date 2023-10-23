resource "local_file" "directory" {
  content  = "Sample Configuration for ${var.environment}"
  filename = "${path.module}/sample-config-${var.environment}.conf"
}

terraform {
  backend "local" {}
}