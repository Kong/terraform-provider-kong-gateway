// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateAiprompttemplatePluginRequest struct {
	// ID of the Plugin to lookup
	PluginID               string                              `pathParam:"style=simple,explode=false,name=PluginId"`
	AiPromptTemplatePlugin *shared.AiPromptTemplatePluginInput `request:"mediaType=application/json"`
}

func (o *UpdateAiprompttemplatePluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAiprompttemplatePluginRequest) GetAiPromptTemplatePlugin() *shared.AiPromptTemplatePluginInput {
	if o == nil {
		return nil
	}
	return o.AiPromptTemplatePlugin
}

type UpdateAiprompttemplatePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AiPromptTemplate plugin
	AiPromptTemplatePlugin *shared.AiPromptTemplatePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAiprompttemplatePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAiprompttemplatePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAiprompttemplatePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAiprompttemplatePluginResponse) GetAiPromptTemplatePlugin() *shared.AiPromptTemplatePlugin {
	if o == nil {
		return nil
	}
	return o.AiPromptTemplatePlugin
}

func (o *UpdateAiprompttemplatePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
