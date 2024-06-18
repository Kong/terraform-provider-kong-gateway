resource "kong-gateway_vault" "my_vault" {
  name   = "env"
  prefix = "my-env"
  config = jsonencode({
    prefix = "abc"
  })
}
