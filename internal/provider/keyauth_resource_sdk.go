// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *KeyAuthResourceModel) ToSharedKeyAuthInput() *shared.KeyAuthInput {
	var consumer *shared.KeyAuthConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.KeyAuthConsumer{
			ID: id,
		}
	}
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	var key string
	key = r.Key.ValueString()

	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.KeyAuthInput{
		Consumer: consumer,
		ID:       id1,
		Key:      key,
		Tags:     tags,
	}
	return &out
}

func (r *KeyAuthResourceModel) RefreshFromSharedKeyAuth(resp *shared.KeyAuth) {
	if resp != nil {
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringValue(resp.Key)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
