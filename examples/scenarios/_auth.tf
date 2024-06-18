terraform {
  required_providers {
    kong-gateway = {
      source = "kong/kong-gateway"
    }
  }
}

provider "kong-gateway" {
  server_url = "http://localhost:8001"
}
