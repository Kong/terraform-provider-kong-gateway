// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTargetWithUpstreamRequest struct {
	// ID or target of the Target to lookup
	UpstreamIDForTarget string `pathParam:"style=simple,explode=false,name=UpstreamIdForTarget"`
	// ID or target of the Target to lookup
	TargetIDOrTarget string `pathParam:"style=simple,explode=false,name=TargetIdOrTarget"`
}

func (o *DeleteTargetWithUpstreamRequest) GetUpstreamIDForTarget() string {
	if o == nil {
		return ""
	}
	return o.UpstreamIDForTarget
}

func (o *DeleteTargetWithUpstreamRequest) GetTargetIDOrTarget() string {
	if o == nil {
		return ""
	}
	return o.TargetIDOrTarget
}

type DeleteTargetWithUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteTargetWithUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTargetWithUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTargetWithUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
