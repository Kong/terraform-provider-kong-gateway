// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ClientCertificate - Certificate to be used as client certificate while TLS handshaking to the upstream server.
type ClientCertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *ClientCertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// Protocol - The protocol used to communicate with the upstream.
type Protocol string

const (
	ProtocolGrpc           Protocol = "grpc"
	ProtocolGrpcs          Protocol = "grpcs"
	ProtocolHTTP           Protocol = "http"
	ProtocolHTTPS          Protocol = "https"
	ProtocolTCP            Protocol = "tcp"
	ProtocolTLS            Protocol = "tls"
	ProtocolTLSPassthrough Protocol = "tls_passthrough"
	ProtocolUDP            Protocol = "udp"
	ProtocolWs             Protocol = "ws"
	ProtocolWss            Protocol = "wss"
)

func (e Protocol) ToPointer() *Protocol {
	return &e
}
func (e *Protocol) UnmarshalJSON(data []byte) error {
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
		*e = Protocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Protocol: %v", v)
	}
}

type Service struct {
	// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
	// The timeout in milliseconds for establishing a connection to the upstream server.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// The host of the upstream server. Note that the host value is case sensitive.
	Host *string `json:"host,omitempty"`
	ID   *string `json:"id,omitempty"`
	// The Service name.
	Name *string `json:"name,omitempty"`
	// The path to be used in requests to the upstream server.
	Path *string `json:"path,omitempty"`
	// The upstream server port.
	Port *int64 `json:"port,omitempty"`
	// The protocol used to communicate with the upstream.
	Protocol *Protocol `json:"protocol,omitempty"`
	// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// The number of retries to execute upon failure to proxy.
	Retries *int64 `json:"retries,omitempty"`
	// An optional set of strings associated with the Service for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerify *bool `json:"tls_verify,omitempty"`
	// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerifyDepth *int64 `json:"tls_verify_depth,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
	WriteTimeout *int64 `json:"write_timeout,omitempty"`
}

func (o *Service) GetCaCertificates() []string {
	if o == nil {
		return nil
	}
	return o.CaCertificates
}

func (o *Service) GetClientCertificate() *ClientCertificate {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *Service) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *Service) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Service) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Service) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *Service) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Service) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Service) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *Service) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Service) GetProtocol() *Protocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *Service) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *Service) GetRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *Service) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Service) GetTLSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.TLSVerify
}

func (o *Service) GetTLSVerifyDepth() *int64 {
	if o == nil {
		return nil
	}
	return o.TLSVerifyDepth
}

func (o *Service) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Service) GetWriteTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteTimeout
}

type ServiceInput struct {
	// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
	// The timeout in milliseconds for establishing a connection to the upstream server.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// The host of the upstream server. Note that the host value is case sensitive.
	Host *string `json:"host,omitempty"`
	// The Service name.
	Name *string `json:"name,omitempty"`
	// The path to be used in requests to the upstream server.
	Path *string `json:"path,omitempty"`
	// The upstream server port.
	Port *int64 `json:"port,omitempty"`
	// The protocol used to communicate with the upstream.
	Protocol *Protocol `json:"protocol,omitempty"`
	// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// The number of retries to execute upon failure to proxy.
	Retries *int64 `json:"retries,omitempty"`
	// An optional set of strings associated with the Service for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerify *bool `json:"tls_verify,omitempty"`
	// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerifyDepth *int64 `json:"tls_verify_depth,omitempty"`
	// Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.
	URL *string `json:"url,omitempty"`
	// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
	WriteTimeout *int64 `json:"write_timeout,omitempty"`
}

func (o *ServiceInput) GetCaCertificates() []string {
	if o == nil {
		return nil
	}
	return o.CaCertificates
}

func (o *ServiceInput) GetClientCertificate() *ClientCertificate {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *ServiceInput) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *ServiceInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ServiceInput) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *ServiceInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ServiceInput) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *ServiceInput) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ServiceInput) GetProtocol() *Protocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *ServiceInput) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *ServiceInput) GetRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *ServiceInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ServiceInput) GetTLSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.TLSVerify
}

func (o *ServiceInput) GetTLSVerifyDepth() *int64 {
	if o == nil {
		return nil
	}
	return o.TLSVerifyDepth
}

func (o *ServiceInput) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *ServiceInput) GetWriteTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteTimeout
}
