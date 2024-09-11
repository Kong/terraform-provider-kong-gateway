// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginDatadogDataSourceModel) RefreshFromSharedDatadogPlugin(resp *shared.DatadogPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateDatadogPluginConfig{}
			r.Config.ConsumerTag = types.StringPointerValue(resp.Config.ConsumerTag)
			if resp.Config.FlushTimeout != nil {
				r.Config.FlushTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.FlushTimeout)))
			} else {
				r.Config.FlushTimeout = types.NumberNull()
			}
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.Metrics = []tfTypes.Metrics{}
			if len(r.Config.Metrics) > len(resp.Config.Metrics) {
				r.Config.Metrics = r.Config.Metrics[:len(resp.Config.Metrics)]
			}
			for metricsCount, metricsItem := range resp.Config.Metrics {
				var metrics1 tfTypes.Metrics
				if metricsItem.ConsumerIdentifier != nil {
					metrics1.ConsumerIdentifier = types.StringValue(string(*metricsItem.ConsumerIdentifier))
				} else {
					metrics1.ConsumerIdentifier = types.StringNull()
				}
				metrics1.Name = types.StringValue(string(metricsItem.Name))
				if metricsItem.SampleRate != nil {
					metrics1.SampleRate = types.NumberValue(big.NewFloat(float64(*metricsItem.SampleRate)))
				} else {
					metrics1.SampleRate = types.NumberNull()
				}
				metrics1.StatType = types.StringValue(string(metricsItem.StatType))
				metrics1.Tags = []types.String{}
				for _, v := range metricsItem.Tags {
					metrics1.Tags = append(metrics1.Tags, types.StringValue(v))
				}
				if metricsCount+1 > len(r.Config.Metrics) {
					r.Config.Metrics = append(r.Config.Metrics, metrics1)
				} else {
					r.Config.Metrics[metricsCount].ConsumerIdentifier = metrics1.ConsumerIdentifier
					r.Config.Metrics[metricsCount].Name = metrics1.Name
					r.Config.Metrics[metricsCount].SampleRate = metrics1.SampleRate
					r.Config.Metrics[metricsCount].StatType = metrics1.StatType
					r.Config.Metrics[metricsCount].Tags = metrics1.Tags
				}
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			r.Config.Prefix = types.StringPointerValue(resp.Config.Prefix)
			if resp.Config.Queue == nil {
				r.Config.Queue = nil
			} else {
				r.Config.Queue = &tfTypes.Queue{}
				if resp.Config.Queue.ConcurrencyLimit != nil {
					r.Config.Queue.ConcurrencyLimit = types.Int64Value(int64(*resp.Config.Queue.ConcurrencyLimit))
				} else {
					r.Config.Queue.ConcurrencyLimit = types.Int64Null()
				}
				if resp.Config.Queue.InitialRetryDelay != nil {
					r.Config.Queue.InitialRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.InitialRetryDelay)))
				} else {
					r.Config.Queue.InitialRetryDelay = types.NumberNull()
				}
				r.Config.Queue.MaxBatchSize = types.Int64PointerValue(resp.Config.Queue.MaxBatchSize)
				r.Config.Queue.MaxBytes = types.Int64PointerValue(resp.Config.Queue.MaxBytes)
				if resp.Config.Queue.MaxCoalescingDelay != nil {
					r.Config.Queue.MaxCoalescingDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxCoalescingDelay)))
				} else {
					r.Config.Queue.MaxCoalescingDelay = types.NumberNull()
				}
				r.Config.Queue.MaxEntries = types.Int64PointerValue(resp.Config.Queue.MaxEntries)
				if resp.Config.Queue.MaxRetryDelay != nil {
					r.Config.Queue.MaxRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryDelay)))
				} else {
					r.Config.Queue.MaxRetryDelay = types.NumberNull()
				}
				if resp.Config.Queue.MaxRetryTime != nil {
					r.Config.Queue.MaxRetryTime = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryTime)))
				} else {
					r.Config.Queue.MaxRetryTime = types.NumberNull()
				}
			}
			r.Config.QueueSize = types.Int64PointerValue(resp.Config.QueueSize)
			r.Config.RetryCount = types.Int64PointerValue(resp.Config.RetryCount)
			r.Config.ServiceNameTag = types.StringPointerValue(resp.Config.ServiceNameTag)
			r.Config.StatusTag = types.StringPointerValue(resp.Config.StatusTag)
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
			r.Ordering = types.StringNull()
		} else {
			orderingResult, _ := json.Marshal(resp.Ordering)
			r.Ordering = types.StringValue(string(orderingResult))
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