// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginProxyCacheResourceModel) ToSharedProxyCachePlugin() *shared.ProxyCachePlugin {
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
	var ordering *shared.ProxyCachePluginOrdering
	if r.Ordering != nil {
		var after *shared.ProxyCachePluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ProxyCachePluginAfter{
				Access: access,
			}
		}
		var before *shared.ProxyCachePluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ProxyCachePluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ProxyCachePluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.ProxyCachePluginPartials = []shared.ProxyCachePluginPartials{}
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
		partials = append(partials, shared.ProxyCachePluginPartials{
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
	var config *shared.ProxyCachePluginConfig
	if r.Config != nil {
		cacheControl := new(bool)
		if !r.Config.CacheControl.IsUnknown() && !r.Config.CacheControl.IsNull() {
			*cacheControl = r.Config.CacheControl.ValueBool()
		} else {
			cacheControl = nil
		}
		cacheTTL := new(int64)
		if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
			*cacheTTL = r.Config.CacheTTL.ValueInt64()
		} else {
			cacheTTL = nil
		}
		var contentType []string = []string{}
		for _, contentTypeItem := range r.Config.ContentType {
			contentType = append(contentType, contentTypeItem.ValueString())
		}
		ignoreURICase := new(bool)
		if !r.Config.IgnoreURICase.IsUnknown() && !r.Config.IgnoreURICase.IsNull() {
			*ignoreURICase = r.Config.IgnoreURICase.ValueBool()
		} else {
			ignoreURICase = nil
		}
		var memory *shared.ProxyCachePluginMemory
		if r.Config.Memory != nil {
			dictionaryName := new(string)
			if !r.Config.Memory.DictionaryName.IsUnknown() && !r.Config.Memory.DictionaryName.IsNull() {
				*dictionaryName = r.Config.Memory.DictionaryName.ValueString()
			} else {
				dictionaryName = nil
			}
			memory = &shared.ProxyCachePluginMemory{
				DictionaryName: dictionaryName,
			}
		}
		var requestMethod []shared.RequestMethod = []shared.RequestMethod{}
		for _, requestMethodItem := range r.Config.RequestMethod {
			requestMethod = append(requestMethod, shared.RequestMethod(requestMethodItem.ValueString()))
		}
		var responseCode []int64 = []int64{}
		for _, responseCodeItem := range r.Config.ResponseCode {
			responseCode = append(responseCode, responseCodeItem.ValueInt64())
		}
		var responseHeaders *shared.ResponseHeaders
		if r.Config.ResponseHeaders != nil {
			xCacheKey := new(bool)
			if !r.Config.ResponseHeaders.XCacheKey.IsUnknown() && !r.Config.ResponseHeaders.XCacheKey.IsNull() {
				*xCacheKey = r.Config.ResponseHeaders.XCacheKey.ValueBool()
			} else {
				xCacheKey = nil
			}
			xCacheStatus := new(bool)
			if !r.Config.ResponseHeaders.XCacheStatus.IsUnknown() && !r.Config.ResponseHeaders.XCacheStatus.IsNull() {
				*xCacheStatus = r.Config.ResponseHeaders.XCacheStatus.ValueBool()
			} else {
				xCacheStatus = nil
			}
			age := new(bool)
			if !r.Config.ResponseHeaders.Age.IsUnknown() && !r.Config.ResponseHeaders.Age.IsNull() {
				*age = r.Config.ResponseHeaders.Age.ValueBool()
			} else {
				age = nil
			}
			responseHeaders = &shared.ResponseHeaders{
				XCacheKey:    xCacheKey,
				XCacheStatus: xCacheStatus,
				Age:          age,
			}
		}
		storageTTL := new(int64)
		if !r.Config.StorageTTL.IsUnknown() && !r.Config.StorageTTL.IsNull() {
			*storageTTL = r.Config.StorageTTL.ValueInt64()
		} else {
			storageTTL = nil
		}
		strategy := new(shared.ProxyCachePluginStrategy)
		if !r.Config.Strategy.IsUnknown() && !r.Config.Strategy.IsNull() {
			*strategy = shared.ProxyCachePluginStrategy(r.Config.Strategy.ValueString())
		} else {
			strategy = nil
		}
		var varyHeaders []string = []string{}
		for _, varyHeadersItem := range r.Config.VaryHeaders {
			varyHeaders = append(varyHeaders, varyHeadersItem.ValueString())
		}
		var varyQueryParams []string = []string{}
		for _, varyQueryParamsItem := range r.Config.VaryQueryParams {
			varyQueryParams = append(varyQueryParams, varyQueryParamsItem.ValueString())
		}
		config = &shared.ProxyCachePluginConfig{
			CacheControl:    cacheControl,
			CacheTTL:        cacheTTL,
			ContentType:     contentType,
			IgnoreURICase:   ignoreURICase,
			Memory:          memory,
			RequestMethod:   requestMethod,
			ResponseCode:    responseCode,
			ResponseHeaders: responseHeaders,
			StorageTTL:      storageTTL,
			Strategy:        strategy,
			VaryHeaders:     varyHeaders,
			VaryQueryParams: varyQueryParams,
		}
	}
	var consumer *shared.ProxyCachePluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.ProxyCachePluginConsumer{
			ID: id2,
		}
	}
	var consumerGroup *shared.ProxyCachePluginConsumerGroup
	if r.ConsumerGroup != nil {
		id3 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id3 = r.ConsumerGroup.ID.ValueString()
		} else {
			id3 = nil
		}
		consumerGroup = &shared.ProxyCachePluginConsumerGroup{
			ID: id3,
		}
	}
	var protocols []shared.ProxyCachePluginProtocols = []shared.ProxyCachePluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ProxyCachePluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ProxyCachePluginRoute
	if r.Route != nil {
		id4 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id4 = r.Route.ID.ValueString()
		} else {
			id4 = nil
		}
		route = &shared.ProxyCachePluginRoute{
			ID: id4,
		}
	}
	var service *shared.ProxyCachePluginService
	if r.Service != nil {
		id5 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id5 = r.Service.ID.ValueString()
		} else {
			id5 = nil
		}
		service = &shared.ProxyCachePluginService{
			ID: id5,
		}
	}
	out := shared.ProxyCachePlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Partials:      partials,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *PluginProxyCacheResourceModel) RefreshFromSharedProxyCachePlugin(resp *shared.ProxyCachePlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.ProxyCachePluginConfig{}
			r.Config.CacheControl = types.BoolPointerValue(resp.Config.CacheControl)
			r.Config.CacheTTL = types.Int64PointerValue(resp.Config.CacheTTL)
			r.Config.ContentType = make([]types.String, 0, len(resp.Config.ContentType))
			for _, v := range resp.Config.ContentType {
				r.Config.ContentType = append(r.Config.ContentType, types.StringValue(v))
			}
			r.Config.IgnoreURICase = types.BoolPointerValue(resp.Config.IgnoreURICase)
			if resp.Config.Memory == nil {
				r.Config.Memory = nil
			} else {
				r.Config.Memory = &tfTypes.Memory{}
				r.Config.Memory.DictionaryName = types.StringPointerValue(resp.Config.Memory.DictionaryName)
			}
			r.Config.RequestMethod = make([]types.String, 0, len(resp.Config.RequestMethod))
			for _, v := range resp.Config.RequestMethod {
				r.Config.RequestMethod = append(r.Config.RequestMethod, types.StringValue(string(v)))
			}
			r.Config.ResponseCode = make([]types.Int64, 0, len(resp.Config.ResponseCode))
			for _, v := range resp.Config.ResponseCode {
				r.Config.ResponseCode = append(r.Config.ResponseCode, types.Int64Value(v))
			}
			if resp.Config.ResponseHeaders == nil {
				r.Config.ResponseHeaders = nil
			} else {
				r.Config.ResponseHeaders = &tfTypes.ResponseHeaders{}
				r.Config.ResponseHeaders.Age = types.BoolPointerValue(resp.Config.ResponseHeaders.Age)
				r.Config.ResponseHeaders.XCacheKey = types.BoolPointerValue(resp.Config.ResponseHeaders.XCacheKey)
				r.Config.ResponseHeaders.XCacheStatus = types.BoolPointerValue(resp.Config.ResponseHeaders.XCacheStatus)
			}
			r.Config.StorageTTL = types.Int64PointerValue(resp.Config.StorageTTL)
			if resp.Config.Strategy != nil {
				r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
			} else {
				r.Config.Strategy = types.StringNull()
			}
			r.Config.VaryHeaders = make([]types.String, 0, len(resp.Config.VaryHeaders))
			for _, v := range resp.Config.VaryHeaders {
				r.Config.VaryHeaders = append(r.Config.VaryHeaders, types.StringValue(v))
			}
			r.Config.VaryQueryParams = make([]types.String, 0, len(resp.Config.VaryQueryParams))
			for _, v := range resp.Config.VaryQueryParams {
				r.Config.VaryQueryParams = append(r.Config.VaryQueryParams, types.StringValue(v))
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
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
