// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *ConsumerResourceModel) ToSharedConsumer() *shared.Consumer {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	customID := new(string)
	if !r.CustomID.IsUnknown() && !r.CustomID.IsNull() {
		*customID = r.CustomID.ValueString()
	} else {
		customID = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	username := new(string)
	if !r.Username.IsUnknown() && !r.Username.IsNull() {
		*username = r.Username.ValueString()
	} else {
		username = nil
	}
	out := shared.Consumer{
		CreatedAt: createdAt,
		CustomID:  customID,
		ID:        id,
		Tags:      tags,
		UpdatedAt: updatedAt,
		Username:  username,
	}
	return &out
}

func (r *ConsumerResourceModel) RefreshFromSharedConsumer(resp *shared.Consumer) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.CustomID = types.StringPointerValue(resp.CustomID)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.Username = types.StringPointerValue(resp.Username)
	}
}
