resource "kong-gateway_plugin_basic_auth" "basic_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  service = {
    id = kong-gateway_service.httpbin.id
  }
  config = {
    hide_credentials = false
  }
}

resource "kong-gateway_plugin_key_auth" "my_key_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    key_names = ["apikey"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
resource "kong-gateway_plugin_key_auth_enc" "my_key_auth_enc" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    key_names    = ["apikey"]
    key_in_query = false
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_hmac_auth" "my_hmac_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    enforce_headers       = ["date", "@request-target"]
    algorithms            = ["hmac-sha1", "hmac-sha256"]
    validate_request_body = true
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_header_cert_auth" "my_header_cert_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    ca_certificates           = ["322dce96-d434-4e0d-9038-311b3520f0a3"]
    certificate_header_name   = "my-header-name"
    certificate_header_format = "base64_encoded"
    secure_source             = false
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_jwt" "my_jwt" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    cookie_names = ["jwt"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_jwt_signer" "my_jwt_signer" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    access_token_introspection_scopes_claim = ["scope"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ldap_auth" "my_ldap_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    ldap_host        = "my-ldap-host"
    ldap_port        = 636
    ldaps            = true
    base_dn          = "dc=example,dc=com"
    verify_ldap_host = false
    attribute        = "cn"
    cache_ttl        = 60
    header_type      = "ldap"
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ldap_auth_advanced" "my_ldap_auth_advanced" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    ldap_host        = "my-ldap-host"
    ldap_port        = 389
    start_tls        = true
    base_dn          = "dc=example,dc=com"
    verify_ldap_host = false
    attribute        = "cn"
    cache_ttl        = 60
    header_type      = "ldap"
    consumer_by      = ["username"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_mtls_auth" "my_mtls_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    ca_certificates = [kong-gateway_ca_certificate.my_ca_cert.id]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_oauth2" "my_oauth2" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    scopes                    = ["email"]
    provision_key             = "my-provision-key"
    enable_client_credentials = true
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_oauth2_introspection" "my_oauth2_introspection" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    introspection_url     = "https://example.com/oauth2/introspect"
    authorization_value   = "Bearer my-access-token"
    consumer_by           = "client_id"
    custom_claims_forward = ["my_custom_claim"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_openid_connect" "my_openid_connect" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    issuer              = "https://issuer.example.com"
    client_id           = ["client_id_123"]
    client_secret       = ["client_secret_123"]
    client_auth         = ["client_secret_post"]
    auth_methods        = ["authorization_code", "session"]
    response_mode       = "form_post"
    preserve_query_args = true
    login_action        = "redirect"
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_saml" "my_saml" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    anonymous                    = "anonymous"
    issuer                       = "my-issuer"
    idp_sso_url                  = "https://example.com/sso"
    assertion_consumer_path      = "/consume"
    validate_assertion_signature = false
    session_secret               = "uwcLGoTJCWnHWZdVpbLYKlztNOyoGJ07"
    idp_certificate              = "my-certificate"
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_session" "my_session" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    storage = "cookie"
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_upstream_oauth" "my_upstream_oauth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {

    oauth = {
      token_endpoint = "https://example.com/oauth2/token"
      grant_type     = "client_credentials"
      client_id      = "my-client-id"
      client_secret  = "my-client-secret"
      scopes         = ["openid", "profile"]
    }

    behavior = {
      upstream_access_token_header_name = "X-Custom-Auth"
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_vault_auth" "my_vault_auth" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    vault = {
      id = kong-gateway_vault.my_vault.id
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
