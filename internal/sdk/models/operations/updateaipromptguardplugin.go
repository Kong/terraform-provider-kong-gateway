// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateAipromptguardPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID            string                           `pathParam:"style=simple,explode=false,name=PluginId"`
	AiPromptGuardPlugin *shared.AiPromptGuardPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateAipromptguardPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAipromptguardPluginRequest) GetAiPromptGuardPlugin() *shared.AiPromptGuardPluginInput {
	if o == nil {
		return nil
	}
	return o.AiPromptGuardPlugin
}

type UpdateAipromptguardPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AiPromptGuard plugin
	AiPromptGuardPlugin *shared.AiPromptGuardPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAipromptguardPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAipromptguardPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAipromptguardPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAipromptguardPluginResponse) GetAiPromptGuardPlugin() *shared.AiPromptGuardPlugin {
	if o == nil {
		return nil
	}
	return o.AiPromptGuardPlugin
}

func (o *UpdateAipromptguardPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
