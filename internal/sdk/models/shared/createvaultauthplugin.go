// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreateVaultAuthPluginConfig struct {
	// Describes an array of comma-separated parameter names where the plugin looks for an access token. The client must send the access token in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].
	AccessTokenName *string `json:"access_token_name,omitempty"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin will strip the credential from the request (i.e. the header or querystring containing the key) before proxying it.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests will always be allowed.
	RunOnPreflight *bool `json:"run_on_preflight,omitempty"`
	// Describes an array of comma-separated parameter names where the plugin looks for a secret token. The client must send the secret in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].
	SecretTokenName *string `json:"secret_token_name,omitempty"`
	// If enabled, the plugin will read the request body (if said request has one and its MIME type is supported) and try to find the key in it. Supported MIME types are `application/www-form-urlencoded`, `application/json`, and `multipart/form-data`.
	TokensInBody *bool `json:"tokens_in_body,omitempty"`
	// A reference to an existing `vault` object within the database. `vault` entities define the connection and authentication parameters used to connect to a Vault HTTP(S) API.
	Vault *string `json:"vault,omitempty"`
}

func (o *CreateVaultAuthPluginConfig) GetAccessTokenName() *string {
	if o == nil {
		return nil
	}
	return o.AccessTokenName
}

func (o *CreateVaultAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateVaultAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *CreateVaultAuthPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

func (o *CreateVaultAuthPluginConfig) GetSecretTokenName() *string {
	if o == nil {
		return nil
	}
	return o.SecretTokenName
}

func (o *CreateVaultAuthPluginConfig) GetTokensInBody() *bool {
	if o == nil {
		return nil
	}
	return o.TokensInBody
}

func (o *CreateVaultAuthPluginConfig) GetVault() *string {
	if o == nil {
		return nil
	}
	return o.Vault
}

type CreateVaultAuthPluginProtocols string

const (
	CreateVaultAuthPluginProtocolsGrpc           CreateVaultAuthPluginProtocols = "grpc"
	CreateVaultAuthPluginProtocolsGrpcs          CreateVaultAuthPluginProtocols = "grpcs"
	CreateVaultAuthPluginProtocolsHTTP           CreateVaultAuthPluginProtocols = "http"
	CreateVaultAuthPluginProtocolsHTTPS          CreateVaultAuthPluginProtocols = "https"
	CreateVaultAuthPluginProtocolsTCP            CreateVaultAuthPluginProtocols = "tcp"
	CreateVaultAuthPluginProtocolsTLS            CreateVaultAuthPluginProtocols = "tls"
	CreateVaultAuthPluginProtocolsTLSPassthrough CreateVaultAuthPluginProtocols = "tls_passthrough"
	CreateVaultAuthPluginProtocolsUDP            CreateVaultAuthPluginProtocols = "udp"
	CreateVaultAuthPluginProtocolsWs             CreateVaultAuthPluginProtocols = "ws"
	CreateVaultAuthPluginProtocolsWss            CreateVaultAuthPluginProtocols = "wss"
)

func (e CreateVaultAuthPluginProtocols) ToPointer() *CreateVaultAuthPluginProtocols {
	return &e
}
func (e *CreateVaultAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateVaultAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateVaultAuthPluginProtocols: %v", v)
	}
}

// CreateVaultAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateVaultAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateVaultAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateVaultAuthPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateVaultAuthPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateVaultAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateVaultAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateVaultAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateVaultAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateVaultAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateVaultAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateVaultAuthPlugin struct {
	Config *CreateVaultAuthPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"vault-auth" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateVaultAuthPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateVaultAuthPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateVaultAuthPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateVaultAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateVaultAuthPluginService `json:"service,omitempty"`
}

func (c CreateVaultAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateVaultAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateVaultAuthPlugin) GetConfig() *CreateVaultAuthPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateVaultAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateVaultAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateVaultAuthPlugin) GetName() *string {
	return types.String("vault-auth")
}

func (o *CreateVaultAuthPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateVaultAuthPlugin) GetProtocols() []CreateVaultAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateVaultAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateVaultAuthPlugin) GetConsumer() *CreateVaultAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateVaultAuthPlugin) GetConsumerGroup() *CreateVaultAuthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateVaultAuthPlugin) GetRoute() *CreateVaultAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateVaultAuthPlugin) GetService() *CreateVaultAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}