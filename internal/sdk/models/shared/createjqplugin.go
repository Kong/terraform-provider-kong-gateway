// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type RequestJqProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *RequestJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *RequestJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *RequestJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *RequestJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *RequestJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type ResponseJqProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *ResponseJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *ResponseJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *ResponseJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *ResponseJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *ResponseJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type CreateJQPluginConfig struct {
	RequestIfMediaType       []string                  `json:"request_if_media_type,omitempty"`
	RequestJqProgram         *string                   `json:"request_jq_program,omitempty"`
	RequestJqProgramOptions  *RequestJqProgramOptions  `json:"request_jq_program_options,omitempty"`
	ResponseIfMediaType      []string                  `json:"response_if_media_type,omitempty"`
	ResponseIfStatusCode     []int64                   `json:"response_if_status_code,omitempty"`
	ResponseJqProgram        *string                   `json:"response_jq_program,omitempty"`
	ResponseJqProgramOptions *ResponseJqProgramOptions `json:"response_jq_program_options,omitempty"`
}

func (o *CreateJQPluginConfig) GetRequestIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.RequestIfMediaType
}

func (o *CreateJQPluginConfig) GetRequestJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.RequestJqProgram
}

func (o *CreateJQPluginConfig) GetRequestJqProgramOptions() *RequestJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.RequestJqProgramOptions
}

func (o *CreateJQPluginConfig) GetResponseIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.ResponseIfMediaType
}

func (o *CreateJQPluginConfig) GetResponseIfStatusCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseIfStatusCode
}

func (o *CreateJQPluginConfig) GetResponseJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgram
}

func (o *CreateJQPluginConfig) GetResponseJqProgramOptions() *ResponseJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgramOptions
}

type CreateJQPluginProtocols string

const (
	CreateJQPluginProtocolsGrpc           CreateJQPluginProtocols = "grpc"
	CreateJQPluginProtocolsGrpcs          CreateJQPluginProtocols = "grpcs"
	CreateJQPluginProtocolsHTTP           CreateJQPluginProtocols = "http"
	CreateJQPluginProtocolsHTTPS          CreateJQPluginProtocols = "https"
	CreateJQPluginProtocolsTCP            CreateJQPluginProtocols = "tcp"
	CreateJQPluginProtocolsTLS            CreateJQPluginProtocols = "tls"
	CreateJQPluginProtocolsTLSPassthrough CreateJQPluginProtocols = "tls_passthrough"
	CreateJQPluginProtocolsUDP            CreateJQPluginProtocols = "udp"
	CreateJQPluginProtocolsWs             CreateJQPluginProtocols = "ws"
	CreateJQPluginProtocolsWss            CreateJQPluginProtocols = "wss"
)

func (e CreateJQPluginProtocols) ToPointer() *CreateJQPluginProtocols {
	return &e
}
func (e *CreateJQPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateJQPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateJQPluginProtocols: %v", v)
	}
}

// CreateJQPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateJQPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJQPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJQPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateJQPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJQPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateJQPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJQPlugin struct {
	Config *CreateJQPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"jq" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateJQPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateJQPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateJQPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateJQPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateJQPluginService `json:"service,omitempty"`
}

func (c CreateJQPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateJQPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateJQPlugin) GetConfig() *CreateJQPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateJQPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateJQPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateJQPlugin) GetName() *string {
	return types.String("jq")
}

func (o *CreateJQPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateJQPlugin) GetProtocols() []CreateJQPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateJQPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateJQPlugin) GetConsumer() *CreateJQPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateJQPlugin) GetConsumerGroup() *CreateJQPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateJQPlugin) GetRoute() *CreateJQPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateJQPlugin) GetService() *CreateJQPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
