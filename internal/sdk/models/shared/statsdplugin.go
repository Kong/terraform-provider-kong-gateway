// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type StatsdPluginConsumerIdentifierDefault string

const (
	StatsdPluginConsumerIdentifierDefaultConsumerID StatsdPluginConsumerIdentifierDefault = "consumer_id"
	StatsdPluginConsumerIdentifierDefaultCustomID   StatsdPluginConsumerIdentifierDefault = "custom_id"
	StatsdPluginConsumerIdentifierDefaultUsername   StatsdPluginConsumerIdentifierDefault = "username"
)

func (e StatsdPluginConsumerIdentifierDefault) ToPointer() *StatsdPluginConsumerIdentifierDefault {
	return &e
}
func (e *StatsdPluginConsumerIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = StatsdPluginConsumerIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginConsumerIdentifierDefault: %v", v)
	}
}

// StatsdPluginConsumerIdentifier - Authenticated user detail.
type StatsdPluginConsumerIdentifier string

const (
	StatsdPluginConsumerIdentifierConsumerID StatsdPluginConsumerIdentifier = "consumer_id"
	StatsdPluginConsumerIdentifierCustomID   StatsdPluginConsumerIdentifier = "custom_id"
	StatsdPluginConsumerIdentifierUsername   StatsdPluginConsumerIdentifier = "username"
)

func (e StatsdPluginConsumerIdentifier) ToPointer() *StatsdPluginConsumerIdentifier {
	return &e
}
func (e *StatsdPluginConsumerIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = StatsdPluginConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginConsumerIdentifier: %v", v)
	}
}

// StatsdPluginName - StatsD metric’s name.
type StatsdPluginName string

const (
	StatsdPluginNameKongLatency                StatsdPluginName = "kong_latency"
	StatsdPluginNameLatency                    StatsdPluginName = "latency"
	StatsdPluginNameRequestCount               StatsdPluginName = "request_count"
	StatsdPluginNameRequestPerUser             StatsdPluginName = "request_per_user"
	StatsdPluginNameRequestSize                StatsdPluginName = "request_size"
	StatsdPluginNameResponseSize               StatsdPluginName = "response_size"
	StatsdPluginNameStatusCount                StatsdPluginName = "status_count"
	StatsdPluginNameStatusCountPerUser         StatsdPluginName = "status_count_per_user"
	StatsdPluginNameUniqueUsers                StatsdPluginName = "unique_users"
	StatsdPluginNameUpstreamLatency            StatsdPluginName = "upstream_latency"
	StatsdPluginNameStatusCountPerWorkspace    StatsdPluginName = "status_count_per_workspace"
	StatsdPluginNameStatusCountPerUserPerRoute StatsdPluginName = "status_count_per_user_per_route"
	StatsdPluginNameShdictUsage                StatsdPluginName = "shdict_usage"
	StatsdPluginNameCacheDatastoreHitsTotal    StatsdPluginName = "cache_datastore_hits_total"
	StatsdPluginNameCacheDatastoreMissesTotal  StatsdPluginName = "cache_datastore_misses_total"
)

func (e StatsdPluginName) ToPointer() *StatsdPluginName {
	return &e
}
func (e *StatsdPluginName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong_latency":
		fallthrough
	case "latency":
		fallthrough
	case "request_count":
		fallthrough
	case "request_per_user":
		fallthrough
	case "request_size":
		fallthrough
	case "response_size":
		fallthrough
	case "status_count":
		fallthrough
	case "status_count_per_user":
		fallthrough
	case "unique_users":
		fallthrough
	case "upstream_latency":
		fallthrough
	case "status_count_per_workspace":
		fallthrough
	case "status_count_per_user_per_route":
		fallthrough
	case "shdict_usage":
		fallthrough
	case "cache_datastore_hits_total":
		fallthrough
	case "cache_datastore_misses_total":
		*e = StatsdPluginName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginName: %v", v)
	}
}

// StatsdPluginServiceIdentifier - Service detail.
type StatsdPluginServiceIdentifier string

