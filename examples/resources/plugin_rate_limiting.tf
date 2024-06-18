resource "kong-gateway_plugin_rate_limiting" "my_rate_limiting_plugin" {
  enabled = true
  config = {
    minute = 5
    policy = "local"
  }

  protocols = ["http", "https"]
  route = {
    id = kong-gateway_route.hello.id
  }
}
