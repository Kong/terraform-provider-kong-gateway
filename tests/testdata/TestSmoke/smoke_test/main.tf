provider "kong-gateway" {
  server_url = "http://localhost:8001"
}

resource "kong-gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "http"
  host     = "httpbin.org"
  port     = 443
  path     = "/"
}

resource "kong-gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_basic_auth" "basic_auth" {
  enabled = true
  service = {
    id = kong-gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}

resource "kong-gateway_consumer" "alice" {
  username  = "alice"
  custom_id = "alice"
}

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

