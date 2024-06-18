resource "kong-gateway_basic_auth" "my_basicauth" {
  username = "alice-test"
  password = "demo"

  consumer = {
    id = kong-gateway_consumer.alice.id
  }
}