const (
	StatsdPluginServiceIdentifierServiceID         StatsdPluginServiceIdentifier = "service_id"
	StatsdPluginServiceIdentifierServiceName       StatsdPluginServiceIdentifier = "service_name"
	StatsdPluginServiceIdentifierServiceHost       StatsdPluginServiceIdentifier = "service_host"
	StatsdPluginServiceIdentifierServiceNameOrHost StatsdPluginServiceIdentifier = "service_name_or_host"
)

func (e StatsdPluginServiceIdentifier) ToPointer() *StatsdPluginServiceIdentifier {
	return &e
}
func (e *StatsdPluginServiceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = StatsdPluginServiceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginServiceIdentifier: %v", v)
	}
}

// StatsdPluginStatType - Determines what sort of event a metric represents.
type StatsdPluginStatType string

const (
	StatsdPluginStatTypeCounter   StatsdPluginStatType = "counter"
	StatsdPluginStatTypeGauge     StatsdPluginStatType = "gauge"
	StatsdPluginStatTypeHistogram StatsdPluginStatType = "histogram"
	StatsdPluginStatTypeMeter     StatsdPluginStatType = "meter"
	StatsdPluginStatTypeSet       StatsdPluginStatType = "set"
	StatsdPluginStatTypeTimer     StatsdPluginStatType = "timer"
)

func (e StatsdPluginStatType) ToPointer() *StatsdPluginStatType {
	return &e
}
func (e *StatsdPluginStatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "counter":
		fallthrough
	case "gauge":
		fallthrough
	case "histogram":
		fallthrough
	case "meter":
		fallthrough
	case "set":
		fallthrough
	case "timer":
		*e = StatsdPluginStatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginStatType: %v", v)
	}
}

// StatsdPluginWorkspaceIdentifier - Workspace detail.
type StatsdPluginWorkspaceIdentifier string

const (
	StatsdPluginWorkspaceIdentifierWorkspaceID   StatsdPluginWorkspaceIdentifier = "workspace_id"
	StatsdPluginWorkspaceIdentifierWorkspaceName StatsdPluginWorkspaceIdentifier = "workspace_name"
)

func (e StatsdPluginWorkspaceIdentifier) ToPointer() *StatsdPluginWorkspaceIdentifier {
	return &e
}
func (e *StatsdPluginWorkspaceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = StatsdPluginWorkspaceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginWorkspaceIdentifier: %v", v)
	}
}

type StatsdPluginMetrics struct {
	// Authenticated user detail.
	ConsumerIdentifier *StatsdPluginConsumerIdentifier `json:"consumer_identifier,omitempty"`
	// StatsD metric’s name.
	Name StatsdPluginName `json:"name"`
	// Sampling rate
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// Service detail.
	ServiceIdentifier *StatsdPluginServiceIdentifier `json:"service_identifier,omitempty"`
	// Determines what sort of event a metric represents.
	StatType StatsdPluginStatType `json:"stat_type"`
	// Workspace detail.
	WorkspaceIdentifier *StatsdPluginWorkspaceIdentifier `json:"workspace_identifier,omitempty"`
}

func (o *StatsdPluginMetrics) GetConsumerIdentifier() *StatsdPluginConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *StatsdPluginMetrics) GetName() StatsdPluginName {
	if o == nil {
		return StatsdPluginName("")
	}
	return o.Name
}

func (o *StatsdPluginMetrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *StatsdPluginMetrics) GetServiceIdentifier() *StatsdPluginServiceIdentifier {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifier
}

func (o *StatsdPluginMetrics) GetStatType() StatsdPluginStatType {
	if o == nil {
		return StatsdPluginStatType("")
	}
	return o.StatType
}

func (o *StatsdPluginMetrics) GetWorkspaceIdentifier() *StatsdPluginWorkspaceIdentifier {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifier
}

// StatsdPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type StatsdPluginConcurrencyLimit int64

const (
	StatsdPluginConcurrencyLimitMinus1 StatsdPluginConcurrencyLimit = -1
	StatsdPluginConcurrencyLimitOne    StatsdPluginConcurrencyLimit = 1
)

func (e StatsdPluginConcurrencyLimit) ToPointer() *StatsdPluginConcurrencyLimit {
	return &e
}
func (e *StatsdPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = StatsdPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginConcurrencyLimit: %v", v)
	}
}

type StatsdPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *StatsdPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
	// Time in seconds before the initial retry is made for a failing batch.
	InitialRetryDelay *float64 `json:"initial_retry_delay,omitempty"`
	// Maximum number of entries that can be processed at a time.
	MaxBatchSize *int64 `json:"max_batch_size,omitempty"`
	// Maximum number of bytes that can be waiting on a queue, requires string content.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
	MaxCoalescingDelay *float64 `json:"max_coalescing_delay,omitempty"`
	// Maximum number of entries that can be waiting on the queue.
	MaxEntries *int64 `json:"max_entries,omitempty"`
	// Maximum time in seconds between retries, caps exponential backoff.
	MaxRetryDelay *float64 `json:"max_retry_delay,omitempty"`
	// Time in seconds before the queue gives up calling a failed handler for a batch.
	MaxRetryTime *float64 `json:"max_retry_time,omitempty"`
}

func (o *StatsdPluginQueue) GetConcurrencyLimit() *StatsdPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *StatsdPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *StatsdPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *StatsdPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *StatsdPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *StatsdPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *StatsdPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *StatsdPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type StatsdPluginServiceIdentifierDefault string

const (
	StatsdPluginServiceIdentifierDefaultServiceID         StatsdPluginServiceIdentifierDefault = "service_id"
	StatsdPluginServiceIdentifierDefaultServiceName       StatsdPluginServiceIdentifierDefault = "service_name"
	StatsdPluginServiceIdentifierDefaultServiceHost       StatsdPluginServiceIdentifierDefault = "service_host"
	StatsdPluginServiceIdentifierDefaultServiceNameOrHost StatsdPluginServiceIdentifierDefault = "service_name_or_host"
)

func (e StatsdPluginServiceIdentifierDefault) ToPointer() *StatsdPluginServiceIdentifierDefault {
	return &e
}
func (e *StatsdPluginServiceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = StatsdPluginServiceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginServiceIdentifierDefault: %v", v)
	}
}

type StatsdPluginTagStyle string

const (
	StatsdPluginTagStyleDogstatsd StatsdPluginTagStyle = "dogstatsd"
	StatsdPluginTagStyleInfluxdb  StatsdPluginTagStyle = "influxdb"
	StatsdPluginTagStyleLibrato   StatsdPluginTagStyle = "librato"
	StatsdPluginTagStyleSignalfx  StatsdPluginTagStyle = "signalfx"
)

func (e StatsdPluginTagStyle) ToPointer() *StatsdPluginTagStyle {
	return &e
}
func (e *StatsdPluginTagStyle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dogstatsd":
		fallthrough
	case "influxdb":
		fallthrough
	case "librato":
		fallthrough
	case "signalfx":
		*e = StatsdPluginTagStyle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginTagStyle: %v", v)
	}
}

type StatsdPluginWorkspaceIdentifierDefault string

const (
	StatsdPluginWorkspaceIdentifierDefaultWorkspaceID   StatsdPluginWorkspaceIdentifierDefault = "workspace_id"
	StatsdPluginWorkspaceIdentifierDefaultWorkspaceName StatsdPluginWorkspaceIdentifierDefault = "workspace_name"
)

func (e StatsdPluginWorkspaceIdentifierDefault) ToPointer() *StatsdPluginWorkspaceIdentifierDefault {
	return &e
}
func (e *StatsdPluginWorkspaceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = StatsdPluginWorkspaceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginWorkspaceIdentifierDefault: %v", v)
	}
}

type StatsdPluginConfig struct {
	// List of status code ranges that are allowed to be logged in metrics.
	AllowStatusCodes          []string                               `json:"allow_status_codes,omitempty"`
	ConsumerIdentifierDefault *StatsdPluginConsumerIdentifierDefault `json:"consumer_identifier_default,omitempty"`
	FlushTimeout              *float64                               `json:"flush_timeout,omitempty"`
	// The IP address or hostname of StatsD server to send data to.
	Host             *string `json:"host,omitempty"`
	HostnameInPrefix *bool   `json:"hostname_in_prefix,omitempty"`
	// List of metrics to be logged.
	Metrics []StatsdPluginMetrics `json:"metrics,omitempty"`
	// The port of StatsD server to send data to.
	Port *int64 `json:"port,omitempty"`
	// String to prefix to each metric's name.
	Prefix                     *string                                 `json:"prefix,omitempty"`
	Queue                      *StatsdPluginQueue                      `json:"queue,omitempty"`
	QueueSize                  *int64                                  `json:"queue_size,omitempty"`
	RetryCount                 *int64                                  `json:"retry_count,omitempty"`
	ServiceIdentifierDefault   *StatsdPluginServiceIdentifierDefault   `json:"service_identifier_default,omitempty"`
	TagStyle                   *StatsdPluginTagStyle                   `json:"tag_style,omitempty"`
	UDPPacketSize              *float64                                `json:"udp_packet_size,omitempty"`
	UseTCP                     *bool                                   `json:"use_tcp,omitempty"`
	WorkspaceIdentifierDefault *StatsdPluginWorkspaceIdentifierDefault `json:"workspace_identifier_default,omitempty"`
}

