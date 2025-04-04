// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteHmacAuthWithConsumerRequest struct {
	// Consumer ID for nested entities
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the HMAC-auth credential to lookup
	HMACAuthID string `pathParam:"style=simple,explode=false,name=HMACAuthId"`
}

func (o *DeleteHmacAuthWithConsumerRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *DeleteHmacAuthWithConsumerRequest) GetHMACAuthID() string {
	if o == nil {
		return ""
	}
	return o.HMACAuthID
}

type DeleteHmacAuthWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteHmacAuthWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteHmacAuthWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteHmacAuthWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
