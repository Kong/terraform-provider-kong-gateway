// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Vectordb struct {
	Dimensions     types.Int64                `tfsdk:"dimensions"`
	DistanceMetric types.String               `tfsdk:"distance_metric"`
	Pgvector       Pgvector                   `tfsdk:"pgvector"`
	Redis          AiProxyAdvancedPluginRedis `tfsdk:"redis"`
	Strategy       types.String               `tfsdk:"strategy"`
	Threshold      types.Float64              `tfsdk:"threshold"`
}
