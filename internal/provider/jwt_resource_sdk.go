// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *JwtResourceModel) ToSharedJWTInput() *shared.JWTInput {
	algorithm := new(shared.Algorithm)
	if !r.Algorithm.IsUnknown() && !r.Algorithm.IsNull() {
		*algorithm = shared.Algorithm(r.Algorithm.ValueString())
	} else {
		algorithm = nil
	}
	var consumer *shared.JWTConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.JWTConsumer{
			ID: id,
		}
	}
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	key := new(string)
	if !r.Key.IsUnknown() && !r.Key.IsNull() {
		*key = r.Key.ValueString()
	} else {
		key = nil
	}
	rsaPublicKey := new(string)
	if !r.RsaPublicKey.IsUnknown() && !r.RsaPublicKey.IsNull() {
		*rsaPublicKey = r.RsaPublicKey.ValueString()
	} else {
		rsaPublicKey = nil
	}
	secret := new(string)
	if !r.Secret.IsUnknown() && !r.Secret.IsNull() {
		*secret = r.Secret.ValueString()
	} else {
		secret = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.JWTInput{
		Algorithm:    algorithm,
		Consumer:     consumer,
		ID:           id1,
		Key:          key,
		RsaPublicKey: rsaPublicKey,
		Secret:       secret,
		Tags:         tags,
	}
	return &out
}

func (r *JwtResourceModel) RefreshFromSharedJwt(resp *shared.Jwt) {
	if resp != nil {
		if resp.Algorithm != nil {
			r.Algorithm = types.StringValue(string(*resp.Algorithm))
		} else {
			r.Algorithm = types.StringNull()
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.RsaPublicKey = types.StringPointerValue(resp.RsaPublicKey)
		r.Secret = types.StringPointerValue(resp.Secret)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
