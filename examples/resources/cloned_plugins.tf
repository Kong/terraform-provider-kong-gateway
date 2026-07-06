resource "kong-gateway_cloned_plugin" "my_cloned_plugin" {
  name = "custom-acl"
  ref  = "acl"
}

resource "kong-gateway_custom_plugin" "acl" {
  name = "custom-acl"
  instance_name = "custom-acl-instance"
  config = {
    allow = ["mygroup"]
  }

  depends_on = [ kong-gateway_cloned_plugin.my_cloned_plugin ]
}