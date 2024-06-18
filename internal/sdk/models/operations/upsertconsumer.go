// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpsertConsumerRequest struct {
	// ID or username of the Consumer to lookup
	ConsumerIDOrUsername string `pathParam:"style=simple,explode=false,name=ConsumerIdOrUsername"`
	// Description of the Consumer
	Consumer shared.ConsumerInput `request:"mediaType=application/json"`
}

func (o *UpsertConsumerRequest) GetConsumerIDOrUsername() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDOrUsername
}

func (o *UpsertConsumerRequest) GetConsumer() shared.ConsumerInput {
	if o == nil {
		return shared.ConsumerInput{}
	}
	return o.Consumer
}

type UpsertConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Consumer
	Consumer *shared.Consumer
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertConsumerResponse) GetConsumer() *shared.Consumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *UpsertConsumerResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}