provider "kong-gateway" {
  server_url = "http://localhost:8001"
}
resource "kong-gateway_service" "httpbin" {
  name     = "httpbin-partials"
  protocol = "http"
  host     = "httpbin.org"
  port     = 443
  path     = "/"
}

resource "kong-gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  protocols = [
    "http",
    "https"
  ]

  strip_path = false

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_rate_limiting" "my_rate_limiting_plugin" {
  enabled = true
  config = {
    minute = 5
    policy = "local"
  }

  protocols = [
    "http"
  ]

  service = {
    id = kong-gateway_service.httpbin.id
  }

}

