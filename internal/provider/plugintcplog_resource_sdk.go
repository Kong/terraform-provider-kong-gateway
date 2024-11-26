// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginTCPLogResourceModel) ToSharedTCPLogPluginInput() *shared.TCPLogPluginInput {
	customFieldsByLua := make(map[string]interface{})
	for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
		var customFieldsByLuaInst interface{}
		_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
		customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
	}
	host := new(string)
	if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
		*host = r.Config.Host.ValueString()
	} else {
		host = nil
	}
	keepalive := new(float64)
	if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
		*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
	} else {
		keepalive = nil
	}
	port := new(int64)
	if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
		*port = r.Config.Port.ValueInt64()
	} else {
		port = nil
	}
	timeout := new(float64)
	if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
		*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
	} else {
		timeout = nil
	}
	tls := new(bool)
	if !r.Config.TLS.IsUnknown() && !r.Config.TLS.IsNull() {
		*tls = r.Config.TLS.ValueBool()
	} else {
		tls = nil
	}
	tlsSni := new(string)
	if !r.Config.TLSSni.IsUnknown() && !r.Config.TLSSni.IsNull() {
		*tlsSni = r.Config.TLSSni.ValueString()
	} else {
		tlsSni = nil
	}
	config := shared.TCPLogPluginConfig{
		CustomFieldsByLua: customFieldsByLua,
		Host:              host,
		Keepalive:         keepalive,
		Port:              port,
		Timeout:           timeout,
		TLS:               tls,
		TLSSni:            tlsSni,
	}
	var consumer *shared.TCPLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.TCPLogPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.TCPLogPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.TCPLogPluginConsumerGroup{
			ID: id1,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.TCPLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.TCPLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.TCPLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.TCPLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.TCPLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.TCPLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.TCPLogPluginProtocols = []shared.TCPLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.TCPLogPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.TCPLogPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.TCPLogPluginRoute{
			ID: id3,
		}
	}
	var service *shared.TCPLogPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.TCPLogPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.TCPLogPluginInput{
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Enabled:       enabled,
		ID:            id2,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
		Tags:          tags,
	}
	return &out
}

func (r *PluginTCPLogResourceModel) RefreshFromSharedTCPLogPlugin(resp *shared.TCPLogPlugin) {
	if resp != nil {
		if len(resp.Config.CustomFieldsByLua) > 0 {
			r.Config.CustomFieldsByLua = make(map[string]types.String)
			for key, value := range resp.Config.CustomFieldsByLua {
				result, _ := json.Marshal(value)
				r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
			}
		}
		r.Config.Host = types.StringPointerValue(resp.Config.Host)
		if resp.Config.Keepalive != nil {
			r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
		} else {
			r.Config.Keepalive = types.NumberNull()
		}
		r.Config.Port = types.Int64PointerValue(resp.Config.Port)
		if resp.Config.Timeout != nil {
			r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
		} else {
			r.Config.Timeout = types.NumberNull()
		}
		r.Config.TLS = types.BoolPointerValue(resp.Config.TLS)
		r.Config.TLSSni = types.StringPointerValue(resp.Config.TLSSni)
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
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
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
