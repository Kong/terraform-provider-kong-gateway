// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginJwtSignerDataSourceModel) RefreshFromSharedJwtSignerPlugin(resp *shared.JwtSignerPlugin) {
	if resp != nil {
		r.Config.AccessTokenConsumerBy = []types.String{}
		for _, v := range resp.Config.AccessTokenConsumerBy {
			r.Config.AccessTokenConsumerBy = append(r.Config.AccessTokenConsumerBy, types.StringValue(string(v)))
		}
		r.Config.AccessTokenConsumerClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenConsumerClaim {
			r.Config.AccessTokenConsumerClaim = append(r.Config.AccessTokenConsumerClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionAuthorization = types.StringPointerValue(resp.Config.AccessTokenIntrospectionAuthorization)
		r.Config.AccessTokenIntrospectionBodyArgs = types.StringPointerValue(resp.Config.AccessTokenIntrospectionBodyArgs)
		r.Config.AccessTokenIntrospectionConsumerBy = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionConsumerBy {
			r.Config.AccessTokenIntrospectionConsumerBy = append(r.Config.AccessTokenIntrospectionConsumerBy, types.StringValue(string(v)))
		}
		r.Config.AccessTokenIntrospectionConsumerClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionConsumerClaim {
			r.Config.AccessTokenIntrospectionConsumerClaim = append(r.Config.AccessTokenIntrospectionConsumerClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionEndpoint = types.StringPointerValue(resp.Config.AccessTokenIntrospectionEndpoint)
		r.Config.AccessTokenIntrospectionHint = types.StringPointerValue(resp.Config.AccessTokenIntrospectionHint)
		r.Config.AccessTokenIntrospectionJwtClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionJwtClaim {
			r.Config.AccessTokenIntrospectionJwtClaim = append(r.Config.AccessTokenIntrospectionJwtClaim, types.StringValue(v))
		}
		if resp.Config.AccessTokenIntrospectionLeeway != nil {
			r.Config.AccessTokenIntrospectionLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenIntrospectionLeeway)))
		} else {
			r.Config.AccessTokenIntrospectionLeeway = types.NumberNull()
		}
		r.Config.AccessTokenIntrospectionScopesClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionScopesClaim {
			r.Config.AccessTokenIntrospectionScopesClaim = append(r.Config.AccessTokenIntrospectionScopesClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionScopesRequired = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionScopesRequired {
			r.Config.AccessTokenIntrospectionScopesRequired = append(r.Config.AccessTokenIntrospectionScopesRequired, types.StringValue(v))
		}
		if resp.Config.AccessTokenIntrospectionTimeout != nil {
			r.Config.AccessTokenIntrospectionTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenIntrospectionTimeout)))
		} else {
			r.Config.AccessTokenIntrospectionTimeout = types.NumberNull()
		}
		r.Config.AccessTokenIssuer = types.StringPointerValue(resp.Config.AccessTokenIssuer)
		r.Config.AccessTokenJwksURI = types.StringPointerValue(resp.Config.AccessTokenJwksURI)
		r.Config.AccessTokenJwksURIClientCertificate = types.StringPointerValue(resp.Config.AccessTokenJwksURIClientCertificate)
		r.Config.AccessTokenJwksURIClientPassword = types.StringPointerValue(resp.Config.AccessTokenJwksURIClientPassword)
		r.Config.AccessTokenJwksURIClientUsername = types.StringPointerValue(resp.Config.AccessTokenJwksURIClientUsername)
		if resp.Config.AccessTokenJwksURIRotatePeriod != nil {
			r.Config.AccessTokenJwksURIRotatePeriod = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenJwksURIRotatePeriod)))
		} else {
			r.Config.AccessTokenJwksURIRotatePeriod = types.NumberNull()
		}
		r.Config.AccessTokenKeyset = types.StringPointerValue(resp.Config.AccessTokenKeyset)
		r.Config.AccessTokenKeysetClientCertificate = types.StringPointerValue(resp.Config.AccessTokenKeysetClientCertificate)
		r.Config.AccessTokenKeysetClientPassword = types.StringPointerValue(resp.Config.AccessTokenKeysetClientPassword)
		r.Config.AccessTokenKeysetClientUsername = types.StringPointerValue(resp.Config.AccessTokenKeysetClientUsername)
		if resp.Config.AccessTokenKeysetRotatePeriod != nil {
			r.Config.AccessTokenKeysetRotatePeriod = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenKeysetRotatePeriod)))
		} else {
			r.Config.AccessTokenKeysetRotatePeriod = types.NumberNull()
		}
		if resp.Config.AccessTokenLeeway != nil {
			r.Config.AccessTokenLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenLeeway)))
		} else {
			r.Config.AccessTokenLeeway = types.NumberNull()
		}
		r.Config.AccessTokenOptional = types.BoolPointerValue(resp.Config.AccessTokenOptional)
		r.Config.AccessTokenRequestHeader = types.StringPointerValue(resp.Config.AccessTokenRequestHeader)
		r.Config.AccessTokenScopesClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenScopesClaim {
			r.Config.AccessTokenScopesClaim = append(r.Config.AccessTokenScopesClaim, types.StringValue(v))
		}
		r.Config.AccessTokenScopesRequired = []types.String{}
		for _, v := range resp.Config.AccessTokenScopesRequired {
			r.Config.AccessTokenScopesRequired = append(r.Config.AccessTokenScopesRequired, types.StringValue(v))
		}
		if resp.Config.AccessTokenSigningAlgorithm != nil {
			r.Config.AccessTokenSigningAlgorithm = types.StringValue(string(*resp.Config.AccessTokenSigningAlgorithm))
		} else {
			r.Config.AccessTokenSigningAlgorithm = types.StringNull()
		}
		r.Config.AccessTokenUpstreamHeader = types.StringPointerValue(resp.Config.AccessTokenUpstreamHeader)
		if resp.Config.AccessTokenUpstreamLeeway != nil {
			r.Config.AccessTokenUpstreamLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenUpstreamLeeway)))
		} else {
			r.Config.AccessTokenUpstreamLeeway = types.NumberNull()
		}
		if len(resp.Config.AddAccessTokenClaims) > 0 {
			r.Config.AddAccessTokenClaims = make(map[string]types.String)
			for key, value := range resp.Config.AddAccessTokenClaims {
				result, _ := json.Marshal(value)
				r.Config.AddAccessTokenClaims[key] = types.StringValue(string(result))
			}
		}
		if len(resp.Config.AddChannelTokenClaims) > 0 {
			r.Config.AddChannelTokenClaims = make(map[string]types.String)
			for key1, value1 := range resp.Config.AddChannelTokenClaims {
				result1, _ := json.Marshal(value1)
				r.Config.AddChannelTokenClaims[key1] = types.StringValue(string(result1))
			}
		}
		if len(resp.Config.AddClaims) > 0 {
			r.Config.AddClaims = make(map[string]types.String)
			for key2, value2 := range resp.Config.AddClaims {
				result2, _ := json.Marshal(value2)
				r.Config.AddClaims[key2] = types.StringValue(string(result2))
			}
		}
		r.Config.CacheAccessTokenIntrospection = types.BoolPointerValue(resp.Config.CacheAccessTokenIntrospection)
		r.Config.CacheChannelTokenIntrospection = types.BoolPointerValue(resp.Config.CacheChannelTokenIntrospection)
		r.Config.ChannelTokenConsumerBy = []types.String{}
		for _, v := range resp.Config.ChannelTokenConsumerBy {
			r.Config.ChannelTokenConsumerBy = append(r.Config.ChannelTokenConsumerBy, types.StringValue(string(v)))
		}
		r.Config.ChannelTokenConsumerClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenConsumerClaim {
			r.Config.ChannelTokenConsumerClaim = append(r.Config.ChannelTokenConsumerClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionAuthorization = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionAuthorization)
		r.Config.ChannelTokenIntrospectionBodyArgs = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionBodyArgs)
		r.Config.ChannelTokenIntrospectionConsumerBy = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionConsumerBy {
			r.Config.ChannelTokenIntrospectionConsumerBy = append(r.Config.ChannelTokenIntrospectionConsumerBy, types.StringValue(string(v)))
		}
		r.Config.ChannelTokenIntrospectionConsumerClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionConsumerClaim {
			r.Config.ChannelTokenIntrospectionConsumerClaim = append(r.Config.ChannelTokenIntrospectionConsumerClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionEndpoint = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionEndpoint)
		r.Config.ChannelTokenIntrospectionHint = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionHint)
		r.Config.ChannelTokenIntrospectionJwtClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionJwtClaim {
			r.Config.ChannelTokenIntrospectionJwtClaim = append(r.Config.ChannelTokenIntrospectionJwtClaim, types.StringValue(v))
		}
		if resp.Config.ChannelTokenIntrospectionLeeway != nil {
			r.Config.ChannelTokenIntrospectionLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenIntrospectionLeeway)))
		} else {
			r.Config.ChannelTokenIntrospectionLeeway = types.NumberNull()
		}
		r.Config.ChannelTokenIntrospectionScopesClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionScopesClaim {
			r.Config.ChannelTokenIntrospectionScopesClaim = append(r.Config.ChannelTokenIntrospectionScopesClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionScopesRequired = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionScopesRequired {
			r.Config.ChannelTokenIntrospectionScopesRequired = append(r.Config.ChannelTokenIntrospectionScopesRequired, types.StringValue(v))
		}
		if resp.Config.ChannelTokenIntrospectionTimeout != nil {
			r.Config.ChannelTokenIntrospectionTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenIntrospectionTimeout)))
		} else {
			r.Config.ChannelTokenIntrospectionTimeout = types.NumberNull()
		}
		r.Config.ChannelTokenIssuer = types.StringPointerValue(resp.Config.ChannelTokenIssuer)
		r.Config.ChannelTokenJwksURI = types.StringPointerValue(resp.Config.ChannelTokenJwksURI)
		r.Config.ChannelTokenJwksURIClientCertificate = types.StringPointerValue(resp.Config.ChannelTokenJwksURIClientCertificate)
		r.Config.ChannelTokenJwksURIClientPassword = types.StringPointerValue(resp.Config.ChannelTokenJwksURIClientPassword)
		r.Config.ChannelTokenJwksURIClientUsername = types.StringPointerValue(resp.Config.ChannelTokenJwksURIClientUsername)
		if resp.Config.ChannelTokenJwksURIRotatePeriod != nil {
			r.Config.ChannelTokenJwksURIRotatePeriod = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenJwksURIRotatePeriod)))
		} else {
			r.Config.ChannelTokenJwksURIRotatePeriod = types.NumberNull()
		}
		r.Config.ChannelTokenKeyset = types.StringPointerValue(resp.Config.ChannelTokenKeyset)
		r.Config.ChannelTokenKeysetClientCertificate = types.StringPointerValue(resp.Config.ChannelTokenKeysetClientCertificate)
		r.Config.ChannelTokenKeysetClientPassword = types.StringPointerValue(resp.Config.ChannelTokenKeysetClientPassword)
		r.Config.ChannelTokenKeysetClientUsername = types.StringPointerValue(resp.Config.ChannelTokenKeysetClientUsername)
		if resp.Config.ChannelTokenKeysetRotatePeriod != nil {
			r.Config.ChannelTokenKeysetRotatePeriod = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenKeysetRotatePeriod)))
		} else {
			r.Config.ChannelTokenKeysetRotatePeriod = types.NumberNull()
		}
		if resp.Config.ChannelTokenLeeway != nil {
			r.Config.ChannelTokenLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenLeeway)))
		} else {
			r.Config.ChannelTokenLeeway = types.NumberNull()
		}
		r.Config.ChannelTokenOptional = types.BoolPointerValue(resp.Config.ChannelTokenOptional)
		r.Config.ChannelTokenRequestHeader = types.StringPointerValue(resp.Config.ChannelTokenRequestHeader)
		r.Config.ChannelTokenScopesClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenScopesClaim {
			r.Config.ChannelTokenScopesClaim = append(r.Config.ChannelTokenScopesClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenScopesRequired = []types.String{}
		for _, v := range resp.Config.ChannelTokenScopesRequired {
			r.Config.ChannelTokenScopesRequired = append(r.Config.ChannelTokenScopesRequired, types.StringValue(v))
		}
		if resp.Config.ChannelTokenSigningAlgorithm != nil {
			r.Config.ChannelTokenSigningAlgorithm = types.StringValue(string(*resp.Config.ChannelTokenSigningAlgorithm))
		} else {
			r.Config.ChannelTokenSigningAlgorithm = types.StringNull()
		}
		r.Config.ChannelTokenUpstreamHeader = types.StringPointerValue(resp.Config.ChannelTokenUpstreamHeader)
		if resp.Config.ChannelTokenUpstreamLeeway != nil {
			r.Config.ChannelTokenUpstreamLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenUpstreamLeeway)))
		} else {
			r.Config.ChannelTokenUpstreamLeeway = types.NumberNull()
		}
		r.Config.EnableAccessTokenIntrospection = types.BoolPointerValue(resp.Config.EnableAccessTokenIntrospection)
		r.Config.EnableChannelTokenIntrospection = types.BoolPointerValue(resp.Config.EnableChannelTokenIntrospection)
		r.Config.EnableHsSignatures = types.BoolPointerValue(resp.Config.EnableHsSignatures)
		r.Config.EnableInstrumentation = types.BoolPointerValue(resp.Config.EnableInstrumentation)
		r.Config.OriginalAccessTokenUpstreamHeader = types.StringPointerValue(resp.Config.OriginalAccessTokenUpstreamHeader)
		r.Config.OriginalChannelTokenUpstreamHeader = types.StringPointerValue(resp.Config.OriginalChannelTokenUpstreamHeader)
		r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
		r.Config.RemoveAccessTokenClaims = []types.String{}
		for _, v := range resp.Config.RemoveAccessTokenClaims {
			r.Config.RemoveAccessTokenClaims = append(r.Config.RemoveAccessTokenClaims, types.StringValue(v))
		}
		r.Config.RemoveChannelTokenClaims = []types.String{}
		for _, v := range resp.Config.RemoveChannelTokenClaims {
			r.Config.RemoveChannelTokenClaims = append(r.Config.RemoveChannelTokenClaims, types.StringValue(v))
		}
		if len(resp.Config.SetAccessTokenClaims) > 0 {
			r.Config.SetAccessTokenClaims = make(map[string]types.String)
			for key3, value3 := range resp.Config.SetAccessTokenClaims {
				result3, _ := json.Marshal(value3)
				r.Config.SetAccessTokenClaims[key3] = types.StringValue(string(result3))
			}
		}
		if len(resp.Config.SetChannelTokenClaims) > 0 {
			r.Config.SetChannelTokenClaims = make(map[string]types.String)
			for key4, value4 := range resp.Config.SetChannelTokenClaims {
				result4, _ := json.Marshal(value4)
				r.Config.SetChannelTokenClaims[key4] = types.StringValue(string(result4))
			}
		}
		if len(resp.Config.SetClaims) > 0 {
			r.Config.SetClaims = make(map[string]types.String)
			for key5, value5 := range resp.Config.SetClaims {
				result5, _ := json.Marshal(value5)
				r.Config.SetClaims[key5] = types.StringValue(string(result5))
			}
		}
		r.Config.TrustAccessTokenIntrospection = types.BoolPointerValue(resp.Config.TrustAccessTokenIntrospection)
		r.Config.TrustChannelTokenIntrospection = types.BoolPointerValue(resp.Config.TrustChannelTokenIntrospection)
		r.Config.VerifyAccessTokenExpiry = types.BoolPointerValue(resp.Config.VerifyAccessTokenExpiry)
		r.Config.VerifyAccessTokenIntrospectionExpiry = types.BoolPointerValue(resp.Config.VerifyAccessTokenIntrospectionExpiry)
		r.Config.VerifyAccessTokenIntrospectionScopes = types.BoolPointerValue(resp.Config.VerifyAccessTokenIntrospectionScopes)
		r.Config.VerifyAccessTokenScopes = types.BoolPointerValue(resp.Config.VerifyAccessTokenScopes)
		r.Config.VerifyAccessTokenSignature = types.BoolPointerValue(resp.Config.VerifyAccessTokenSignature)
		r.Config.VerifyChannelTokenExpiry = types.BoolPointerValue(resp.Config.VerifyChannelTokenExpiry)
		r.Config.VerifyChannelTokenIntrospectionExpiry = types.BoolPointerValue(resp.Config.VerifyChannelTokenIntrospectionExpiry)
		r.Config.VerifyChannelTokenIntrospectionScopes = types.BoolPointerValue(resp.Config.VerifyChannelTokenIntrospectionScopes)
		r.Config.VerifyChannelTokenScopes = types.BoolPointerValue(resp.Config.VerifyChannelTokenScopes)
		r.Config.VerifyChannelTokenSignature = types.BoolPointerValue(resp.Config.VerifyChannelTokenSignature)
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
