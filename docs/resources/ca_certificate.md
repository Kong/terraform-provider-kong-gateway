---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_ca_certificate Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  CACertificate Resource
---

# kong-gateway_ca_certificate (Resource)

CACertificate Resource

## Example Usage

```terraform
resource "kong-gateway_ca_certificate" "my_cacertificate" {
  cert        = "...my_cert..."
  cert_digest = "...my_cert_digest..."
  created_at  = 10
  id          = "...my_id..."
  tags = [
    "..."
  ]
  updated_at = 7
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cert` (String) PEM-encoded public certificate of the CA.

### Optional

- `cert_digest` (String) SHA256 hex digest of the public certificate. This field is read-only and it cannot be set by the caller, the value is automatically computed.
- `created_at` (Number) Unix epoch when the resource was created.
- `tags` (List of String) An optional set of strings associated with the Certificate for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import kong-gateway_ca_certificate.my_kong-gateway_ca_certificate ""
```
