// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *SniResourceModel) ToSharedSni() *shared.Sni {
	id := new(string)
	if !r.Certificate.ID.IsUnknown() && !r.Certificate.ID.IsNull() {
		*id = r.Certificate.ID.ValueString()
	} else {
		id = nil
	}
	certificate := shared.SNICertificate{
		ID: id,
	}
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	var name string
	name = r.Name.ValueString()

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
	out := shared.Sni{
		Certificate: certificate,
		CreatedAt:   createdAt,
		ID:          id1,
		Name:        name,
		Tags:        tags,
		UpdatedAt:   updatedAt,
	}
	return &out
}

func (r *SniResourceModel) RefreshFromSharedSni(resp *shared.Sni) {
	if resp != nil {
		r.Certificate.ID = types.StringPointerValue(resp.Certificate.ID)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
