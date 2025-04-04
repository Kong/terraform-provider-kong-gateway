// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type CreateHmacAuthWithConsumerRequest struct {
	// Consumer ID for nested entities
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// Description of new HMAC-auth credential for creation
	HMACAuthWithoutParents shared.HMACAuthWithoutParents `request:"mediaType=application/json"`
}

func (o *CreateHmacAuthWithConsumerRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *CreateHmacAuthWithConsumerRequest) GetHMACAuthWithoutParents() shared.HMACAuthWithoutParents {
	if o == nil {
		return shared.HMACAuthWithoutParents{}
	}
	return o.HMACAuthWithoutParents
}

type CreateHmacAuthWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created HMAC-auth credential
	HMACAuth *shared.HMACAuth
}

func (o *CreateHmacAuthWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateHmacAuthWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateHmacAuthWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateHmacAuthWithConsumerResponse) GetHMACAuth() *shared.HMACAuth {
	if o == nil {
		return nil
	}
	return o.HMACAuth
}
