// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetMtlsAuthRequest struct {
	// ID of the MTLS-auth credential to lookup
	MTLSAuthID string `pathParam:"style=simple,explode=false,name=MTLSAuthId"`
}

func (o *GetMtlsAuthRequest) GetMTLSAuthID() string {
	if o == nil {
		return ""
	}
	return o.MTLSAuthID
}

type GetMtlsAuthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched MTLS-auth credential
	MTLSAuth *shared.MTLSAuth
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetMtlsAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMtlsAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMtlsAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMtlsAuthResponse) GetMTLSAuth() *shared.MTLSAuth {
	if o == nil {
		return nil
	}
	return o.MTLSAuth
}

func (o *GetMtlsAuthResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
