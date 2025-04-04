// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type RequestSizeLimitingPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *RequestSizeLimitingPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RequestSizeLimitingPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *RequestSizeLimitingPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RequestSizeLimitingPluginOrdering struct {
	After  *RequestSizeLimitingPluginAfter  `json:"after,omitempty"`
	Before *RequestSizeLimitingPluginBefore `json:"before,omitempty"`
}

func (o *RequestSizeLimitingPluginOrdering) GetAfter() *RequestSizeLimitingPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *RequestSizeLimitingPluginOrdering) GetBefore() *RequestSizeLimitingPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type RequestSizeLimitingPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *RequestSizeLimitingPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestSizeLimitingPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RequestSizeLimitingPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// SizeUnit - Size unit can be set either in `bytes`, `kilobytes`, or `megabytes` (default). This configuration is not available in versions prior to Kong Gateway 1.3 and Kong Gateway (OSS) 2.0.
type SizeUnit string

const (
	SizeUnitBytes     SizeUnit = "bytes"
	SizeUnitKilobytes SizeUnit = "kilobytes"
	SizeUnitMegabytes SizeUnit = "megabytes"
)

func (e SizeUnit) ToPointer() *SizeUnit {
	return &e
}
func (e *SizeUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bytes":
		fallthrough
	case "kilobytes":
		fallthrough
	case "megabytes":
		*e = SizeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SizeUnit: %v", v)
	}
}

type RequestSizeLimitingPluginConfig struct {
	// Allowed request payload size in megabytes. Default is `128` megabytes (128000000 bytes).
	AllowedPayloadSize *int64 `json:"allowed_payload_size,omitempty"`
	// Set to `true` to ensure a valid `Content-Length` header exists before reading the request body.
	RequireContentLength *bool `json:"require_content_length,omitempty"`
	// Size unit can be set either in `bytes`, `kilobytes`, or `megabytes` (default). This configuration is not available in versions prior to Kong Gateway 1.3 and Kong Gateway (OSS) 2.0.
	SizeUnit *SizeUnit `json:"size_unit,omitempty"`
}

func (o *RequestSizeLimitingPluginConfig) GetAllowedPayloadSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AllowedPayloadSize
}

func (o *RequestSizeLimitingPluginConfig) GetRequireContentLength() *bool {
	if o == nil {
		return nil
	}
	return o.RequireContentLength
}

func (o *RequestSizeLimitingPluginConfig) GetSizeUnit() *SizeUnit {
	if o == nil {
		return nil
	}
	return o.SizeUnit
}

// RequestSizeLimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RequestSizeLimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestSizeLimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestSizeLimitingPluginProtocols string

const (
	RequestSizeLimitingPluginProtocolsGrpc  RequestSizeLimitingPluginProtocols = "grpc"
	RequestSizeLimitingPluginProtocolsGrpcs RequestSizeLimitingPluginProtocols = "grpcs"
	RequestSizeLimitingPluginProtocolsHTTP  RequestSizeLimitingPluginProtocols = "http"
	RequestSizeLimitingPluginProtocolsHTTPS RequestSizeLimitingPluginProtocols = "https"
)

func (e RequestSizeLimitingPluginProtocols) ToPointer() *RequestSizeLimitingPluginProtocols {
	return &e
}
func (e *RequestSizeLimitingPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		*e = RequestSizeLimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSizeLimitingPluginProtocols: %v", v)
	}
}

// RequestSizeLimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type RequestSizeLimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestSizeLimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestSizeLimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RequestSizeLimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestSizeLimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestSizeLimitingPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RequestSizeLimitingPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                               `json:"enabled,omitempty"`
	ID           *string                             `json:"id,omitempty"`
	InstanceName *string                             `json:"instance_name,omitempty"`
	name         string                              `const:"request-size-limiting" json:"name"`
	Ordering     *RequestSizeLimitingPluginOrdering  `json:"ordering,omitempty"`
	Partials     []RequestSizeLimitingPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                           `json:"updated_at,omitempty"`
	Config    *RequestSizeLimitingPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RequestSizeLimitingPluginConsumer `json:"consumer,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []RequestSizeLimitingPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *RequestSizeLimitingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestSizeLimitingPluginService `json:"service,omitempty"`
}

func (r RequestSizeLimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestSizeLimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestSizeLimitingPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestSizeLimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestSizeLimitingPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestSizeLimitingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestSizeLimitingPlugin) GetName() string {
	return "request-size-limiting"
}

func (o *RequestSizeLimitingPlugin) GetOrdering() *RequestSizeLimitingPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestSizeLimitingPlugin) GetPartials() []RequestSizeLimitingPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *RequestSizeLimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestSizeLimitingPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestSizeLimitingPlugin) GetConfig() *RequestSizeLimitingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *RequestSizeLimitingPlugin) GetConsumer() *RequestSizeLimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestSizeLimitingPlugin) GetProtocols() []RequestSizeLimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestSizeLimitingPlugin) GetRoute() *RequestSizeLimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestSizeLimitingPlugin) GetService() *RequestSizeLimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
