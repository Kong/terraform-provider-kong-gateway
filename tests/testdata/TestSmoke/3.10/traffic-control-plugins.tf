resource "kong-gateway_plugin_acl" "my_acl" {
  enabled = true

  config = {
    include_consumer_groups = true
    allow                   = ["dev", "admin"]
  }
}

resource "kong-gateway_plugin_canary" "my_canary" {
  enabled = true

  config = {
    groups        = ["example-group"]
    upstream_host = "example-upstream-host"
    upstream_port = 80
    hash          = "allow"
  }
}

resource "kong-gateway_plugin_forward_proxy" "my_forward_proxy" {
  enabled = true

  config = {
    http_proxy_host = "example-proxy-hostname"
    http_proxy_port = 80
    proxy_scheme    = "http"
    https_verify    = false
    x_headers       = "transparent"
    auth_username   = "example-proxy-username"
    auth_password   = "example-proxy-password"
  }
}

resource "kong-gateway_plugin_graphql_proxy_cache_advanced" "my_graphql_proxy_cache_advanced" {
  enabled = true

  config = {
    strategy  = "memory"
    cache_ttl = 500
  }
}

resource "kong-gateway_plugin_graphql_rate_limiting_advanced" "my_graphql_rate_limiting_advanced" {
  enabled = true

  config = {
    limit         = [100]
    window_size   = [60]
    window_type   = "fixed"
    cost_strategy = "node_quantifier"
    max_cost      = 5000
    sync_rate     = 0
  }
}

resource "kong-gateway_plugin_mocking" "my_mocking" {
  enabled = true

  config = {
    api_specification = <<EOF
openapi: 3.0.2
servers:
  - url: /v3
info:
  description: |-
    This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about
    Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!
    You can now help us improve the API whether it's by making changes to the definition itself or to the code.
    That way, with time, we can improve the API in general, and expose some of the new features in OAS3.
    Some useful links:
    - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
    - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
  version: 1.0.20-SNAPSHOT
  title: Swagger Petstore - OpenAPI 3.0
  termsOfService: 'http://swagger.io/terms/'
  contact:
    email: apiteam@swagger.io
  license:
    name: Apache 2.0
    url: 'http://www.apache.org/licenses/LICENSE-2.0.html'
tags:
  - name: pet
    description: Everything about your Pets
    externalDocs:
      description: Find out more
      url: 'http://swagger.io'
  - name: store
    description: Access to Petstore orders
    externalDocs:
      description: Find out more about our store
      url: 'http://swagger.io'
  - name: user
    description: Operations about user
paths:
  /pet:
    post:
      tags:
        - pet
      summary: Add a new pet to the store
      description: Add a new pet to the store
      operationId: addPet
      responses:
        '200':
          description: Successful operation
          content:
            application/xml:
              schema:
                $ref: '#/components/schemas/Pet'
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        '405':
          description: Invalid input
      security:
        - petstore_auth:
            - 'write:pets'
            - 'read:pets'
      requestBody:
        description: Create a new pet in the store
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet'
          application/xml:
            schema:
              $ref: '#/components/schemas/Pet'
          application/x-www-form-urlencoded:
            schema:
              $ref: '#/components/schemas/Pet'
    put:
      tags:
        - pet
      summary: Update an existing pet
      description: Update an existing pet by Id
      operationId: updatePet
      responses:
        '200':
          description: Successful operation
          content:
            application/xml:
              schema:
                $ref: '#/components/schemas/Pet'
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        '400':
          description: Invalid ID supplied
        '404':
          description: Pet not found
        '405':
          description: Validation exception
      security:
        - petstore_auth:
            - 'write:pets'
            - 'read:pets'
      requestBody:
        description: Update an existent pet in the store
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet'
          application/xml:
            schema:
              $ref: '#/components/schemas/Pet'
          application/x-www-form-urlencoded:
            schema:
              $ref: '#/components/schemas/Pet'
EOF
  }
}

