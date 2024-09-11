// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// MtlsAuthPluginAuthenticatedGroupBy - Certificate property to use as the authenticated group. Valid values are `CN` (Common Name) or `DN` (Distinguished Name). Once `skip_consumer_lookup` is applied, any client with a valid certificate can access the Service/API. To restrict usage to only some of the authenticated users, also add the ACL plugin (not covered here) and create allowed or denied groups of users.
type MtlsAuthPluginAuthenticatedGroupBy string

const (
	MtlsAuthPluginAuthenticatedGroupByCn MtlsAuthPluginAuthenticatedGroupBy = "CN"
	MtlsAuthPluginAuthenticatedGroupByDn MtlsAuthPluginAuthenticatedGroupBy = "DN"
)

func (e MtlsAuthPluginAuthenticatedGroupBy) ToPointer() *MtlsAuthPluginAuthenticatedGroupBy {
	return &e
}
func (e *MtlsAuthPluginAuthenticatedGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CN":
		fallthrough
	case "DN":
		*e = MtlsAuthPluginAuthenticatedGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MtlsAuthPluginAuthenticatedGroupBy: %v", v)
	}
}

type MtlsAuthPluginConsumerBy string

const (
	MtlsAuthPluginConsumerByUsername MtlsAuthPluginConsumerBy = "username"
	MtlsAuthPluginConsumerByCustomID MtlsAuthPluginConsumerBy = "custom_id"
)

func (e MtlsAuthPluginConsumerBy) ToPointer() *MtlsAuthPluginConsumerBy {
	return &e
}
func (e *MtlsAuthPluginConsumerBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "username":
		fallthrough
	case "custom_id":
		*e = MtlsAuthPluginConsumerBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MtlsAuthPluginConsumerBy: %v", v)
	}
}

// MtlsAuthPluginRevocationCheckMode - Controls client certificate revocation check behavior. If set to `SKIP`, no revocation check is performed. If set to `IGNORE_CA_ERROR`, the plugin respects the revocation status when either OCSP or CRL URL is set, and doesn't fail on network issues. If set to `STRICT`, the plugin only treats the certificate as valid when it's able to verify the revocation status.
type MtlsAuthPluginRevocationCheckMode string

const (
	MtlsAuthPluginRevocationCheckModeSkip          MtlsAuthPluginRevocationCheckMode = "SKIP"
	MtlsAuthPluginRevocationCheckModeIgnoreCaError MtlsAuthPluginRevocationCheckMode = "IGNORE_CA_ERROR"
	MtlsAuthPluginRevocationCheckModeStrict        MtlsAuthPluginRevocationCheckMode = "STRICT"
)

func (e MtlsAuthPluginRevocationCheckMode) ToPointer() *MtlsAuthPluginRevocationCheckMode {
	return &e
}
func (e *MtlsAuthPluginRevocationCheckMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SKIP":
		fallthrough
	case "IGNORE_CA_ERROR":
		fallthrough
	case "STRICT":
		*e = MtlsAuthPluginRevocationCheckMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MtlsAuthPluginRevocationCheckMode: %v", v)
	}
}

