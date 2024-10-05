---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_mocking Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginMocking Resource
---

# kong-gateway_plugin_mocking (Resource)

PluginMocking Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_mocking" "my_pluginmocking" {
  config = {
    api_specification          = "...my_api_specification..."
    api_specification_filename = "...my_api_specification_filename..."
    custom_base_path           = "...my_custom_base_path..."
    include_base_path          = true
    included_status_codes = [
      1
    ]
    max_delay_time     = 8.63
    min_delay_time     = 6.83
    random_delay       = false
    random_examples    = false
    random_status_code = false
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

- `api_specification` (String) The contents of the specification file. You must use this option for hybrid or DB-less mode. You can include the full specification as part of the configuration. In Kong Manager, you can copy and paste the contents of the spec directly into the `Config.Api Specification` text field.
- `api_specification_filename` (String) The path and name of the specification file loaded into Kong Gateway's database. You cannot use this option for DB-less or hybrid mode.
- `custom_base_path` (String) The base path to be used for path match evaluation. This value is ignored if `include_base_path` is set to `false`.
- `include_base_path` (Boolean) Indicates whether to include the base path when performing path match evaluation.
- `included_status_codes` (List of Number) A global list of the HTTP status codes that can only be selected and returned.
- `max_delay_time` (Number) The maximum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be greater than the `min_delay_time`.
- `min_delay_time` (Number) The minimum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be less than the `max_delay_time`.
- `random_delay` (Boolean) Enables a random delay in the mocked response. Introduces delays to simulate real-time response times by APIs.
- `random_examples` (Boolean) Randomly selects one example and returns it. This parameter requires the spec to have multiple examples configured.
- `random_status_code` (Boolean) Determines whether to randomly select an HTTP status code from the responses of the corresponding API method. The default value is `false`, which means the minimum HTTP status code is always selected and returned.


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
terraform import kong-gateway_plugin_mocking.my_kong-gateway_plugin_mocking ""
```