// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetVaultauthPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
}

func (o *GetVaultauthPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

type GetVaultauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// VaultAuth plugin
	VaultAuthPlugin *shared.VaultAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetVaultauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetVaultauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetVaultauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetVaultauthPluginResponse) GetVaultAuthPlugin() *shared.VaultAuthPlugin {
	if o == nil {
		return nil
	}
	return o.VaultAuthPlugin
}

func (o *GetVaultauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
