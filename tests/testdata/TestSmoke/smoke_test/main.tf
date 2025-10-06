provider "kong-gateway" {
  server_url  = "http://localhost:8001"
  admin_token = "kongFTW"
}

resource "kong-gateway_workspace" "demo" {
  name = "demo"
}

resource "kong-gateway_service" "httpbin_in_workspace" {
  name     = "HTTPBin"
  protocol = "http"
  host     = "httpbin.org"
  port     = 443
  path     = "/"

  workspace = kong-gateway_workspace.demo.name
}

resource "kong-gateway_service" "httpbin" {
  name     = "HTTPBin"
  protocol = "http"
  host     = "httpbin.org"
  port     = 443
  path     = "/"
}

resource "kong-gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  protocols = [
    "http",
    "https"
  ]

  strip_path = false

  service = {
    id = kong-gateway_service.httpbin.id
  }

}

resource "kong-gateway_consumer" "alice" {
  username  = "alice"
  custom_id = "alice"
}

resource "kong-gateway_consumer_group" "alpha" {
  name = "alpha"
  tags = ["demo1234"]
}

resource "kong-gateway_consumer_group_member" "my_consumergroupmember" {
  consumer_group_id = kong-gateway_consumer_group.alpha.id
  consumer_id       = kong-gateway_consumer.alice.id
}

resource "kong-gateway_acl" "my_acl" {
  group       = "demo"
  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_basic_auth" "my_basic_auth" {
  username    = "alice"
  password    = "password"
  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_key_auth" "my_key_auth" {
  key         = "demo123"
  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_ca_certificate" "my_ca_cert" {
  cert = <<EOT
-----BEGIN CERTIFICATE-----
MIIByDCCAU6gAwIBAgIUTiH53ur+PTurNUfIav6840EPkwQwCgYIKoZIzj0EAwIw
GzEZMBcGA1UEAwwQZGVtby5leGFtcGxlLmNvbTAeFw0yNTEwMDMxNjE0MDdaFw0y
ODEwMDIxNjE0MDdaMBsxGTAXBgNVBAMMEGRlbW8uZXhhbXBsZS5jb20wdjAQBgcq
hkjOPQIBBgUrgQQAIgNiAASpCbbkPRiGKR/nutM0gG8uIK1xKaEQMpEKqmbgn9mM
gqg2q9yLnQfnIpG7G/VE67WGS0PoonSN8v2nMj9XLYipK7oNYoOMlAigW2p4LQLO
iDkx7/hRXKDCcCHEPFOvVzWjUzBRMB0GA1UdDgQWBBQ/2XvmGDSh+Db5v+rm5VEl
2hoUODAfBgNVHSMEGDAWgBQ/2XvmGDSh+Db5v+rm5VEl2hoUODAPBgNVHRMBAf8E
BTADAQH/MAoGCCqGSM49BAMCA2gAMGUCMQCLHZGPqT0SVcSVYtHp3CefjItqG+7N
kJbts2v550RmZxW8xbwyJ6gAt0Z52WGgU1cCMC2vzHu4Xfthl4WwgM51t/Ghj7gP
QpSKm0HO+CfngrnYjJgLmzIFFRLZ44pR3YgQWA==
-----END CERTIFICATE-----
EOT
}

resource "kong-gateway_certificate" "my_cert" {
  cert = <<EOT
-----BEGIN CERTIFICATE-----
MIIByDCCAU6gAwIBAgIUTiH53ur+PTurNUfIav6840EPkwQwCgYIKoZIzj0EAwIw
GzEZMBcGA1UEAwwQZGVtby5leGFtcGxlLmNvbTAeFw0yNTEwMDMxNjE0MDdaFw0y
ODEwMDIxNjE0MDdaMBsxGTAXBgNVBAMMEGRlbW8uZXhhbXBsZS5jb20wdjAQBgcq
hkjOPQIBBgUrgQQAIgNiAASpCbbkPRiGKR/nutM0gG8uIK1xKaEQMpEKqmbgn9mM
gqg2q9yLnQfnIpG7G/VE67WGS0PoonSN8v2nMj9XLYipK7oNYoOMlAigW2p4LQLO
iDkx7/hRXKDCcCHEPFOvVzWjUzBRMB0GA1UdDgQWBBQ/2XvmGDSh+Db5v+rm5VEl
2hoUODAfBgNVHSMEGDAWgBQ/2XvmGDSh+Db5v+rm5VEl2hoUODAPBgNVHRMBAf8E
BTADAQH/MAoGCCqGSM49BAMCA2gAMGUCMQCLHZGPqT0SVcSVYtHp3CefjItqG+7N
kJbts2v550RmZxW8xbwyJ6gAt0Z52WGgU1cCMC2vzHu4Xfthl4WwgM51t/Ghj7gP
QpSKm0HO+CfngrnYjJgLmzIFFRLZ44pR3YgQWA==
-----END CERTIFICATE-----
EOT

  key = <<EOT
-----BEGIN PRIVATE KEY-----
MIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDBOQyJxGencXyR0GwH9
JYvZVIrEGCgN+YmEGOaulEH3ihFc31uNxDqUkzM+uXrsPoihZANiAASpCbbkPRiG
KR/nutM0gG8uIK1xKaEQMpEKqmbgn9mMgqg2q9yLnQfnIpG7G/VE67WGS0PoonSN
8v2nMj9XLYipK7oNYoOMlAigW2p4LQLOiDkx7/hRXKDCcCHEPFOvVzU=
-----END PRIVATE KEY-----
EOT
}

