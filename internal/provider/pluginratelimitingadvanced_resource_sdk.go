// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginRateLimitingAdvancedResourceModel) ToSharedCreateRateLimitingAdvancedPlugin() *shared.CreateRateLimitingAdvancedPlugin {
	var config *shared.CreateRateLimitingAdvancedPluginConfig
	if r.Config != nil {
		var consumerGroups []string = []string{}
		for _, consumerGroupsItem := range r.Config.ConsumerGroups {
			consumerGroups = append(consumerGroups, consumerGroupsItem.ValueString())
		}
		dictionaryName := new(string)
		if !r.Config.DictionaryName.IsUnknown() && !r.Config.DictionaryName.IsNull() {
			*dictionaryName = r.Config.DictionaryName.ValueString()
		} else {
			dictionaryName = nil
		}
		disablePenalty := new(bool)
		if !r.Config.DisablePenalty.IsUnknown() && !r.Config.DisablePenalty.IsNull() {
			*disablePenalty = r.Config.DisablePenalty.ValueBool()
		} else {
			disablePenalty = nil
		}
		enforceConsumerGroups := new(bool)
		if !r.Config.EnforceConsumerGroups.IsUnknown() && !r.Config.EnforceConsumerGroups.IsNull() {
			*enforceConsumerGroups = r.Config.EnforceConsumerGroups.ValueBool()
		} else {
			enforceConsumerGroups = nil
		}
		errorCode := new(float64)
		if !r.Config.ErrorCode.IsUnknown() && !r.Config.ErrorCode.IsNull() {
			*errorCode, _ = r.Config.ErrorCode.ValueBigFloat().Float64()
		} else {
			errorCode = nil
		}
		errorMessage := new(string)
		if !r.Config.ErrorMessage.IsUnknown() && !r.Config.ErrorMessage.IsNull() {
			*errorMessage = r.Config.ErrorMessage.ValueString()
		} else {
			errorMessage = nil
		}
		headerName := new(string)
		if !r.Config.HeaderName.IsUnknown() && !r.Config.HeaderName.IsNull() {
			*headerName = r.Config.HeaderName.ValueString()
		} else {
			headerName = nil
		}
		hideClientHeaders := new(bool)
		if !r.Config.HideClientHeaders.IsUnknown() && !r.Config.HideClientHeaders.IsNull() {
			*hideClientHeaders = r.Config.HideClientHeaders.ValueBool()
		} else {
			hideClientHeaders = nil
		}
		identifier := new(shared.Identifier)
		if !r.Config.Identifier.IsUnknown() && !r.Config.Identifier.IsNull() {
			*identifier = shared.Identifier(r.Config.Identifier.ValueString())
		} else {
			identifier = nil
		}
		var limit []float64 = []float64{}
		for _, limitItem := range r.Config.Limit {
			limitTmp, _ := limitItem.ValueBigFloat().Float64()
			limit = append(limit, limitTmp)
		}
		namespace := new(string)
		if !r.Config.Namespace.IsUnknown() && !r.Config.Namespace.IsNull() {
			*namespace = r.Config.Namespace.ValueString()
		} else {
			namespace = nil
		}
		path := new(string)
		if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
			*path = r.Config.Path.ValueString()
		} else {
			path = nil
		}
		var redis *shared.CreateRateLimitingAdvancedPluginRedis
		if r.Config.Redis != nil {
			var clusterAddresses []string = []string{}
			for _, clusterAddressesItem := range r.Config.Redis.ClusterAddresses {
				clusterAddresses = append(clusterAddresses, clusterAddressesItem.ValueString())
			}
			connectTimeout := new(int64)
			if !r.Config.Redis.ConnectTimeout.IsUnknown() && !r.Config.Redis.ConnectTimeout.IsNull() {
				*connectTimeout = r.Config.Redis.ConnectTimeout.ValueInt64()
			} else {
				connectTimeout = nil
			}
			database := new(int64)
			if !r.Config.Redis.Database.IsUnknown() && !r.Config.Redis.Database.IsNull() {
				*database = r.Config.Redis.Database.ValueInt64()
			} else {
				database = nil
			}
			host := new(string)
			if !r.Config.Redis.Host.IsUnknown() && !r.Config.Redis.Host.IsNull() {
				*host = r.Config.Redis.Host.ValueString()
			} else {
				host = nil
			}
			keepaliveBacklog := new(int64)
			if !r.Config.Redis.KeepaliveBacklog.IsUnknown() && !r.Config.Redis.KeepaliveBacklog.IsNull() {
				*keepaliveBacklog = r.Config.Redis.KeepaliveBacklog.ValueInt64()
			} else {
				keepaliveBacklog = nil
			}
			keepalivePoolSize := new(int64)
			if !r.Config.Redis.KeepalivePoolSize.IsUnknown() && !r.Config.Redis.KeepalivePoolSize.IsNull() {
				*keepalivePoolSize = r.Config.Redis.KeepalivePoolSize.ValueInt64()
			} else {
				keepalivePoolSize = nil
			}
			password := new(string)
			if !r.Config.Redis.Password.IsUnknown() && !r.Config.Redis.Password.IsNull() {
				*password = r.Config.Redis.Password.ValueString()
			} else {
				password = nil
			}
			port := new(int64)
			if !r.Config.Redis.Port.IsUnknown() && !r.Config.Redis.Port.IsNull() {
				*port = r.Config.Redis.Port.ValueInt64()
			} else {
				port = nil
			}
			readTimeout := new(int64)
			if !r.Config.Redis.ReadTimeout.IsUnknown() && !r.Config.Redis.ReadTimeout.IsNull() {
				*readTimeout = r.Config.Redis.ReadTimeout.ValueInt64()
			} else {
				readTimeout = nil
			}
			sendTimeout := new(int64)
			if !r.Config.Redis.SendTimeout.IsUnknown() && !r.Config.Redis.SendTimeout.IsNull() {
				*sendTimeout = r.Config.Redis.SendTimeout.ValueInt64()
			} else {
				sendTimeout = nil
			}
			var sentinelAddresses []string = []string{}
			for _, sentinelAddressesItem := range r.Config.Redis.SentinelAddresses {
				sentinelAddresses = append(sentinelAddresses, sentinelAddressesItem.ValueString())
			}
			sentinelMaster := new(string)
			if !r.Config.Redis.SentinelMaster.IsUnknown() && !r.Config.Redis.SentinelMaster.IsNull() {
				*sentinelMaster = r.Config.Redis.SentinelMaster.ValueString()
			} else {
				sentinelMaster = nil
			}
			sentinelPassword := new(string)
			if !r.Config.Redis.SentinelPassword.IsUnknown() && !r.Config.Redis.SentinelPassword.IsNull() {
				*sentinelPassword = r.Config.Redis.SentinelPassword.ValueString()
			} else {
				sentinelPassword = nil
			}
			sentinelRole := new(shared.SentinelRole)
			if !r.Config.Redis.SentinelRole.IsUnknown() && !r.Config.Redis.SentinelRole.IsNull() {
				*sentinelRole = shared.SentinelRole(r.Config.Redis.SentinelRole.ValueString())
			} else {
				sentinelRole = nil
			}
			sentinelUsername := new(string)
			if !r.Config.Redis.SentinelUsername.IsUnknown() && !r.Config.Redis.SentinelUsername.IsNull() {
				*sentinelUsername = r.Config.Redis.SentinelUsername.ValueString()
			} else {
				sentinelUsername = nil
			}
			serverName := new(string)
			if !r.Config.Redis.ServerName.IsUnknown() && !r.Config.Redis.ServerName.IsNull() {
				*serverName = r.Config.Redis.ServerName.ValueString()
			} else {
				serverName = nil
			}
			ssl := new(bool)
			if !r.Config.Redis.Ssl.IsUnknown() && !r.Config.Redis.Ssl.IsNull() {
				*ssl = r.Config.Redis.Ssl.ValueBool()
			} else {
				ssl = nil
			}
			sslVerify := new(bool)
			if !r.Config.Redis.SslVerify.IsUnknown() && !r.Config.Redis.SslVerify.IsNull() {
				*sslVerify = r.Config.Redis.SslVerify.ValueBool()
			} else {
				sslVerify = nil
			}
			timeout := new(int64)
			if !r.Config.Redis.Timeout.IsUnknown() && !r.Config.Redis.Timeout.IsNull() {
				*timeout = r.Config.Redis.Timeout.ValueInt64()
			} else {
				timeout = nil
			}
			username := new(string)
			if !r.Config.Redis.Username.IsUnknown() && !r.Config.Redis.Username.IsNull() {
				*username = r.Config.Redis.Username.ValueString()
			} else {
				username = nil
			}
			redis = &shared.CreateRateLimitingAdvancedPluginRedis{
				ClusterAddresses:  clusterAddresses,
				ConnectTimeout:    connectTimeout,
				Database:          database,
				Host:              host,
				KeepaliveBacklog:  keepaliveBacklog,
				KeepalivePoolSize: keepalivePoolSize,
				Password:          password,
				Port:              port,
				ReadTimeout:       readTimeout,
				SendTimeout:       sendTimeout,
				SentinelAddresses: sentinelAddresses,
				SentinelMaster:    sentinelMaster,
				SentinelPassword:  sentinelPassword,
				SentinelRole:      sentinelRole,
				SentinelUsername:  sentinelUsername,
				ServerName:        serverName,
				Ssl:               ssl,
				SslVerify:         sslVerify,
				Timeout:           timeout,
				Username:          username,
			}
		}
		retryAfterJitterMax := new(float64)
		if !r.Config.RetryAfterJitterMax.IsUnknown() && !r.Config.RetryAfterJitterMax.IsNull() {
			*retryAfterJitterMax, _ = r.Config.RetryAfterJitterMax.ValueBigFloat().Float64()
		} else {
			retryAfterJitterMax = nil
		}
		strategy := new(shared.CreateRateLimitingAdvancedPluginStrategy)
		if !r.Config.Strategy.IsUnknown() && !r.Config.Strategy.IsNull() {
			*strategy = shared.CreateRateLimitingAdvancedPluginStrategy(r.Config.Strategy.ValueString())
		} else {
			strategy = nil
		}
		syncRate := new(float64)
		if !r.Config.SyncRate.IsUnknown() && !r.Config.SyncRate.IsNull() {
			*syncRate, _ = r.Config.SyncRate.ValueBigFloat().Float64()
		} else {
			syncRate = nil
		}
		var windowSize []float64 = []float64{}
		for _, windowSizeItem := range r.Config.WindowSize {
			windowSizeTmp, _ := windowSizeItem.ValueBigFloat().Float64()
			windowSize = append(windowSize, windowSizeTmp)
		}
		windowType := new(shared.WindowType)
		if !r.Config.WindowType.IsUnknown() && !r.Config.WindowType.IsNull() {
			*windowType = shared.WindowType(r.Config.WindowType.ValueString())
		} else {
			windowType = nil
		}
		config = &shared.CreateRateLimitingAdvancedPluginConfig{
			ConsumerGroups:        consumerGroups,
			DictionaryName:        dictionaryName,
			DisablePenalty:        disablePenalty,
			EnforceConsumerGroups: enforceConsumerGroups,
			ErrorCode:             errorCode,
			ErrorMessage:          errorMessage,
			HeaderName:            headerName,
			HideClientHeaders:     hideClientHeaders,
			Identifier:            identifier,
			Limit:                 limit,
			Namespace:             namespace,
			Path:                  path,
			Redis:                 redis,
			RetryAfterJitterMax:   retryAfterJitterMax,
			Strategy:              strategy,
			SyncRate:              syncRate,
			WindowSize:            windowSize,
			WindowType:            windowType,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering interface{}
	if !r.Ordering.IsUnknown() && !r.Ordering.IsNull() {
		_ = json.Unmarshal([]byte(r.Ordering.ValueString()), &ordering)
	}
	var protocols []shared.CreateRateLimitingAdvancedPluginProtocols = []shared.CreateRateLimitingAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateRateLimitingAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateRateLimitingAdvancedPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateRateLimitingAdvancedPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateRateLimitingAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateRateLimitingAdvancedPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateRateLimitingAdvancedPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateRateLimitingAdvancedPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateRateLimitingAdvancedPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateRateLimitingAdvancedPluginService{
			ID: id3,
		}
	}
	out := shared.CreateRateLimitingAdvancedPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *PluginRateLimitingAdvancedResourceModel) RefreshFromSharedRateLimitingAdvancedPlugin(resp *shared.RateLimitingAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateRateLimitingAdvancedPluginConfig{}
			r.Config.ConsumerGroups = []types.String{}
			for _, v := range resp.Config.ConsumerGroups {
				r.Config.ConsumerGroups = append(r.Config.ConsumerGroups, types.StringValue(v))
			}
			r.Config.DictionaryName = types.StringPointerValue(resp.Config.DictionaryName)
			r.Config.DisablePenalty = types.BoolPointerValue(resp.Config.DisablePenalty)
			r.Config.EnforceConsumerGroups = types.BoolPointerValue(resp.Config.EnforceConsumerGroups)
			if resp.Config.ErrorCode != nil {
				r.Config.ErrorCode = types.NumberValue(big.NewFloat(float64(*resp.Config.ErrorCode)))
			} else {
				r.Config.ErrorCode = types.NumberNull()
			}
			r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
			r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.Identifier != nil {
				r.Config.Identifier = types.StringValue(string(*resp.Config.Identifier))
			} else {
				r.Config.Identifier = types.StringNull()
			}
			r.Config.Limit = []types.Number{}
			for _, v := range resp.Config.Limit {
				r.Config.Limit = append(r.Config.Limit, types.NumberValue(big.NewFloat(float64(v))))
			}
			r.Config.Namespace = types.StringPointerValue(resp.Config.Namespace)
			r.Config.Path = types.StringPointerValue(resp.Config.Path)
			if resp.Config.Redis == nil {
				r.Config.Redis = nil
			} else {
				r.Config.Redis = &tfTypes.CreateRateLimitingAdvancedPluginRedis{}
				r.Config.Redis.ClusterAddresses = []types.String{}
				for _, v := range resp.Config.Redis.ClusterAddresses {
					r.Config.Redis.ClusterAddresses = append(r.Config.Redis.ClusterAddresses, types.StringValue(v))
				}
				r.Config.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Redis.ConnectTimeout)
				r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
				r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
				r.Config.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Redis.KeepaliveBacklog)
				r.Config.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Redis.KeepalivePoolSize)
				r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
				r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
				r.Config.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Redis.ReadTimeout)
				r.Config.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Redis.SendTimeout)
				r.Config.Redis.SentinelAddresses = []types.String{}
				for _, v := range resp.Config.Redis.SentinelAddresses {
					r.Config.Redis.SentinelAddresses = append(r.Config.Redis.SentinelAddresses, types.StringValue(v))
				}
				r.Config.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Redis.SentinelMaster)
				r.Config.Redis.SentinelPassword = types.StringPointerValue(resp.Config.Redis.SentinelPassword)
				if resp.Config.Redis.SentinelRole != nil {
					r.Config.Redis.SentinelRole = types.StringValue(string(*resp.Config.Redis.SentinelRole))
				} else {
					r.Config.Redis.SentinelRole = types.StringNull()
				}
				r.Config.Redis.SentinelUsername = types.StringPointerValue(resp.Config.Redis.SentinelUsername)
				r.Config.Redis.ServerName = types.StringPointerValue(resp.Config.Redis.ServerName)
				r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
				r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
				r.Config.Redis.Timeout = types.Int64PointerValue(resp.Config.Redis.Timeout)
				r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
			}
			if resp.Config.RetryAfterJitterMax != nil {
				r.Config.RetryAfterJitterMax = types.NumberValue(big.NewFloat(float64(*resp.Config.RetryAfterJitterMax)))
			} else {
				r.Config.RetryAfterJitterMax = types.NumberNull()
			}
			if resp.Config.Strategy != nil {
				r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
			} else {
				r.Config.Strategy = types.StringNull()
			}
			if resp.Config.SyncRate != nil {
				r.Config.SyncRate = types.NumberValue(big.NewFloat(float64(*resp.Config.SyncRate)))
			} else {
				r.Config.SyncRate = types.NumberNull()
			}
			r.Config.WindowSize = []types.Number{}
			for _, v := range resp.Config.WindowSize {
				r.Config.WindowSize = append(r.Config.WindowSize, types.NumberValue(big.NewFloat(float64(v))))
			}
			if resp.Config.WindowType != nil {
				r.Config.WindowType = types.StringValue(string(*resp.Config.WindowType))
			} else {
				r.Config.WindowType = types.StringNull()
			}
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