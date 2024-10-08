// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginRequestTerminationResourceModel) ToSharedCreateRequestTerminationPlugin() *shared.CreateRequestTerminationPlugin {
	var config *shared.CreateRequestTerminationPluginConfig
	if r.Config != nil {
		body := new(string)
		if !r.Config.Body.IsUnknown() && !r.Config.Body.IsNull() {
			*body = r.Config.Body.ValueString()
		} else {
			body = nil
		}
		contentType := new(string)
		if !r.Config.ContentType.IsUnknown() && !r.Config.ContentType.IsNull() {
			*contentType = r.Config.ContentType.ValueString()
		} else {
			contentType = nil
		}
		echo := new(bool)
		if !r.Config.Echo.IsUnknown() && !r.Config.Echo.IsNull() {
			*echo = r.Config.Echo.ValueBool()
		} else {
			echo = nil
		}
		message := new(string)
		if !r.Config.Message.IsUnknown() && !r.Config.Message.IsNull() {
			*message = r.Config.Message.ValueString()
		} else {
			message = nil
		}
		statusCode := new(int64)
		if !r.Config.StatusCode.IsUnknown() && !r.Config.StatusCode.IsNull() {
			*statusCode = r.Config.StatusCode.ValueInt64()
		} else {
			statusCode = nil
		}
		trigger := new(string)
		if !r.Config.Trigger.IsUnknown() && !r.Config.Trigger.IsNull() {
			*trigger = r.Config.Trigger.ValueString()
		} else {
			trigger = nil
		}
		config = &shared.CreateRequestTerminationPluginConfig{
			Body:        body,
			ContentType: contentType,
			Echo:        echo,
			Message:     message,
			StatusCode:  statusCode,
			Trigger:     trigger,
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
	var protocols []shared.CreateRequestTerminationPluginProtocols = []shared.CreateRequestTerminationPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateRequestTerminationPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateRequestTerminationPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateRequestTerminationPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateRequestTerminationPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateRequestTerminationPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateRequestTerminationPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateRequestTerminationPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateRequestTerminationPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateRequestTerminationPluginService{
			ID: id3,
		}
	}
	out := shared.CreateRequestTerminationPlugin{
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

func (r *PluginRequestTerminationResourceModel) RefreshFromSharedRequestTerminationPlugin(resp *shared.RequestTerminationPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateRequestTerminationPluginConfig{}
			r.Config.Body = types.StringPointerValue(resp.Config.Body)
			r.Config.ContentType = types.StringPointerValue(resp.Config.ContentType)
			r.Config.Echo = types.BoolPointerValue(resp.Config.Echo)
			r.Config.Message = types.StringPointerValue(resp.Config.Message)
			r.Config.StatusCode = types.Int64PointerValue(resp.Config.StatusCode)
			r.Config.Trigger = types.StringPointerValue(resp.Config.Trigger)
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
