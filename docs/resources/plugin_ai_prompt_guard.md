---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_ai_prompt_guard Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginAiPromptGuard Resource
---

# kong-gateway_plugin_ai_prompt_guard (Resource)

PluginAiPromptGuard Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_ai_prompt_guard" "my_pluginaipromptguard" {
  config = {
    allow_all_conversation_history = true
    allow_patterns = [
      "..."
    ]
    deny_patterns = [
      "..."
    ]
    llm_format            = "openai"
    match_all_roles       = true
    max_request_body_size = 0
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  created_at    = 1
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
  updated_at = 4
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups (see [below for nested schema](#nestedatt--consumer_group))
- `created_at` (Number) Unix epoch when the resource was created.
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `partials` (Attributes List) (see [below for nested schema](#nestedatt--partials))
- `protocols` (List of String) A set of strings representing HTTP protocols.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `allow_all_conversation_history` (Boolean) If true, will ignore all previous chat prompts from the conversation history.
- `allow_patterns` (List of String) Array of valid regex patterns, or valid questions from the 'user' role in chat.
- `deny_patterns` (List of String) Array of invalid regex patterns, or invalid questions from the 'user' role in chat.
- `llm_format` (String) LLM input and output format and schema to use. must be one of ["bedrock", "gemini", "openai"]
- `match_all_roles` (Boolean) If true, will match all roles in addition to 'user' role in conversation history.
- `max_request_body_size` (Number) max allowed body size allowed to be introspected


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
terraform import kong-gateway_plugin_ai_prompt_guard.my_kong-gateway_plugin_ai_prompt_guard ""
```
