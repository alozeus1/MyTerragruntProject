remote_state {
  backend = "local"
  config = {
    path = "${get_terragrunt_dir()}/../../mock_remote_state/terraform.tfstate"
  }
}

inputs = {
  environment = "prod"
}
