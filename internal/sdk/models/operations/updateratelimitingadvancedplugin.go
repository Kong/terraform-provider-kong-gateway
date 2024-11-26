// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateRatelimitingadvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                   string                                  `pathParam:"style=simple,explode=false,name=PluginId"`
	RateLimitingAdvancedPlugin *shared.RateLimitingAdvancedPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateRatelimitingadvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRatelimitingadvancedPluginRequest) GetRateLimitingAdvancedPlugin() *shared.RateLimitingAdvancedPluginInput {
	if o == nil {
		return nil
	}
	return o.RateLimitingAdvancedPlugin
}

type UpdateRatelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RateLimitingAdvanced plugin
	RateLimitingAdvancedPlugin *shared.RateLimitingAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetRateLimitingAdvancedPlugin() *shared.RateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.RateLimitingAdvancedPlugin
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
