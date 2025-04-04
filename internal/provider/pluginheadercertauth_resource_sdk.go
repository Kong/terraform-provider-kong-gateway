// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginHeaderCertAuthResourceModel) ToSharedHeaderCertAuthPlugin() *shared.HeaderCertAuthPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.HeaderCertAuthPluginOrdering
	if r.Ordering != nil {
		var after *shared.HeaderCertAuthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.HeaderCertAuthPluginAfter{
				Access: access,
			}
		}
		var before *shared.HeaderCertAuthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.HeaderCertAuthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.HeaderCertAuthPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.HeaderCertAuthPluginPartials = []shared.HeaderCertAuthPluginPartials{}
	for _, partialsItem := range r.Partials {
		id1 := new(string)
		if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
			*id1 = partialsItem.ID.ValueString()
		} else {
			id1 = nil
		}
		name := new(string)
		if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
			*name = partialsItem.Name.ValueString()
		} else {
			name = nil
		}
		path := new(string)
		if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
			*path = partialsItem.Path.ValueString()
		} else {
			path = nil
		}
		partials = append(partials, shared.HeaderCertAuthPluginPartials{
			ID:   id1,
			Name: name,
			Path: path,
		})
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.HeaderCertAuthPluginConfig
	if r.Config != nil {
		allowPartialChain := new(bool)
		if !r.Config.AllowPartialChain.IsUnknown() && !r.Config.AllowPartialChain.IsNull() {
			*allowPartialChain = r.Config.AllowPartialChain.ValueBool()
		} else {
			allowPartialChain = nil
		}
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		authenticatedGroupBy := new(shared.AuthenticatedGroupBy)
		if !r.Config.AuthenticatedGroupBy.IsUnknown() && !r.Config.AuthenticatedGroupBy.IsNull() {
			*authenticatedGroupBy = shared.AuthenticatedGroupBy(r.Config.AuthenticatedGroupBy.ValueString())
		} else {
			authenticatedGroupBy = nil
		}
		var caCertificates []string = []string{}
		for _, caCertificatesItem := range r.Config.CaCertificates {
			caCertificates = append(caCertificates, caCertificatesItem.ValueString())
		}
		cacheTTL := new(float64)
		if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
			*cacheTTL, _ = r.Config.CacheTTL.ValueBigFloat().Float64()
		} else {
			cacheTTL = nil
		}
		certCacheTTL := new(float64)
		if !r.Config.CertCacheTTL.IsUnknown() && !r.Config.CertCacheTTL.IsNull() {
			*certCacheTTL, _ = r.Config.CertCacheTTL.ValueBigFloat().Float64()
		} else {
			certCacheTTL = nil
		}
		certificateHeaderFormat := new(shared.CertificateHeaderFormat)
		if !r.Config.CertificateHeaderFormat.IsUnknown() && !r.Config.CertificateHeaderFormat.IsNull() {
			*certificateHeaderFormat = shared.CertificateHeaderFormat(r.Config.CertificateHeaderFormat.ValueString())
		} else {
			certificateHeaderFormat = nil
		}
		certificateHeaderName := new(string)
		if !r.Config.CertificateHeaderName.IsUnknown() && !r.Config.CertificateHeaderName.IsNull() {
			*certificateHeaderName = r.Config.CertificateHeaderName.ValueString()
		} else {
			certificateHeaderName = nil
		}
		var consumerBy []shared.ConsumerBy = []shared.ConsumerBy{}
		for _, consumerByItem := range r.Config.ConsumerBy {
			consumerBy = append(consumerBy, shared.ConsumerBy(consumerByItem.ValueString()))
		}
		defaultConsumer := new(string)
		if !r.Config.DefaultConsumer.IsUnknown() && !r.Config.DefaultConsumer.IsNull() {
			*defaultConsumer = r.Config.DefaultConsumer.ValueString()
		} else {
			defaultConsumer = nil
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
		httpTimeout := new(float64)
		if !r.Config.HTTPTimeout.IsUnknown() && !r.Config.HTTPTimeout.IsNull() {
			*httpTimeout, _ = r.Config.HTTPTimeout.ValueBigFloat().Float64()
		} else {
			httpTimeout = nil
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
		revocationCheckMode := new(shared.RevocationCheckMode)
		if !r.Config.RevocationCheckMode.IsUnknown() && !r.Config.RevocationCheckMode.IsNull() {
			*revocationCheckMode = shared.RevocationCheckMode(r.Config.RevocationCheckMode.ValueString())
		} else {
			revocationCheckMode = nil
		}
		secureSource := new(bool)
		if !r.Config.SecureSource.IsUnknown() && !r.Config.SecureSource.IsNull() {
			*secureSource = r.Config.SecureSource.ValueBool()
		} else {
			secureSource = nil
		}
		skipConsumerLookup := new(bool)
		if !r.Config.SkipConsumerLookup.IsUnknown() && !r.Config.SkipConsumerLookup.IsNull() {
			*skipConsumerLookup = r.Config.SkipConsumerLookup.ValueBool()
		} else {
			skipConsumerLookup = nil
		}
		config = &shared.HeaderCertAuthPluginConfig{
			AllowPartialChain:       allowPartialChain,
			Anonymous:               anonymous,
			AuthenticatedGroupBy:    authenticatedGroupBy,
			CaCertificates:          caCertificates,
			CacheTTL:                cacheTTL,
			CertCacheTTL:            certCacheTTL,
			CertificateHeaderFormat: certificateHeaderFormat,
			CertificateHeaderName:   certificateHeaderName,
			ConsumerBy:              consumerBy,
			DefaultConsumer:         defaultConsumer,
			HTTPProxyHost:           httpProxyHost,
			HTTPProxyPort:           httpProxyPort,
			HTTPTimeout:             httpTimeout,
			HTTPSProxyHost:          httpsProxyHost,
			HTTPSProxyPort:          httpsProxyPort,
			RevocationCheckMode:     revocationCheckMode,
			SecureSource:            secureSource,
			SkipConsumerLookup:      skipConsumerLookup,
		}
	}
	var protocols []shared.HeaderCertAuthPluginProtocols = []shared.HeaderCertAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.HeaderCertAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.HeaderCertAuthPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.HeaderCertAuthPluginRoute{
			ID: id2,
		}
	}
	var service *shared.HeaderCertAuthPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.HeaderCertAuthPluginService{
			ID: id3,
		}
	}
	out := shared.HeaderCertAuthPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Partials:     partials,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *PluginHeaderCertAuthResourceModel) RefreshFromSharedHeaderCertAuthPlugin(resp *shared.HeaderCertAuthPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.HeaderCertAuthPluginConfig{}
			r.Config.AllowPartialChain = types.BoolPointerValue(resp.Config.AllowPartialChain)
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			if resp.Config.AuthenticatedGroupBy != nil {
				r.Config.AuthenticatedGroupBy = types.StringValue(string(*resp.Config.AuthenticatedGroupBy))
			} else {
				r.Config.AuthenticatedGroupBy = types.StringNull()
			}
			r.Config.CaCertificates = make([]types.String, 0, len(resp.Config.CaCertificates))
			for _, v := range resp.Config.CaCertificates {
				r.Config.CaCertificates = append(r.Config.CaCertificates, types.StringValue(v))
			}
			if resp.Config.CacheTTL != nil {
				r.Config.CacheTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.CacheTTL)))
			} else {
				r.Config.CacheTTL = types.NumberNull()
			}
			if resp.Config.CertCacheTTL != nil {
				r.Config.CertCacheTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.CertCacheTTL)))
			} else {
				r.Config.CertCacheTTL = types.NumberNull()
			}
			if resp.Config.CertificateHeaderFormat != nil {
				r.Config.CertificateHeaderFormat = types.StringValue(string(*resp.Config.CertificateHeaderFormat))
			} else {
				r.Config.CertificateHeaderFormat = types.StringNull()
			}
			r.Config.CertificateHeaderName = types.StringPointerValue(resp.Config.CertificateHeaderName)
			r.Config.ConsumerBy = make([]types.String, 0, len(resp.Config.ConsumerBy))
			for _, v := range resp.Config.ConsumerBy {
				r.Config.ConsumerBy = append(r.Config.ConsumerBy, types.StringValue(string(v)))
			}
			r.Config.DefaultConsumer = types.StringPointerValue(resp.Config.DefaultConsumer)
			r.Config.HTTPProxyHost = types.StringPointerValue(resp.Config.HTTPProxyHost)
			r.Config.HTTPProxyPort = types.Int64PointerValue(resp.Config.HTTPProxyPort)
			if resp.Config.HTTPTimeout != nil {
				r.Config.HTTPTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.HTTPTimeout)))
			} else {
				r.Config.HTTPTimeout = types.NumberNull()
			}
			r.Config.HTTPSProxyHost = types.StringPointerValue(resp.Config.HTTPSProxyHost)
			r.Config.HTTPSProxyPort = types.Int64PointerValue(resp.Config.HTTPSProxyPort)
			if resp.Config.RevocationCheckMode != nil {
				r.Config.RevocationCheckMode = types.StringValue(string(*resp.Config.RevocationCheckMode))
			} else {
				r.Config.RevocationCheckMode = types.StringNull()
			}
			r.Config.SecureSource = types.BoolPointerValue(resp.Config.SecureSource)
			r.Config.SkipConsumerLookup = types.BoolPointerValue(resp.Config.SkipConsumerLookup)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.Ordering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.After{}
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.After{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials1 tfTypes.Partials
				partials1.ID = types.StringPointerValue(partialsItem.ID)
				partials1.Name = types.StringPointerValue(partialsItem.Name)
				partials1.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials1)
				} else {
					r.Partials[partialsCount].ID = partials1.ID
					r.Partials[partialsCount].Name = partials1.Name
					r.Partials[partialsCount].Path = partials1.Path
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
