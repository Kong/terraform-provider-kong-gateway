// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateJqPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string          `pathParam:"style=simple,explode=false,name=PluginId"`
	JqPlugin shared.JqPlugin `request:"mediaType=application/json"`
}

func (o *UpdateJqPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateJqPluginRequest) GetJqPlugin() shared.JqPlugin {
	if o == nil {
		return shared.JqPlugin{}
	}
	return o.JqPlugin
}

type UpdateJqPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Jq plugin
	JqPlugin *shared.JqPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateJqPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJqPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJqPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJqPluginResponse) GetJqPlugin() *shared.JqPlugin {
	if o == nil {
		return nil
	}
	return o.JqPlugin
}

func (o *UpdateJqPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
