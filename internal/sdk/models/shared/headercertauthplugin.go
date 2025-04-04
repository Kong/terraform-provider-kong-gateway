// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type HeaderCertAuthPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *HeaderCertAuthPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type HeaderCertAuthPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *HeaderCertAuthPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type HeaderCertAuthPluginOrdering struct {
	After  *HeaderCertAuthPluginAfter  `json:"after,omitempty"`
	Before *HeaderCertAuthPluginBefore `json:"before,omitempty"`
}

func (o *HeaderCertAuthPluginOrdering) GetAfter() *HeaderCertAuthPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *HeaderCertAuthPluginOrdering) GetBefore() *HeaderCertAuthPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type HeaderCertAuthPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *HeaderCertAuthPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HeaderCertAuthPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HeaderCertAuthPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// AuthenticatedGroupBy - Certificate property to use as the authenticated group. Valid values are `CN` (Common Name) or `DN` (Distinguished Name). Once `skip_consumer_lookup` is applied, any client with a valid certificate can access the Service/API. To restrict usage to only some of the authenticated users, also add the ACL plugin (not covered here) and create allowed or denied groups of users.
type AuthenticatedGroupBy string

const (
	AuthenticatedGroupByCn AuthenticatedGroupBy = "CN"
	AuthenticatedGroupByDn AuthenticatedGroupBy = "DN"
)

func (e AuthenticatedGroupBy) ToPointer() *AuthenticatedGroupBy {
	return &e
}
func (e *AuthenticatedGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CN":
		fallthrough
	case "DN":
		*e = AuthenticatedGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthenticatedGroupBy: %v", v)
	}
}

// CertificateHeaderFormat - Format of the certificate header. Supported formats: `base64_encoded`, `url_encoded`.
type CertificateHeaderFormat string

const (
	CertificateHeaderFormatBase64Encoded CertificateHeaderFormat = "base64_encoded"
	CertificateHeaderFormatURLEncoded    CertificateHeaderFormat = "url_encoded"
)

func (e CertificateHeaderFormat) ToPointer() *CertificateHeaderFormat {
	return &e
}
func (e *CertificateHeaderFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "base64_encoded":
		fallthrough
	case "url_encoded":
		*e = CertificateHeaderFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CertificateHeaderFormat: %v", v)
	}
}

type ConsumerBy string

const (
	ConsumerByCustomID ConsumerBy = "custom_id"
	ConsumerByUsername ConsumerBy = "username"
)

func (e ConsumerBy) ToPointer() *ConsumerBy {
	return &e
}
func (e *ConsumerBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom_id":
		fallthrough
	case "username":
		*e = ConsumerBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsumerBy: %v", v)
	}
}

// RevocationCheckMode - Controls client certificate revocation check behavior. If set to `SKIP`, no revocation check is performed. If set to `IGNORE_CA_ERROR`, the plugin respects the revocation status when either OCSP or CRL URL is set, and doesn't fail on network issues. If set to `STRICT`, the plugin only treats the certificate as valid when it's able to verify the revocation status.
type RevocationCheckMode string

const (
	RevocationCheckModeIgnoreCaError RevocationCheckMode = "IGNORE_CA_ERROR"
	RevocationCheckModeSkip          RevocationCheckMode = "SKIP"
	RevocationCheckModeStrict        RevocationCheckMode = "STRICT"
)

func (e RevocationCheckMode) ToPointer() *RevocationCheckMode {
	return &e
}
func (e *RevocationCheckMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IGNORE_CA_ERROR":
		fallthrough
	case "SKIP":
		fallthrough
	case "STRICT":
		*e = RevocationCheckMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RevocationCheckMode: %v", v)
	}
}

