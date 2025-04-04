// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginSessionDataSourceModel) RefreshFromSharedSessionPlugin(resp *shared.SessionPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.SessionPluginConfig{}
			if resp.Config.AbsoluteTimeout != nil {
				r.Config.AbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.AbsoluteTimeout)))
			} else {
				r.Config.AbsoluteTimeout = types.NumberNull()
			}
			r.Config.Audience = types.StringPointerValue(resp.Config.Audience)
			r.Config.CookieDomain = types.StringPointerValue(resp.Config.CookieDomain)
			r.Config.CookieHTTPOnly = types.BoolPointerValue(resp.Config.CookieHTTPOnly)
			r.Config.CookieName = types.StringPointerValue(resp.Config.CookieName)
			r.Config.CookiePath = types.StringPointerValue(resp.Config.CookiePath)
			if resp.Config.CookieSameSite != nil {
				r.Config.CookieSameSite = types.StringValue(string(*resp.Config.CookieSameSite))
			} else {
				r.Config.CookieSameSite = types.StringNull()
			}
			r.Config.CookieSecure = types.BoolPointerValue(resp.Config.CookieSecure)
			if resp.Config.IdlingTimeout != nil {
				r.Config.IdlingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.IdlingTimeout)))
			} else {
				r.Config.IdlingTimeout = types.NumberNull()
			}
			r.Config.LogoutMethods = make([]types.String, 0, len(resp.Config.LogoutMethods))
			for _, v := range resp.Config.LogoutMethods {
				r.Config.LogoutMethods = append(r.Config.LogoutMethods, types.StringValue(string(v)))
			}
			r.Config.LogoutPostArg = types.StringPointerValue(resp.Config.LogoutPostArg)
			r.Config.LogoutQueryArg = types.StringPointerValue(resp.Config.LogoutQueryArg)
			r.Config.ReadBodyForLogout = types.BoolPointerValue(resp.Config.ReadBodyForLogout)
			r.Config.Remember = types.BoolPointerValue(resp.Config.Remember)
			if resp.Config.RememberAbsoluteTimeout != nil {
				r.Config.RememberAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RememberAbsoluteTimeout)))
			} else {
				r.Config.RememberAbsoluteTimeout = types.NumberNull()
			}
			r.Config.RememberCookieName = types.StringPointerValue(resp.Config.RememberCookieName)
			if resp.Config.RememberRollingTimeout != nil {
				r.Config.RememberRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RememberRollingTimeout)))
			} else {
				r.Config.RememberRollingTimeout = types.NumberNull()
			}
			r.Config.RequestHeaders = make([]types.String, 0, len(resp.Config.RequestHeaders))
			for _, v := range resp.Config.RequestHeaders {
				r.Config.RequestHeaders = append(r.Config.RequestHeaders, types.StringValue(string(v)))
			}
			r.Config.ResponseHeaders = make([]types.String, 0, len(resp.Config.ResponseHeaders))
			for _, v := range resp.Config.ResponseHeaders {
				r.Config.ResponseHeaders = append(r.Config.ResponseHeaders, types.StringValue(string(v)))
			}
			if resp.Config.RollingTimeout != nil {
				r.Config.RollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RollingTimeout)))
			} else {
				r.Config.RollingTimeout = types.NumberNull()
			}
			r.Config.Secret = types.StringPointerValue(resp.Config.Secret)
			if resp.Config.StaleTTL != nil {
				r.Config.StaleTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.StaleTTL)))
			} else {
				r.Config.StaleTTL = types.NumberNull()
			}
			if resp.Config.Storage != nil {
				r.Config.Storage = types.StringValue(string(*resp.Config.Storage))
			} else {
				r.Config.Storage = types.StringNull()
			}
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
