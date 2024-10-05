---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_degraphql Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginDegraphql Resource
---

# kong-gateway_plugin_degraphql (Resource)

PluginDegraphql Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_degraphql" "my_plugindegraphql" {
  config = {
    graphql_server_path = "...my_graphql_server_path..."
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  enabled       = true
  instance_name = "...my_instance_name..."
  ordering      = "{ \"see\": \"documentation\" }"
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
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) (see [below for nested schema](#nestedatt--consumer_group))
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (String) Parsed as JSON.
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

- `graphql_server_path` (String) A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).


<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Optional:

- `id` (String)


<a id="nestedatt--consumer_group"></a>
### Nested Schema for `consumer_group`

Optional:

- `id` (String)


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
terraform import kong-gateway_plugin_degraphql.my_kong-gateway_plugin_degraphql ""
```