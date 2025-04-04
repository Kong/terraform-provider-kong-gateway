// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *JwtDataSourceModel) RefreshFromSharedJwt(resp *shared.Jwt) {
	if resp != nil {
		if resp.Algorithm != nil {
			r.Algorithm = types.StringValue(string(*resp.Algorithm))
		} else {
			r.Algorithm = types.StringNull()
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.RsaPublicKey = types.StringPointerValue(resp.RsaPublicKey)
		r.Secret = types.StringPointerValue(resp.Secret)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
