---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_service Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  Service Resource
---

# kong-gateway_service (Resource)

Service Resource

## Example Usage

```terraform
resource "kong-gateway_service" "my_service" {
  ca_certificates = [
    "..."
  ]
  client_certificate = {
    id = "...my_id..."
  }
  connect_timeout = 8
  enabled         = false
  host            = "...my_host..."
  name            = "...my_name..."
  path            = "...my_path..."
  port            = 1
  protocol        = "grpcs"
  read_timeout    = 7
  retries         = 5
  tags = [
    "..."
  ]
  tls_verify       = true
  tls_verify_depth = 0
  url              = "...my_url..."
  write_timeout    = 4
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ca_certificates` (List of String) Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
- `client_certificate` (Attributes) Certificate to be used as client certificate while TLS handshaking to the upstream server. (see [below for nested schema](#nestedatt--client_certificate))
- `connect_timeout` (Number) The timeout in milliseconds for establishing a connection to the upstream server.
- `enabled` (Boolean) Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
- `host` (String) The host of the upstream server. Note that the host value is case sensitive.
- `name` (String) The Service name.
- `path` (String) The path to be used in requests to the upstream server.
- `port` (Number) The upstream server port.
- `protocol` (String) The protocol used to communicate with the upstream. must be one of ["grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp", "ws", "wss"]
- `read_timeout` (Number) The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
- `retries` (Number) The number of retries to execute upon failure to proxy.
- `tags` (List of String) An optional set of strings associated with the Service for grouping and filtering.
- `tls_verify` (Boolean) Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
- `tls_verify_depth` (Number) Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
- `url` (String) Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.
- `write_timeout` (Number) The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.

### Read-Only

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.
- `updated_at` (Number) Unix epoch when the resource was last updated.

<a id="nestedatt--client_certificate"></a>
### Nested Schema for `client_certificate`

Optional:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import kong-gateway_service.my_kong-gateway_service ""
```