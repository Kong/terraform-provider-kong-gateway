// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetServiceRequest struct {
	// ID or name of the Service to lookup
	ServiceIDOrName string `pathParam:"style=simple,explode=false,name=ServiceIdOrName"`
}

func (o *GetServiceRequest) GetServiceIDOrName() string {
	if o == nil {
		return ""
	}
	return o.ServiceIDOrName
}

type GetServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Service
	Service *shared.Service
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServiceResponse) GetService() *shared.Service {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *GetServiceResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
