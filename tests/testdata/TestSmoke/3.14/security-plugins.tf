resource "kong-gateway_plugin_bot_detection" "my_bot_detection" {
  enabled = true

  config = {
    allow = []
    deny  = []
  }
}

resource "kong-gateway_plugin_cors" "my_cors" {
  enabled = true

  config = {
    origins = ["http://mockbin.com"]
    methods = ["GET", "POST"]
  }
}

resource "kong-gateway_plugin_ip_restriction" "my_ip_restriction" {
  enabled = true

  config = {
    allow = ["203.0.113.1", "192.0.2.0/24"]
  }
}

resource "kong-gateway_plugin_injection_protection" "my_injection_protection" {
  enabled = true

  config = {
    injection_types   = ["sql", "java_exception", "js", "ssi", "xpath_abbreviated", "xpath_extended"]
    locations         = ["path_and_query"]
    enforcement_mode  = "block"
    error_status_code = 400
    error_message     = "Bad Request"
  }
}

resource "kong-gateway_plugin_json_threat_protection" "my_json_threat_protection" {
  enabled = true

  config = {
    max_body_size                = 1024
    max_container_depth          = 2
    max_object_entry_count       = 4
    max_object_entry_name_length = 7
    max_array_element_count      = 2
    max_string_value_length      = 6
    enforcement_mode             = "block"
    error_status_code            = 400
    error_message                = "Incorrect request format"
  }
}

resource "kong-gateway_plugin_opa" "my_opa" {
  enabled = true

  config = {
    opa_host = "example.com"
    opa_port = 8182
    opa_path = "/v1/data/my_policies/header"
  }
}

resource "kong-gateway_plugin_tls_handshake_modifier" "my_tls_handshake_modifier" {
  enabled = true
}

resource "kong-gateway_plugin_tls_metadata_headers" "my_tls_metadata_headers" {
  enabled = true

  config = {
    inject_client_cert_details = true
    client_cert_header_name    = "X-Forwarded-Client-Cert"
  }
}

resource "kong-gateway_plugin_acme" "my_acme" {
  config = {
    account_email = "user@example.com"
    account_key = {
      key_id  = "my_key_id"
      key_set = "my_key_set"
    }
    domains      = ["example.com"]
    tos_accepted = true
    storage      = "kong"
  }
}
