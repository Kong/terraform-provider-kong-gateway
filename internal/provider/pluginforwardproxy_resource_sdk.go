// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginForwardProxyResourceModel) ToSharedCreateForwardProxyPlugin() *shared.CreateForwardProxyPlugin {
	var config *shared.CreateForwardProxyPluginConfig
	if r.Config != nil {
		authPassword := new(string)
		if !r.Config.AuthPassword.IsUnknown() && !r.Config.AuthPassword.IsNull() {
			*authPassword = r.Config.AuthPassword.ValueString()
		} else {
			authPassword = nil
		}
		authUsername := new(string)
		if !r.Config.AuthUsername.IsUnknown() && !r.Config.AuthUsername.IsNull() {
			*authUsername = r.Config.AuthUsername.ValueString()
		} else {
			authUsername = nil
		}
		httpProxyHost := new(string)
		if !r.Config.HTTPProxyHost.IsUnknown() && !r.Config.HTTPProxyHost.IsNull() {
			*httpProxyHost = r.Config.HTTPProxyHost.ValueString()
		} else {
			httpProxyHost = nil
		}
		httpProxyPort := new(int64)
		if !r.Config.HTTPProxyPort.IsUnknown() && !r.Config.HTTPProxyPort.IsNull() {
			*httpProxyPort = r.Config.HTTPProxyPort.ValueInt64()
		} else {
			httpProxyPort = nil
		}
		httpsProxyHost := new(string)
		if !r.Config.HTTPSProxyHost.IsUnknown() && !r.Config.HTTPSProxyHost.IsNull() {
			*httpsProxyHost = r.Config.HTTPSProxyHost.ValueString()
		} else {
			httpsProxyHost = nil
		}
		httpsProxyPort := new(int64)
		if !r.Config.HTTPSProxyPort.IsUnknown() && !r.Config.HTTPSProxyPort.IsNull() {
			*httpsProxyPort = r.Config.HTTPSProxyPort.ValueInt64()
		} else {
			httpsProxyPort = nil
		}
		httpsVerify := new(bool)
		if !r.Config.HTTPSVerify.IsUnknown() && !r.Config.HTTPSVerify.IsNull() {
			*httpsVerify = r.Config.HTTPSVerify.ValueBool()
		} else {
			httpsVerify = nil
		}
		proxyScheme := new(shared.ProxyScheme)
		if !r.Config.ProxyScheme.IsUnknown() && !r.Config.ProxyScheme.IsNull() {
			*proxyScheme = shared.ProxyScheme(r.Config.ProxyScheme.ValueString())
		} else {
			proxyScheme = nil
		}
		xHeaders := new(shared.XHeaders)
		if !r.Config.XHeaders.IsUnknown() && !r.Config.XHeaders.IsNull() {
			*xHeaders = shared.XHeaders(r.Config.XHeaders.ValueString())
		} else {
			xHeaders = nil
		}
		config = &shared.CreateForwardProxyPluginConfig{
			AuthPassword:   authPassword,
			AuthUsername:   authUsername,
			HTTPProxyHost:  httpProxyHost,
			HTTPProxyPort:  httpProxyPort,
			HTTPSProxyHost: httpsProxyHost,
			HTTPSProxyPort: httpsProxyPort,
			HTTPSVerify:    httpsVerify,
			ProxyScheme:    proxyScheme,
			XHeaders:       xHeaders,
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
	var protocols []shared.CreateForwardProxyPluginProtocols = []shared.CreateForwardProxyPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateForwardProxyPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateForwardProxyPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateForwardProxyPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateForwardProxyPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateForwardProxyPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateForwardProxyPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateForwardProxyPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateForwardProxyPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateForwardProxyPluginService{
			ID: id3,
		}
	}
	out := shared.CreateForwardProxyPlugin{
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

func (r *PluginForwardProxyResourceModel) RefreshFromSharedForwardProxyPlugin(resp *shared.ForwardProxyPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateForwardProxyPluginConfig{}
			r.Config.AuthPassword = types.StringPointerValue(resp.Config.AuthPassword)
			r.Config.AuthUsername = types.StringPointerValue(resp.Config.AuthUsername)
			r.Config.HTTPProxyHost = types.StringPointerValue(resp.Config.HTTPProxyHost)
			r.Config.HTTPProxyPort = types.Int64PointerValue(resp.Config.HTTPProxyPort)
			r.Config.HTTPSProxyHost = types.StringPointerValue(resp.Config.HTTPSProxyHost)
			r.Config.HTTPSProxyPort = types.Int64PointerValue(resp.Config.HTTPSProxyPort)
			r.Config.HTTPSVerify = types.BoolPointerValue(resp.Config.HTTPSVerify)
			if resp.Config.ProxyScheme != nil {
				r.Config.ProxyScheme = types.StringValue(string(*resp.Config.ProxyScheme))
			} else {
				r.Config.ProxyScheme = types.StringNull()
			}
			if resp.Config.XHeaders != nil {
				r.Config.XHeaders = types.StringValue(string(*resp.Config.XHeaders))
			} else {
				r.Config.XHeaders = types.StringNull()
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
