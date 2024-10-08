// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// CreateResponseRatelimitingPluginLimitBy - The entity that will be used when aggregating the limits: `consumer`, `credential`, `ip`. If the `consumer` or the `credential` cannot be determined, the system will always fallback to `ip`.
type CreateResponseRatelimitingPluginLimitBy string

const (
	CreateResponseRatelimitingPluginLimitByConsumer   CreateResponseRatelimitingPluginLimitBy = "consumer"
	CreateResponseRatelimitingPluginLimitByCredential CreateResponseRatelimitingPluginLimitBy = "credential"
	CreateResponseRatelimitingPluginLimitByIP         CreateResponseRatelimitingPluginLimitBy = "ip"
)

func (e CreateResponseRatelimitingPluginLimitBy) ToPointer() *CreateResponseRatelimitingPluginLimitBy {
	return &e
}
func (e *CreateResponseRatelimitingPluginLimitBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "credential":
		fallthrough
	case "ip":
		*e = CreateResponseRatelimitingPluginLimitBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseRatelimitingPluginLimitBy: %v", v)
	}
}

// CreateResponseRatelimitingPluginPolicy - The rate-limiting policies to use for retrieving and incrementing the limits.
type CreateResponseRatelimitingPluginPolicy string

const (
	CreateResponseRatelimitingPluginPolicyLocal   CreateResponseRatelimitingPluginPolicy = "local"
	CreateResponseRatelimitingPluginPolicyCluster CreateResponseRatelimitingPluginPolicy = "cluster"
	CreateResponseRatelimitingPluginPolicyRedis   CreateResponseRatelimitingPluginPolicy = "redis"
)

func (e CreateResponseRatelimitingPluginPolicy) ToPointer() *CreateResponseRatelimitingPluginPolicy {
	return &e
}
func (e *CreateResponseRatelimitingPluginPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		fallthrough
	case "cluster":
		fallthrough
	case "redis":
		*e = CreateResponseRatelimitingPluginPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseRatelimitingPluginPolicy: %v", v)
	}
}

