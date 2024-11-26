// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type ProxyCachePluginMemory struct {
	// The name of the shared dictionary in which to hold cache entities when the memory strategy is selected. Note that this dictionary currently must be defined manually in the Kong Nginx template.
	DictionaryName *string `json:"dictionary_name,omitempty"`
}

func (o *ProxyCachePluginMemory) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

type RequestMethod string

const (
	RequestMethodHead  RequestMethod = "HEAD"
	RequestMethodGet   RequestMethod = "GET"
	RequestMethodPost  RequestMethod = "POST"
	RequestMethodPatch RequestMethod = "PATCH"
	RequestMethodPut   RequestMethod = "PUT"
)

func (e RequestMethod) ToPointer() *RequestMethod {
	return &e
}
func (e *RequestMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HEAD":
		fallthrough
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PATCH":
		fallthrough
	case "PUT":
		*e = RequestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestMethod: %v", v)
	}
}

// ResponseHeaders - Caching related diagnostic headers that should be included in cached responses
type ResponseHeaders struct {
	XCacheKey    *bool `json:"X-Cache-Key,omitempty"`
	XCacheStatus *bool `json:"X-Cache-Status,omitempty"`
	Age          *bool `json:"age,omitempty"`
}

func (o *ResponseHeaders) GetXCacheKey() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheKey
}

func (o *ResponseHeaders) GetXCacheStatus() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheStatus
}

func (o *ResponseHeaders) GetAge() *bool {
	if o == nil {
		return nil
	}
	return o.Age
}

// ProxyCachePluginStrategy - The backing data store in which to hold cache entities.
type ProxyCachePluginStrategy string

const (
	ProxyCachePluginStrategyMemory ProxyCachePluginStrategy = "memory"
)

func (e ProxyCachePluginStrategy) ToPointer() *ProxyCachePluginStrategy {
	return &e
}
func (e *ProxyCachePluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		*e = ProxyCachePluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProxyCachePluginStrategy: %v", v)
	}
}

type ProxyCachePluginConfig struct {
	// When enabled, respect the Cache-Control behaviors defined in RFC7234.
	CacheControl *bool `json:"cache_control,omitempty"`
	// TTL, in seconds, of cache entities.
	CacheTTL *int64 `json:"cache_ttl,omitempty"`
	// Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value.
	ContentType   []string                `json:"content_type,omitempty"`
	IgnoreURICase *bool                   `json:"ignore_uri_case,omitempty"`
	Memory        *ProxyCachePluginMemory `json:"memory,omitempty"`
	// Downstream request methods considered cacheable.
	RequestMethod []RequestMethod `json:"request_method,omitempty"`
	// Upstream response status code considered cacheable.
	ResponseCode []int64 `json:"response_code,omitempty"`
	// Caching related diagnostic headers that should be included in cached responses
	ResponseHeaders *ResponseHeaders `json:"response_headers,omitempty"`
	// Number of seconds to keep resources in the storage backend. This value is independent of `cache_ttl` or resource TTLs defined by Cache-Control behaviors.
	StorageTTL *int64 `json:"storage_ttl,omitempty"`
	// The backing data store in which to hold cache entities.
	Strategy *ProxyCachePluginStrategy `json:"strategy,omitempty"`
	// Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.
	VaryHeaders []string `json:"vary_headers,omitempty"`
	// Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.
	VaryQueryParams []string `json:"vary_query_params,omitempty"`
}

func (o *ProxyCachePluginConfig) GetCacheControl() *bool {
	if o == nil {
		return nil
	}
	return o.CacheControl
}

func (o *ProxyCachePluginConfig) GetCacheTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *ProxyCachePluginConfig) GetContentType() []string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *ProxyCachePluginConfig) GetIgnoreURICase() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreURICase
}

func (o *ProxyCachePluginConfig) GetMemory() *ProxyCachePluginMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *ProxyCachePluginConfig) GetRequestMethod() []RequestMethod {
	if o == nil {
		return nil
	}
	return o.RequestMethod
}

func (o *ProxyCachePluginConfig) GetResponseCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseCode
}

func (o *ProxyCachePluginConfig) GetResponseHeaders() *ResponseHeaders {
	if o == nil {
		return nil
	}
	return o.ResponseHeaders
}

func (o *ProxyCachePluginConfig) GetStorageTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageTTL
}

func (o *ProxyCachePluginConfig) GetStrategy() *ProxyCachePluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *ProxyCachePluginConfig) GetVaryHeaders() []string {
	if o == nil {
		return nil
	}
	return o.VaryHeaders
}

