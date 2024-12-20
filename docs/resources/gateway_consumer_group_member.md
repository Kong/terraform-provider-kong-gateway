---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kong-gateway_gateway_consumer_group_member Resource - terraform-provider-kong-gateway"
subcategory: ""
description: |-
  GatewayConsumerGroupMember Resource
---

# kong-gateway_gateway_consumer_group_member (Resource)

GatewayConsumerGroupMember Resource

## Example Usage

```terraform
resource "kong-gateway_gateway_consumer_group_member" "my_gatewayconsumergroupmember" {
  consumer_group_id       = "...my_consumer_group_id..."
  consumer_id             = "cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b"
  consumer_id_or_username = "...my_consumer_id_or_username..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `consumer_group_id` (String) Requires replacement if changed.
- `consumer_id_or_username` (String)

### Optional

- `consumer_id` (String) Requires replacement if changed.

### Read-Only

- `consumer_group` (Attributes) (see [below for nested schema](#nestedatt--consumer_group))
- `consumers` (Attributes List) (see [below for nested schema](#nestedatt--consumers))

<a id="nestedatt--consumer_group"></a>
### Nested Schema for `consumer_group`

Read-Only:

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String)
- `name` (String)
- `tags` (List of String)
- `updated_at` (Number) Unix epoch when the resource was last updated.


<a id="nestedatt--consumers"></a>
### Nested Schema for `consumers`

Read-Only:

- `created_at` (Number) Unix epoch when the resource was created.
- `custom_id` (String) Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or `username` with the request.
- `id` (String)
- `tags` (List of String) An optional set of strings associated with the Consumer for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.
- `username` (String) The unique username of the Consumer. You must send either this field or `custom_id` with the request.
