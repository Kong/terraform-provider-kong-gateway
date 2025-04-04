// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateJwtsignerPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID        string                 `pathParam:"style=simple,explode=false,name=PluginId"`
	JwtSignerPlugin shared.JwtSignerPlugin `request:"mediaType=application/json"`
}

func (o *UpdateJwtsignerPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateJwtsignerPluginRequest) GetJwtSignerPlugin() shared.JwtSignerPlugin {
	if o == nil {
		return shared.JwtSignerPlugin{}
	}
	return o.JwtSignerPlugin
}

type UpdateJwtsignerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// JwtSigner plugin
	JwtSignerPlugin *shared.JwtSignerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateJwtsignerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJwtsignerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJwtsignerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJwtsignerPluginResponse) GetJwtSignerPlugin() *shared.JwtSignerPlugin {
	if o == nil {
		return nil
	}
	return o.JwtSignerPlugin
}

func (o *UpdateJwtsignerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
