// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginHTTPLogResourceModel) ToSharedHTTPLogPlugin() *shared.HTTPLogPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.HTTPLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.HTTPLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.HTTPLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.HTTPLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.HTTPLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.HTTPLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.HTTPLogPluginPartials = []shared.HTTPLogPluginPartials{}
	for _, partialsItem := range r.Partials {
		id1 := new(string)
		if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
			*id1 = partialsItem.ID.ValueString()
		} else {
			id1 = nil
		}
		name := new(string)
		if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
			*name = partialsItem.Name.ValueString()
		} else {
			name = nil
		}
		path := new(string)
		if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
			*path = partialsItem.Path.ValueString()
		} else {
			path = nil
		}
		partials = append(partials, shared.HTTPLogPluginPartials{
			ID:   id1,
			Name: name,
			Path: path,
		})
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
	var config *shared.HTTPLogPluginConfig
	if r.Config != nil {
		contentType := new(shared.ContentType)
		if !r.Config.ContentType.IsUnknown() && !r.Config.ContentType.IsNull() {
			*contentType = shared.ContentType(r.Config.ContentType.ValueString())
		} else {
			contentType = nil
		}
		customFieldsByLua := make(map[string]interface{})
		for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
			var customFieldsByLuaInst interface{}
			_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
			customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
		}
		flushTimeout := new(float64)
		if !r.Config.FlushTimeout.IsUnknown() && !r.Config.FlushTimeout.IsNull() {
			*flushTimeout, _ = r.Config.FlushTimeout.ValueBigFloat().Float64()
		} else {
			flushTimeout = nil
		}
		headers := make(map[string]interface{})
		for headersKey, headersValue := range r.Config.Headers {
			var headersInst interface{}
			_ = json.Unmarshal([]byte(headersValue.ValueString()), &headersInst)
			headers[headersKey] = headersInst
		}
		httpEndpoint := new(string)
		if !r.Config.HTTPEndpoint.IsUnknown() && !r.Config.HTTPEndpoint.IsNull() {
			*httpEndpoint = r.Config.HTTPEndpoint.ValueString()
		} else {
			httpEndpoint = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
		} else {
			keepalive = nil
		}
		method := new(shared.Method)
		if !r.Config.Method.IsUnknown() && !r.Config.Method.IsNull() {
			*method = shared.Method(r.Config.Method.ValueString())
		} else {
			method = nil
		}
		var queue *shared.HTTPLogPluginQueue
		if r.Config.Queue != nil {
			concurrencyLimit := new(shared.HTTPLogPluginConcurrencyLimit)
			if !r.Config.Queue.ConcurrencyLimit.IsUnknown() && !r.Config.Queue.ConcurrencyLimit.IsNull() {
				*concurrencyLimit = shared.HTTPLogPluginConcurrencyLimit(r.Config.Queue.ConcurrencyLimit.ValueInt64())
			} else {
				concurrencyLimit = nil
			}
			initialRetryDelay := new(float64)
			if !r.Config.Queue.InitialRetryDelay.IsUnknown() && !r.Config.Queue.InitialRetryDelay.IsNull() {
				*initialRetryDelay, _ = r.Config.Queue.InitialRetryDelay.ValueBigFloat().Float64()
			} else {
				initialRetryDelay = nil
			}
			maxBatchSize := new(int64)
			if !r.Config.Queue.MaxBatchSize.IsUnknown() && !r.Config.Queue.MaxBatchSize.IsNull() {
				*maxBatchSize = r.Config.Queue.MaxBatchSize.ValueInt64()
			} else {
				maxBatchSize = nil
			}
			maxBytes := new(int64)
			if !r.Config.Queue.MaxBytes.IsUnknown() && !r.Config.Queue.MaxBytes.IsNull() {
				*maxBytes = r.Config.Queue.MaxBytes.ValueInt64()
			} else {
				maxBytes = nil
			}
			maxCoalescingDelay := new(float64)
			if !r.Config.Queue.MaxCoalescingDelay.IsUnknown() && !r.Config.Queue.MaxCoalescingDelay.IsNull() {
				*maxCoalescingDelay, _ = r.Config.Queue.MaxCoalescingDelay.ValueBigFloat().Float64()
			} else {
				maxCoalescingDelay = nil
			}
			maxEntries := new(int64)
			if !r.Config.Queue.MaxEntries.IsUnknown() && !r.Config.Queue.MaxEntries.IsNull() {
				*maxEntries = r.Config.Queue.MaxEntries.ValueInt64()
			} else {
				maxEntries = nil
			}
			maxRetryDelay := new(float64)
			if !r.Config.Queue.MaxRetryDelay.IsUnknown() && !r.Config.Queue.MaxRetryDelay.IsNull() {
				*maxRetryDelay, _ = r.Config.Queue.MaxRetryDelay.ValueBigFloat().Float64()
			} else {
				maxRetryDelay = nil
			}
			maxRetryTime := new(float64)
			if !r.Config.Queue.MaxRetryTime.IsUnknown() && !r.Config.Queue.MaxRetryTime.IsNull() {
				*maxRetryTime, _ = r.Config.Queue.MaxRetryTime.ValueBigFloat().Float64()
			} else {
				maxRetryTime = nil
			}
			queue = &shared.HTTPLogPluginQueue{
				ConcurrencyLimit:   concurrencyLimit,
				InitialRetryDelay:  initialRetryDelay,
				MaxBatchSize:       maxBatchSize,
				MaxBytes:           maxBytes,
				MaxCoalescingDelay: maxCoalescingDelay,
				MaxEntries:         maxEntries,
				MaxRetryDelay:      maxRetryDelay,
				MaxRetryTime:       maxRetryTime,
			}
		}
		queueSize := new(int64)
		if !r.Config.QueueSize.IsUnknown() && !r.Config.QueueSize.IsNull() {
			*queueSize = r.Config.QueueSize.ValueInt64()
		} else {
			queueSize = nil
		}
		retryCount := new(int64)
		if !r.Config.RetryCount.IsUnknown() && !r.Config.RetryCount.IsNull() {
			*retryCount = r.Config.RetryCount.ValueInt64()
		} else {
			retryCount = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		config = &shared.HTTPLogPluginConfig{
			ContentType:       contentType,
			CustomFieldsByLua: customFieldsByLua,
			FlushTimeout:      flushTimeout,
			Headers:           headers,
			HTTPEndpoint:      httpEndpoint,
			Keepalive:         keepalive,
			Method:            method,
			Queue:             queue,
			QueueSize:         queueSize,
			RetryCount:        retryCount,
			Timeout:           timeout,
		}
	}
	var consumer *shared.HTTPLogPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.HTTPLogPluginConsumer{
			ID: id2,
		}
	}
	var protocols []shared.HTTPLogPluginProtocols = []shared.HTTPLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.HTTPLogPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.HTTPLogPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.HTTPLogPluginRoute{
			ID: id3,
		}
	}
	var service *shared.HTTPLogPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.HTTPLogPluginService{
			ID: id4,
		}
	}
	out := shared.HTTPLogPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Partials:     partials,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *PluginHTTPLogResourceModel) RefreshFromSharedHTTPLogPlugin(resp *shared.HTTPLogPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.HTTPLogPluginConfig{}
			if resp.Config.ContentType != nil {
				r.Config.ContentType = types.StringValue(string(*resp.Config.ContentType))
			} else {
				r.Config.ContentType = types.StringNull()
			}
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String, len(resp.Config.CustomFieldsByLua))
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			if resp.Config.FlushTimeout != nil {
				r.Config.FlushTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.FlushTimeout)))
			} else {
				r.Config.FlushTimeout = types.NumberNull()
			}
			if len(resp.Config.Headers) > 0 {
				r.Config.Headers = make(map[string]types.String, len(resp.Config.Headers))
				for key1, value1 := range resp.Config.Headers {
					result1, _ := json.Marshal(value1)
					r.Config.Headers[key1] = types.StringValue(string(result1))
				}
			}
			r.Config.HTTPEndpoint = types.StringPointerValue(resp.Config.HTTPEndpoint)
			if resp.Config.Keepalive != nil {
				r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
			} else {
				r.Config.Keepalive = types.NumberNull()
			}
			if resp.Config.Method != nil {
				r.Config.Method = types.StringValue(string(*resp.Config.Method))
			} else {
				r.Config.Method = types.StringNull()
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
			r.Config.QueueSize = types.Int64PointerValue(resp.Config.QueueSize)
			r.Config.RetryCount = types.Int64PointerValue(resp.Config.RetryCount)
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
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
