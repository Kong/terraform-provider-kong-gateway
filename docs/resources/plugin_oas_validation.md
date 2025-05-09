---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_oas_validation Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginOasValidation Resource
---

# kong-gateway_plugin_oas_validation (Resource)

PluginOasValidation Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_oas_validation" "my_pluginoasvalidation" {
  config = {
    allowed_header_parameters                    = "...my_allowed_header_parameters..."
    api_spec                                     = "...my_api_spec..."
    api_spec_encoded                             = false
    custom_base_path                             = "...my_custom_base_path..."
    header_parameter_check                       = false
    include_base_path                            = true
    notify_only_request_validation_failure       = true
    notify_only_response_body_validation_failure = false
    query_parameter_check                        = false
    validate_request_body                        = true
    validate_request_header_params               = true
    validate_request_query_params                = true
    validate_request_uri_params                  = false
    validate_response_body                       = false
    verbose_response                             = true
  }
  consumer = {
    id = "...my_id..."
  }
  created_at    = 1
  enabled       = false
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
    "https"
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
  updated_at = 3
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
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

- `allowed_header_parameters` (String) List of header parameters in the request that will be ignored when performing HTTP header validation. These are additional headers added to an API request beyond those defined in the API specification.  For example, you might include the HTTP header `User-Agent`, which lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.
- `api_spec` (String) The API specification defined using either Swagger or the OpenAPI. This can be either a JSON or YAML based file. If using a YAML file, the spec needs to be URI-Encoded to preserve the YAML format.
- `api_spec_encoded` (Boolean) Indicates whether the api_spec is URI-Encoded.
- `custom_base_path` (String) The base path to be used for path match evaluation. This value is ignored if `include_base_path` is set to `false`.
- `header_parameter_check` (Boolean) If set to true, checks if HTTP header parameters in the request exist in the API specification.
- `include_base_path` (Boolean) Indicates whether to include the base path when performing path match evaluation.
- `notify_only_request_validation_failure` (Boolean) If set to true, notifications via event hooks are enabled, but request based validation failures don't affect the request flow.
- `notify_only_response_body_validation_failure` (Boolean) If set to true, notifications via event hooks are enabled, but response validation failures don't affect the response flow.
- `query_parameter_check` (Boolean) If set to true, checks if query parameters in the request exist in the API specification.
- `validate_request_body` (Boolean) If set to true, validates the request body content against the API specification.
- `validate_request_header_params` (Boolean) If set to true, validates HTTP header parameters against the API specification.
- `validate_request_query_params` (Boolean) If set to true, validates query parameters against the API specification.
- `validate_request_uri_params` (Boolean) If set to true, validates URI parameters in the request against the API specification.
- `validate_response_body` (Boolean) If set to true, validates the response from the upstream services against the API specification. If validation fails, it results in an `HTTP 406 Not Acceptable` status code.
- `verbose_response` (Boolean) If set to true, returns a detailed error message for invalid requests & responses. This is useful while testing.


<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

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
terraform import kong-gateway_plugin_oas_validation.my_kong-gateway_plugin_oas_validation ""
```
