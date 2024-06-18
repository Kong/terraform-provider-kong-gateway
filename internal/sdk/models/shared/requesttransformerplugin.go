// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type RequestTransformerPluginAdd struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerPluginAdd) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerPluginAdd) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerPluginAppend struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerPluginAppend) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerPluginAppend) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerPluginRemove struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerPluginRemove) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerPluginRemove) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerPluginRename struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerPluginRename) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerPluginRename) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerPluginReplace struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
	URI         *string  `json:"uri,omitempty"`
}

func (o *RequestTransformerPluginReplace) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerPluginReplace) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

func (o *RequestTransformerPluginReplace) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

type RequestTransformerPluginConfig struct {
	Add    *RequestTransformerPluginAdd    `json:"add,omitempty"`
	Append *RequestTransformerPluginAppend `json:"append,omitempty"`
	// A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.
	HTTPMethod *string                          `json:"http_method,omitempty"`
	Remove     *RequestTransformerPluginRemove  `json:"remove,omitempty"`
	Rename     *RequestTransformerPluginRename  `json:"rename,omitempty"`
	Replace    *RequestTransformerPluginReplace `json:"replace,omitempty"`
}

func (o *RequestTransformerPluginConfig) GetAdd() *RequestTransformerPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *RequestTransformerPluginConfig) GetAppend() *RequestTransformerPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *RequestTransformerPluginConfig) GetHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *RequestTransformerPluginConfig) GetRemove() *RequestTransformerPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *RequestTransformerPluginConfig) GetRename() *RequestTransformerPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *RequestTransformerPluginConfig) GetReplace() *RequestTransformerPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

type RequestTransformerPluginProtocols string

const (
	RequestTransformerPluginProtocolsGrpc           RequestTransformerPluginProtocols = "grpc"
	RequestTransformerPluginProtocolsGrpcs          RequestTransformerPluginProtocols = "grpcs"
	RequestTransformerPluginProtocolsHTTP           RequestTransformerPluginProtocols = "http"
	RequestTransformerPluginProtocolsHTTPS          RequestTransformerPluginProtocols = "https"
	RequestTransformerPluginProtocolsTCP            RequestTransformerPluginProtocols = "tcp"
	RequestTransformerPluginProtocolsTLS            RequestTransformerPluginProtocols = "tls"
	RequestTransformerPluginProtocolsTLSPassthrough RequestTransformerPluginProtocols = "tls_passthrough"
	RequestTransformerPluginProtocolsUDP            RequestTransformerPluginProtocols = "udp"
	RequestTransformerPluginProtocolsWs             RequestTransformerPluginProtocols = "ws"
	RequestTransformerPluginProtocolsWss            RequestTransformerPluginProtocols = "wss"
)

func (e RequestTransformerPluginProtocols) ToPointer() *RequestTransformerPluginProtocols {
	return &e
}
func (e *RequestTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RequestTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTransformerPluginProtocols: %v", v)
	}
}

// RequestTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RequestTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type RequestTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RequestTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTransformerPlugin struct {
	Config *RequestTransformerPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"request-transformer" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RequestTransformerPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *RequestTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *RequestTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RequestTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTransformerPluginService `json:"service,omitempty"`
}

func (r RequestTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTransformerPlugin) GetConfig() *RequestTransformerPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *RequestTransformerPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTransformerPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTransformerPlugin) GetName() *string {
	return types.String("request-transformer")
}

func (o *RequestTransformerPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestTransformerPlugin) GetProtocols() []RequestTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestTransformerPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestTransformerPlugin) GetConsumer() *RequestTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTransformerPlugin) GetConsumerGroup() *RequestTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTransformerPlugin) GetRoute() *RequestTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTransformerPlugin) GetService() *RequestTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
