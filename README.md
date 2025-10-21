# terraform-provider-kong-gateway

This repository contains a Terraform provider for [Kong Gateway](https://docs.konghq.com/gateway/?utm_source=github&utm_campaign=terraform-provider-kong-gateway).

## Usage

The provider can be installed from the Terraform registry. Before using the provider, you must configure a `server_url`.

```hcl
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
  # http_headers = {
  #   "Other-Header" = "demo"
  # }
}
```

## Examples

The examples directory contains sample usage for all supported resources. For a full list of supported parameters for each resource, see the [provider documentation](https://registry.terraform.io/providers/Kong/kong-gateway/latest/docs).
