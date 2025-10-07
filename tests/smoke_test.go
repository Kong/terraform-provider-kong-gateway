package tests

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type TestCase struct {
	name    string
	runWhen *string
}

func TestSmoke(t *testing.T) {
	t.Parallel()

	var testCases = []TestCase{
		//{
		//	name:    "3.10",
		//	runWhen: ptr("3.10"),
		//},
		//{
		//	name:    "3.11",
		//	runWhen: ptr("3.11"),
		//},
		{
			name:    "3.12",
			runWhen: ptr("3.12"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			currentVersion := os.Getenv("KONG_VERSION")
			if tc.runWhen != nil && *tc.runWhen != currentVersion {
				t.Skip("Skipping " + tc.name + " (" + *tc.runWhen + ") because KONG_VERSION is " + currentVersion)
			}

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
}

func ptr(s string) *string {
	return &s
}