func (o *StatsdPluginConfig) GetAllowStatusCodes() []string {
	if o == nil {
		return nil
	}
	return o.AllowStatusCodes
}

func (o *StatsdPluginConfig) GetConsumerIdentifierDefault() *StatsdPluginConsumerIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifierDefault
}

func (o *StatsdPluginConfig) GetFlushTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushTimeout
}

func (o *StatsdPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *StatsdPluginConfig) GetHostnameInPrefix() *bool {
	if o == nil {
		return nil
	}
	return o.HostnameInPrefix
}

func (o *StatsdPluginConfig) GetMetrics() []StatsdPluginMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *StatsdPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *StatsdPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *StatsdPluginConfig) GetQueue() *StatsdPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *StatsdPluginConfig) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *StatsdPluginConfig) GetRetryCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryCount
}

func (o *StatsdPluginConfig) GetServiceIdentifierDefault() *StatsdPluginServiceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifierDefault
}

func (o *StatsdPluginConfig) GetTagStyle() *StatsdPluginTagStyle {
	if o == nil {
		return nil
	}
	return o.TagStyle
}

func (o *StatsdPluginConfig) GetUDPPacketSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPPacketSize
}

func (o *StatsdPluginConfig) GetUseTCP() *bool {
	if o == nil {
		return nil
	}
	return o.UseTCP
}

func (o *StatsdPluginConfig) GetWorkspaceIdentifierDefault() *StatsdPluginWorkspaceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifierDefault
}

type StatsdPluginProtocols string

const (
	StatsdPluginProtocolsGrpc           StatsdPluginProtocols = "grpc"
	StatsdPluginProtocolsGrpcs          StatsdPluginProtocols = "grpcs"
	StatsdPluginProtocolsHTTP           StatsdPluginProtocols = "http"
	StatsdPluginProtocolsHTTPS          StatsdPluginProtocols = "https"
	StatsdPluginProtocolsTCP            StatsdPluginProtocols = "tcp"
	StatsdPluginProtocolsTLS            StatsdPluginProtocols = "tls"
	StatsdPluginProtocolsTLSPassthrough StatsdPluginProtocols = "tls_passthrough"
	StatsdPluginProtocolsUDP            StatsdPluginProtocols = "udp"
	StatsdPluginProtocolsWs             StatsdPluginProtocols = "ws"
	StatsdPluginProtocolsWss            StatsdPluginProtocols = "wss"
)

func (e StatsdPluginProtocols) ToPointer() *StatsdPluginProtocols {
	return &e
}
func (e *StatsdPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = StatsdPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginProtocols: %v", v)
	}
}

// StatsdPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type StatsdPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type StatsdPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type StatsdPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type StatsdPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type StatsdPlugin struct {
	Config *StatsdPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"statsd" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []StatsdPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *StatsdPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *StatsdPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *StatsdPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *StatsdPluginService `json:"service,omitempty"`
}

func (s StatsdPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StatsdPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StatsdPlugin) GetConfig() *StatsdPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *StatsdPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *StatsdPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *StatsdPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StatsdPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *StatsdPlugin) GetName() *string {
	return types.String("statsd")
}

func (o *StatsdPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *StatsdPlugin) GetProtocols() []StatsdPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *StatsdPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *StatsdPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *StatsdPlugin) GetConsumer() *StatsdPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *StatsdPlugin) GetConsumerGroup() *StatsdPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *StatsdPlugin) GetRoute() *StatsdPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *StatsdPlugin) GetService() *StatsdPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
