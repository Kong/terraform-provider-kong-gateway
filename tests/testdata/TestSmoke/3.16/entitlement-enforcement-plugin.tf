resource "kong-gateway_plugin_entitlement_enforcement" "my_entitlement_enforcement" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    api_token                   = "my-api-token"
    entitlement_access_endpoint = "https://example.com/entitlements"

    feature = {
      key = "my-feature-key"
    }

    fail_policy             = "block"
    credit_balance_required = false
    deny_unknown_customers  = false
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
