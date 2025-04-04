// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetRouteRequest struct {
	// ID or name of the Route to lookup
	RouteIDOrName string `pathParam:"style=simple,explode=false,name=RouteIdOrName"`
}

func (o *GetRouteRequest) GetRouteIDOrName() string {
	if o == nil {
		return ""
	}
	return o.RouteIDOrName
}

type GetRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Route
	RouteJSON *shared.RouteJSON
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRouteResponse) GetRouteJSON() *shared.RouteJSON {
	if o == nil {
		return nil
	}
	return o.RouteJSON
}

func (o *GetRouteResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
