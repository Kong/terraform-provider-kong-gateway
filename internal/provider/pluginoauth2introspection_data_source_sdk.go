// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginOauth2IntrospectionDataSourceModel) RefreshFromSharedOauth2IntrospectionPlugin(resp *shared.Oauth2IntrospectionPlugin) {
	if resp != nil {
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AuthorizationValue = types.StringPointerValue(resp.Config.AuthorizationValue)
		if resp.Config.ConsumerBy != nil {
			r.Config.ConsumerBy = types.StringValue(string(*resp.Config.ConsumerBy))
		} else {
			r.Config.ConsumerBy = types.StringNull()
		}
		r.Config.CustomClaimsForward = []types.String{}
		for _, v := range resp.Config.CustomClaimsForward {
			r.Config.CustomClaimsForward = append(r.Config.CustomClaimsForward, types.StringValue(v))
		}
		if len(resp.Config.CustomIntrospectionHeaders) > 0 {
			r.Config.CustomIntrospectionHeaders = make(map[string]types.String)
			for key, value := range resp.Config.CustomIntrospectionHeaders {
				result, _ := json.Marshal(value)
				r.Config.CustomIntrospectionHeaders[key] = types.StringValue(string(result))
			}
		}
		r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
		r.Config.IntrospectRequest = types.BoolPointerValue(resp.Config.IntrospectRequest)
		r.Config.IntrospectionURL = types.StringPointerValue(resp.Config.IntrospectionURL)
		r.Config.Keepalive = types.Int64PointerValue(resp.Config.Keepalive)
		r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
		r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
		r.Config.TokenTypeHint = types.StringPointerValue(resp.Config.TokenTypeHint)
		if resp.Config.TTL != nil {
			r.Config.TTL = types.NumberValue(big.NewFloat(float64(*resp.Config.TTL)))
		} else {
			r.Config.TTL = types.NumberNull()
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
