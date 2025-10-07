resource "kong-gateway_plugin_aws_lambda" "my_aws_lambda" {
  enabled = true

  config = {
    aws_key             = "my-aws-key"
    aws_secret          = "my-aws-secret"
    aws_region          = "us-east-1"
    aws_assume_role_arn = "arn:aws:iam::123456789012:role/my-role"
    function_name       = "my-lambda-function"
    invocation_type     = "RequestResponse"
  }
}

resource "kong-gateway_plugin_azure_functions" "my_azure_functions" {
  enabled = true

  config = {
    functionname = "my-azure-function"
    appname      = "my-azure-app"
    hostdomain   = "azurewebsites.net"
    apikey       = "my-azure-api-key"
  }
}

resource "kong-gateway_plugin_post_function" "my_post_function" {
  enabled = true

  config = {
    access = ["ngx.var.upstream_x_forwarded_host=nil"]
  }
}

resource "kong-gateway_plugin_pre_function" "my_pre_function" {
  enabled = true

  config = {
    access = [<<EOF
local existing_request_id = kong.request.get_header("Request-Id")

if(existing_request_id ~= nil) then
  kong.service.request.set_header("New-Request-Id", existing_request_id)
end
EOF
    ]
  }
}
