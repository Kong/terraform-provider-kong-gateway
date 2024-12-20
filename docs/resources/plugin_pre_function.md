---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_pre_function Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginPreFunction Resource
---

# kong-gateway_plugin_pre_function (Resource)

PluginPreFunction Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_pre_function" "my_pluginprefunction" {
  config = {
    access = [
      "..."
    ]
    body_filter = [
      "..."
    ]
    certificate = [
      "..."
    ]
    header_filter = [
      "..."
    ]
    log = [
      "..."
    ]
    rewrite = [
      "..."
    ]
    ws_client_frame = [
      "..."
    ]
    ws_close = [
      "..."
    ]
    ws_handshake = [
      "..."
    ]
    ws_upstream_frame = [
      "..."
    ]
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
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
  protocols = [
    "ws"
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
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))

### Optional

- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) (see [below for nested schema](#nestedatt--consumer_group))
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `protocols` (List of String) A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.

### Read-Only

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.
- `updated_at` (Number) Unix epoch when the resource was last updated.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `access` (List of String)
- `body_filter` (List of String)
- `certificate` (List of String)
- `header_filter` (List of String)
- `log` (List of String)
- `rewrite` (List of String)
- `ws_client_frame` (List of String)
- `ws_close` (List of String)
- `ws_handshake` (List of String)
- `ws_upstream_frame` (List of String)


<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Optional:

- `id` (String)


<a id="nestedatt--consumer_group"></a>
### Nested Schema for `consumer_group`

Optional:

- `id` (String)


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
terraform import kong-gateway_plugin_pre_function.my_kong-gateway_plugin_pre_function ""
```
