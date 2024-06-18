resource "kong-gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "http"
  host     = "httpbin.org"
  port     = 443
  path     = "/"
}
