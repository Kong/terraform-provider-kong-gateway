---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_plugin_response_ratelimiting Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  PluginResponseRatelimiting Resource
---

# kong-gateway_plugin_response_ratelimiting (Resource)

PluginResponseRatelimiting Resource

## Example Usage

```terraform
resource "kong-gateway_plugin_response_ratelimiting" "my_pluginresponseratelimiting" {
  config = {
    block_on_first_violation = false
    fault_tolerant           = true
    header_name              = "...my_header_name..."
    hide_client_headers      = true
    limit_by                 = "consumer"
    limits = {
      "see" : jsonencode("documentation"),
    }
    policy = "redis"
    redis = {
      database    = 10
      host        = "...my_host..."
      password    = "...my_password..."
      port        = 22322
      server_name = "...my_server_name..."
      ssl         = true
      ssl_verify  = false
      timeout     = 320380033
      username    = "...my_username..."
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  enabled       = false
  instance_name = "...my_instance_name..."
  ordering      = "{ \"see\": \"documentation\" }"
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

- `block_on_first_violation` (Boolean) A boolean value that determines if the requests should be blocked as soon as one limit is being exceeded. This will block requests that are supposed to consume other limits too.
- `fault_tolerant` (Boolean) A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party datastore. If `true`, requests will be proxied anyway, effectively disabling the rate-limiting function until the datastore is working again. If `false`, then the clients will see `500` errors.
- `header_name` (String) The name of the response header used to increment the counters.
- `hide_client_headers` (Boolean) Optionally hide informative response headers.
- `limit_by` (String) The entity that will be used when aggregating the limits: `consumer`, `credential`, `ip`. If the `consumer` or the `credential` cannot be determined, the system will always fallback to `ip`. must be one of ["consumer", "credential", "ip"]
- `limits` (Map of String) A map that defines rate limits for the plugin.
- `policy` (String) The rate-limiting policies to use for retrieving and incrementing the limits. must be one of ["local", "cluster", "redis"]
- `redis` (Attributes) Redis configuration (see [below for nested schema](#nestedatt--config--redis))

<a id="nestedatt--config--redis"></a>
### Nested Schema for `config.redis`

Optional:

- `database` (Number) Database to use for the Redis connection when using the `redis` strategy
- `host` (String) A string representing a host name, such as example.com.
- `password` (String) Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
- `port` (Number) An integer representing a port number between 0 and 65535, inclusive.
- `server_name` (String) A string representing an SNI (server name indication) value for TLS.
- `ssl` (Boolean) If set to true, uses SSL to connect to Redis.
- `ssl_verify` (Boolean) If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
- `timeout` (Number) An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
- `username` (String) Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.



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
terraform import kong-gateway_plugin_response_ratelimiting.my_kong-gateway_plugin_response_ratelimiting ""
```