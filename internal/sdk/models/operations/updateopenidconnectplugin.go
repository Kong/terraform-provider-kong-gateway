// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateOpenidconnectPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                  string                            `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateOpenidConnectPlugin *shared.CreateOpenidConnectPlugin `request:"mediaType=application/json"`
}

func (o *UpdateOpenidconnectPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateOpenidconnectPluginRequest) GetCreateOpenidConnectPlugin() *shared.CreateOpenidConnectPlugin {
	if o == nil {
		return nil
	}
	return o.CreateOpenidConnectPlugin
}

type UpdateOpenidconnectPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OpenidConnect plugin
	OpenidConnectPlugin *shared.OpenidConnectPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateOpenidconnectPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateOpenidconnectPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateOpenidconnectPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateOpenidconnectPluginResponse) GetOpenidConnectPlugin() *shared.OpenidConnectPlugin {
	if o == nil {
		return nil
	}
	return o.OpenidConnectPlugin
}

func (o *UpdateOpenidconnectPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}