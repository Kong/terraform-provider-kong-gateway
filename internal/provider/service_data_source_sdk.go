// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *ServiceDataSourceModel) RefreshFromSharedService(resp *shared.Service) {
	if resp != nil {
		r.CaCertificates = []types.String{}
		for _, v := range resp.CaCertificates {
			r.CaCertificates = append(r.CaCertificates, types.StringValue(v))
		}
		if resp.ClientCertificate == nil {
			r.ClientCertificate = nil
		} else {
			r.ClientCertificate = &tfTypes.ACLConsumer{}
			r.ClientCertificate.ID = types.StringPointerValue(resp.ClientCertificate.ID)
		}
		r.ConnectTimeout = types.Int64PointerValue(resp.ConnectTimeout)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.Host = types.StringValue(resp.Host)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.Path = types.StringPointerValue(resp.Path)
		r.Port = types.Int64Value(resp.Port)
		r.Protocol = types.StringValue(string(resp.Protocol))
		r.ReadTimeout = types.Int64PointerValue(resp.ReadTimeout)
		r.Retries = types.Int64PointerValue(resp.Retries)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.TLSVerify = types.BoolPointerValue(resp.TLSVerify)
		r.TLSVerifyDepth = types.Int64PointerValue(resp.TLSVerifyDepth)
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.WriteTimeout = types.Int64PointerValue(resp.WriteTimeout)
	}
}
