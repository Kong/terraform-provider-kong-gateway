// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Keys struct {
	Alg           types.String   `tfsdk:"alg"`
	Crv           types.String   `tfsdk:"crv"`
	D             types.String   `tfsdk:"d"`
	Dp            types.String   `tfsdk:"dp"`
	Dq            types.String   `tfsdk:"dq"`
	E             types.String   `tfsdk:"e"`
	Issuer        types.String   `tfsdk:"issuer"`
	K             types.String   `tfsdk:"k"`
	KeyOps        []types.String `tfsdk:"key_ops"`
	Kid           types.String   `tfsdk:"kid"`
	Kty           types.String   `tfsdk:"kty"`
	N             types.String   `tfsdk:"n"`
	Oth           types.String   `tfsdk:"oth"`
	P             types.String   `tfsdk:"p"`
	Q             types.String   `tfsdk:"q"`
	Qi            types.String   `tfsdk:"qi"`
	R             types.String   `tfsdk:"r"`
	T             types.String   `tfsdk:"t"`
	Use           types.String   `tfsdk:"use"`
	X             types.String   `tfsdk:"x"`
	X5c           []types.String `tfsdk:"x5c"`
	X5t           types.String   `tfsdk:"x5t"`
	X5tNumberS256 types.String   `tfsdk:"x5t_number_s256"`
	X5u           types.String   `tfsdk:"x5u"`
	Y             types.String   `tfsdk:"y"`
}
