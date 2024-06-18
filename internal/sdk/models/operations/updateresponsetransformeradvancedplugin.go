// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateResponsetransformeradvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                                string                                          `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateResponseTransformerAdvancedPlugin *shared.CreateResponseTransformerAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateResponsetransformeradvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateResponsetransformeradvancedPluginRequest) GetCreateResponseTransformerAdvancedPlugin() *shared.CreateResponseTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateResponseTransformerAdvancedPlugin
}

type UpdateResponsetransformeradvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ResponseTransformerAdvanced plugin
	ResponseTransformerAdvancedPlugin *shared.ResponseTransformerAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetResponseTransformerAdvancedPlugin() *shared.ResponseTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerAdvancedPlugin
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
