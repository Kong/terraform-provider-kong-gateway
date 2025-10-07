package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCustomPlugin(t *testing.T) {
	t.Run("can be configured", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("kong-gateway_custom_plugin.rl", "name", "rate-limiting"),
						resource.TestCheckResourceAttr("kong-gateway_custom_plugin.rl", "config.minute", "5"),
					),
				},
			},
		})
	})
}
