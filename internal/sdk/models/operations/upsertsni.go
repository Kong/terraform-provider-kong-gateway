// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpsertSniRequest struct {
	// ID or name of the SNI to lookup
	SNIIDOrName string `pathParam:"style=simple,explode=false,name=SNIIdOrName"`
	// Description of the SNI
	Sni shared.SNIInput `request:"mediaType=application/json"`
}

func (o *UpsertSniRequest) GetSNIIDOrName() string {
	if o == nil {
		return ""
	}
	return o.SNIIDOrName
}

func (o *UpsertSniRequest) GetSni() shared.SNIInput {
	if o == nil {
		return shared.SNIInput{}
	}
	return o.Sni
}

type UpsertSniResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted SNI
	Sni *shared.Sni
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertSniResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertSniResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertSniResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertSniResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *UpsertSniResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}