resource "kong-gateway_plugin_oas_validation" "my_oas_validation" {
  enabled = true

  config = {
    api_spec         = <<EOF
openapi: 3.0.2
servers:
  - url: /v3
info:
  description: |-
    This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about
    Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!
    You can now help us improve the API whether it's by making changes to the definition itself or to the code.
    That way, with time, we can improve the API in general, and expose some of the new features in OAS3.
    Some useful links:
    - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
    - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
  version: 1.0.20-SNAPSHOT
  title: Swagger Petstore - OpenAPI 3.0
  termsOfService: 'http://swagger.io/terms/'
  contact:
    email: apiteam@swagger.io
  license:
    name: Apache 2.0
    url: 'http://www.apache.org/licenses/LICENSE-2.0.html'
tags:
  - name: pet
    description: Everything about your Pets
    externalDocs:
      description: Find out more
      url: 'http://swagger.io'
  - name: store
    description: Access to Petstore orders
    externalDocs:
      description: Find out more about our store
      url: 'http://swagger.io'
  - name: user
    description: Operations about user
paths:
  /pet:
    post:
      tags:
        - pet
      summary: Add a new pet to the store
      description: Add a new pet to the store
      operationId: addPet
      responses:
        '200':
          description: Successful operation
          content:
            application/xml:
              schema:
                $ref: '#/components/schemas/Pet'
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        '405':
          description: Invalid input
      security:
        - petstore_auth:
            - 'write:pets'
            - 'read:pets'
      requestBody:
        description: Create a new pet in the store
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet'
          application/xml:
            schema:
              $ref: '#/components/schemas/Pet'
          application/x-www-form-urlencoded:
            schema:
              $ref: '#/components/schemas/Pet'
    put:
      tags:
        - pet
      summary: Update an existing pet
      description: Update an existing pet by Id
      operationId: updatePet
      responses:
        '200':
          description: Successful operation
          content:
            application/xml:
              schema:
                $ref: '#/components/schemas/Pet'
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        '400':
          description: Invalid ID supplied
        '404':
          description: Pet not found
        '405':
          description: Validation exception
      security:
        - petstore_auth:
            - 'write:pets'
            - 'read:pets'
      requestBody:
        description: Update an existent pet in the store
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet'
          application/xml:
            schema:
              $ref: '#/components/schemas/Pet'
          application/x-www-form-urlencoded:
            schema:
              $ref: '#/components/schemas/Pet'
EOF
    verbose_response = true
  }
}

resource "kong-gateway_plugin_proxy_cache" "my_proxy_cache" {
  enabled = true

  config = {
    response_code  = [200]
    request_method = ["GET"]
    content_type   = ["text/plain", "application/json"]
    cache_ttl      = 300
    strategy       = "memory"
  }
}

resource "kong-gateway_plugin_proxy_cache_advanced" "my_proxy_cache_advanced" {
  enabled = true

  config = {
    response_code  = [200]
    request_method = ["GET"]
    content_type   = ["text/plain", "application/json"]
    cache_ttl      = 300
    strategy       = "memory"
  }
}

resource "kong-gateway_plugin_rate_limiting" "my_rate_limiting" {
  enabled = true

  config = {
    second = 5
    hour   = 1000
    policy = "local"
  }
}

resource "kong-gateway_plugin_rate_limiting_advanced" "my_rate_limiting_advanced" {
  enabled = true

  config = {
    limit       = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace   = "my-namespace"
  }
}

resource "kong-gateway_plugin_redirect" "my_redirect" {
  enabled = true

  config = {
    location           = "https://example-redirect-url.com"
    keep_incoming_path = true
  }
}

resource "kong-gateway_plugin_request_size_limiting" "my_request_size_limiting" {
  enabled = true

  config = {
    allowed_payload_size = 256000000
    size_unit            = "bytes"
  }
}

resource "kong-gateway_plugin_request_termination" "my_request_termination" {
  enabled = true

  config = {
    status_code = 401
    message     = "\"Error - Authentication required\""
  }
}

resource "kong-gateway_plugin_request_validator" "my_request_validator" {
  enabled = true

  config = {
    version = "draft4"
    parameter_schema = [
      {
        name     = "status_code"
        in       = "path"
        required = true
        schema   = "{\"type\": \"number\"}"
        style    = "simple"
        explode  = false
    }]
  }
}

resource "kong-gateway_plugin_response_ratelimiting" "my_response_ratelimiting" {
  enabled = true

  config = {
    limits = {
      demo-limit = {
        minute = 10
      }
    }
    policy = "local"
  }
}

resource "kong-gateway_plugin_route_by_header" "my_route_by_header" {
  enabled = true

  config = {
    rules = [
      {
        upstream_name = "east-upstream"

        condition = {
          location = "us-east"
        }
      },

      {
        upstream_name = "west-upstream"

        condition = {
          location = "us-west"
        }
    }]
  }
}

resource "kong-gateway_plugin_service_protection" "my_service_protection" {
  enabled = true

  config = {
    window_size = [60]
    window_type = "sliding"
    limit       = [5]
    namespace   = "example_namespace"
  }
}

resource "kong-gateway_plugin_standard_webhooks" "my_standard_webhooks" {
  enabled = true

  config = {
    secret_v1 = "example-webhook-secret-key"
  }
}

resource "kong-gateway_plugin_upstream_timeout" "my_upstream_timeout" {
  enabled = true

  config = {
    connect_timeout = 5000
    read_timeout    = 5000
    send_timeout    = 5000
  }
}

resource "kong-gateway_plugin_websocket_size_limit" "my_websocket_size_limit" {
  enabled = true

  config = {
    client_max_payload   = 4096
    upstream_max_payload = 1048576
  }
}

resource "kong-gateway_plugin_websocket_validator" "my_websocket_validator" {
  enabled = true

  config = {

    client = {

      text = {
        schema = "{ \"type\": \"object\", \"required\": [\"name\"] }"
        type   = "draft4"
      }
    }
  }
}

resource "kong-gateway_plugin_xml_threat_protection" "my_xml_threat_protection" {
  enabled = true

  config = {
    max_depth    = 50
    localname    = 512
    prefix       = 512
    namespaceuri = 1024
  }
}
