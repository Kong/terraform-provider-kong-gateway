// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateACLPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID  string           `pathParam:"style=simple,explode=false,name=PluginId"`
	ACLPlugin shared.ACLPlugin `request:"mediaType=application/json"`
}

func (o *UpdateACLPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateACLPluginRequest) GetACLPlugin() shared.ACLPlugin {
	if o == nil {
		return shared.ACLPlugin{}
	}
	return o.ACLPlugin
}

type UpdateACLPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ACL plugin
	ACLPlugin *shared.ACLPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateACLPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateACLPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateACLPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateACLPluginResponse) GetACLPlugin() *shared.ACLPlugin {
	if o == nil {
		return nil
	}
	return o.ACLPlugin
}

func (o *UpdateACLPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
