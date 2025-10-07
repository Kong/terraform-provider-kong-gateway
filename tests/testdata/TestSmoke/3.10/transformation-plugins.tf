resource "kong-gateway_plugin_confluent" "my_confluent" {
  enabled = true

  config = {
    bootstrap_servers = [
      {
        host = "example.com"
        port = 9092
    }]
    producer_async     = false
    topic              = "kong-test"
    cluster_api_key    = "your-static-cluster-api-key"
    cluster_api_secret = "your-static-cluster-api-secret"
  }
}

resource "kong-gateway_plugin_correlation_id" "my_correlation_id" {
  enabled = true

  config = {
    header_name     = "Kong-Request-ID"
    generator       = "uuid#counter"
    echo_downstream = false
  }
}

resource "kong-gateway_plugin_degraphql" "my_degraphql" {
  enabled = true
}

resource "kong-gateway_plugin_exit_transformer" "my_exit_transformer" {
  enabled = true

  config = {
    handle_unknown = true
    functions      = ["return function(status, body, headers) if status == 404 then local new_body = { error = true, status = status, message = \"This is not the Route you are looking for\", } return status, new_body, headers else return status, body, headers end end"]
  }
}

resource "kong-gateway_plugin_kafka_upstream" "my_kafka_upstream" {
  enabled = true

  config = {
    topic = "your-static-kafka-topic"
  }
}

resource "kong-gateway_plugin_request_callout" "my_request_callout" {
  enabled = true

  config = {
    callouts = [
      {
        name = "c1"

        request = {
          url    = "http://httpbin.org/uuid"
          method = "GET"
        }

        response = {

          body = {
            decode = true
          }
        }
      },

      {
        name = "c2"

        request = {
          url    = "http://httpbin.org/anything"
          method = "GET"
        }

        response = {

          body = {
            decode = true
          }
        }
    }]

    upstream = {
      by_lua = "kong.response.exit(200, { uuid = kong.ctx.shared.callouts.c1.response.body.uuid, origin = kong.ctx.shared.callouts.c2.response.body.url})"
    }
  }
}

resource "kong-gateway_plugin_request_transformer" "my_request_transformer" {
  enabled = true

  config = {

    add = {
      headers = ["New-Header:Header Value"]
    }
  }
}

resource "kong-gateway_plugin_request_transformer_advanced" "my_request_transformer_advanced" {
  enabled = true

  config = {

    replace = {
      uri = "/status/\\$(uri_captures[\"status\"])"
    }
  }
}

resource "kong-gateway_plugin_response_transformer" "my_response_transformer" {
  enabled = true

  config = {

    add = {
      headers = ["New-Header:Header Value"]
    }
  }
}

resource "kong-gateway_plugin_response_transformer_advanced" "my_response_transformer_advanced" {
  enabled = true

  config = {

    remove = {
      json = ["customers.info.phone"]
    }
    dots_in_keys = false
  }
}

resource "kong-gateway_plugin_route_transformer_advanced" "my_route_transformer_advanced" {
  enabled = true

  config = {
    path = "/$(headers[\"Custom-Path\"])"
  }
}

resource "kong-gateway_plugin_grpc_gateway" "my_grpc_gateway" {
  enabled = true

  config = {
    proto = "path/to/hello-gateway.proto"
  }
}

resource "kong-gateway_plugin_grpc_web" "my_grpc_web" {
  enabled = true

  config = {
    proto = "path/to/hello.proto"
  }
}

resource "kong-gateway_plugin_jq" "my_jq" {
  enabled = true

  config = {
    request_jq_program = <<EOF
select(.name == "James Dean").name = "John Doe"
EOF
  }
}