type MtlsAuthPluginConfig struct {
	// Allow certificate verification with only an intermediate certificate. When this is enabled, you don't need to upload the full chain to Kong Certificates.
	AllowPartialChain *bool `json:"allow_partial_chain,omitempty"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// Certificate property to use as the authenticated group. Valid values are `CN` (Common Name) or `DN` (Distinguished Name). Once `skip_consumer_lookup` is applied, any client with a valid certificate can access the Service/API. To restrict usage to only some of the authenticated users, also add the ACL plugin (not covered here) and create allowed or denied groups of users.
	AuthenticatedGroupBy *MtlsAuthPluginAuthenticatedGroupBy `json:"authenticated_group_by,omitempty"`
	// List of CA Certificates strings to use as Certificate Authorities (CA) when validating a client certificate. At least one is required but you can specify as many as needed. The value of this array is comprised of primary keys (`id`).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Cache expiry time in seconds.
	CacheTTL *float64 `json:"cache_ttl,omitempty"`
	// The length of time in milliseconds between refreshes of the revocation check status cache.
	CertCacheTTL *float64 `json:"cert_cache_ttl,omitempty"`
	// Whether to match the subject name of the client-supplied certificate against consumer's `username` and/or `custom_id` attribute. If set to `[]` (the empty array), then auto-matching is disabled.
	ConsumerBy []MtlsAuthPluginConsumerBy `json:"consumer_by,omitempty"`
	// The UUID or username of the consumer to use when a trusted client certificate is presented but no consumer matches. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	DefaultConsumer *string `json:"default_consumer,omitempty"`
	// A string representing a host name, such as example.com.
	HTTPProxyHost *string `json:"http_proxy_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	HTTPProxyPort *int64 `json:"http_proxy_port,omitempty"`
	// HTTP timeout threshold in milliseconds when communicating with the OCSP server or downloading CRL.
	HTTPTimeout *float64 `json:"http_timeout,omitempty"`
	// A string representing a host name, such as example.com.
	HTTPSProxyHost *string `json:"https_proxy_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	HTTPSProxyPort *int64 `json:"https_proxy_port,omitempty"`
	// Controls client certificate revocation check behavior. If set to `SKIP`, no revocation check is performed. If set to `IGNORE_CA_ERROR`, the plugin respects the revocation status when either OCSP or CRL URL is set, and doesn't fail on network issues. If set to `STRICT`, the plugin only treats the certificate as valid when it's able to verify the revocation status.
	RevocationCheckMode *MtlsAuthPluginRevocationCheckMode `json:"revocation_check_mode,omitempty"`
	// Sends the distinguished names (DN) of the configured CA list in the TLS handshake message.
	SendCaDn *bool `json:"send_ca_dn,omitempty"`
	// Skip consumer lookup once certificate is trusted against the configured CA list.
	SkipConsumerLookup *bool `json:"skip_consumer_lookup,omitempty"`
}

func (o *MtlsAuthPluginConfig) GetAllowPartialChain() *bool {
	if o == nil {
		return nil
	}
	return o.AllowPartialChain
}

func (o *MtlsAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *MtlsAuthPluginConfig) GetAuthenticatedGroupBy() *MtlsAuthPluginAuthenticatedGroupBy {
	if o == nil {
		return nil
	}
	return o.AuthenticatedGroupBy
}

func (o *MtlsAuthPluginConfig) GetCaCertificates() []string {
	if o == nil {
		return nil
	}
	return o.CaCertificates
}

func (o *MtlsAuthPluginConfig) GetCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *MtlsAuthPluginConfig) GetCertCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CertCacheTTL
}

func (o *MtlsAuthPluginConfig) GetConsumerBy() []MtlsAuthPluginConsumerBy {
	if o == nil {
		return nil
	}
	return o.ConsumerBy
}

func (o *MtlsAuthPluginConfig) GetDefaultConsumer() *string {
	if o == nil {
		return nil
	}
	return o.DefaultConsumer
}

func (o *MtlsAuthPluginConfig) GetHTTPProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPProxyHost
}

func (o *MtlsAuthPluginConfig) GetHTTPProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPProxyPort
}

func (o *MtlsAuthPluginConfig) GetHTTPTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.HTTPTimeout
}

func (o *MtlsAuthPluginConfig) GetHTTPSProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyHost
}

func (o *MtlsAuthPluginConfig) GetHTTPSProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyPort
}

func (o *MtlsAuthPluginConfig) GetRevocationCheckMode() *MtlsAuthPluginRevocationCheckMode {
	if o == nil {
		return nil
	}
	return o.RevocationCheckMode
}

func (o *MtlsAuthPluginConfig) GetSendCaDn() *bool {
	if o == nil {
		return nil
	}
	return o.SendCaDn
}

func (o *MtlsAuthPluginConfig) GetSkipConsumerLookup() *bool {
	if o == nil {
		return nil
	}
	return o.SkipConsumerLookup
}

type MtlsAuthPluginProtocols string

const (
	MtlsAuthPluginProtocolsGrpc           MtlsAuthPluginProtocols = "grpc"
	MtlsAuthPluginProtocolsGrpcs          MtlsAuthPluginProtocols = "grpcs"
	MtlsAuthPluginProtocolsHTTP           MtlsAuthPluginProtocols = "http"
	MtlsAuthPluginProtocolsHTTPS          MtlsAuthPluginProtocols = "https"
	MtlsAuthPluginProtocolsTCP            MtlsAuthPluginProtocols = "tcp"
	MtlsAuthPluginProtocolsTLS            MtlsAuthPluginProtocols = "tls"
	MtlsAuthPluginProtocolsTLSPassthrough MtlsAuthPluginProtocols = "tls_passthrough"
	MtlsAuthPluginProtocolsUDP            MtlsAuthPluginProtocols = "udp"
	MtlsAuthPluginProtocolsWs             MtlsAuthPluginProtocols = "ws"
	MtlsAuthPluginProtocolsWss            MtlsAuthPluginProtocols = "wss"
)

func (e MtlsAuthPluginProtocols) ToPointer() *MtlsAuthPluginProtocols {
	return &e
}
func (e *MtlsAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = MtlsAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MtlsAuthPluginProtocols: %v", v)
	}
}

// MtlsAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type MtlsAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *MtlsAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type MtlsAuthPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *MtlsAuthPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// MtlsAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type MtlsAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *MtlsAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// MtlsAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type MtlsAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *MtlsAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type MtlsAuthPlugin struct {
	Config *MtlsAuthPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"mtls-auth" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []MtlsAuthPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *MtlsAuthPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *MtlsAuthPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *MtlsAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *MtlsAuthPluginService `json:"service,omitempty"`
}

func (m MtlsAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MtlsAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MtlsAuthPlugin) GetConfig() *MtlsAuthPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *MtlsAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MtlsAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *MtlsAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MtlsAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *MtlsAuthPlugin) GetName() *string {
	return types.String("mtls-auth")
}

func (o *MtlsAuthPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *MtlsAuthPlugin) GetProtocols() []MtlsAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *MtlsAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *MtlsAuthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *MtlsAuthPlugin) GetConsumer() *MtlsAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *MtlsAuthPlugin) GetConsumerGroup() *MtlsAuthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *MtlsAuthPlugin) GetRoute() *MtlsAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *MtlsAuthPlugin) GetService() *MtlsAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}