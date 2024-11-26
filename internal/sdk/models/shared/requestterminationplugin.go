// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type RequestTerminationPluginConfig struct {
	// The raw response body to send. This is mutually exclusive with the `config.message` field.
	Body *string `json:"body,omitempty"`
	// Content type of the raw response configured with `config.body`.
	ContentType *string `json:"content_type,omitempty"`
	// When set, the plugin will echo a copy of the request back to the client. The main usecase for this is debugging. It can be combined with `trigger` in order to debug requests on live systems without disturbing real traffic.
	Echo *bool `json:"echo,omitempty"`
	// The message to send, if using the default response generator.
	Message *string `json:"message,omitempty"`
	// The response code to send. Must be an integer between 100 and 599.
	StatusCode *int64 `json:"status_code,omitempty"`
	// A string representing an HTTP header name.
	Trigger *string `json:"trigger,omitempty"`
}

func (o *RequestTerminationPluginConfig) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTerminationPluginConfig) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *RequestTerminationPluginConfig) GetEcho() *bool {
	if o == nil {
		return nil
	}
	return o.Echo
}

func (o *RequestTerminationPluginConfig) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RequestTerminationPluginConfig) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *RequestTerminationPluginConfig) GetTrigger() *string {
	if o == nil {
		return nil
	}
	return o.Trigger
}

// RequestTerminationPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RequestTerminationPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTerminationPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTerminationPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTerminationPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTerminationPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *RequestTerminationPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RequestTerminationPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *RequestTerminationPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RequestTerminationPluginOrdering struct {
	After  *RequestTerminationPluginAfter  `json:"after,omitempty"`
	Before *RequestTerminationPluginBefore `json:"before,omitempty"`
}

func (o *RequestTerminationPluginOrdering) GetAfter() *RequestTerminationPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *RequestTerminationPluginOrdering) GetBefore() *RequestTerminationPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type RequestTerminationPluginProtocols string

const (
	RequestTerminationPluginProtocolsGrpc           RequestTerminationPluginProtocols = "grpc"
	RequestTerminationPluginProtocolsGrpcs          RequestTerminationPluginProtocols = "grpcs"
	RequestTerminationPluginProtocolsHTTP           RequestTerminationPluginProtocols = "http"
	RequestTerminationPluginProtocolsHTTPS          RequestTerminationPluginProtocols = "https"
	RequestTerminationPluginProtocolsTCP            RequestTerminationPluginProtocols = "tcp"
	RequestTerminationPluginProtocolsTLS            RequestTerminationPluginProtocols = "tls"
	RequestTerminationPluginProtocolsTLSPassthrough RequestTerminationPluginProtocols = "tls_passthrough"
	RequestTerminationPluginProtocolsUDP            RequestTerminationPluginProtocols = "udp"
	RequestTerminationPluginProtocolsWs             RequestTerminationPluginProtocols = "ws"
	RequestTerminationPluginProtocolsWss            RequestTerminationPluginProtocols = "wss"
)

func (e RequestTerminationPluginProtocols) ToPointer() *RequestTerminationPluginProtocols {
	return &e
}
func (e *RequestTerminationPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RequestTerminationPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTerminationPluginProtocols: %v", v)
	}
}

// RequestTerminationPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type RequestTerminationPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTerminationPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTerminationPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RequestTerminationPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTerminationPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTerminationPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RequestTerminationPlugin struct {
	Config RequestTerminationPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *RequestTerminationPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *RequestTerminationPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                             `json:"enabled,omitempty"`
	ID           *string                           `json:"id,omitempty"`
	InstanceName *string                           `json:"instance_name,omitempty"`
	name         string                            `const:"request-termination" json:"name"`
	Ordering     *RequestTerminationPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RequestTerminationPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RequestTerminationPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTerminationPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (r RequestTerminationPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTerminationPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTerminationPlugin) GetConfig() RequestTerminationPluginConfig {
	if o == nil {
		return RequestTerminationPluginConfig{}
	}
	return o.Config
}

func (o *RequestTerminationPlugin) GetConsumer() *RequestTerminationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTerminationPlugin) GetConsumerGroup() *RequestTerminationPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTerminationPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestTerminationPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTerminationPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTerminationPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTerminationPlugin) GetName() string {
	return "request-termination"
}

func (o *RequestTerminationPlugin) GetOrdering() *RequestTerminationPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestTerminationPlugin) GetProtocols() []RequestTerminationPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTerminationPlugin) GetRoute() *RequestTerminationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTerminationPlugin) GetService() *RequestTerminationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *RequestTerminationPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestTerminationPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// RequestTerminationPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RequestTerminationPluginInput struct {
	Config RequestTerminationPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *RequestTerminationPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *RequestTerminationPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                             `json:"enabled,omitempty"`
	ID           *string                           `json:"id,omitempty"`
	InstanceName *string                           `json:"instance_name,omitempty"`
	name         string                            `const:"request-termination" json:"name"`
	Ordering     *RequestTerminationPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RequestTerminationPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RequestTerminationPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTerminationPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (r RequestTerminationPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTerminationPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTerminationPluginInput) GetConfig() RequestTerminationPluginConfig {
	if o == nil {
		return RequestTerminationPluginConfig{}
	}
	return o.Config
}

func (o *RequestTerminationPluginInput) GetConsumer() *RequestTerminationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTerminationPluginInput) GetConsumerGroup() *RequestTerminationPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTerminationPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTerminationPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTerminationPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTerminationPluginInput) GetName() string {
	return "request-termination"
}

func (o *RequestTerminationPluginInput) GetOrdering() *RequestTerminationPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestTerminationPluginInput) GetProtocols() []RequestTerminationPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTerminationPluginInput) GetRoute() *RequestTerminationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTerminationPluginInput) GetService() *RequestTerminationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *RequestTerminationPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
