// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetGrpcwebPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
}

func (o *GetGrpcwebPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

type GetGrpcwebPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// GrpcWeb plugin
	GrpcWebPlugin *shared.GrpcWebPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetGrpcwebPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGrpcwebPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGrpcwebPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGrpcwebPluginResponse) GetGrpcWebPlugin() *shared.GrpcWebPlugin {
	if o == nil {
		return nil
	}
	return o.GrpcWebPlugin
}

func (o *GetGrpcwebPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}