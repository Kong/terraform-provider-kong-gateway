resource "kong-gateway_plugin_datadog" "my_datadog" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    host          = "127.0.0.1"
    port          = 8125
    flush_timeout = 2
    retry_count   = 10
    metrics = [
      {
        name                = "request_count"
        stat_type           = "counter"
        consumer_identifier = "consumer_id"
        sample_rate         = 1
    }]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_opentelemetry" "my_opentelemetry" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    traces_endpoint = "http://localhost:4318/v1/traces"
    logs_endpoint   = "http://localhost:4318/v1/logs"

    headers = {
      X-Auth-Token = "secret-token"
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_prometheus" "my_prometheus" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    per_consumer = true
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_statsd" "my_statsd" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    host               = "localhost"
    port               = 8125
    allow_status_codes = ["200-205", "400-499"]
    flush_timeout      = 2
    retry_count        = 10
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_zipkin" "my_zipkin" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {

    propagation = {
      extract        = ["w3c", "b3", "jaeger", "ot", "aws", "datadog"]
      clear          = ["b3", "uber-trace-id"]
      inject         = ["w3c"]
      default_format = "w3c"
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
