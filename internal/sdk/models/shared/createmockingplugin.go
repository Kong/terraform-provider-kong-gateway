// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreateMockingPluginConfig struct {
	// The contents of the specification file. You must use this option for hybrid or DB-less mode. You can include the full specification as part of the configuration. In Kong Manager, you can copy and paste the contents of the spec directly into the `Config.Api Specification` text field.
	APISpecification *string `json:"api_specification,omitempty"`
	// The path and name of the specification file loaded into Kong Gateway's database. You cannot use this option for DB-less or hybrid mode.
	APISpecificationFilename *string `json:"api_specification_filename,omitempty"`
	// The base path to be used for path match evaluation. This value is ignored if `include_base_path` is set to `false`.
	CustomBasePath *string `json:"custom_base_path,omitempty"`
	// Indicates whether to include the base path when performing path match evaluation.
	IncludeBasePath *bool `json:"include_base_path,omitempty"`
	// A global list of the HTTP status codes that can only be selected and returned.
	IncludedStatusCodes []int64 `json:"included_status_codes,omitempty"`
	// The maximum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be greater than the `min_delay_time`.
	MaxDelayTime *float64 `json:"max_delay_time,omitempty"`
	// The minimum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be less than the `max_delay_time`.
	MinDelayTime *float64 `json:"min_delay_time,omitempty"`
	// Enables a random delay in the mocked response. Introduces delays to simulate real-time response times by APIs.
	RandomDelay *bool `json:"random_delay,omitempty"`
	// Randomly selects one example and returns it. This parameter requires the spec to have multiple examples configured.
	RandomExamples *bool `json:"random_examples,omitempty"`
	// Determines whether to randomly select an HTTP status code from the responses of the corresponding API method. The default value is `false`, which means the minimum HTTP status code is always selected and returned.
	RandomStatusCode *bool `json:"random_status_code,omitempty"`
}

func (o *CreateMockingPluginConfig) GetAPISpecification() *string {
	if o == nil {
		return nil
	}
	return o.APISpecification
}

func (o *CreateMockingPluginConfig) GetAPISpecificationFilename() *string {
	if o == nil {
		return nil
	}
	return o.APISpecificationFilename
}

func (o *CreateMockingPluginConfig) GetCustomBasePath() *string {
	if o == nil {
		return nil
	}
	return o.CustomBasePath
}

func (o *CreateMockingPluginConfig) GetIncludeBasePath() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeBasePath
}

func (o *CreateMockingPluginConfig) GetIncludedStatusCodes() []int64 {
	if o == nil {
		return nil
	}
	return o.IncludedStatusCodes
}

func (o *CreateMockingPluginConfig) GetMaxDelayTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxDelayTime
}

func (o *CreateMockingPluginConfig) GetMinDelayTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MinDelayTime
}

func (o *CreateMockingPluginConfig) GetRandomDelay() *bool {
	if o == nil {
		return nil
	}
	return o.RandomDelay
}

func (o *CreateMockingPluginConfig) GetRandomExamples() *bool {
	if o == nil {
		return nil
	}
	return o.RandomExamples
}

func (o *CreateMockingPluginConfig) GetRandomStatusCode() *bool {
	if o == nil {
		return nil
	}
	return o.RandomStatusCode
}

type CreateMockingPluginProtocols string

const (
	CreateMockingPluginProtocolsGrpc           CreateMockingPluginProtocols = "grpc"
	CreateMockingPluginProtocolsGrpcs          CreateMockingPluginProtocols = "grpcs"
	CreateMockingPluginProtocolsHTTP           CreateMockingPluginProtocols = "http"
	CreateMockingPluginProtocolsHTTPS          CreateMockingPluginProtocols = "https"
	CreateMockingPluginProtocolsTCP            CreateMockingPluginProtocols = "tcp"
	CreateMockingPluginProtocolsTLS            CreateMockingPluginProtocols = "tls"
	CreateMockingPluginProtocolsTLSPassthrough CreateMockingPluginProtocols = "tls_passthrough"
	CreateMockingPluginProtocolsUDP            CreateMockingPluginProtocols = "udp"
	CreateMockingPluginProtocolsWs             CreateMockingPluginProtocols = "ws"
	CreateMockingPluginProtocolsWss            CreateMockingPluginProtocols = "wss"
)

func (e CreateMockingPluginProtocols) ToPointer() *CreateMockingPluginProtocols {
	return &e
}
func (e *CreateMockingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateMockingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMockingPluginProtocols: %v", v)
	}
}

// CreateMockingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateMockingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateMockingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateMockingPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateMockingPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateMockingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateMockingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateMockingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateMockingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateMockingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateMockingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateMockingPlugin struct {
	Config *CreateMockingPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"mocking" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateMockingPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateMockingPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateMockingPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateMockingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateMockingPluginService `json:"service,omitempty"`
}

func (c CreateMockingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateMockingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateMockingPlugin) GetConfig() *CreateMockingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateMockingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateMockingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateMockingPlugin) GetName() *string {
	return types.String("mocking")
}

func (o *CreateMockingPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateMockingPlugin) GetProtocols() []CreateMockingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateMockingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateMockingPlugin) GetConsumer() *CreateMockingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateMockingPlugin) GetConsumerGroup() *CreateMockingPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateMockingPlugin) GetRoute() *CreateMockingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateMockingPlugin) GetService() *CreateMockingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}