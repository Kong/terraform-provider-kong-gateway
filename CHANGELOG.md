# Changelog

## 0.5.2
> Released on ?

### Bug Fixes

* The provider now works with terraform v1.12.1

## 0.5.1
> Released on 2025/06/09

### Bug Fixes

* Fixed a bug in the `file_log` plugin `path` validation that disallowed `/dev/stdout` as a path

## 0.5.0
> Released on 2025/05/01

### BREAKING CHANGES

* The properties `limit` and `window_size` now accept array of values in `ai-rate-limiting-advanced` plugin according to the Gateway 3.10 schema

### Features

* Add support for new Gateway 3.10 plugins:
  * `gateway_plugin_ai_rag_injector`
  * `gateway_plugin_ai_sanitizer`
  * `gateway_plugin_confluent_consume`
  * `gateway_plugin_kafka_consume`
  * `gateway_plugin_request_callout`

### Bug Fixes

* The `partials` field on Plugins is no longer sent by default, enabling support for versions of Gateway < 3.10

## 0.4.0
> Released on 2025/04/11

### BREAKING CHANGES

* Credentials (e.g. `basic_auth`, `key_auth`) now require a `consumer_id` to be provided

### Features
* Add support for `route_expression` `oidc_jwk`, `partial` and `target` resouces
* Add support for the following plugins
  * AI Azure Content Safety
  * AI Proxy Advanced
  * AI Rate Limiting Advanced
  * AI Semantic Cache
  * AI Semantic Prompt Guard
  * Confluent
  * Datadog Tracing
  * Header Cert Auth
  * Injection Protection
  * JSON Threat Protection
  * Redirect
  * Service Protection
  * Standard Webhooks
  * Upstream OAuth


## 0.3.0
> Released on 2024/11/26

### Features
* Add support for `admin_token` in the provider configuration to access secured admin APIs

## 0.2.1
> Released on 2024/09/16

### Bug Fixes
* Publish resource documentation on Terraform registry

## 0.2.0
> Released on 2024/06/11

### Features
* Add support for all plugins

## 0.1.0
> Released on 2024/06/18

* Initial release