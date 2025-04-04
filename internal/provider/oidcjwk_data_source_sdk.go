// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *OidcJwkDataSourceModel) RefreshFromSharedOidcJwk(resp *shared.OidcJwk) {
	if resp != nil {
		r.ID = types.StringPointerValue(resp.ID)
		if resp.Jwks == nil {
			r.Jwks = nil
		} else {
			r.Jwks = &tfTypes.Jwks{}
			r.Jwks.Keys = []tfTypes.Keys{}
			if len(r.Jwks.Keys) > len(resp.Jwks.Keys) {
				r.Jwks.Keys = r.Jwks.Keys[:len(resp.Jwks.Keys)]
			}
			for keysCount, keysItem := range resp.Jwks.Keys {
				var keys1 tfTypes.Keys
				keys1.Alg = types.StringPointerValue(keysItem.Alg)
				keys1.Crv = types.StringPointerValue(keysItem.Crv)
				keys1.D = types.StringPointerValue(keysItem.D)
				keys1.Dp = types.StringPointerValue(keysItem.Dp)
				keys1.Dq = types.StringPointerValue(keysItem.Dq)
				keys1.E = types.StringPointerValue(keysItem.E)
				keys1.Issuer = types.StringPointerValue(keysItem.Issuer)
				keys1.K = types.StringPointerValue(keysItem.K)
				keys1.KeyOps = make([]types.String, 0, len(keysItem.KeyOps))
				for _, v := range keysItem.KeyOps {
					keys1.KeyOps = append(keys1.KeyOps, types.StringValue(v))
				}
				keys1.Kid = types.StringPointerValue(keysItem.Kid)
				keys1.Kty = types.StringPointerValue(keysItem.Kty)
				keys1.N = types.StringPointerValue(keysItem.N)
				keys1.Oth = types.StringPointerValue(keysItem.Oth)
				keys1.P = types.StringPointerValue(keysItem.P)
				keys1.Q = types.StringPointerValue(keysItem.Q)
				keys1.Qi = types.StringPointerValue(keysItem.Qi)
				keys1.R = types.StringPointerValue(keysItem.R)
				keys1.T = types.StringPointerValue(keysItem.T)
				keys1.Use = types.StringPointerValue(keysItem.Use)
				keys1.X = types.StringPointerValue(keysItem.X)
				keys1.X5c = make([]types.String, 0, len(keysItem.X5c))
				for _, v := range keysItem.X5c {
					keys1.X5c = append(keys1.X5c, types.StringValue(v))
				}
				keys1.X5t = types.StringPointerValue(keysItem.X5t)
				keys1.X5tNumberS256 = types.StringPointerValue(keysItem.X5tNumberS256)
				keys1.X5u = types.StringPointerValue(keysItem.X5u)
				keys1.Y = types.StringPointerValue(keysItem.Y)
				if keysCount+1 > len(r.Jwks.Keys) {
					r.Jwks.Keys = append(r.Jwks.Keys, keys1)
				} else {
					r.Jwks.Keys[keysCount].Alg = keys1.Alg
					r.Jwks.Keys[keysCount].Crv = keys1.Crv
					r.Jwks.Keys[keysCount].D = keys1.D
					r.Jwks.Keys[keysCount].Dp = keys1.Dp
					r.Jwks.Keys[keysCount].Dq = keys1.Dq
					r.Jwks.Keys[keysCount].E = keys1.E
					r.Jwks.Keys[keysCount].Issuer = keys1.Issuer
					r.Jwks.Keys[keysCount].K = keys1.K
					r.Jwks.Keys[keysCount].KeyOps = keys1.KeyOps
					r.Jwks.Keys[keysCount].Kid = keys1.Kid
					r.Jwks.Keys[keysCount].Kty = keys1.Kty
					r.Jwks.Keys[keysCount].N = keys1.N
					r.Jwks.Keys[keysCount].Oth = keys1.Oth
					r.Jwks.Keys[keysCount].P = keys1.P
					r.Jwks.Keys[keysCount].Q = keys1.Q
					r.Jwks.Keys[keysCount].Qi = keys1.Qi
					r.Jwks.Keys[keysCount].R = keys1.R
					r.Jwks.Keys[keysCount].T = keys1.T
					r.Jwks.Keys[keysCount].Use = keys1.Use
					r.Jwks.Keys[keysCount].X = keys1.X
					r.Jwks.Keys[keysCount].X5c = keys1.X5c
					r.Jwks.Keys[keysCount].X5t = keys1.X5t
					r.Jwks.Keys[keysCount].X5tNumberS256 = keys1.X5tNumberS256
					r.Jwks.Keys[keysCount].X5u = keys1.X5u
					r.Jwks.Keys[keysCount].Y = keys1.Y
				}
			}
		}
	}
}