resource "kong-gateway_hmac_auth" "my_hmac_auth" {
  username    = "alice"
  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_jwt" "my_jwt" {
  algorithm   = "HS512"
  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_key" "my_key" {
  name = "example-key"
  kid  = "example-key-id"
  jwk = jsonencode({
    kty = "RSA"
    use = "enc"
    kid = "example-key-id"
    d   = "jpq-eutUC5uNpVfYdsEZacbC9w0C3tPwl6jCLa2WB2yj1WcQRLRR5TJwCoPHUXucsKhtG8oHcvknmXgo1TzMWxUiSP2fhnqr9GEA4SSCvMqMvSazbgTLKtq1ZLyCBbyjlEBg2Leo9H4rsnh8p09WRQAbkq9S3Aq5kmUTLScWMCZLD5WZ95TxBJa7jRq8Ij1J69WI0v0Omb_jNbXfCYMZHaGxQrIYifwYYMtcrn70VxF2n0jh6TR5MnYggZdr84JdjQ564C-9ENYmAwyfKcWJ1yqMkLpRmy4dXV-MpBKuCarG1JdCu-r15595YtzObNd84-4B9JvaoJdy3hUXBsYTsQ"
    n   = "9mXXIzrcNUohBgRU7GWsFd5rrToLEVZtY7kQY-M_ASpXBoMpxsUmfp5fk39YHRGThwiVYFw-c3h97qOlHDWggq0PhjA_TxNp8ZcLNGybyDSnmJIBFbGU2JxCyX4AJm9RY4ZHCWlyzmMNu3uL22s6ydirSdfWt5dKBSW2STRUVXTslGKH3VE3zpMR4J2T81jhQsuwhrdXg3My6G90FJ5ltSaksVgiErIjqFiu1y5cEG5Gvhz99QoomHY0enKaX7mrT9XfQVtUWkbsf8Pwi3W-zsZsHQsjZZ8u9F0AdNRkCIheH3NCw46H1ouzARgxT3mTxEn8dcFzbRFOlOtoTw98nw"
    e   = "AQAB"
    p   = "_qAspCgjxg-3eICW8V6wUgN61KZVGRKHCHCK9JqDmOj9d09y14zuQ9j10WjvxK4NPzdQUygJlzDRhB9Dbk4AHVj9eIIoCH1kpYk0SYWOhgqgdtQzbzJjM3X_03geghifuQc9VnjZxGymXAukS7EIGWTNQzurcrM70_TKDBGoErs"
    q   = "97pMGP30n_e-DMMFuOoKaNKn_NKrh8KBJJXu3Bux5gV4YXh_IDDGI7qSu_7lvP24vrzBq9pj4zoMviBpfSz9YKrA9-Rs2_S6RngWxYAIie2gbqsbJ2ZvqiehLmuo3oUJBknUqDJ6b_K8Hy42e70ZXH27PHEwNWvWuXpZIPtt2W0"
    dp  = "9GRa1KjuRTFaoR-TQVLoG5_ZanfH4AvHbdNPnB0eSEsA1V59VOSg4KBCuN9mmzmP32hBAb_BDMu_nXfAagQV2hVLHDqZICTy0GvTsums9X0HrWZZg9YyHveYN6noZmgqDhcjyXavVfgO6PQHmtrtciotVeXU1n-v4e3nbBQaZPc"
    dq  = "2tCTpv-qtCIAnQUmaM9RooVwHMF5AdGsgMRu170exi7OxknJAIYUfjquoZ_lDaqPJOtVppag5HTCDK5Uf1zd8iThjhUWkrL4VoZ8lrcg07QxoY9BzOuOdp3KoVY3M1YPQp60WF0-COQ_hssrFOFTJX9pg1n3WziF0g9f6uIrhYE"
    qi  = "AgC9YSgEgkwr0ucJiO-wk7PSXcmsbxh-tR_Zs3oVH32yKa0GbRw68ocneLrL9_3lfK3EfvEbbhNuWFdpi6szExx9tWz9y4xKdax5AmkceY-yYaCFV353NgEipmeJL7-olT-YS93zjomhVzATZNl400uGJ3TC90rt6Be5rDGWhNg"
  })

  set = {
    "id" = kong-gateway_key_set.my_key_set.id
  }
}

resource "kong-gateway_key_set" "my_key_set" {
  name = "example-key-set"
}

resource "kong-gateway_mtls_auth" "my_mtls_auth" {
  ca_certificate = {
    id = kong-gateway_ca_certificate.my_ca_cert.id
  }
  subject_name = "CA_Subject_Name"

  consumer_id = kong-gateway_consumer.alice.id
}

resource "kong-gateway_partial" "my_partial" {
  redis_ce = {
    config = {
      host = "host.docker.internal"
      port = 6379
    }
  }
}

resource "kong-gateway_sni" "my_sni" {
  name = "example-sni"
  certificate = {
    id = kong-gateway_certificate.my_cert.id
  }
}

resource "kong-gateway_upstream" "my_upstream" {
  name      = "my-upstream"
  algorithm = "round-robin"
}

resource "kong-gateway_target" "my_target" {
  upstream_id = kong-gateway_upstream.my_upstream.id
  target      = "example.com:80"
  weight      = 100
}

resource "kong-gateway_vault" "my_vault" {
  name        = "env"
  description = "ENV vault for secrets"
  prefix      = "my-env-vault"
  config = jsonencode({
    prefix        = "MY_SECRET_"
    base64_decode = false
  })
}




#
#resource "kong-gateway_plugin_rate_limiting" "my_rate_limiting_plugin" {
#  enabled = true
#  config = {
#    minute = 5
#    policy = "local"
#  }
#
#  protocols = ["http", "https"]
#  route = {
#    id = kong-gateway_route.hello.id
#  }
#}
#
#
