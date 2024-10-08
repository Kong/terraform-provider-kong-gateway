// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginPostFunctionResourceModel) ToSharedCreatePostFunctionPlugin() *shared.CreatePostFunctionPlugin {
	var config *shared.CreatePostFunctionPluginConfig
	if r.Config != nil {
		var access []string = []string{}
		for _, accessItem := range r.Config.Access {
			access = append(access, accessItem.ValueString())
		}
		var bodyFilter []string = []string{}
		for _, bodyFilterItem := range r.Config.BodyFilter {
			bodyFilter = append(bodyFilter, bodyFilterItem.ValueString())
		}
		var certificate []string = []string{}
		for _, certificateItem := range r.Config.Certificate {
			certificate = append(certificate, certificateItem.ValueString())
		}
		var headerFilter []string = []string{}
		for _, headerFilterItem := range r.Config.HeaderFilter {
			headerFilter = append(headerFilter, headerFilterItem.ValueString())
		}
		var log []string = []string{}
		for _, logItem := range r.Config.Log {
			log = append(log, logItem.ValueString())
		}
		var rewrite []string = []string{}
		for _, rewriteItem := range r.Config.Rewrite {
			rewrite = append(rewrite, rewriteItem.ValueString())
		}
		var wsClientFrame []string = []string{}
		for _, wsClientFrameItem := range r.Config.WsClientFrame {
			wsClientFrame = append(wsClientFrame, wsClientFrameItem.ValueString())
		}
		var wsClose []string = []string{}
		for _, wsCloseItem := range r.Config.WsClose {
			wsClose = append(wsClose, wsCloseItem.ValueString())
		}
		var wsHandshake []string = []string{}
		for _, wsHandshakeItem := range r.Config.WsHandshake {
			wsHandshake = append(wsHandshake, wsHandshakeItem.ValueString())
		}
		var wsUpstreamFrame []string = []string{}
		for _, wsUpstreamFrameItem := range r.Config.WsUpstreamFrame {
			wsUpstreamFrame = append(wsUpstreamFrame, wsUpstreamFrameItem.ValueString())
		}
		config = &shared.CreatePostFunctionPluginConfig{
			Access:          access,
			BodyFilter:      bodyFilter,
			Certificate:     certificate,
			HeaderFilter:    headerFilter,
			Log:             log,
			Rewrite:         rewrite,
			WsClientFrame:   wsClientFrame,
			WsClose:         wsClose,
			WsHandshake:     wsHandshake,
			WsUpstreamFrame: wsUpstreamFrame,
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
	var protocols []shared.CreatePostFunctionPluginProtocols = []shared.CreatePostFunctionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreatePostFunctionPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreatePostFunctionPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreatePostFunctionPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreatePostFunctionPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreatePostFunctionPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreatePostFunctionPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreatePostFunctionPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreatePostFunctionPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreatePostFunctionPluginService{
			ID: id3,
		}
	}
	out := shared.CreatePostFunctionPlugin{
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

func (r *PluginPostFunctionResourceModel) RefreshFromSharedPostFunctionPlugin(resp *shared.PostFunctionPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreatePostFunctionPluginConfig{}
			r.Config.Access = []types.String{}
			for _, v := range resp.Config.Access {
				r.Config.Access = append(r.Config.Access, types.StringValue(v))
			}
			r.Config.BodyFilter = []types.String{}
			for _, v := range resp.Config.BodyFilter {
				r.Config.BodyFilter = append(r.Config.BodyFilter, types.StringValue(v))
			}
			r.Config.Certificate = []types.String{}
			for _, v := range resp.Config.Certificate {
				r.Config.Certificate = append(r.Config.Certificate, types.StringValue(v))
			}
			r.Config.HeaderFilter = []types.String{}
			for _, v := range resp.Config.HeaderFilter {
				r.Config.HeaderFilter = append(r.Config.HeaderFilter, types.StringValue(v))
			}
			r.Config.Log = []types.String{}
			for _, v := range resp.Config.Log {
				r.Config.Log = append(r.Config.Log, types.StringValue(v))
			}
			r.Config.Rewrite = []types.String{}
			for _, v := range resp.Config.Rewrite {
				r.Config.Rewrite = append(r.Config.Rewrite, types.StringValue(v))
			}
			r.Config.WsClientFrame = []types.String{}
			for _, v := range resp.Config.WsClientFrame {
				r.Config.WsClientFrame = append(r.Config.WsClientFrame, types.StringValue(v))
			}
			r.Config.WsClose = []types.String{}
			for _, v := range resp.Config.WsClose {
				r.Config.WsClose = append(r.Config.WsClose, types.StringValue(v))
			}
			r.Config.WsHandshake = []types.String{}
			for _, v := range resp.Config.WsHandshake {
				r.Config.WsHandshake = append(r.Config.WsHandshake, types.StringValue(v))
			}
			r.Config.WsUpstreamFrame = []types.String{}
			for _, v := range resp.Config.WsUpstreamFrame {
				r.Config.WsUpstreamFrame = append(r.Config.WsUpstreamFrame, types.StringValue(v))
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
