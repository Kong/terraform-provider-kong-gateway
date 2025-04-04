---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_hmac_auth Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginHmacAuth Resource
---

# kong-gateway_plugin_hmac_auth (Resource)

PluginHmacAuth Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_hmac_auth" "my_pluginhmacauth" {
  config = {
    algorithms = [
      "hmac-sha256"
    ]
    anonymous  = "...my_anonymous..."
    clock_skew = 8.73
    enforce_headers = [
      "..."
    ]
    hide_credentials      = false
    realm                 = "...my_realm..."
    validate_request_body = false
  }
  created_at    = 2
  enabled       = true
  id            = "...my_id..."
  instance_name = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "grpcs"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
  updated_at = 9
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `created_at` (Number) Unix epoch when the resource was created.
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `partials` (Attributes List) (see [below for nested schema](#nestedatt--partials))
- `protocols` (List of String) A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `algorithms` (List of String) A list of HMAC digest algorithms that the user wants to support. Allowed values are `hmac-sha1`, `hmac-sha256`, `hmac-sha384`, and `hmac-sha512`
- `anonymous` (String) An optional string (Consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
- `clock_skew` (Number) Clock skew in seconds to prevent replay attacks.
- `enforce_headers` (List of String) A list of headers that the client should at least use for HTTP signature creation.
- `hide_credentials` (Boolean) An optional boolean value telling the plugin to show or hide the credential from the upstream service.
- `realm` (String) When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
- `validate_request_body` (Boolean) A boolean value telling the plugin to enable body validation.


<a id="nestedatt--ordering"></a>
### Nested Schema for `ordering`

Optional:

- `after` (Attributes) (see [below for nested schema](#nestedatt--ordering--after))
- `before` (Attributes) (see [below for nested schema](#nestedatt--ordering--before))

<a id="nestedatt--ordering--after"></a>
### Nested Schema for `ordering.after`

Optional:

- `access` (List of String)


<a id="nestedatt--ordering--before"></a>
### Nested Schema for `ordering.before`

Optional:

- `access` (List of String)



<a id="nestedatt--partials"></a>
### Nested Schema for `partials`

Optional:

- `id` (String)
- `name` (String)
- `path` (String)


<a id="nestedatt--route"></a>
### Nested Schema for `route`

Optional:

- `id` (String)


<a id="nestedatt--service"></a>
### Nested Schema for `service`

Optional:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import kong-gateway_plugin_hmac_auth.my_kong-gateway_plugin_hmac_auth ""
```
