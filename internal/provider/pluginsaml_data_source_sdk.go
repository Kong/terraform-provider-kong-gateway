// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginSamlDataSourceModel) RefreshFromSharedSamlPlugin(resp *shared.SamlPlugin) {
	if resp != nil {
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AssertionConsumerPath = types.StringPointerValue(resp.Config.AssertionConsumerPath)
		r.Config.IdpCertificate = types.StringPointerValue(resp.Config.IdpCertificate)
		r.Config.IdpSsoURL = types.StringPointerValue(resp.Config.IdpSsoURL)
		r.Config.Issuer = types.StringPointerValue(resp.Config.Issuer)
		if resp.Config.NameidFormat != nil {
			r.Config.NameidFormat = types.StringValue(string(*resp.Config.NameidFormat))
		} else {
			r.Config.NameidFormat = types.StringNull()
		}
		if resp.Config.Redis == nil {
			r.Config.Redis = nil
		} else {
			r.Config.Redis = &tfTypes.KonnectApplicationAuthPluginRedis{}
			r.Config.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Redis.ClusterMaxRedirections)
			r.Config.Redis.ClusterNodes = []tfTypes.ClusterNodes{}
			if len(r.Config.Redis.ClusterNodes) > len(resp.Config.Redis.ClusterNodes) {
				r.Config.Redis.ClusterNodes = r.Config.Redis.ClusterNodes[:len(resp.Config.Redis.ClusterNodes)]
			}
			for clusterNodesCount, clusterNodesItem := range resp.Config.Redis.ClusterNodes {
				var clusterNodes1 tfTypes.ClusterNodes
				clusterNodes1.IP = types.StringPointerValue(clusterNodesItem.IP)
				clusterNodes1.Port = types.Int64PointerValue(clusterNodesItem.Port)
				if clusterNodesCount+1 > len(r.Config.Redis.ClusterNodes) {
					r.Config.Redis.ClusterNodes = append(r.Config.Redis.ClusterNodes, clusterNodes1)
				} else {
					r.Config.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes1.IP
					r.Config.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes1.Port
				}
			}
			r.Config.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Redis.ConnectTimeout)
			r.Config.Redis.ConnectionIsProxied = types.BoolPointerValue(resp.Config.Redis.ConnectionIsProxied)
			r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
			r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
			r.Config.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Redis.KeepaliveBacklog)
			r.Config.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Redis.KeepalivePoolSize)
			r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
			r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
			r.Config.Redis.Prefix = types.StringPointerValue(resp.Config.Redis.Prefix)
			r.Config.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Redis.ReadTimeout)
			r.Config.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Redis.SendTimeout)
			r.Config.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Redis.SentinelMaster)
			r.Config.Redis.SentinelNodes = []tfTypes.SentinelNodes{}
			if len(r.Config.Redis.SentinelNodes) > len(resp.Config.Redis.SentinelNodes) {
				r.Config.Redis.SentinelNodes = r.Config.Redis.SentinelNodes[:len(resp.Config.Redis.SentinelNodes)]
			}
			for sentinelNodesCount, sentinelNodesItem := range resp.Config.Redis.SentinelNodes {
				var sentinelNodes1 tfTypes.SentinelNodes
				sentinelNodes1.Host = types.StringPointerValue(sentinelNodesItem.Host)
				sentinelNodes1.Port = types.Int64PointerValue(sentinelNodesItem.Port)
				if sentinelNodesCount+1 > len(r.Config.Redis.SentinelNodes) {
					r.Config.Redis.SentinelNodes = append(r.Config.Redis.SentinelNodes, sentinelNodes1)
				} else {
					r.Config.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes1.Host
					r.Config.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes1.Port
				}
			}
			r.Config.Redis.SentinelPassword = types.StringPointerValue(resp.Config.Redis.SentinelPassword)
			if resp.Config.Redis.SentinelRole != nil {
				r.Config.Redis.SentinelRole = types.StringValue(string(*resp.Config.Redis.SentinelRole))
			} else {
				r.Config.Redis.SentinelRole = types.StringNull()
			}
			r.Config.Redis.SentinelUsername = types.StringPointerValue(resp.Config.Redis.SentinelUsername)
			r.Config.Redis.ServerName = types.StringPointerValue(resp.Config.Redis.ServerName)
			r.Config.Redis.Socket = types.StringPointerValue(resp.Config.Redis.Socket)
			r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
			r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
			r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
		}
		if resp.Config.RequestDigestAlgorithm != nil {
			r.Config.RequestDigestAlgorithm = types.StringValue(string(*resp.Config.RequestDigestAlgorithm))
		} else {
			r.Config.RequestDigestAlgorithm = types.StringNull()
		}
		if resp.Config.RequestSignatureAlgorithm != nil {
			r.Config.RequestSignatureAlgorithm = types.StringValue(string(*resp.Config.RequestSignatureAlgorithm))
		} else {
			r.Config.RequestSignatureAlgorithm = types.StringNull()
		}
		r.Config.RequestSigningCertificate = types.StringPointerValue(resp.Config.RequestSigningCertificate)
		r.Config.RequestSigningKey = types.StringPointerValue(resp.Config.RequestSigningKey)
		if resp.Config.ResponseDigestAlgorithm != nil {
			r.Config.ResponseDigestAlgorithm = types.StringValue(string(*resp.Config.ResponseDigestAlgorithm))
		} else {
			r.Config.ResponseDigestAlgorithm = types.StringNull()
		}
		r.Config.ResponseEncryptionKey = types.StringPointerValue(resp.Config.ResponseEncryptionKey)
		if resp.Config.ResponseSignatureAlgorithm != nil {
			r.Config.ResponseSignatureAlgorithm = types.StringValue(string(*resp.Config.ResponseSignatureAlgorithm))
		} else {
			r.Config.ResponseSignatureAlgorithm = types.StringNull()
		}
		if resp.Config.SessionAbsoluteTimeout != nil {
			r.Config.SessionAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionAbsoluteTimeout)))
		} else {
			r.Config.SessionAbsoluteTimeout = types.NumberNull()
		}
		r.Config.SessionAudience = types.StringPointerValue(resp.Config.SessionAudience)
		r.Config.SessionCookieDomain = types.StringPointerValue(resp.Config.SessionCookieDomain)
		r.Config.SessionCookieHTTPOnly = types.BoolPointerValue(resp.Config.SessionCookieHTTPOnly)
		r.Config.SessionCookieName = types.StringPointerValue(resp.Config.SessionCookieName)
		r.Config.SessionCookiePath = types.StringPointerValue(resp.Config.SessionCookiePath)
		if resp.Config.SessionCookieSameSite != nil {
			r.Config.SessionCookieSameSite = types.StringValue(string(*resp.Config.SessionCookieSameSite))
		} else {
			r.Config.SessionCookieSameSite = types.StringNull()
		}
		r.Config.SessionCookieSecure = types.BoolPointerValue(resp.Config.SessionCookieSecure)
		r.Config.SessionEnforceSameSubject = types.BoolPointerValue(resp.Config.SessionEnforceSameSubject)
		r.Config.SessionHashStorageKey = types.BoolPointerValue(resp.Config.SessionHashStorageKey)
		r.Config.SessionHashSubject = types.BoolPointerValue(resp.Config.SessionHashSubject)
		if resp.Config.SessionIdlingTimeout != nil {
			r.Config.SessionIdlingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionIdlingTimeout)))
		} else {
			r.Config.SessionIdlingTimeout = types.NumberNull()
		}
		r.Config.SessionMemcachedHost = types.StringPointerValue(resp.Config.SessionMemcachedHost)
		r.Config.SessionMemcachedPort = types.Int64PointerValue(resp.Config.SessionMemcachedPort)
		r.Config.SessionMemcachedPrefix = types.StringPointerValue(resp.Config.SessionMemcachedPrefix)
		r.Config.SessionMemcachedSocket = types.StringPointerValue(resp.Config.SessionMemcachedSocket)
		r.Config.SessionRemember = types.BoolPointerValue(resp.Config.SessionRemember)
		if resp.Config.SessionRememberAbsoluteTimeout != nil {
			r.Config.SessionRememberAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberAbsoluteTimeout)))
		} else {
			r.Config.SessionRememberAbsoluteTimeout = types.NumberNull()
		}
		r.Config.SessionRememberCookieName = types.StringPointerValue(resp.Config.SessionRememberCookieName)
		if resp.Config.SessionRememberRollingTimeout != nil {
			r.Config.SessionRememberRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberRollingTimeout)))
		} else {
			r.Config.SessionRememberRollingTimeout = types.NumberNull()
		}
		r.Config.SessionRequestHeaders = []types.String{}
		for _, v := range resp.Config.SessionRequestHeaders {
			r.Config.SessionRequestHeaders = append(r.Config.SessionRequestHeaders, types.StringValue(string(v)))
		}
		r.Config.SessionResponseHeaders = []types.String{}
		for _, v := range resp.Config.SessionResponseHeaders {
			r.Config.SessionResponseHeaders = append(r.Config.SessionResponseHeaders, types.StringValue(string(v)))
		}
		if resp.Config.SessionRollingTimeout != nil {
			r.Config.SessionRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRollingTimeout)))
		} else {
			r.Config.SessionRollingTimeout = types.NumberNull()
		}
		r.Config.SessionSecret = types.StringPointerValue(resp.Config.SessionSecret)
		if resp.Config.SessionStorage != nil {
			r.Config.SessionStorage = types.StringValue(string(*resp.Config.SessionStorage))
		} else {
			r.Config.SessionStorage = types.StringNull()
		}
		r.Config.SessionStoreMetadata = types.BoolPointerValue(resp.Config.SessionStoreMetadata)
		r.Config.ValidateAssertionSignature = types.BoolPointerValue(resp.Config.ValidateAssertionSignature)
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
