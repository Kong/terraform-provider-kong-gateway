// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *MTLSAuthDataSourceModel) RefreshFromSharedMTLSAuth(resp *shared.MTLSAuth) {
	if resp != nil {
		if resp.CaCertificate == nil {
			r.CaCertificate = nil
		} else {
			r.CaCertificate = &tfTypes.ACLConsumer{}
			r.CaCertificate.ID = types.StringPointerValue(resp.CaCertificate.ID)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.SubjectName = types.StringValue(resp.SubjectName)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}