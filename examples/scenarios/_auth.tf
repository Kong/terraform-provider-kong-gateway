terraform {
  required_providers {
    kong-gateway = {
      source = "kong/kong-gateway"
    }
  }
}

provider "kong-gateway" {
  server_url = "http://localhost:8001"
  # admin_token = "your_rbac_token_here__enterprise_only"
}
