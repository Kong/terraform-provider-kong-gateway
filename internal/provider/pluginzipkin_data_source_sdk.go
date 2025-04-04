// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginZipkinDataSourceModel) RefreshFromSharedZipkinPlugin(resp *shared.ZipkinPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.ZipkinPluginConfig{}
			r.Config.ConnectTimeout = types.Int64PointerValue(resp.Config.ConnectTimeout)
			if resp.Config.DefaultHeaderType != nil {
				r.Config.DefaultHeaderType = types.StringValue(string(*resp.Config.DefaultHeaderType))
			} else {
				r.Config.DefaultHeaderType = types.StringNull()
			}
			r.Config.DefaultServiceName = types.StringPointerValue(resp.Config.DefaultServiceName)
			if resp.Config.HeaderType != nil {
				r.Config.HeaderType = types.StringValue(string(*resp.Config.HeaderType))
			} else {
				r.Config.HeaderType = types.StringNull()
			}
			r.Config.HTTPEndpoint = types.StringPointerValue(resp.Config.HTTPEndpoint)
			r.Config.HTTPResponseHeaderForTraceid = types.StringPointerValue(resp.Config.HTTPResponseHeaderForTraceid)
			if resp.Config.HTTPSpanName != nil {
				r.Config.HTTPSpanName = types.StringValue(string(*resp.Config.HTTPSpanName))
			} else {
				r.Config.HTTPSpanName = types.StringNull()
			}
			r.Config.IncludeCredential = types.BoolPointerValue(resp.Config.IncludeCredential)
			r.Config.LocalServiceName = types.StringPointerValue(resp.Config.LocalServiceName)
			if resp.Config.PhaseDurationFlavor != nil {
				r.Config.PhaseDurationFlavor = types.StringValue(string(*resp.Config.PhaseDurationFlavor))
			} else {
				r.Config.PhaseDurationFlavor = types.StringNull()
			}
			if resp.Config.Propagation == nil {
				r.Config.Propagation = nil
			} else {
				r.Config.Propagation = &tfTypes.Propagation{}
				r.Config.Propagation.Clear = make([]types.String, 0, len(resp.Config.Propagation.Clear))
				for _, v := range resp.Config.Propagation.Clear {
					r.Config.Propagation.Clear = append(r.Config.Propagation.Clear, types.StringValue(v))
				}
				r.Config.Propagation.DefaultFormat = types.StringValue(string(resp.Config.Propagation.DefaultFormat))
				r.Config.Propagation.Extract = make([]types.String, 0, len(resp.Config.Propagation.Extract))
				for _, v := range resp.Config.Propagation.Extract {
					r.Config.Propagation.Extract = append(r.Config.Propagation.Extract, types.StringValue(string(v)))
				}
				r.Config.Propagation.Inject = make([]types.String, 0, len(resp.Config.Propagation.Inject))
				for _, v := range resp.Config.Propagation.Inject {
					r.Config.Propagation.Inject = append(r.Config.Propagation.Inject, types.StringValue(string(v)))
				}
			}
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
			r.Config.ReadTimeout = types.Int64PointerValue(resp.Config.ReadTimeout)
			if resp.Config.SampleRatio != nil {
				r.Config.SampleRatio = types.NumberValue(big.NewFloat(float64(*resp.Config.SampleRatio)))
			} else {
				r.Config.SampleRatio = types.NumberNull()
			}
			r.Config.SendTimeout = types.Int64PointerValue(resp.Config.SendTimeout)
			r.Config.StaticTags = []tfTypes.StaticTags{}
			if len(r.Config.StaticTags) > len(resp.Config.StaticTags) {
				r.Config.StaticTags = r.Config.StaticTags[:len(resp.Config.StaticTags)]
			}
			for staticTagsCount, staticTagsItem := range resp.Config.StaticTags {
				var staticTags1 tfTypes.StaticTags
				staticTags1.Name = types.StringValue(staticTagsItem.Name)
				staticTags1.Value = types.StringValue(staticTagsItem.Value)
				if staticTagsCount+1 > len(r.Config.StaticTags) {
					r.Config.StaticTags = append(r.Config.StaticTags, staticTags1)
				} else {
					r.Config.StaticTags[staticTagsCount].Name = staticTags1.Name
					r.Config.StaticTags[staticTagsCount].Value = staticTags1.Value
				}
			}
			r.Config.TagsHeader = types.StringPointerValue(resp.Config.TagsHeader)
			if resp.Config.TraceidByteCount != nil {
				r.Config.TraceidByteCount = types.Int64Value(int64(*resp.Config.TraceidByteCount))
			} else {
				r.Config.TraceidByteCount = types.Int64Null()
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.Ordering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.After{}
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.After{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials1 tfTypes.Partials
				partials1.ID = types.StringPointerValue(partialsItem.ID)
				partials1.Name = types.StringPointerValue(partialsItem.Name)
				partials1.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials1)
				} else {
					r.Partials[partialsCount].ID = partials1.ID
					r.Partials[partialsCount].Name = partials1.Name
					r.Partials[partialsCount].Path = partials1.Path
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
