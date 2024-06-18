// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FetchPluginSchemaRequest struct {
	// The name of the plugin
	PluginName string `pathParam:"style=simple,explode=false,name=pluginName"`
}

func (o *FetchPluginSchemaRequest) GetPluginName() string {
	if o == nil {
		return ""
	}
	return o.PluginName
}

// FetchPluginSchemaResponseBody - The schema for the plugin
type FetchPluginSchemaResponseBody struct {
	Fields []map[string]any `json:"fields,omitempty"`
}

func (o *FetchPluginSchemaResponseBody) GetFields() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Fields
}

type FetchPluginSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The schema for the plugin
	Object *FetchPluginSchemaResponseBody
}

func (o *FetchPluginSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchPluginSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchPluginSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *FetchPluginSchemaResponse) GetObject() *FetchPluginSchemaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
