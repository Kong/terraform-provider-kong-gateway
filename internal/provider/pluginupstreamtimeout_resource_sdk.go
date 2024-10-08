// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginUpstreamTimeoutResourceModel) ToSharedCreateUpstreamTimeoutPlugin() *shared.CreateUpstreamTimeoutPlugin {
	var config *shared.CreateUpstreamTimeoutPluginConfig
	if r.Config != nil {
		connectTimeout := new(int64)
		if !r.Config.ConnectTimeout.IsUnknown() && !r.Config.ConnectTimeout.IsNull() {
			*connectTimeout = r.Config.ConnectTimeout.ValueInt64()
		} else {
			connectTimeout = nil
		}
		readTimeout := new(int64)
		if !r.Config.ReadTimeout.IsUnknown() && !r.Config.ReadTimeout.IsNull() {
			*readTimeout = r.Config.ReadTimeout.ValueInt64()
		} else {
			readTimeout = nil
		}
		sendTimeout := new(int64)
		if !r.Config.SendTimeout.IsUnknown() && !r.Config.SendTimeout.IsNull() {
			*sendTimeout = r.Config.SendTimeout.ValueInt64()
		} else {
			sendTimeout = nil
		}
		config = &shared.CreateUpstreamTimeoutPluginConfig{
			ConnectTimeout: connectTimeout,
			ReadTimeout:    readTimeout,
			SendTimeout:    sendTimeout,
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
	var protocols []shared.CreateUpstreamTimeoutPluginProtocols = []shared.CreateUpstreamTimeoutPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateUpstreamTimeoutPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateUpstreamTimeoutPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateUpstreamTimeoutPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateUpstreamTimeoutPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateUpstreamTimeoutPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateUpstreamTimeoutPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateUpstreamTimeoutPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateUpstreamTimeoutPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateUpstreamTimeoutPluginService{
			ID: id3,
		}
	}
	out := shared.CreateUpstreamTimeoutPlugin{
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

func (r *PluginUpstreamTimeoutResourceModel) RefreshFromSharedUpstreamTimeoutPlugin(resp *shared.UpstreamTimeoutPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateUpstreamTimeoutPluginConfig{}
			r.Config.ConnectTimeout = types.Int64PointerValue(resp.Config.ConnectTimeout)
			r.Config.ReadTimeout = types.Int64PointerValue(resp.Config.ReadTimeout)
			r.Config.SendTimeout = types.Int64PointerValue(resp.Config.SendTimeout)
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
