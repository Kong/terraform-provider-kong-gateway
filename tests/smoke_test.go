package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSmoke(t *testing.T) {
	t.Parallel()
	t.Run("smoke test", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
				},
			},
		})
	})
}
