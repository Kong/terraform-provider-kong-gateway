resource "kong-gateway_plugin_basic_auth" "basic_auth" {
  enabled = true
  protocols = [
    "http",
    "https"
  ]

  service = {
    id = kong-gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}

resource "kong-gateway_plugin_key_auth" "my_key_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    key_names = ["apikey"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
resource "kong-gateway_plugin_key_auth_enc" "my_key_auth_enc" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    key_names    = ["apikey"]
    key_in_query = false
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
