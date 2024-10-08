// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type AiPromptDecoratorPluginRole string

const (
	AiPromptDecoratorPluginRoleSystem    AiPromptDecoratorPluginRole = "system"
	AiPromptDecoratorPluginRoleAssistant AiPromptDecoratorPluginRole = "assistant"
	AiPromptDecoratorPluginRoleUser      AiPromptDecoratorPluginRole = "user"
)

func (e AiPromptDecoratorPluginRole) ToPointer() *AiPromptDecoratorPluginRole {
	return &e
}
func (e *AiPromptDecoratorPluginRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = AiPromptDecoratorPluginRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptDecoratorPluginRole: %v", v)
	}
}

type AiPromptDecoratorPluginAppend struct {
	Content string                       `json:"content"`
	Role    *AiPromptDecoratorPluginRole `json:"role,omitempty"`
}

func (o *AiPromptDecoratorPluginAppend) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *AiPromptDecoratorPluginAppend) GetRole() *AiPromptDecoratorPluginRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type AiPromptDecoratorPluginConfigRole string

const (
	AiPromptDecoratorPluginConfigRoleSystem    AiPromptDecoratorPluginConfigRole = "system"
	AiPromptDecoratorPluginConfigRoleAssistant AiPromptDecoratorPluginConfigRole = "assistant"
	AiPromptDecoratorPluginConfigRoleUser      AiPromptDecoratorPluginConfigRole = "user"
)

func (e AiPromptDecoratorPluginConfigRole) ToPointer() *AiPromptDecoratorPluginConfigRole {
	return &e
}
func (e *AiPromptDecoratorPluginConfigRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = AiPromptDecoratorPluginConfigRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptDecoratorPluginConfigRole: %v", v)
	}
}

type AiPromptDecoratorPluginPrepend struct {
	Content string                             `json:"content"`
	Role    *AiPromptDecoratorPluginConfigRole `json:"role,omitempty"`
}

func (o *AiPromptDecoratorPluginPrepend) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *AiPromptDecoratorPluginPrepend) GetRole() *AiPromptDecoratorPluginConfigRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type AiPromptDecoratorPluginPrompts struct {
	// Insert chat messages at the end of the chat message array. This array preserves exact order when adding messages.
	Append []AiPromptDecoratorPluginAppend `json:"append,omitempty"`
	// Insert chat messages at the beginning of the chat message array. This array preserves exact order when adding messages.
	Prepend []AiPromptDecoratorPluginPrepend `json:"prepend,omitempty"`
}

func (o *AiPromptDecoratorPluginPrompts) GetAppend() []AiPromptDecoratorPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *AiPromptDecoratorPluginPrompts) GetPrepend() []AiPromptDecoratorPluginPrepend {
	if o == nil {
		return nil
	}
	return o.Prepend
}

type AiPromptDecoratorPluginConfig struct {
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64                          `json:"max_request_body_size,omitempty"`
	Prompts            *AiPromptDecoratorPluginPrompts `json:"prompts,omitempty"`
}

func (o *AiPromptDecoratorPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

func (o *AiPromptDecoratorPluginConfig) GetPrompts() *AiPromptDecoratorPluginPrompts {
	if o == nil {
		return nil
	}
	return o.Prompts
}

type AiPromptDecoratorPluginProtocols string

const (
	AiPromptDecoratorPluginProtocolsGrpc           AiPromptDecoratorPluginProtocols = "grpc"
	AiPromptDecoratorPluginProtocolsGrpcs          AiPromptDecoratorPluginProtocols = "grpcs"
	AiPromptDecoratorPluginProtocolsHTTP           AiPromptDecoratorPluginProtocols = "http"
	AiPromptDecoratorPluginProtocolsHTTPS          AiPromptDecoratorPluginProtocols = "https"
	AiPromptDecoratorPluginProtocolsTCP            AiPromptDecoratorPluginProtocols = "tcp"
	AiPromptDecoratorPluginProtocolsTLS            AiPromptDecoratorPluginProtocols = "tls"
	AiPromptDecoratorPluginProtocolsTLSPassthrough AiPromptDecoratorPluginProtocols = "tls_passthrough"
	AiPromptDecoratorPluginProtocolsUDP            AiPromptDecoratorPluginProtocols = "udp"
	AiPromptDecoratorPluginProtocolsWs             AiPromptDecoratorPluginProtocols = "ws"
	AiPromptDecoratorPluginProtocolsWss            AiPromptDecoratorPluginProtocols = "wss"
)

func (e AiPromptDecoratorPluginProtocols) ToPointer() *AiPromptDecoratorPluginProtocols {
	return &e
}
func (e *AiPromptDecoratorPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiPromptDecoratorPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptDecoratorPluginProtocols: %v", v)
	}
}

// AiPromptDecoratorPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiPromptDecoratorPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptDecoratorPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptDecoratorPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptDecoratorPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptDecoratorPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AiPromptDecoratorPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptDecoratorPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptDecoratorPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiPromptDecoratorPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptDecoratorPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptDecoratorPlugin struct {
	Config *AiPromptDecoratorPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"ai-prompt-decorator" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AiPromptDecoratorPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AiPromptDecoratorPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AiPromptDecoratorPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AiPromptDecoratorPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiPromptDecoratorPluginService `json:"service,omitempty"`
}

func (a AiPromptDecoratorPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiPromptDecoratorPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiPromptDecoratorPlugin) GetConfig() *AiPromptDecoratorPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *AiPromptDecoratorPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiPromptDecoratorPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiPromptDecoratorPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptDecoratorPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiPromptDecoratorPlugin) GetName() *string {
	return types.String("ai-prompt-decorator")
}

func (o *AiPromptDecoratorPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiPromptDecoratorPlugin) GetProtocols() []AiPromptDecoratorPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiPromptDecoratorPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiPromptDecoratorPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiPromptDecoratorPlugin) GetConsumer() *AiPromptDecoratorPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiPromptDecoratorPlugin) GetConsumerGroup() *AiPromptDecoratorPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiPromptDecoratorPlugin) GetRoute() *AiPromptDecoratorPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiPromptDecoratorPlugin) GetService() *AiPromptDecoratorPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
