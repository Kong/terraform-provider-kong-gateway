// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type FileLogPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *FileLogPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type FileLogPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *FileLogPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type FileLogPluginOrdering struct {
	After  *FileLogPluginAfter  `json:"after,omitempty"`
	Before *FileLogPluginBefore `json:"before,omitempty"`
}

func (o *FileLogPluginOrdering) GetAfter() *FileLogPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *FileLogPluginOrdering) GetBefore() *FileLogPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type FileLogPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *FileLogPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FileLogPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FileLogPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

type FileLogPluginConfig struct {
	// Lua code as a key-value map
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	// The file path of the output log file. The plugin creates the log file if it doesn't exist yet.
	Path *string `json:"path,omitempty"`
	// Determines whether the log file is closed and reopened on every request.
	Reopen *bool `json:"reopen,omitempty"`
}

func (o *FileLogPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *FileLogPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *FileLogPluginConfig) GetReopen() *bool {
	if o == nil {
		return nil
	}
	return o.Reopen
}

// FileLogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type FileLogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// FileLogPluginProtocols - A string representing a protocol, such as HTTP or HTTPS.
type FileLogPluginProtocols string

const (
	FileLogPluginProtocolsGrpc           FileLogPluginProtocols = "grpc"
	FileLogPluginProtocolsGrpcs          FileLogPluginProtocols = "grpcs"
	FileLogPluginProtocolsHTTP           FileLogPluginProtocols = "http"
	FileLogPluginProtocolsHTTPS          FileLogPluginProtocols = "https"
	FileLogPluginProtocolsTCP            FileLogPluginProtocols = "tcp"
	FileLogPluginProtocolsTLS            FileLogPluginProtocols = "tls"
	FileLogPluginProtocolsTLSPassthrough FileLogPluginProtocols = "tls_passthrough"
	FileLogPluginProtocolsUDP            FileLogPluginProtocols = "udp"
	FileLogPluginProtocolsWs             FileLogPluginProtocols = "ws"
	FileLogPluginProtocolsWss            FileLogPluginProtocols = "wss"
)

func (e FileLogPluginProtocols) ToPointer() *FileLogPluginProtocols {
	return &e
}
func (e *FileLogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = FileLogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileLogPluginProtocols: %v", v)
	}
}

// FileLogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type FileLogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// FileLogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type FileLogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// FileLogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type FileLogPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                   `json:"enabled,omitempty"`
	ID           *string                 `json:"id,omitempty"`
	InstanceName *string                 `json:"instance_name,omitempty"`
	name         string                  `const:"file-log" json:"name"`
	Ordering     *FileLogPluginOrdering  `json:"ordering,omitempty"`
	Partials     []FileLogPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64               `json:"updated_at,omitempty"`
	Config    *FileLogPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *FileLogPluginConsumer `json:"consumer,omitempty"`
	// A set of strings representing protocols.
	Protocols []FileLogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *FileLogPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *FileLogPluginService `json:"service,omitempty"`
}

func (f FileLogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileLogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileLogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *FileLogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *FileLogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FileLogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *FileLogPlugin) GetName() string {
	return "file-log"
}

func (o *FileLogPlugin) GetOrdering() *FileLogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *FileLogPlugin) GetPartials() []FileLogPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *FileLogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *FileLogPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *FileLogPlugin) GetConfig() *FileLogPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *FileLogPlugin) GetConsumer() *FileLogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *FileLogPlugin) GetProtocols() []FileLogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *FileLogPlugin) GetRoute() *FileLogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *FileLogPlugin) GetService() *FileLogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
