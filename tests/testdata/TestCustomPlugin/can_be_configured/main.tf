provider "kong-gateway" {
  server_url = "http://localhost:8001"
}

resource "kong-gateway_custom_plugin" "rl" {
  name = "rate-limiting"
  config = {
    minute = 5
  }
}
