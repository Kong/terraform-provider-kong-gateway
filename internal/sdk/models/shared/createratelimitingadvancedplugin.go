// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// Identifier - The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`.
type Identifier string

const (
	IdentifierIP            Identifier = "ip"
	IdentifierCredential    Identifier = "credential"
	IdentifierConsumer      Identifier = "consumer"
	IdentifierService       Identifier = "service"
	IdentifierHeader        Identifier = "header"
	IdentifierPath          Identifier = "path"
	IdentifierConsumerGroup Identifier = "consumer-group"
)

func (e Identifier) ToPointer() *Identifier {
	return &e
}
func (e *Identifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ip":
		fallthrough
	case "credential":
		fallthrough
	case "consumer":
		fallthrough
	case "service":
		fallthrough
	case "header":
		fallthrough
	case "path":
		fallthrough
	case "consumer-group":
		*e = Identifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Identifier: %v", v)
	}
}

// SentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type SentinelRole string

const (
	SentinelRoleMaster SentinelRole = "master"
	SentinelRoleSlave  SentinelRole = "slave"
	SentinelRoleAny    SentinelRole = "any"
)

func (e SentinelRole) ToPointer() *SentinelRole {
	return &e
}
func (e *SentinelRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "master":
		fallthrough
	case "slave":
		fallthrough
	case "any":
		*e = SentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SentinelRole: %v", v)
	}
}

type CreateRateLimitingAdvancedPluginRedis struct {
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Cluster. Each string element must be a hostname. The minimum length of the array is 1 element.
	ClusterAddresses []string `json:"cluster_addresses,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
	KeepalivePoolSize *int64 `json:"keepalive_pool_size,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel. Each string element must be a hostname. The minimum length of the array is 1 element.
	SentinelAddresses []string `json:"sentinel_addresses,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *SentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
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

func (o *CreateRateLimitingAdvancedPluginRedis) GetClusterAddresses() []string {
	if o == nil {
		return nil
	}
	return o.ClusterAddresses
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSentinelAddresses() []string {
	if o == nil {
		return nil
	}
	return o.SentinelAddresses
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSentinelRole() *SentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateRateLimitingAdvancedPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// CreateRateLimitingAdvancedPluginStrategy - The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`.
type CreateRateLimitingAdvancedPluginStrategy string

const (
	CreateRateLimitingAdvancedPluginStrategyCluster CreateRateLimitingAdvancedPluginStrategy = "cluster"
	CreateRateLimitingAdvancedPluginStrategyRedis   CreateRateLimitingAdvancedPluginStrategy = "redis"
	CreateRateLimitingAdvancedPluginStrategyLocal   CreateRateLimitingAdvancedPluginStrategy = "local"
)

func (e CreateRateLimitingAdvancedPluginStrategy) ToPointer() *CreateRateLimitingAdvancedPluginStrategy {
	return &e
}
func (e *CreateRateLimitingAdvancedPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cluster":
		fallthrough
	case "redis":
		fallthrough
	case "local":
		*e = CreateRateLimitingAdvancedPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRateLimitingAdvancedPluginStrategy: %v", v)
	}
}

// WindowType - Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters.
type WindowType string

const (
	WindowTypeFixed   WindowType = "fixed"
	WindowTypeSliding WindowType = "sliding"
)

func (e WindowType) ToPointer() *WindowType {
	return &e
}
func (e *WindowType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fixed":
		fallthrough
	case "sliding":
		*e = WindowType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WindowType: %v", v)
	}
}

