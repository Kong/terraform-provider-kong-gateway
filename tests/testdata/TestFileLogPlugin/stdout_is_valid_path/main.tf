provider "kong-gateway" {
  server_url = "http://localhost:8001"
}

resource "kong-gateway_plugin_file_log" "my_file_log" {
  enabled = true

  config = {
    path = "/dev/stdout"
  }
}