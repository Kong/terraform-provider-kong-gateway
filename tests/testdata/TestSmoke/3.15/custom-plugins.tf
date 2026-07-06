resource "kong-gateway_custom_plugin_streaming" "my_streaming_plugin" {
  name = "set-header"

  handler = <<-EOT
    return {
      VERSION = "1.0,0",
      PRIORITY = 500,
      access = function(self, config)
        kong.service.request.set_header(config.name, config.value)
      end
    }
  EOT

  schema = <<-EOT
    return {
      name = "set-header",
      fields = {
        { protocols = require("kong.db.schema.typedefs").protocols_http },
        {
          config = {
            type = "record",
            fields = {
              { name = { description = "The name of the header to set.", type = "string", required = true, }, },
              { value = { description = "The value for the header.", type = "string", required = true, }, },
            },
          },
        },
      },
    }
  EOT
}


resource "kong-gateway_custom_plugin" "seth1" {
  name = "set-header"
  instance_name = "setheader-instance"
  config = {
    name = "x-custom-header"
    value = "my-custom-value"
  }

  depends_on = [ kong-gateway_custom_plugin_streaming.my_streaming_plugin ]
}


resource "kong-gateway_cloned_plugin" "my_cloned_plugin" {
  name = "custom-acl"
  ref  = "acl"
}

resource "kong-gateway_custom_plugin" "my_custom_acl" {
  name = "custom-acl"
  config = {
    allow = ["mygroup"]
  }

  depends_on = [ kong-gateway_cloned_plugin.my_cloned_plugin ]
}
