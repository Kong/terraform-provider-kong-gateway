// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateRequestcalloutPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID             string                      `pathParam:"style=simple,explode=false,name=PluginId"`
	RequestCalloutPlugin shared.RequestCalloutPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRequestcalloutPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRequestcalloutPluginRequest) GetRequestCalloutPlugin() shared.RequestCalloutPlugin {
	if o == nil {
		return shared.RequestCalloutPlugin{}
	}
	return o.RequestCalloutPlugin
}

type UpdateRequestcalloutPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RequestCallout plugin
	RequestCalloutPlugin *shared.RequestCalloutPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRequestcalloutPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRequestcalloutPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRequestcalloutPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRequestcalloutPluginResponse) GetRequestCalloutPlugin() *shared.RequestCalloutPlugin {
	if o == nil {
		return nil
	}
	return o.RequestCalloutPlugin
}

func (o *UpdateRequestcalloutPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
