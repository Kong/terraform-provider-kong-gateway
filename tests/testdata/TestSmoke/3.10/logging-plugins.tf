resource "kong-gateway_plugin_http_log" "my_http_log" {
  enabled = true

  config = {
    http_endpoint = "https://example.com"

    headers = {
      Authorization = "Bearer auth_token"
    }
    method      = "POST"
    timeout     = 3000
    retry_count = 1
  }
}

resource "kong-gateway_plugin_kafka_log" "my_kafka_log" {
  enabled = true

  config = {
    topic = "my-topic"

    authentication = {
      strategy  = "sasl"
      mechanism = "SCRAM-SHA-256"
      user      = "token-here"
      password  = "password-here"
      tokenauth = true
    }
  }
}

resource "kong-gateway_plugin_loggly" "my_loggly" {
  enabled = true

  config = {
    key = "service-token"
  }
}

resource "kong-gateway_plugin_syslog" "my_syslog" {
  enabled = true

  config = {
    log_level = "info"
  }
}

resource "kong-gateway_plugin_tcp_log" "my_tcp_log" {
  enabled = true

  config = {
    host = "127.0.0.1"
    port = 8080
  }
}

resource "kong-gateway_plugin_udp_log" "my_udp_log" {
  enabled = true

  config = {
    host = "127.0.0.1"
    port = 8080
  }
}

