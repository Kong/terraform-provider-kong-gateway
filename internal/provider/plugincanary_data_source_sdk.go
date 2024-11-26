// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginCanaryDataSourceModel) RefreshFromSharedCanaryPlugin(resp *shared.CanaryPlugin) {
	if resp != nil {
		r.Config.CanaryByHeaderName = types.StringPointerValue(resp.Config.CanaryByHeaderName)
		if resp.Config.Duration != nil {
			r.Config.Duration = types.NumberValue(big.NewFloat(float64(*resp.Config.Duration)))
		} else {
			r.Config.Duration = types.NumberNull()
		}
		r.Config.Groups = []types.String{}
		for _, v := range resp.Config.Groups {
			r.Config.Groups = append(r.Config.Groups, types.StringValue(v))
		}
		if resp.Config.Hash != nil {
			r.Config.Hash = types.StringValue(string(*resp.Config.Hash))
		} else {
			r.Config.Hash = types.StringNull()
		}
		r.Config.HashHeader = types.StringPointerValue(resp.Config.HashHeader)
		if resp.Config.Percentage != nil {
			r.Config.Percentage = types.NumberValue(big.NewFloat(float64(*resp.Config.Percentage)))
		} else {
			r.Config.Percentage = types.NumberNull()
		}
		if resp.Config.Start != nil {
			r.Config.Start = types.NumberValue(big.NewFloat(float64(*resp.Config.Start)))
		} else {
			r.Config.Start = types.NumberNull()
		}
		if resp.Config.Steps != nil {
			r.Config.Steps = types.NumberValue(big.NewFloat(float64(*resp.Config.Steps)))
		} else {
			r.Config.Steps = types.NumberNull()
		}
		r.Config.UpstreamFallback = types.BoolPointerValue(resp.Config.UpstreamFallback)
		r.Config.UpstreamHost = types.StringPointerValue(resp.Config.UpstreamHost)
		r.Config.UpstreamPort = types.Int64PointerValue(resp.Config.UpstreamPort)
		r.Config.UpstreamURI = types.StringPointerValue(resp.Config.UpstreamURI)
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
