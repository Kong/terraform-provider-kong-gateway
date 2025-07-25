// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Cache struct {
	CacheTTL types.Int64                 `tfsdk:"cache_ttl"`
	Memory   *Memory                     `tfsdk:"memory"`
	Redis    *AiProxyAdvancedPluginRedis `tfsdk:"redis"`
	Strategy types.String                `tfsdk:"strategy"`
}
