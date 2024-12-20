---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_ai_prompt_template Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginAiPromptTemplate Resource
---

# kong-gateway_plugin_ai_prompt_template (Resource)

PluginAiPromptTemplate Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_ai_prompt_template" "my_pluginaiprompttemplate" {
  config = {
    allow_untemplated_requests = false
    log_original_request       = false
    max_request_body_size      = 3
    templates = [
      {
        name     = "...my_name..."
        template = "...my_template..."
      }
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
    "tls"
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

- `allow_untemplated_requests` (Boolean) Set true to allow requests that don't call or match any template.
- `log_original_request` (Boolean) Set true to add the original request to the Kong log plugin(s) output.
- `max_request_body_size` (Number) max allowed body size allowed to be introspected
- `templates` (Attributes List) Array of templates available to the request context. (see [below for nested schema](#nestedatt--config--templates))

<a id="nestedatt--config--templates"></a>
### Nested Schema for `config.templates`

Optional:

- `name` (String) Unique name for the template, can be called with `{template://NAME}`. Not Null
- `template` (String) Template string for this request, supports mustache-style `{{"{{"}}placeholders{{"}}"}}`. Not Null



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
terraform import kong-gateway_plugin_ai_prompt_template.my_kong-gateway_plugin_ai_prompt_template ""
```
