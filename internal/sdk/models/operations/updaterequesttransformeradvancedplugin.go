// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateRequesttransformeradvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                               string                                         `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateRequestTransformerAdvancedPlugin *shared.CreateRequestTransformerAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRequesttransformeradvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRequesttransformeradvancedPluginRequest) GetCreateRequestTransformerAdvancedPlugin() *shared.CreateRequestTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateRequestTransformerAdvancedPlugin
}

type UpdateRequesttransformeradvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RequestTransformerAdvanced plugin
	RequestTransformerAdvancedPlugin *shared.RequestTransformerAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRequesttransformeradvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRequesttransformeradvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRequesttransformeradvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRequesttransformeradvancedPluginResponse) GetRequestTransformerAdvancedPlugin() *shared.RequestTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.RequestTransformerAdvancedPlugin
}

func (o *UpdateRequesttransformeradvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
