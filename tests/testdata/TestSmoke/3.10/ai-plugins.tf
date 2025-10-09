resource "kong-gateway_plugin_ai_azure_content_safety" "my_ai_azure_content_safety" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    content_safety_url = "https://example.com"
    content_safety_key = "sample-key"
    categories = [
      {
        name            = "Hate"
        rejection_level = 2
      },

      {
        name            = "Violence"
        rejection_level = 2
    }]
  }

  route = {
    id = kong-gateway_route.hello.id
  }
}

resource "kong-gateway_plugin_ai_sanitizer" "my_ai_sanitizer" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    anonymize = ["email", "phone", "ssn", "creditcard", "custom"]
    custom_patterns = [
      {
        name  = "aws_api_key"
        regex = "AKIA[0-9A-Z]{16}"
        score = 0.95
      },

      {
        name  = "github_token"
        regex = "ghp_[A-Za-z0-9]{36}"
        score = 0.9
      },

      {
        name  = "jwt_token"
        regex = "eyJ[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+"
        score = 0.85
    }]
    port             = 1234
    host             = "localhost"
    redact_type      = "placeholder"
    stop_on_error    = true
    recover_redacted = false
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_prompt_decorator" "my_ai_prompt_decorator" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {

    prompts = {
      prepend = [
        {
          role    = "system"
          content = "You are data scientist, specialising in survey analytics."
        },

        {
          role    = "user"
          content = "Classify this test result set as positive, negative, or neutral."
        },

        {
          role    = "assistant"
          content = "These tests are NEUTRAL."
      }]
      append = [
        {
          role    = "user"
          content = "Do not mention any real participant names in your justification."
      }]
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_prompt_guard" "my_ai_prompt_guard" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]


  config = {
    allow_patterns = ["(?i).*what is .*", "(?i).*how do i .*", "(?i).*install .*", "(?i).*configure .*", "(?i).*reset .*", "(?i).*troubleshoot .*"]
    deny_patterns  = ["(?i).*bypass.*(login|password|auth).*", "(?i).*hack.*", "(?i).*phish.*", "(?i).*malware.*", "(?i).*cve.*", "(?i).*exploit.*", "(?i).*social engineering.*", "(?i).*pentest.*", "(?i).*impersonate.*", "(?i).*dating.*"]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_prompt_template" "my_ai_prompt_template" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]


  config = {
    allow_untemplated_requests = false
    templates = [
      {
        name     = "developer-chat"
        template = <<EOF
{
    "messages": [
    {
        "role": "system",
        "content": "You are a {{program}} expert, in {{language}} programming language."
    },
    {
        "role": "user",
        "content": "Write me a {{program}} program."
    }
  ]
}
EOF
    }]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_proxy" "my_ai_proxy" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    route_type = "llm/v1/chat"

    auth = {
      header_name  = "Authorization"
      header_value = "Bearer var.openai_api_key"
    }

    model = {
      provider = "openai"
      name     = "gpt-4"

      options = {
        max_tokens  = 512
        temperature = 1.0
      }
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_proxy_advanced" "my_ai_proxy_advanced" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    targets = [
      {
        route_type = "llm/v1/chat"

        auth = {
          header_name  = "Authorization"
          header_value = "Bearer var.openai_api_key"
        }

        model = {
          provider = "openai"
          name     = "gpt-4"

          options = {
            max_tokens  = 512
            temperature = 1.0
          }
        }
    }]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_rag_injector" "my_ai_rag_injector" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]


  config = {
    inject_template = <<EOF
Only use the following information surrounded by <RAG></RAG>to and your existing knowledge to provide the best possible answer to the user.
<RAG><CONTEXT></RAG>
User's question: <PROMPT>
EOF

    embeddings = {

      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }

      model = {
        provider = "openai"
        name     = "text-embedding-3-large"
      }
    }

    vectordb = {
      strategy = "redis"

      redis = {
        host = "localhost"
        port = 6379
      }
      distance_metric = "cosine"
      dimensions      = 76
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_rate_limiting_advanced" "my_ai_rate_limiting_advanced" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    llm_providers = [
      {
        name        = "openai"
        limit       = [100, 1000]
        window_size = [60, 3600]
    }]
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_request_transformer" "my_ai_request_transformer" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    prompt                         = "In my JSON message, anywhere there is a JSON tag for a city, also add a country tag with the name of the country that city is in."
    transformation_extract_pattern = "{((.|\n)*)}"

    llm = {
      route_type = "llm/v1/chat"

      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }

      model = {
        provider = "openai"
        name     = "gpt-4"
      }
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_response_transformer" "my_ai_response_transformer" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {
    prompt                         = "In my JSON message, anywhere there is a JSON tag for a city, also add a country tag with the name of the country that city is in."
    transformation_extract_pattern = "{((.|\n)*)}"

    llm = {
      route_type = "llm/v1/chat"

      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }

      model = {
        provider = "openai"
        name     = "gpt-4"
      }
    }
  }

  route = {
    id = kong-gateway_route.hello.id
  }
}

resource "kong-gateway_plugin_ai_semantic_cache" "my_ai_semantic_cache" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {

    embeddings = {

      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.mistral_api_key"
      }

      model = {
        provider = "mistral"
        name     = "mistral-embed"

        options = {
          upstream_url = "https://api.mistral.ai/v1/embeddings"
        }
      }
    }

    vectordb = {
      dimensions      = 1024
      distance_metric = "cosine"
      strategy        = "redis"
      threshold       = 0.1

      redis = {
        host = "redis-stack.redis.svc.cluster.local"
        port = 6379
      }
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}

resource "kong-gateway_plugin_ai_semantic_prompt_guard" "my_ai_semantic_prompt_guard" {
  enabled = true

  protocols = [
    "http",
    "https"
  ]

  config = {

    embeddings = {

      auth = {
        header_name  = "Authorization"
        header_value = "Bearer var.openai_api_key"
      }

      model = {
        name     = "text-embedding-3-small"
        provider = "openai"
      }
    }

    search = {
      threshold = 0.7
    }

    vectordb = {
      strategy        = "redis"
      distance_metric = "cosine"
      threshold       = 0.5
      dimensions      = 1024

      redis = {
        host = "localhost"
        port = 6379
      }
    }

    rules = {
      match_all_conversation_history = true
      allow_prompts                  = ["Network troubleshooting and diagnostics", "Cloud infrastructure management (AWS, Azure, GCP)", "Cybersecurity best practices and incident response", "DevOps workflows and automation", "Programming concepts and language usage", "IT policy and compliance guidance", "Software development lifecycle and CI/CD", "Documentation writing and technical explanation", "System administration and configuration", "Productivity and collaboration tools usage"]
      deny_prompts                   = ["Hacking techniques or penetration testing without authorization", "Bypassing software licensing or digital rights management", "Instructions on exploiting vulnerabilities or writing malware", "Circumventing security controls or access restrictions", "Gathering personal or confidential employee information", "Using AI to impersonate or phish others", "Social engineering tactics or manipulation techniques", "Guidance on violating company IT policies", "Content unrelated to work, such as entertainment or dating", "Political, religious, or sensitive non-work-related discussions"]
    }
  }

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
