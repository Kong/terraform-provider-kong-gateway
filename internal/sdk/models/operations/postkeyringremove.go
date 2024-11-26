// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

// PostKeyringRemoveRequestBody - The request body contains a key ID that can be generated from the `/keyring/generate` endpoint.
type PostKeyringRemoveRequestBody struct {
	// The key ID
	Key *string `json:"key,omitempty"`
}

func (o *PostKeyringRemoveRequestBody) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

type PostKeyringRemoveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *PostKeyringRemoveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostKeyringRemoveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostKeyringRemoveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostKeyringRemoveResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}