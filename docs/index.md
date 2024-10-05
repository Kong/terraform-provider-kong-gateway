---
page_title: "kong-gateway Provider"
---

# Kong-Gateway Provider

The Kong-Gateway provider is used to interact with the resources supported by [Kong Gateway](https://docs.konghq.com/gateway/latest/).


## Example Usage

Configure the kong-gateway Provider
```terraform
provider "kong-gateway" {
    server_url = "http://localhost:8001"
}
```
Create consumer resource

```terraform
resource "kong-gateway_consumer" "example_consumer" {
  custom_id = "consumer1"
  tags      = [ "consumer1" ]
}
```
