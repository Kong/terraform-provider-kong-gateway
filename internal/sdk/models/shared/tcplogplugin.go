// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type TCPLogPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *TCPLogPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TCPLogPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *TCPLogPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TCPLogPluginOrdering struct {
	After  *TCPLogPluginAfter  `json:"after,omitempty"`
	Before *TCPLogPluginBefore `json:"before,omitempty"`
}

func (o *TCPLogPluginOrdering) GetAfter() *TCPLogPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *TCPLogPluginOrdering) GetBefore() *TCPLogPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type TCPLogPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *TCPLogPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TCPLogPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TCPLogPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

type TCPLogPluginConfig struct {
	// A list of key-value pairs, where the key is the name of a log field and the value is a chunk of Lua code, whose return value sets or replaces the log field value.
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	// The IP address or host name to send data to.
	Host *string `json:"host,omitempty"`
	// An optional value in milliseconds that defines how long an idle connection lives before being closed.
	Keepalive *float64 `json:"keepalive,omitempty"`
	// The port to send data to on the upstream server.
	Port *int64 `json:"port,omitempty"`
	// An optional timeout in milliseconds when sending data to the upstream server.
	Timeout *float64 `json:"timeout,omitempty"`
	// Indicates whether to perform a TLS handshake against the remote server.
	TLS *bool `json:"tls,omitempty"`
	// An optional string that defines the SNI (Server Name Indication) hostname to send in the TLS handshake.
	TLSSni *string `json:"tls_sni,omitempty"`
}

func (o *TCPLogPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *TCPLogPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *TCPLogPluginConfig) GetKeepalive() *float64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *TCPLogPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *TCPLogPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *TCPLogPluginConfig) GetTLS() *bool {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *TCPLogPluginConfig) GetTLSSni() *string {
	if o == nil {
		return nil
	}
	return o.TLSSni
}

// TCPLogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type TCPLogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *TCPLogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TCPLogPluginProtocols - A string representing a protocol, such as HTTP or HTTPS.
type TCPLogPluginProtocols string

const (
	TCPLogPluginProtocolsGrpc           TCPLogPluginProtocols = "grpc"
	TCPLogPluginProtocolsGrpcs          TCPLogPluginProtocols = "grpcs"
	TCPLogPluginProtocolsHTTP           TCPLogPluginProtocols = "http"
	TCPLogPluginProtocolsHTTPS          TCPLogPluginProtocols = "https"
	TCPLogPluginProtocolsTCP            TCPLogPluginProtocols = "tcp"
	TCPLogPluginProtocolsTLS            TCPLogPluginProtocols = "tls"
	TCPLogPluginProtocolsTLSPassthrough TCPLogPluginProtocols = "tls_passthrough"
	TCPLogPluginProtocolsUDP            TCPLogPluginProtocols = "udp"
	TCPLogPluginProtocolsWs             TCPLogPluginProtocols = "ws"
	TCPLogPluginProtocolsWss            TCPLogPluginProtocols = "wss"
)

func (e TCPLogPluginProtocols) ToPointer() *TCPLogPluginProtocols {
	return &e
}
func (e *TCPLogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = TCPLogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TCPLogPluginProtocols: %v", v)
	}
}

// TCPLogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type TCPLogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *TCPLogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TCPLogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type TCPLogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *TCPLogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TCPLogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type TCPLogPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         string                 `const:"tcp-log" json:"name"`
	Ordering     *TCPLogPluginOrdering  `json:"ordering,omitempty"`
	Partials     []TCPLogPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64              `json:"updated_at,omitempty"`
	Config    *TCPLogPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *TCPLogPluginConsumer `json:"consumer,omitempty"`
	// A set of strings representing protocols.
	Protocols []TCPLogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *TCPLogPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *TCPLogPluginService `json:"service,omitempty"`
}

func (t TCPLogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TCPLogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TCPLogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TCPLogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *TCPLogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TCPLogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *TCPLogPlugin) GetName() string {
	return "tcp-log"
}

func (o *TCPLogPlugin) GetOrdering() *TCPLogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *TCPLogPlugin) GetPartials() []TCPLogPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *TCPLogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TCPLogPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TCPLogPlugin) GetConfig() *TCPLogPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *TCPLogPlugin) GetConsumer() *TCPLogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *TCPLogPlugin) GetProtocols() []TCPLogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *TCPLogPlugin) GetRoute() *TCPLogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *TCPLogPlugin) GetService() *TCPLogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
