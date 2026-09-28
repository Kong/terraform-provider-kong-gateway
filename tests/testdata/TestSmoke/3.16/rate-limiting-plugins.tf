resource "kong-gateway_plugin_rate_limiting" "my_rate_limiting_expressions" {
  enabled = true

  config = {
    policy = "local"
    hour = 1000

    expressions = {
      custom_key = "net.src.ip"
      hour   = "1000"
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_rate_limiting_advanced" "my_rate_limiting_advanced_expressions" {
  enabled = true

  config = {
    custom_key   = "net.src.ip"
    limit       = [200]
    window_size = [1800]
    window_type = "fixed"
    namespace   = "my-namespace"
    header_name = "X-RateLimit-Limit"
    redis = {
      host = "redis.example.com"
      port = 6379
    }
  }

  expressions = {
    custom_key = "net.src.ip"
    limit      = ["200"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }

}
