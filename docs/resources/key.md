---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_key Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  Key Resource
---

# kong-gateway_key (Resource)

Key Resource

## Example Usage

```terraform
resource "kong-gateway_key" "my_key" {
  jwk  = "...my_jwk..."
  kid  = "...my_kid..."
  name = "...my_name..."
  pem = {
    private_key = "...my_private_key..."
    public_key  = "...my_public_key..."
  }
  set = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `jwk` (String) A JSON Web Key represented as a string.
- `kid` (String) A unique identifier for a key.
- `name` (String) The name to associate with the given keys.
- `pem` (Attributes) A keypair in PEM format. (see [below for nested schema](#nestedatt--pem))
- `set` (Attributes) The id (an UUID) of the key-set with which to associate the key. (see [below for nested schema](#nestedatt--set))
- `tags` (List of String) An optional set of strings associated with the Key for grouping and filtering.

### Read-Only

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.
- `updated_at` (Number) Unix epoch when the resource was last updated.

<a id="nestedatt--pem"></a>
### Nested Schema for `pem`

Optional:

- `private_key` (String)
- `public_key` (String)


<a id="nestedatt--set"></a>
### Nested Schema for `set`

Optional:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import kong-gateway_key.my_kong-gateway_key ""
```