// CreateResponseRatelimitingPluginRedis - Redis configuration
type CreateResponseRatelimitingPluginRedis struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	Timeout *int64 `json:"timeout,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *CreateResponseRatelimitingPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *CreateResponseRatelimitingPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateResponseRatelimitingPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateResponseRatelimitingPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateResponseRatelimitingPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *CreateResponseRatelimitingPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *CreateResponseRatelimitingPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *CreateResponseRatelimitingPluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateResponseRatelimitingPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type CreateResponseRatelimitingPluginConfig struct {
	// A boolean value that determines if the requests should be blocked as soon as one limit is being exceeded. This will block requests that are supposed to consume other limits too.
	BlockOnFirstViolation *bool `json:"block_on_first_violation,omitempty"`
	// A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party datastore. If `true`, requests will be proxied anyway, effectively disabling the rate-limiting function until the datastore is working again. If `false`, then the clients will see `500` errors.
	FaultTolerant *bool `json:"fault_tolerant,omitempty"`
	// The name of the response header used to increment the counters.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers.
	HideClientHeaders *bool `json:"hide_client_headers,omitempty"`
	// The entity that will be used when aggregating the limits: `consumer`, `credential`, `ip`. If the `consumer` or the `credential` cannot be determined, the system will always fallback to `ip`.
	LimitBy *CreateResponseRatelimitingPluginLimitBy `json:"limit_by,omitempty"`
	// A map that defines rate limits for the plugin.
	Limits map[string]any `json:"limits,omitempty"`
	// The rate-limiting policies to use for retrieving and incrementing the limits.
	Policy *CreateResponseRatelimitingPluginPolicy `json:"policy,omitempty"`
	// Redis configuration
	Redis *CreateResponseRatelimitingPluginRedis `json:"redis,omitempty"`
}

func (o *CreateResponseRatelimitingPluginConfig) GetBlockOnFirstViolation() *bool {
	if o == nil {
		return nil
	}
	return o.BlockOnFirstViolation
}

func (o *CreateResponseRatelimitingPluginConfig) GetFaultTolerant() *bool {
	if o == nil {
		return nil
	}
	return o.FaultTolerant
}

func (o *CreateResponseRatelimitingPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *CreateResponseRatelimitingPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *CreateResponseRatelimitingPluginConfig) GetLimitBy() *CreateResponseRatelimitingPluginLimitBy {
	if o == nil {
		return nil
	}
	return o.LimitBy
}

func (o *CreateResponseRatelimitingPluginConfig) GetLimits() map[string]any {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *CreateResponseRatelimitingPluginConfig) GetPolicy() *CreateResponseRatelimitingPluginPolicy {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *CreateResponseRatelimitingPluginConfig) GetRedis() *CreateResponseRatelimitingPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

type CreateResponseRatelimitingPluginProtocols string

const (
	CreateResponseRatelimitingPluginProtocolsGrpc           CreateResponseRatelimitingPluginProtocols = "grpc"
	CreateResponseRatelimitingPluginProtocolsGrpcs          CreateResponseRatelimitingPluginProtocols = "grpcs"
	CreateResponseRatelimitingPluginProtocolsHTTP           CreateResponseRatelimitingPluginProtocols = "http"
	CreateResponseRatelimitingPluginProtocolsHTTPS          CreateResponseRatelimitingPluginProtocols = "https"
	CreateResponseRatelimitingPluginProtocolsTCP            CreateResponseRatelimitingPluginProtocols = "tcp"
	CreateResponseRatelimitingPluginProtocolsTLS            CreateResponseRatelimitingPluginProtocols = "tls"
	CreateResponseRatelimitingPluginProtocolsTLSPassthrough CreateResponseRatelimitingPluginProtocols = "tls_passthrough"
	CreateResponseRatelimitingPluginProtocolsUDP            CreateResponseRatelimitingPluginProtocols = "udp"
	CreateResponseRatelimitingPluginProtocolsWs             CreateResponseRatelimitingPluginProtocols = "ws"
	CreateResponseRatelimitingPluginProtocolsWss            CreateResponseRatelimitingPluginProtocols = "wss"
)

func (e CreateResponseRatelimitingPluginProtocols) ToPointer() *CreateResponseRatelimitingPluginProtocols {
	return &e
}
func (e *CreateResponseRatelimitingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateResponseRatelimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseRatelimitingPluginProtocols: %v", v)
	}
}

// CreateResponseRatelimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateResponseRatelimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseRatelimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseRatelimitingPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseRatelimitingPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseRatelimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateResponseRatelimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseRatelimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseRatelimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateResponseRatelimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseRatelimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseRatelimitingPlugin struct {
	Config *CreateResponseRatelimitingPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"response-ratelimiting" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateResponseRatelimitingPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateResponseRatelimitingPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateResponseRatelimitingPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateResponseRatelimitingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateResponseRatelimitingPluginService `json:"service,omitempty"`
}

func (c CreateResponseRatelimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateResponseRatelimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateResponseRatelimitingPlugin) GetConfig() *CreateResponseRatelimitingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateResponseRatelimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateResponseRatelimitingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateResponseRatelimitingPlugin) GetName() *string {
	return types.String("response-ratelimiting")
}

func (o *CreateResponseRatelimitingPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateResponseRatelimitingPlugin) GetProtocols() []CreateResponseRatelimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateResponseRatelimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateResponseRatelimitingPlugin) GetConsumer() *CreateResponseRatelimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateResponseRatelimitingPlugin) GetConsumerGroup() *CreateResponseRatelimitingPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateResponseRatelimitingPlugin) GetRoute() *CreateResponseRatelimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateResponseRatelimitingPlugin) GetService() *CreateResponseRatelimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
