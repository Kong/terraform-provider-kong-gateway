resource "kong-gateway_route" "hello" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  service = {
    id = kong-gateway_service.httpbin.id
  }
}
