provider "kong-gateway" {
  server_url = "http://localhost:8001"
}

resource "kong-gateway_route" "my_route" {
  name = "route-headers"
  headers = {
    "X-Header-Goes-Here" = ["Value-1", "Value-2"]
  }
}
