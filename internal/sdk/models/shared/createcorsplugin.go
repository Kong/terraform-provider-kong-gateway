// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type Methods string

const (
	MethodsGet     Methods = "GET"
	MethodsHead    Methods = "HEAD"
	MethodsPut     Methods = "PUT"
	MethodsPatch   Methods = "PATCH"
	MethodsPost    Methods = "POST"
	MethodsDelete  Methods = "DELETE"
	MethodsOptions Methods = "OPTIONS"
	MethodsTrace   Methods = "TRACE"
	MethodsConnect Methods = "CONNECT"
)

func (e Methods) ToPointer() *Methods {
	return &e
}
func (e *Methods) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "HEAD":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "POST":
		fallthrough
	case "DELETE":
		fallthrough
	case "OPTIONS":
		fallthrough
	case "TRACE":
		fallthrough
	case "CONNECT":
		*e = Methods(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Methods: %v", v)
	}
}

type CreateCORSPluginConfig struct {
	// Flag to determine whether the `Access-Control-Allow-Credentials` header should be sent with `true` as the value.
	Credentials *bool `json:"credentials,omitempty"`
	// Value for the `Access-Control-Expose-Headers` header. If not specified, no custom headers are exposed.
	ExposedHeaders []string `json:"exposed_headers,omitempty"`
	// Value for the `Access-Control-Allow-Headers` header.
	Headers []string `json:"headers,omitempty"`
	// Indicates how long the results of the preflight request can be cached, in `seconds`.
	MaxAge *float64 `json:"max_age,omitempty"`
	// 'Value for the `Access-Control-Allow-Methods` header. Available options include `GET`, `HEAD`, `PUT`, `PATCH`, `POST`, `DELETE`, `OPTIONS`, `TRACE`, `CONNECT`. By default, all options are allowed.'
	Methods []Methods `json:"methods,omitempty"`
	// List of allowed domains for the `Access-Control-Allow-Origin` header. If you want to allow all origins, add `*` as a single value to this configuration field. The accepted values can either be flat strings or PCRE regexes.
	Origins []string `json:"origins,omitempty"`
	// A boolean value that instructs the plugin to proxy the `OPTIONS` preflight request to the Upstream service.
	PreflightContinue *bool `json:"preflight_continue,omitempty"`
	// Flag to determine whether the `Access-Control-Allow-Private-Network` header should be sent with `true` as the value.
	PrivateNetwork *bool `json:"private_network,omitempty"`
}

func (o *CreateCORSPluginConfig) GetCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *CreateCORSPluginConfig) GetExposedHeaders() []string {
	if o == nil {
		return nil
	}
	return o.ExposedHeaders
}

func (o *CreateCORSPluginConfig) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateCORSPluginConfig) GetMaxAge() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxAge
}

func (o *CreateCORSPluginConfig) GetMethods() []Methods {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *CreateCORSPluginConfig) GetOrigins() []string {
	if o == nil {
		return nil
	}
	return o.Origins
}

func (o *CreateCORSPluginConfig) GetPreflightContinue() *bool {
	if o == nil {
		return nil
	}
	return o.PreflightContinue
}

func (o *CreateCORSPluginConfig) GetPrivateNetwork() *bool {
	if o == nil {
		return nil
	}
	return o.PrivateNetwork
}

type CreateCORSPluginProtocols string

const (
	CreateCORSPluginProtocolsGrpc           CreateCORSPluginProtocols = "grpc"
	CreateCORSPluginProtocolsGrpcs          CreateCORSPluginProtocols = "grpcs"
	CreateCORSPluginProtocolsHTTP           CreateCORSPluginProtocols = "http"
	CreateCORSPluginProtocolsHTTPS          CreateCORSPluginProtocols = "https"
	CreateCORSPluginProtocolsTCP            CreateCORSPluginProtocols = "tcp"
	CreateCORSPluginProtocolsTLS            CreateCORSPluginProtocols = "tls"
	CreateCORSPluginProtocolsTLSPassthrough CreateCORSPluginProtocols = "tls_passthrough"
	CreateCORSPluginProtocolsUDP            CreateCORSPluginProtocols = "udp"
	CreateCORSPluginProtocolsWs             CreateCORSPluginProtocols = "ws"
	CreateCORSPluginProtocolsWss            CreateCORSPluginProtocols = "wss"
)

func (e CreateCORSPluginProtocols) ToPointer() *CreateCORSPluginProtocols {
	return &e
}
func (e *CreateCORSPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateCORSPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCORSPluginProtocols: %v", v)
	}
}

// CreateCORSPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateCORSPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCORSPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateCORSPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCORSPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCORSPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateCORSPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCORSPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCORSPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateCORSPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCORSPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateCORSPlugin struct {
	Config *CreateCORSPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"cors" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateCORSPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateCORSPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateCORSPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateCORSPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateCORSPluginService `json:"service,omitempty"`
}

func (c CreateCORSPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCORSPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCORSPlugin) GetConfig() *CreateCORSPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateCORSPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateCORSPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateCORSPlugin) GetName() *string {
	return types.String("cors")
}

func (o *CreateCORSPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateCORSPlugin) GetProtocols() []CreateCORSPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateCORSPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateCORSPlugin) GetConsumer() *CreateCORSPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateCORSPlugin) GetConsumerGroup() *CreateCORSPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateCORSPlugin) GetRoute() *CreateCORSPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateCORSPlugin) GetService() *CreateCORSPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
