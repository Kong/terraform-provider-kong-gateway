
# RBAC Resources
resource "kong-gateway_rbac_role" "my_rbacrole" {
  name      = "Test Role 17"
  comment   = "Some Comment"
  workspace = kong-gateway_workspace.demo.name
}

resource "kong-gateway_rbac_role_entity" "my_rbacroleentity" {
  entity_id   = kong-gateway_service.httpbin_in_workspace.id
  entity_type = "services"
  actions     = ["create", "update", "read"]
  role_id     = kong-gateway_rbac_role.my_rbacrole.id
  workspace   = kong-gateway_workspace.demo.name
}

resource "kong-gateway_rbac_role_endpoint" "my_rbacroleendpoint" {
  actions   = ["create", "delete"]
  endpoint  = "/*"
  role_id   = kong-gateway_rbac_role.my_rbacrole.id
  workspace = kong-gateway_workspace.demo.name
}

resource "kong-gateway_rbac_user" "my_rbacuser" {
  enabled    = true
  name       = "alice"
  user_token = "alice1"
  comment    = "This is a test user"
  workspace  = kong-gateway_workspace.demo.name
}

resource "kong-gateway_rbac_user_role" "my_userrole" {
  user_id   = kong-gateway_rbac_user.my_rbacuser.id
  role      = kong-gateway_rbac_role.my_rbacrole.name
  workspace = kong-gateway_workspace.demo.name
}

resource "kong-gateway_group" "my_group" {
  name = "mytestgroup"
}

resource "kong-gateway_group_role" "my_grouprole" {
  group_id  = kong-gateway_group.my_group.id
  role_id   = kong-gateway_rbac_role.my_rbacrole.id
  workspace = kong-gateway_workspace.demo.id
}