type CreateRateLimitingAdvancedPluginConfig struct {
	// List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if `enforce_consumer_groups` is set to `true`.
	ConsumerGroups []string `json:"consumer_groups,omitempty"`
	// The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is `config.strategy` is `cluster` or `redis` and `config.sync_rate` isn't `-1`), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.
	DictionaryName *string `json:"dictionary_name,omitempty"`
	// If set to `true`, this doesn't count denied requests (status = `429`). If set to `false`, all requests, including denied ones, are counted. This parameter only affects the `sliding` window_type.
	DisablePenalty *bool `json:"disable_penalty,omitempty"`
	// Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping `enforce_consumer_groups` from `true` to `false` disables the group override, but does not clear the list of consumer groups. You can then flip `enforce_consumer_groups` to `true` to re-enforce the groups.
	EnforceConsumerGroups *bool `json:"enforce_consumer_groups,omitempty"`
	// Set a custom error code to return when the rate limit is exceeded.
	ErrorCode *float64 `json:"error_code,omitempty"`
	// Set a custom error message to return when the rate limit is exceeded.
	ErrorMessage *string `json:"error_message,omitempty"`
	// A string representing an HTTP header name.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.
	HideClientHeaders *bool `json:"hide_client_headers,omitempty"`
	// The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`.
	Identifier *Identifier `json:"identifier,omitempty"`
	// One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.
	Limit []float64 `json:"limit,omitempty"`
	// The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. `strategy`, `redis`, `sync_rate`, `window_size`, `dictionary_name`, need to be the same.
	Namespace *string `json:"namespace,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	Path  *string                                `json:"path,omitempty"`
	Redis *CreateRateLimitingAdvancedPluginRedis `json:"redis,omitempty"`
	// The upper bound of a jitter (random delay) in seconds to be added to the `Retry-After` header of denied requests (status = `429`) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is `0`; in this case, the `Retry-After` header is equal to the `RateLimit-Reset` header.
	RetryAfterJitterMax *float64 `json:"retry_after_jitter_max,omitempty"`
	// The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`.
	Strategy *CreateRateLimitingAdvancedPluginStrategy `json:"strategy,omitempty"`
	// How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).
	SyncRate *float64 `json:"sync_rate,omitempty"`
	// One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.
	WindowSize []float64 `json:"window_size,omitempty"`
	// Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters.
	WindowType *WindowType `json:"window_type,omitempty"`
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetConsumerGroups() []string {
	if o == nil {
		return nil
	}
	return o.ConsumerGroups
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetDisablePenalty() *bool {
	if o == nil {
		return nil
	}
	return o.DisablePenalty
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetEnforceConsumerGroups() *bool {
	if o == nil {
		return nil
	}
	return o.EnforceConsumerGroups
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetErrorCode() *float64 {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetIdentifier() *Identifier {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetLimit() []float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetRedis() *CreateRateLimitingAdvancedPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetRetryAfterJitterMax() *float64 {
	if o == nil {
		return nil
	}
	return o.RetryAfterJitterMax
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetStrategy() *CreateRateLimitingAdvancedPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetSyncRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SyncRate
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetWindowSize() []float64 {
	if o == nil {
		return nil
	}
	return o.WindowSize
}

func (o *CreateRateLimitingAdvancedPluginConfig) GetWindowType() *WindowType {
	if o == nil {
		return nil
	}
	return o.WindowType
}

type CreateRateLimitingAdvancedPluginProtocols string

const (
	CreateRateLimitingAdvancedPluginProtocolsGrpc           CreateRateLimitingAdvancedPluginProtocols = "grpc"
	CreateRateLimitingAdvancedPluginProtocolsGrpcs          CreateRateLimitingAdvancedPluginProtocols = "grpcs"
	CreateRateLimitingAdvancedPluginProtocolsHTTP           CreateRateLimitingAdvancedPluginProtocols = "http"
	CreateRateLimitingAdvancedPluginProtocolsHTTPS          CreateRateLimitingAdvancedPluginProtocols = "https"
	CreateRateLimitingAdvancedPluginProtocolsTCP            CreateRateLimitingAdvancedPluginProtocols = "tcp"
	CreateRateLimitingAdvancedPluginProtocolsTLS            CreateRateLimitingAdvancedPluginProtocols = "tls"
	CreateRateLimitingAdvancedPluginProtocolsTLSPassthrough CreateRateLimitingAdvancedPluginProtocols = "tls_passthrough"
	CreateRateLimitingAdvancedPluginProtocolsUDP            CreateRateLimitingAdvancedPluginProtocols = "udp"
	CreateRateLimitingAdvancedPluginProtocolsWs             CreateRateLimitingAdvancedPluginProtocols = "ws"
	CreateRateLimitingAdvancedPluginProtocolsWss            CreateRateLimitingAdvancedPluginProtocols = "wss"
)

func (e CreateRateLimitingAdvancedPluginProtocols) ToPointer() *CreateRateLimitingAdvancedPluginProtocols {
	return &e
}
func (e *CreateRateLimitingAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRateLimitingAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRateLimitingAdvancedPluginProtocols: %v", v)
	}
}

// CreateRateLimitingAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRateLimitingAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRateLimitingAdvancedPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingAdvancedPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRateLimitingAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRateLimitingAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRateLimitingAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRateLimitingAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRateLimitingAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRateLimitingAdvancedPlugin struct {
	Config *CreateRateLimitingAdvancedPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"rate-limiting-advanced" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRateLimitingAdvancedPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateRateLimitingAdvancedPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateRateLimitingAdvancedPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRateLimitingAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRateLimitingAdvancedPluginService `json:"service,omitempty"`
}

func (c CreateRateLimitingAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRateLimitingAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRateLimitingAdvancedPlugin) GetConfig() *CreateRateLimitingAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateRateLimitingAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRateLimitingAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateRateLimitingAdvancedPlugin) GetName() *string {
	return types.String("rate-limiting-advanced")
}

func (o *CreateRateLimitingAdvancedPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateRateLimitingAdvancedPlugin) GetProtocols() []CreateRateLimitingAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRateLimitingAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRateLimitingAdvancedPlugin) GetConsumer() *CreateRateLimitingAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRateLimitingAdvancedPlugin) GetConsumerGroup() *CreateRateLimitingAdvancedPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateRateLimitingAdvancedPlugin) GetRoute() *CreateRateLimitingAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRateLimitingAdvancedPlugin) GetService() *CreateRateLimitingAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}