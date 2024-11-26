// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateResponsetransformerPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                  string                                 `pathParam:"style=simple,explode=false,name=PluginId"`
	ResponseTransformerPlugin *shared.ResponseTransformerPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateResponsetransformerPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateResponsetransformerPluginRequest) GetResponseTransformerPlugin() *shared.ResponseTransformerPluginInput {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerPlugin
}

type UpdateResponsetransformerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ResponseTransformer plugin
	ResponseTransformerPlugin *shared.ResponseTransformerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateResponsetransformerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResponsetransformerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResponsetransformerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResponsetransformerPluginResponse) GetResponseTransformerPlugin() *shared.ResponseTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerPlugin
}

func (o *UpdateResponsetransformerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