func (o *ProxyCachePluginConfig) GetVaryQueryParams() []string {
	if o == nil {
		return nil
	}
	return o.VaryQueryParams
}

// ProxyCachePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type ProxyCachePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ProxyCachePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ProxyCachePluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *ProxyCachePluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ProxyCachePluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *ProxyCachePluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ProxyCachePluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *ProxyCachePluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ProxyCachePluginOrdering struct {
	After  *ProxyCachePluginAfter  `json:"after,omitempty"`
	Before *ProxyCachePluginBefore `json:"before,omitempty"`
}

func (o *ProxyCachePluginOrdering) GetAfter() *ProxyCachePluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ProxyCachePluginOrdering) GetBefore() *ProxyCachePluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type ProxyCachePluginProtocols string

const (
	ProxyCachePluginProtocolsGrpc           ProxyCachePluginProtocols = "grpc"
	ProxyCachePluginProtocolsGrpcs          ProxyCachePluginProtocols = "grpcs"
	ProxyCachePluginProtocolsHTTP           ProxyCachePluginProtocols = "http"
	ProxyCachePluginProtocolsHTTPS          ProxyCachePluginProtocols = "https"
	ProxyCachePluginProtocolsTCP            ProxyCachePluginProtocols = "tcp"
	ProxyCachePluginProtocolsTLS            ProxyCachePluginProtocols = "tls"
	ProxyCachePluginProtocolsTLSPassthrough ProxyCachePluginProtocols = "tls_passthrough"
	ProxyCachePluginProtocolsUDP            ProxyCachePluginProtocols = "udp"
	ProxyCachePluginProtocolsWs             ProxyCachePluginProtocols = "ws"
	ProxyCachePluginProtocolsWss            ProxyCachePluginProtocols = "wss"
)

func (e ProxyCachePluginProtocols) ToPointer() *ProxyCachePluginProtocols {
	return &e
}
func (e *ProxyCachePluginProtocols) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = ProxyCachePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProxyCachePluginProtocols: %v", v)
	}
}

// ProxyCachePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type ProxyCachePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *ProxyCachePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ProxyCachePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type ProxyCachePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *ProxyCachePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ProxyCachePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ProxyCachePlugin struct {
	Config ProxyCachePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *ProxyCachePluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *ProxyCachePluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	ID           *string                   `json:"id,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         string                    `const:"proxy-cache" json:"name"`
	Ordering     *ProxyCachePluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []ProxyCachePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *ProxyCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ProxyCachePluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (p ProxyCachePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProxyCachePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProxyCachePlugin) GetConfig() ProxyCachePluginConfig {
	if o == nil {
		return ProxyCachePluginConfig{}
	}
	return o.Config
}

func (o *ProxyCachePlugin) GetConsumer() *ProxyCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ProxyCachePlugin) GetConsumerGroup() *ProxyCachePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *ProxyCachePlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ProxyCachePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ProxyCachePlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ProxyCachePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ProxyCachePlugin) GetName() string {
	return "proxy-cache"
}

func (o *ProxyCachePlugin) GetOrdering() *ProxyCachePluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ProxyCachePlugin) GetProtocols() []ProxyCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ProxyCachePlugin) GetRoute() *ProxyCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ProxyCachePlugin) GetService() *ProxyCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ProxyCachePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ProxyCachePlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// ProxyCachePluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ProxyCachePluginInput struct {
	Config ProxyCachePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *ProxyCachePluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *ProxyCachePluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	ID           *string                   `json:"id,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         string                    `const:"proxy-cache" json:"name"`
	Ordering     *ProxyCachePluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []ProxyCachePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *ProxyCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ProxyCachePluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (p ProxyCachePluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProxyCachePluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProxyCachePluginInput) GetConfig() ProxyCachePluginConfig {
	if o == nil {
		return ProxyCachePluginConfig{}
	}
	return o.Config
}

func (o *ProxyCachePluginInput) GetConsumer() *ProxyCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ProxyCachePluginInput) GetConsumerGroup() *ProxyCachePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *ProxyCachePluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ProxyCachePluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ProxyCachePluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ProxyCachePluginInput) GetName() string {
	return "proxy-cache"
}

func (o *ProxyCachePluginInput) GetOrdering() *ProxyCachePluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ProxyCachePluginInput) GetProtocols() []ProxyCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ProxyCachePluginInput) GetRoute() *ProxyCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ProxyCachePluginInput) GetService() *ProxyCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ProxyCachePluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
