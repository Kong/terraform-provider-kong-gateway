// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginWebsocketValidatorResourceModel) ToSharedCreateWebsocketValidatorPlugin() *shared.CreateWebsocketValidatorPlugin {
	var config *shared.CreateWebsocketValidatorPluginConfig
	if r.Config != nil {
		var client *shared.Client
		if r.Config.Client != nil {
			var binary *shared.Binary
			if r.Config.Client.Binary != nil {
				var schema string
				schema = r.Config.Client.Binary.Schema.ValueString()

				typeVar := shared.CreateWebsocketValidatorPluginConfigClientType(r.Config.Client.Binary.Type.ValueString())
				binary = &shared.Binary{
					Schema: schema,
					Type:   typeVar,
				}
			}
			var text *shared.Text
			if r.Config.Client.Text != nil {
				var schema1 string
				schema1 = r.Config.Client.Text.Schema.ValueString()

				typeVar1 := shared.CreateWebsocketValidatorPluginType(r.Config.Client.Text.Type.ValueString())
				text = &shared.Text{
					Schema: schema1,
					Type:   typeVar1,
				}
			}
			client = &shared.Client{
				Binary: binary,
				Text:   text,
			}
		}
		var upstream *shared.CreateWebsocketValidatorPluginUpstream
		if r.Config.Upstream != nil {
			var binary1 *shared.CreateWebsocketValidatorPluginBinary
			if r.Config.Upstream.Binary != nil {
				var schema2 string
				schema2 = r.Config.Upstream.Binary.Schema.ValueString()

				typeVar2 := shared.CreateWebsocketValidatorPluginConfigType(r.Config.Upstream.Binary.Type.ValueString())
				binary1 = &shared.CreateWebsocketValidatorPluginBinary{
					Schema: schema2,
					Type:   typeVar2,
				}
			}
			var text1 *shared.CreateWebsocketValidatorPluginText
			if r.Config.Upstream.Text != nil {
				var schema3 string
				schema3 = r.Config.Upstream.Text.Schema.ValueString()

				typeVar3 := shared.CreateWebsocketValidatorPluginConfigUpstreamType(r.Config.Upstream.Text.Type.ValueString())
				text1 = &shared.CreateWebsocketValidatorPluginText{
					Schema: schema3,
					Type:   typeVar3,
				}
			}
			upstream = &shared.CreateWebsocketValidatorPluginUpstream{
				Binary: binary1,
				Text:   text1,
			}
		}
		config = &shared.CreateWebsocketValidatorPluginConfig{
			Client:   client,
			Upstream: upstream,
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
	var protocols []shared.CreateWebsocketValidatorPluginProtocols = []shared.CreateWebsocketValidatorPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateWebsocketValidatorPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateWebsocketValidatorPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateWebsocketValidatorPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateWebsocketValidatorPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateWebsocketValidatorPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateWebsocketValidatorPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateWebsocketValidatorPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateWebsocketValidatorPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateWebsocketValidatorPluginService{
			ID: id3,
		}
	}
	out := shared.CreateWebsocketValidatorPlugin{
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

func (r *PluginWebsocketValidatorResourceModel) RefreshFromSharedWebsocketValidatorPlugin(resp *shared.WebsocketValidatorPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateWebsocketValidatorPluginConfig{}
			if resp.Config.Client == nil {
				r.Config.Client = nil
			} else {
				r.Config.Client = &tfTypes.Client{}
				if resp.Config.Client.Binary == nil {
					r.Config.Client.Binary = nil
				} else {
					r.Config.Client.Binary = &tfTypes.Binary{}
					r.Config.Client.Binary.Schema = types.StringValue(resp.Config.Client.Binary.Schema)
					r.Config.Client.Binary.Type = types.StringValue(string(resp.Config.Client.Binary.Type))
				}
				if resp.Config.Client.Text == nil {
					r.Config.Client.Text = nil
				} else {
					r.Config.Client.Text = &tfTypes.Binary{}
					r.Config.Client.Text.Schema = types.StringValue(resp.Config.Client.Text.Schema)
					r.Config.Client.Text.Type = types.StringValue(string(resp.Config.Client.Text.Type))
				}
			}
			if resp.Config.Upstream == nil {
				r.Config.Upstream = nil
			} else {
				r.Config.Upstream = &tfTypes.Client{}
				if resp.Config.Upstream.Binary == nil {
					r.Config.Upstream.Binary = nil
				} else {
					r.Config.Upstream.Binary = &tfTypes.Binary{}
					r.Config.Upstream.Binary.Schema = types.StringValue(resp.Config.Upstream.Binary.Schema)
					r.Config.Upstream.Binary.Type = types.StringValue(string(resp.Config.Upstream.Binary.Type))
				}
				if resp.Config.Upstream.Text == nil {
					r.Config.Upstream.Text = nil
				} else {
					r.Config.Upstream.Text = &tfTypes.Binary{}
					r.Config.Upstream.Text.Schema = types.StringValue(resp.Config.Upstream.Text.Schema)
					r.Config.Upstream.Text.Type = types.StringValue(string(resp.Config.Upstream.Text.Type))
				}
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