type HeaderCertAuthPluginConfig struct {
	// Allow certificate verification with only an intermediate certificate. When this is enabled, you don't need to upload the full chain to Kong Certificates.
	AllowPartialChain *bool `json:"allow_partial_chain,omitempty"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// Certificate property to use as the authenticated group. Valid values are `CN` (Common Name) or `DN` (Distinguished Name). Once `skip_consumer_lookup` is applied, any client with a valid certificate can access the Service/API. To restrict usage to only some of the authenticated users, also add the ACL plugin (not covered here) and create allowed or denied groups of users.
	AuthenticatedGroupBy *AuthenticatedGroupBy `json:"authenticated_group_by,omitempty"`
	// List of CA Certificates strings to use as Certificate Authorities (CA) when validating a client certificate. At least one is required but you can specify as many as needed. The value of this array is comprised of primary keys (`id`).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Cache expiry time in seconds.
	CacheTTL *float64 `json:"cache_ttl,omitempty"`
	// The length of time in milliseconds between refreshes of the revocation check status cache.
	CertCacheTTL *float64 `json:"cert_cache_ttl,omitempty"`
	// Format of the certificate header. Supported formats: `base64_encoded`, `url_encoded`.
	CertificateHeaderFormat *CertificateHeaderFormat `json:"certificate_header_format,omitempty"`
	// Name of the header that contains the certificate, received from the WAF or other L7 downstream proxy.
	CertificateHeaderName *string `json:"certificate_header_name,omitempty"`
	// Whether to match the subject name of the client-supplied certificate against consumer's `username` and/or `custom_id` attribute. If set to `[]` (the empty array), then auto-matching is disabled.
	ConsumerBy []ConsumerBy `json:"consumer_by,omitempty"`
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
	RevocationCheckMode *RevocationCheckMode `json:"revocation_check_mode,omitempty"`
	// Whether to secure the source of the request. If set to `true`, the plugin will only allow requests from trusted IPs (configured by the `trusted_ips` config option).
	SecureSource *bool `json:"secure_source,omitempty"`
	// Skip consumer lookup once certificate is trusted against the configured CA list.
	SkipConsumerLookup *bool `json:"skip_consumer_lookup,omitempty"`
}

func (o *HeaderCertAuthPluginConfig) GetAllowPartialChain() *bool {
	if o == nil {
		return nil
	}
	return o.AllowPartialChain
}

func (o *HeaderCertAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *HeaderCertAuthPluginConfig) GetAuthenticatedGroupBy() *AuthenticatedGroupBy {
	if o == nil {
		return nil
	}
	return o.AuthenticatedGroupBy
}

func (o *HeaderCertAuthPluginConfig) GetCaCertificates() []string {
	if o == nil {
		return nil
	}
	return o.CaCertificates
}

func (o *HeaderCertAuthPluginConfig) GetCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *HeaderCertAuthPluginConfig) GetCertCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CertCacheTTL
}

func (o *HeaderCertAuthPluginConfig) GetCertificateHeaderFormat() *CertificateHeaderFormat {
	if o == nil {
		return nil
	}
	return o.CertificateHeaderFormat
}

func (o *HeaderCertAuthPluginConfig) GetCertificateHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateHeaderName
}

func (o *HeaderCertAuthPluginConfig) GetConsumerBy() []ConsumerBy {
	if o == nil {
		return nil
	}
	return o.ConsumerBy
}

func (o *HeaderCertAuthPluginConfig) GetDefaultConsumer() *string {
	if o == nil {
		return nil
	}
	return o.DefaultConsumer
}

func (o *HeaderCertAuthPluginConfig) GetHTTPProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPProxyHost
}

func (o *HeaderCertAuthPluginConfig) GetHTTPProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPProxyPort
}

func (o *HeaderCertAuthPluginConfig) GetHTTPTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.HTTPTimeout
}

func (o *HeaderCertAuthPluginConfig) GetHTTPSProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyHost
}

func (o *HeaderCertAuthPluginConfig) GetHTTPSProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyPort
}

func (o *HeaderCertAuthPluginConfig) GetRevocationCheckMode() *RevocationCheckMode {
	if o == nil {
		return nil
	}
	return o.RevocationCheckMode
}

func (o *HeaderCertAuthPluginConfig) GetSecureSource() *bool {
	if o == nil {
		return nil
	}
	return o.SecureSource
}

func (o *HeaderCertAuthPluginConfig) GetSkipConsumerLookup() *bool {
	if o == nil {
		return nil
	}
	return o.SkipConsumerLookup
}

type HeaderCertAuthPluginProtocols string

const (
	HeaderCertAuthPluginProtocolsGrpc  HeaderCertAuthPluginProtocols = "grpc"
	HeaderCertAuthPluginProtocolsGrpcs HeaderCertAuthPluginProtocols = "grpcs"
	HeaderCertAuthPluginProtocolsHTTP  HeaderCertAuthPluginProtocols = "http"
	HeaderCertAuthPluginProtocolsHTTPS HeaderCertAuthPluginProtocols = "https"
)

func (e HeaderCertAuthPluginProtocols) ToPointer() *HeaderCertAuthPluginProtocols {
	return &e
}
func (e *HeaderCertAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = HeaderCertAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderCertAuthPluginProtocols: %v", v)
	}
}

// HeaderCertAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type HeaderCertAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *HeaderCertAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HeaderCertAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type HeaderCertAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *HeaderCertAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HeaderCertAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type HeaderCertAuthPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                          `json:"enabled,omitempty"`
	ID           *string                        `json:"id,omitempty"`
	InstanceName *string                        `json:"instance_name,omitempty"`
	name         string                         `const:"header-cert-auth" json:"name"`
	Ordering     *HeaderCertAuthPluginOrdering  `json:"ordering,omitempty"`
	Partials     []HeaderCertAuthPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                      `json:"updated_at,omitempty"`
	Config    *HeaderCertAuthPluginConfig `json:"config,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []HeaderCertAuthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *HeaderCertAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *HeaderCertAuthPluginService `json:"service,omitempty"`
}

func (h HeaderCertAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HeaderCertAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HeaderCertAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HeaderCertAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *HeaderCertAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HeaderCertAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *HeaderCertAuthPlugin) GetName() string {
	return "header-cert-auth"
}

func (o *HeaderCertAuthPlugin) GetOrdering() *HeaderCertAuthPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *HeaderCertAuthPlugin) GetPartials() []HeaderCertAuthPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *HeaderCertAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HeaderCertAuthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *HeaderCertAuthPlugin) GetConfig() *HeaderCertAuthPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *HeaderCertAuthPlugin) GetProtocols() []HeaderCertAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *HeaderCertAuthPlugin) GetRoute() *HeaderCertAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *HeaderCertAuthPlugin) GetService() *HeaderCertAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
