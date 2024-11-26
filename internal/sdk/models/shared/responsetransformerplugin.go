// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type ResponseTransformerPluginJSONTypes string

const (
	ResponseTransformerPluginJSONTypesBoolean ResponseTransformerPluginJSONTypes = "boolean"
	ResponseTransformerPluginJSONTypesNumber  ResponseTransformerPluginJSONTypes = "number"
	ResponseTransformerPluginJSONTypesString  ResponseTransformerPluginJSONTypes = "string"
)

func (e ResponseTransformerPluginJSONTypes) ToPointer() *ResponseTransformerPluginJSONTypes {
	return &e
}
func (e *ResponseTransformerPluginJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = ResponseTransformerPluginJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseTransformerPluginJSONTypes: %v", v)
	}
}

type ResponseTransformerPluginAdd struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []ResponseTransformerPluginJSONTypes `json:"json_types,omitempty"`
}

func (o *ResponseTransformerPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ResponseTransformerPluginAdd) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *ResponseTransformerPluginAdd) GetJSONTypes() []ResponseTransformerPluginJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type ResponseTransformerPluginConfigJSONTypes string

const (
	ResponseTransformerPluginConfigJSONTypesBoolean ResponseTransformerPluginConfigJSONTypes = "boolean"
	ResponseTransformerPluginConfigJSONTypesNumber  ResponseTransformerPluginConfigJSONTypes = "number"
	ResponseTransformerPluginConfigJSONTypesString  ResponseTransformerPluginConfigJSONTypes = "string"
)

func (e ResponseTransformerPluginConfigJSONTypes) ToPointer() *ResponseTransformerPluginConfigJSONTypes {
	return &e
}
func (e *ResponseTransformerPluginConfigJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = ResponseTransformerPluginConfigJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseTransformerPluginConfigJSONTypes: %v", v)
	}
}

type ResponseTransformerPluginAppend struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []ResponseTransformerPluginConfigJSONTypes `json:"json_types,omitempty"`
}

func (o *ResponseTransformerPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ResponseTransformerPluginAppend) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *ResponseTransformerPluginAppend) GetJSONTypes() []ResponseTransformerPluginConfigJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type ResponseTransformerPluginRemove struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
}

func (o *ResponseTransformerPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ResponseTransformerPluginRemove) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type ResponseTransformerPluginRename struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
}

func (o *ResponseTransformerPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ResponseTransformerPluginRename) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type ResponseTransformerPluginConfigReplaceJSONTypes string

const (
	ResponseTransformerPluginConfigReplaceJSONTypesBoolean ResponseTransformerPluginConfigReplaceJSONTypes = "boolean"
	ResponseTransformerPluginConfigReplaceJSONTypesNumber  ResponseTransformerPluginConfigReplaceJSONTypes = "number"
	ResponseTransformerPluginConfigReplaceJSONTypesString  ResponseTransformerPluginConfigReplaceJSONTypes = "string"
)

func (e ResponseTransformerPluginConfigReplaceJSONTypes) ToPointer() *ResponseTransformerPluginConfigReplaceJSONTypes {
	return &e
}
func (e *ResponseTransformerPluginConfigReplaceJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = ResponseTransformerPluginConfigReplaceJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseTransformerPluginConfigReplaceJSONTypes: %v", v)
	}
}

type ResponseTransformerPluginReplace struct {
	Headers []string `json:"headers,omitempty"`
	JSON    []string `json:"json,omitempty"`
	// List of JSON type names. Specify the types of the JSON values returned when appending
	// JSON properties. Each string element can be one of: boolean, number, or string.
	JSONTypes []ResponseTransformerPluginConfigReplaceJSONTypes `json:"json_types,omitempty"`
}

func (o *ResponseTransformerPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ResponseTransformerPluginReplace) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *ResponseTransformerPluginReplace) GetJSONTypes() []ResponseTransformerPluginConfigReplaceJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type ResponseTransformerPluginConfig struct {
	Add     *ResponseTransformerPluginAdd     `json:"add,omitempty"`
	Append  *ResponseTransformerPluginAppend  `json:"append,omitempty"`
	Remove  *ResponseTransformerPluginRemove  `json:"remove,omitempty"`
	Rename  *ResponseTransformerPluginRename  `json:"rename,omitempty"`
	Replace *ResponseTransformerPluginReplace `json:"replace,omitempty"`
}

func (o *ResponseTransformerPluginConfig) GetAdd() *ResponseTransformerPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *ResponseTransformerPluginConfig) GetAppend() *ResponseTransformerPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *ResponseTransformerPluginConfig) GetRemove() *ResponseTransformerPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *ResponseTransformerPluginConfig) GetRename() *ResponseTransformerPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *ResponseTransformerPluginConfig) GetReplace() *ResponseTransformerPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

// ResponseTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type ResponseTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ResponseTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ResponseTransformerPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *ResponseTransformerPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ResponseTransformerPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *ResponseTransformerPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ResponseTransformerPluginOrdering struct {
	After  *ResponseTransformerPluginAfter  `json:"after,omitempty"`
	Before *ResponseTransformerPluginBefore `json:"before,omitempty"`
}

func (o *ResponseTransformerPluginOrdering) GetAfter() *ResponseTransformerPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ResponseTransformerPluginOrdering) GetBefore() *ResponseTransformerPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type ResponseTransformerPluginProtocols string

const (
	ResponseTransformerPluginProtocolsGrpc           ResponseTransformerPluginProtocols = "grpc"
	ResponseTransformerPluginProtocolsGrpcs          ResponseTransformerPluginProtocols = "grpcs"
	ResponseTransformerPluginProtocolsHTTP           ResponseTransformerPluginProtocols = "http"
	ResponseTransformerPluginProtocolsHTTPS          ResponseTransformerPluginProtocols = "https"
	ResponseTransformerPluginProtocolsTCP            ResponseTransformerPluginProtocols = "tcp"
	ResponseTransformerPluginProtocolsTLS            ResponseTransformerPluginProtocols = "tls"
	ResponseTransformerPluginProtocolsTLSPassthrough ResponseTransformerPluginProtocols = "tls_passthrough"
	ResponseTransformerPluginProtocolsUDP            ResponseTransformerPluginProtocols = "udp"
	ResponseTransformerPluginProtocolsWs             ResponseTransformerPluginProtocols = "ws"
	ResponseTransformerPluginProtocolsWss            ResponseTransformerPluginProtocols = "wss"
)

func (e ResponseTransformerPluginProtocols) ToPointer() *ResponseTransformerPluginProtocols {
	return &e
}
func (e *ResponseTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = ResponseTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseTransformerPluginProtocols: %v", v)
	}
}

// ResponseTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type ResponseTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ResponseTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type ResponseTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ResponseTransformerPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ResponseTransformerPlugin struct {
	Config ResponseTransformerPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *ResponseTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *ResponseTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                              `json:"enabled,omitempty"`
	ID           *string                            `json:"id,omitempty"`
	InstanceName *string                            `json:"instance_name,omitempty"`
	name         string                             `const:"response-transformer" json:"name"`
	Ordering     *ResponseTransformerPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []ResponseTransformerPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *ResponseTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ResponseTransformerPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (r ResponseTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseTransformerPlugin) GetConfig() ResponseTransformerPluginConfig {
	if o == nil {
		return ResponseTransformerPluginConfig{}
	}
	return o.Config
}

func (o *ResponseTransformerPlugin) GetConsumer() *ResponseTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ResponseTransformerPlugin) GetConsumerGroup() *ResponseTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *ResponseTransformerPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ResponseTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ResponseTransformerPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResponseTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ResponseTransformerPlugin) GetName() string {
	return "response-transformer"
}

func (o *ResponseTransformerPlugin) GetOrdering() *ResponseTransformerPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ResponseTransformerPlugin) GetProtocols() []ResponseTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ResponseTransformerPlugin) GetRoute() *ResponseTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ResponseTransformerPlugin) GetService() *ResponseTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ResponseTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ResponseTransformerPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// ResponseTransformerPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ResponseTransformerPluginInput struct {
	Config ResponseTransformerPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *ResponseTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *ResponseTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                              `json:"enabled,omitempty"`
	ID           *string                            `json:"id,omitempty"`
	InstanceName *string                            `json:"instance_name,omitempty"`
	name         string                             `const:"response-transformer" json:"name"`
	Ordering     *ResponseTransformerPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []ResponseTransformerPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *ResponseTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ResponseTransformerPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (r ResponseTransformerPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseTransformerPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseTransformerPluginInput) GetConfig() ResponseTransformerPluginConfig {
	if o == nil {
		return ResponseTransformerPluginConfig{}
	}
	return o.Config
}

func (o *ResponseTransformerPluginInput) GetConsumer() *ResponseTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ResponseTransformerPluginInput) GetConsumerGroup() *ResponseTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *ResponseTransformerPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ResponseTransformerPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResponseTransformerPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ResponseTransformerPluginInput) GetName() string {
	return "response-transformer"
}

func (o *ResponseTransformerPluginInput) GetOrdering() *ResponseTransformerPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ResponseTransformerPluginInput) GetProtocols() []ResponseTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ResponseTransformerPluginInput) GetRoute() *ResponseTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ResponseTransformerPluginInput) GetService() *ResponseTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ResponseTransformerPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
