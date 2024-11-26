// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginOauth2IntrospectionResourceModel) ToSharedOauth2IntrospectionPluginInput() *shared.Oauth2IntrospectionPluginInput {
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	authorizationValue := new(string)
	if !r.Config.AuthorizationValue.IsUnknown() && !r.Config.AuthorizationValue.IsNull() {
		*authorizationValue = r.Config.AuthorizationValue.ValueString()
	} else {
		authorizationValue = nil
	}
	consumerBy := new(shared.Oauth2IntrospectionPluginConsumerBy)
	if !r.Config.ConsumerBy.IsUnknown() && !r.Config.ConsumerBy.IsNull() {
		*consumerBy = shared.Oauth2IntrospectionPluginConsumerBy(r.Config.ConsumerBy.ValueString())
	} else {
		consumerBy = nil
	}
	var customClaimsForward []string = []string{}
	for _, customClaimsForwardItem := range r.Config.CustomClaimsForward {
		customClaimsForward = append(customClaimsForward, customClaimsForwardItem.ValueString())
	}
	customIntrospectionHeaders := make(map[string]interface{})
	for customIntrospectionHeadersKey, customIntrospectionHeadersValue := range r.Config.CustomIntrospectionHeaders {
		var customIntrospectionHeadersInst interface{}
		_ = json.Unmarshal([]byte(customIntrospectionHeadersValue.ValueString()), &customIntrospectionHeadersInst)
		customIntrospectionHeaders[customIntrospectionHeadersKey] = customIntrospectionHeadersInst
	}
	hideCredentials := new(bool)
	if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
		*hideCredentials = r.Config.HideCredentials.ValueBool()
	} else {
		hideCredentials = nil
	}
	introspectRequest := new(bool)
	if !r.Config.IntrospectRequest.IsUnknown() && !r.Config.IntrospectRequest.IsNull() {
		*introspectRequest = r.Config.IntrospectRequest.ValueBool()
	} else {
		introspectRequest = nil
	}
	introspectionURL := new(string)
	if !r.Config.IntrospectionURL.IsUnknown() && !r.Config.IntrospectionURL.IsNull() {
		*introspectionURL = r.Config.IntrospectionURL.ValueString()
	} else {
		introspectionURL = nil
	}
	keepalive := new(int64)
	if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
		*keepalive = r.Config.Keepalive.ValueInt64()
	} else {
		keepalive = nil
	}
	runOnPreflight := new(bool)
	if !r.Config.RunOnPreflight.IsUnknown() && !r.Config.RunOnPreflight.IsNull() {
		*runOnPreflight = r.Config.RunOnPreflight.ValueBool()
	} else {
		runOnPreflight = nil
	}
	timeout := new(int64)
	if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
		*timeout = r.Config.Timeout.ValueInt64()
	} else {
		timeout = nil
	}
	tokenTypeHint := new(string)
	if !r.Config.TokenTypeHint.IsUnknown() && !r.Config.TokenTypeHint.IsNull() {
		*tokenTypeHint = r.Config.TokenTypeHint.ValueString()
	} else {
		tokenTypeHint = nil
	}
	ttl := new(float64)
	if !r.Config.TTL.IsUnknown() && !r.Config.TTL.IsNull() {
		*ttl, _ = r.Config.TTL.ValueBigFloat().Float64()
	} else {
		ttl = nil
	}
	config := shared.Oauth2IntrospectionPluginConfig{
		Anonymous:                  anonymous,
		AuthorizationValue:         authorizationValue,
		ConsumerBy:                 consumerBy,
		CustomClaimsForward:        customClaimsForward,
		CustomIntrospectionHeaders: customIntrospectionHeaders,
		HideCredentials:            hideCredentials,
		IntrospectRequest:          introspectRequest,
		IntrospectionURL:           introspectionURL,
		Keepalive:                  keepalive,
		RunOnPreflight:             runOnPreflight,
		Timeout:                    timeout,
		TokenTypeHint:              tokenTypeHint,
		TTL:                        ttl,
	}
	var consumer *shared.Oauth2IntrospectionPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.Oauth2IntrospectionPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.Oauth2IntrospectionPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.Oauth2IntrospectionPluginConsumerGroup{
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
	var ordering *shared.Oauth2IntrospectionPluginOrdering
	if r.Ordering != nil {
		var after *shared.Oauth2IntrospectionPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.Oauth2IntrospectionPluginAfter{
				Access: access,
			}
		}
		var before *shared.Oauth2IntrospectionPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.Oauth2IntrospectionPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.Oauth2IntrospectionPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.Oauth2IntrospectionPluginProtocols = []shared.Oauth2IntrospectionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.Oauth2IntrospectionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.Oauth2IntrospectionPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.Oauth2IntrospectionPluginRoute{
			ID: id3,
		}
	}
	var service *shared.Oauth2IntrospectionPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.Oauth2IntrospectionPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.Oauth2IntrospectionPluginInput{
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

func (r *PluginOauth2IntrospectionResourceModel) RefreshFromSharedOauth2IntrospectionPlugin(resp *shared.Oauth2IntrospectionPlugin) {
	if resp != nil {
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AuthorizationValue = types.StringPointerValue(resp.Config.AuthorizationValue)
		if resp.Config.ConsumerBy != nil {
			r.Config.ConsumerBy = types.StringValue(string(*resp.Config.ConsumerBy))
		} else {
			r.Config.ConsumerBy = types.StringNull()
		}
		r.Config.CustomClaimsForward = []types.String{}
		for _, v := range resp.Config.CustomClaimsForward {
			r.Config.CustomClaimsForward = append(r.Config.CustomClaimsForward, types.StringValue(v))
		}
		if len(resp.Config.CustomIntrospectionHeaders) > 0 {
			r.Config.CustomIntrospectionHeaders = make(map[string]types.String)
			for key, value := range resp.Config.CustomIntrospectionHeaders {
				result, _ := json.Marshal(value)
				r.Config.CustomIntrospectionHeaders[key] = types.StringValue(string(result))
			}
		}
		r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
		r.Config.IntrospectRequest = types.BoolPointerValue(resp.Config.IntrospectRequest)
		r.Config.IntrospectionURL = types.StringPointerValue(resp.Config.IntrospectionURL)
		r.Config.Keepalive = types.Int64PointerValue(resp.Config.Keepalive)
		r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
		r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
		r.Config.TokenTypeHint = types.StringPointerValue(resp.Config.TokenTypeHint)
		if resp.Config.TTL != nil {
			r.Config.TTL = types.NumberValue(big.NewFloat(float64(*resp.Config.TTL)))
		} else {
			r.Config.TTL = types.NumberNull()
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
