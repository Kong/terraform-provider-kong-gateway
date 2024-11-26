// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type PostFunctionPluginConfig struct {
	Access          []string `json:"access,omitempty"`
	BodyFilter      []string `json:"body_filter,omitempty"`
	Certificate     []string `json:"certificate,omitempty"`
	HeaderFilter    []string `json:"header_filter,omitempty"`
	Log             []string `json:"log,omitempty"`
	Rewrite         []string `json:"rewrite,omitempty"`
	WsClientFrame   []string `json:"ws_client_frame,omitempty"`
	WsClose         []string `json:"ws_close,omitempty"`
	WsHandshake     []string `json:"ws_handshake,omitempty"`
	WsUpstreamFrame []string `json:"ws_upstream_frame,omitempty"`
}

func (o *PostFunctionPluginConfig) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

func (o *PostFunctionPluginConfig) GetBodyFilter() []string {
	if o == nil {
		return nil
	}
	return o.BodyFilter
}

func (o *PostFunctionPluginConfig) GetCertificate() []string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *PostFunctionPluginConfig) GetHeaderFilter() []string {
	if o == nil {
		return nil
	}
	return o.HeaderFilter
}

func (o *PostFunctionPluginConfig) GetLog() []string {
	if o == nil {
		return nil
	}
	return o.Log
}

func (o *PostFunctionPluginConfig) GetRewrite() []string {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *PostFunctionPluginConfig) GetWsClientFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsClientFrame
}

func (o *PostFunctionPluginConfig) GetWsClose() []string {
	if o == nil {
		return nil
	}
	return o.WsClose
}

func (o *PostFunctionPluginConfig) GetWsHandshake() []string {
	if o == nil {
		return nil
	}
	return o.WsHandshake
}

func (o *PostFunctionPluginConfig) GetWsUpstreamFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsUpstreamFrame
}

// PostFunctionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type PostFunctionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostFunctionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PostFunctionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostFunctionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PostFunctionPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *PostFunctionPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type PostFunctionPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *PostFunctionPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type PostFunctionPluginOrdering struct {
	After  *PostFunctionPluginAfter  `json:"after,omitempty"`
	Before *PostFunctionPluginBefore `json:"before,omitempty"`
}

func (o *PostFunctionPluginOrdering) GetAfter() *PostFunctionPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *PostFunctionPluginOrdering) GetBefore() *PostFunctionPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type PostFunctionPluginProtocols string

const (
	PostFunctionPluginProtocolsGrpc           PostFunctionPluginProtocols = "grpc"
	PostFunctionPluginProtocolsGrpcs          PostFunctionPluginProtocols = "grpcs"
	PostFunctionPluginProtocolsHTTP           PostFunctionPluginProtocols = "http"
	PostFunctionPluginProtocolsHTTPS          PostFunctionPluginProtocols = "https"
	PostFunctionPluginProtocolsTCP            PostFunctionPluginProtocols = "tcp"
	PostFunctionPluginProtocolsTLS            PostFunctionPluginProtocols = "tls"
	PostFunctionPluginProtocolsTLSPassthrough PostFunctionPluginProtocols = "tls_passthrough"
	PostFunctionPluginProtocolsUDP            PostFunctionPluginProtocols = "udp"
	PostFunctionPluginProtocolsWs             PostFunctionPluginProtocols = "ws"
	PostFunctionPluginProtocolsWss            PostFunctionPluginProtocols = "wss"
)

func (e PostFunctionPluginProtocols) ToPointer() *PostFunctionPluginProtocols {
	return &e
}
func (e *PostFunctionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = PostFunctionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostFunctionPluginProtocols: %v", v)
	}
}

// PostFunctionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type PostFunctionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostFunctionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PostFunctionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type PostFunctionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *PostFunctionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PostFunctionPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type PostFunctionPlugin struct {
	Config PostFunctionPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PostFunctionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *PostFunctionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                       `json:"enabled,omitempty"`
	ID           *string                     `json:"id,omitempty"`
	InstanceName *string                     `json:"instance_name,omitempty"`
	name         string                      `const:"post-function" json:"name"`
	Ordering     *PostFunctionPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []PostFunctionPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PostFunctionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PostFunctionPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (p PostFunctionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostFunctionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostFunctionPlugin) GetConfig() PostFunctionPluginConfig {
	if o == nil {
		return PostFunctionPluginConfig{}
	}
	return o.Config
}

func (o *PostFunctionPlugin) GetConsumer() *PostFunctionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PostFunctionPlugin) GetConsumerGroup() *PostFunctionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PostFunctionPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PostFunctionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PostFunctionPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostFunctionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PostFunctionPlugin) GetName() string {
	return "post-function"
}

func (o *PostFunctionPlugin) GetOrdering() *PostFunctionPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *PostFunctionPlugin) GetProtocols() []PostFunctionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PostFunctionPlugin) GetRoute() *PostFunctionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PostFunctionPlugin) GetService() *PostFunctionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *PostFunctionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PostFunctionPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// PostFunctionPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type PostFunctionPluginInput struct {
	Config PostFunctionPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PostFunctionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *PostFunctionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                       `json:"enabled,omitempty"`
	ID           *string                     `json:"id,omitempty"`
	InstanceName *string                     `json:"instance_name,omitempty"`
	name         string                      `const:"post-function" json:"name"`
	Ordering     *PostFunctionPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []PostFunctionPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PostFunctionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PostFunctionPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (p PostFunctionPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostFunctionPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostFunctionPluginInput) GetConfig() PostFunctionPluginConfig {
	if o == nil {
		return PostFunctionPluginConfig{}
	}
	return o.Config
}

func (o *PostFunctionPluginInput) GetConsumer() *PostFunctionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PostFunctionPluginInput) GetConsumerGroup() *PostFunctionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PostFunctionPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PostFunctionPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostFunctionPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PostFunctionPluginInput) GetName() string {
	return "post-function"
}

func (o *PostFunctionPluginInput) GetOrdering() *PostFunctionPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *PostFunctionPluginInput) GetProtocols() []PostFunctionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PostFunctionPluginInput) GetRoute() *PostFunctionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PostFunctionPluginInput) GetService() *PostFunctionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *PostFunctionPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
