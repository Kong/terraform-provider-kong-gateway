package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRoute(t *testing.T) {
	t.Run("match headers", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("kong-gateway_route.my_route", "name", "route-headers"),
						resource.TestCheckResourceAttr("kong-gateway_route.my_route", "headers.X-Header-Goes-Here.0", "Value-1"),
						resource.TestCheckResourceAttr("kong-gateway_route.my_route", "headers.X-Header-Goes-Here.1", "Value-2"),
					),
				},
			},
		})
	})
}
