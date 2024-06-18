resource "kong-gateway_plugin_basic_auth" "basic_auth" {
  enabled = true
  service = {
    id = kong-gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}
