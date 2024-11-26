// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginRateLimitingDataSourceModel) RefreshFromSharedRateLimitingPlugin(resp *shared.RateLimitingPlugin) {
	if resp != nil {
		if resp.Config.Day != nil {
			r.Config.Day = types.NumberValue(big.NewFloat(float64(*resp.Config.Day)))
		} else {
			r.Config.Day = types.NumberNull()
		}
		if resp.Config.ErrorCode != nil {
			r.Config.ErrorCode = types.NumberValue(big.NewFloat(float64(*resp.Config.ErrorCode)))
		} else {
			r.Config.ErrorCode = types.NumberNull()
		}
		r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
		r.Config.FaultTolerant = types.BoolPointerValue(resp.Config.FaultTolerant)
		r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
		r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
		if resp.Config.Hour != nil {
			r.Config.Hour = types.NumberValue(big.NewFloat(float64(*resp.Config.Hour)))
		} else {
			r.Config.Hour = types.NumberNull()
		}
		if resp.Config.LimitBy != nil {
			r.Config.LimitBy = types.StringValue(string(*resp.Config.LimitBy))
		} else {
			r.Config.LimitBy = types.StringNull()
		}
		if resp.Config.Minute != nil {
			r.Config.Minute = types.NumberValue(big.NewFloat(float64(*resp.Config.Minute)))
		} else {
			r.Config.Minute = types.NumberNull()
		}
		if resp.Config.Month != nil {
			r.Config.Month = types.NumberValue(big.NewFloat(float64(*resp.Config.Month)))
		} else {
			r.Config.Month = types.NumberNull()
		}
		r.Config.Path = types.StringPointerValue(resp.Config.Path)
		if resp.Config.Policy != nil {
			r.Config.Policy = types.StringValue(string(*resp.Config.Policy))
		} else {
			r.Config.Policy = types.StringNull()
		}
		if resp.Config.Redis == nil {
			r.Config.Redis = nil
		} else {
			r.Config.Redis = &tfTypes.RateLimitingPluginRedis{}
			r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
			r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
			r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
			r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
			r.Config.Redis.ServerName = types.StringPointerValue(resp.Config.Redis.ServerName)
			r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
			r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
			r.Config.Redis.Timeout = types.Int64PointerValue(resp.Config.Redis.Timeout)
			r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
		}
		if resp.Config.Second != nil {
			r.Config.Second = types.NumberValue(big.NewFloat(float64(*resp.Config.Second)))
		} else {
			r.Config.Second = types.NumberNull()
		}
		if resp.Config.SyncRate != nil {
			r.Config.SyncRate = types.NumberValue(big.NewFloat(float64(*resp.Config.SyncRate)))
		} else {
			r.Config.SyncRate = types.NumberNull()
		}
		if resp.Config.Year != nil {
			r.Config.Year = types.NumberValue(big.NewFloat(float64(*resp.Config.Year)))
		} else {
			r.Config.Year = types.NumberNull()
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
