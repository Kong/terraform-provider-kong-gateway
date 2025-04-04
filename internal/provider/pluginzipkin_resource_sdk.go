// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginZipkinResourceModel) ToSharedZipkinPlugin() *shared.ZipkinPlugin {
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
	var ordering *shared.ZipkinPluginOrdering
	if r.Ordering != nil {
		var after *shared.ZipkinPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ZipkinPluginAfter{
				Access: access,
			}
		}
		var before *shared.ZipkinPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ZipkinPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ZipkinPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.ZipkinPluginPartials = []shared.ZipkinPluginPartials{}
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
		partials = append(partials, shared.ZipkinPluginPartials{
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
	var config *shared.ZipkinPluginConfig
	if r.Config != nil {
		connectTimeout := new(int64)
		if !r.Config.ConnectTimeout.IsUnknown() && !r.Config.ConnectTimeout.IsNull() {
			*connectTimeout = r.Config.ConnectTimeout.ValueInt64()
		} else {
			connectTimeout = nil
		}
		defaultHeaderType := new(shared.DefaultHeaderType)
		if !r.Config.DefaultHeaderType.IsUnknown() && !r.Config.DefaultHeaderType.IsNull() {
			*defaultHeaderType = shared.DefaultHeaderType(r.Config.DefaultHeaderType.ValueString())
		} else {
			defaultHeaderType = nil
		}
		defaultServiceName := new(string)
		if !r.Config.DefaultServiceName.IsUnknown() && !r.Config.DefaultServiceName.IsNull() {
			*defaultServiceName = r.Config.DefaultServiceName.ValueString()
		} else {
			defaultServiceName = nil
		}
		headerType := new(shared.ZipkinPluginHeaderType)
		if !r.Config.HeaderType.IsUnknown() && !r.Config.HeaderType.IsNull() {
			*headerType = shared.ZipkinPluginHeaderType(r.Config.HeaderType.ValueString())
		} else {
			headerType = nil
		}
		httpEndpoint := new(string)
		if !r.Config.HTTPEndpoint.IsUnknown() && !r.Config.HTTPEndpoint.IsNull() {
			*httpEndpoint = r.Config.HTTPEndpoint.ValueString()
		} else {
			httpEndpoint = nil
		}
		httpResponseHeaderForTraceid := new(string)
		if !r.Config.HTTPResponseHeaderForTraceid.IsUnknown() && !r.Config.HTTPResponseHeaderForTraceid.IsNull() {
			*httpResponseHeaderForTraceid = r.Config.HTTPResponseHeaderForTraceid.ValueString()
		} else {
			httpResponseHeaderForTraceid = nil
		}
		httpSpanName := new(shared.HTTPSpanName)
		if !r.Config.HTTPSpanName.IsUnknown() && !r.Config.HTTPSpanName.IsNull() {
			*httpSpanName = shared.HTTPSpanName(r.Config.HTTPSpanName.ValueString())
		} else {
			httpSpanName = nil
		}
		includeCredential := new(bool)
		if !r.Config.IncludeCredential.IsUnknown() && !r.Config.IncludeCredential.IsNull() {
			*includeCredential = r.Config.IncludeCredential.ValueBool()
		} else {
			includeCredential = nil
		}
		localServiceName := new(string)
		if !r.Config.LocalServiceName.IsUnknown() && !r.Config.LocalServiceName.IsNull() {
			*localServiceName = r.Config.LocalServiceName.ValueString()
		} else {
			localServiceName = nil
		}
		phaseDurationFlavor := new(shared.PhaseDurationFlavor)
		if !r.Config.PhaseDurationFlavor.IsUnknown() && !r.Config.PhaseDurationFlavor.IsNull() {
			*phaseDurationFlavor = shared.PhaseDurationFlavor(r.Config.PhaseDurationFlavor.ValueString())
		} else {
			phaseDurationFlavor = nil
		}
		var propagation *shared.ZipkinPluginPropagation
		if r.Config.Propagation != nil {
			var clear []string = []string{}
			for _, clearItem := range r.Config.Propagation.Clear {
				clear = append(clear, clearItem.ValueString())
			}
			defaultFormat := shared.ZipkinPluginDefaultFormat(r.Config.Propagation.DefaultFormat.ValueString())
			var extract []shared.ZipkinPluginExtract = []shared.ZipkinPluginExtract{}
			for _, extractItem := range r.Config.Propagation.Extract {
				extract = append(extract, shared.ZipkinPluginExtract(extractItem.ValueString()))
			}
			var inject []shared.ZipkinPluginInject = []shared.ZipkinPluginInject{}
			for _, injectItem := range r.Config.Propagation.Inject {
				inject = append(inject, shared.ZipkinPluginInject(injectItem.ValueString()))
			}
			propagation = &shared.ZipkinPluginPropagation{
				Clear:         clear,
				DefaultFormat: defaultFormat,
				Extract:       extract,
				Inject:        inject,
			}
		}
		var queue *shared.ZipkinPluginQueue
		if r.Config.Queue != nil {
			concurrencyLimit := new(shared.ZipkinPluginConcurrencyLimit)
			if !r.Config.Queue.ConcurrencyLimit.IsUnknown() && !r.Config.Queue.ConcurrencyLimit.IsNull() {
				*concurrencyLimit = shared.ZipkinPluginConcurrencyLimit(r.Config.Queue.ConcurrencyLimit.ValueInt64())
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
			queue = &shared.ZipkinPluginQueue{
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
		readTimeout := new(int64)
		if !r.Config.ReadTimeout.IsUnknown() && !r.Config.ReadTimeout.IsNull() {
			*readTimeout = r.Config.ReadTimeout.ValueInt64()
		} else {
			readTimeout = nil
		}
		sampleRatio := new(float64)
		if !r.Config.SampleRatio.IsUnknown() && !r.Config.SampleRatio.IsNull() {
			*sampleRatio, _ = r.Config.SampleRatio.ValueBigFloat().Float64()
		} else {
			sampleRatio = nil
		}
		sendTimeout := new(int64)
		if !r.Config.SendTimeout.IsUnknown() && !r.Config.SendTimeout.IsNull() {
			*sendTimeout = r.Config.SendTimeout.ValueInt64()
		} else {
			sendTimeout = nil
		}
		var staticTags []shared.StaticTags = []shared.StaticTags{}
		for _, staticTagsItem := range r.Config.StaticTags {
			var name1 string
			name1 = staticTagsItem.Name.ValueString()

			var value string
			value = staticTagsItem.Value.ValueString()

			staticTags = append(staticTags, shared.StaticTags{
				Name:  name1,
				Value: value,
			})
		}
		tagsHeader := new(string)
		if !r.Config.TagsHeader.IsUnknown() && !r.Config.TagsHeader.IsNull() {
			*tagsHeader = r.Config.TagsHeader.ValueString()
		} else {
			tagsHeader = nil
		}
		traceidByteCount := new(shared.TraceidByteCount)
		if !r.Config.TraceidByteCount.IsUnknown() && !r.Config.TraceidByteCount.IsNull() {
			*traceidByteCount = shared.TraceidByteCount(r.Config.TraceidByteCount.ValueInt64())
		} else {
			traceidByteCount = nil
		}
		config = &shared.ZipkinPluginConfig{
			ConnectTimeout:               connectTimeout,
			DefaultHeaderType:            defaultHeaderType,
			DefaultServiceName:           defaultServiceName,
			HeaderType:                   headerType,
			HTTPEndpoint:                 httpEndpoint,
			HTTPResponseHeaderForTraceid: httpResponseHeaderForTraceid,
			HTTPSpanName:                 httpSpanName,
			IncludeCredential:            includeCredential,
			LocalServiceName:             localServiceName,
			PhaseDurationFlavor:          phaseDurationFlavor,
			Propagation:                  propagation,
			Queue:                        queue,
			ReadTimeout:                  readTimeout,
			SampleRatio:                  sampleRatio,
			SendTimeout:                  sendTimeout,
			StaticTags:                   staticTags,
			TagsHeader:                   tagsHeader,
			TraceidByteCount:             traceidByteCount,
		}
	}
	var consumer *shared.ZipkinPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.ZipkinPluginConsumer{
			ID: id2,
		}
	}
	var protocols []shared.ZipkinPluginProtocols = []shared.ZipkinPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ZipkinPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ZipkinPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.ZipkinPluginRoute{
			ID: id3,
		}
	}
	var service *shared.ZipkinPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.ZipkinPluginService{
			ID: id4,
		}
	}
	out := shared.ZipkinPlugin{
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

func (r *PluginZipkinResourceModel) RefreshFromSharedZipkinPlugin(resp *shared.ZipkinPlugin) {
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
