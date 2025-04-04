// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PluginKafkaUpstreamDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginKafkaUpstreamDataSource{}

func NewPluginKafkaUpstreamDataSource() datasource.DataSource {
	return &PluginKafkaUpstreamDataSource{}
}

// PluginKafkaUpstreamDataSource is the data source implementation.
type PluginKafkaUpstreamDataSource struct {
	client *sdk.KongGateway
}

// PluginKafkaUpstreamDataSourceModel describes the data model.
type PluginKafkaUpstreamDataSourceModel struct {
	Config       *tfTypes.KafkaUpstreamPluginConfig `tfsdk:"config"`
	Consumer     *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
	CreatedAt    types.Int64                        `tfsdk:"created_at"`
	Enabled      types.Bool                         `tfsdk:"enabled"`
	ID           types.String                       `tfsdk:"id"`
	InstanceName types.String                       `tfsdk:"instance_name"`
	Ordering     *tfTypes.Ordering                  `tfsdk:"ordering"`
	Partials     []tfTypes.Partials                 `tfsdk:"partials"`
	Protocols    []types.String                     `tfsdk:"protocols"`
	Route        *tfTypes.ACLWithoutParentsConsumer `tfsdk:"route"`
	Service      *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service"`
	Tags         []types.String                     `tfsdk:"tags"`
	UpdatedAt    types.Int64                        `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginKafkaUpstreamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_kafka_upstream"
}

// Schema defines the schema for the data source.
func (r *PluginKafkaUpstreamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginKafkaUpstream DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"mechanism": schema.StringAttribute{
								Computed:    true,
								Description: `The SASL authentication mechanism.  Supported options: ` + "`" + `PLAIN` + "`" + `, ` + "`" + `SCRAM-SHA-256` + "`" + `, or ` + "`" + `SCRAM-SHA-512` + "`" + `.`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Description: `Password for SASL authentication.`,
							},
							"strategy": schema.StringAttribute{
								Computed:    true,
								Description: `The authentication strategy for the plugin, the only option for the value is ` + "`" + `sasl` + "`" + `.`,
							},
							"tokenauth": schema.BoolAttribute{
								Computed:    true,
								Description: `Enable this to indicate ` + "`" + `DelegationToken` + "`" + ` authentication.`,
							},
							"user": schema.StringAttribute{
								Computed:    true,
								Description: `Username for SASL authentication.`,
							},
						},
					},
					"bootstrap_servers": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Computed:    true,
									Description: `A string representing a host name, such as example.com.`,
								},
								"port": schema.Int64Attribute{
									Computed:    true,
									Description: `An integer representing a port number between 0 and 65535, inclusive.`,
								},
							},
						},
						Description: `Set of bootstrap brokers in a ` + "`" + `{host: host, port: port}` + "`" + ` list format.`,
					},
					"cluster_name": schema.StringAttribute{
						Computed:    true,
						Description: `An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a ` + "`" + `cluster_name` + "`" + ` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.`,
					},
					"forward_body": schema.BoolAttribute{
						Computed:    true,
						Description: `Include the request body in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_headers": schema.BoolAttribute{
						Computed:    true,
						Description: `Include the request headers in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_method": schema.BoolAttribute{
						Computed:    true,
						Description: `Include the request method in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"forward_uri": schema.BoolAttribute{
						Computed:    true,
						Description: `Include the request URI and URI arguments (as in, query arguments) in the message. At least one of these must be true: ` + "`" + `forward_method` + "`" + `, ` + "`" + `forward_uri` + "`" + `, ` + "`" + `forward_headers` + "`" + `, ` + "`" + `forward_body` + "`" + `.`,
					},
					"keepalive": schema.Int64Attribute{
						Computed:    true,
						Description: `Keepalive timeout in milliseconds.`,
					},
					"keepalive_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"producer_async": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to enable asynchronous mode.`,
					},
					"producer_async_buffering_limits_messages_in_memory": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of messages that can be buffered in memory in asynchronous mode.`,
					},
					"producer_async_flush_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum time interval in milliseconds between buffer flushes in asynchronous mode.`,
					},
					"producer_request_acks": schema.Int64Attribute{
						Computed:    true,
						Description: `The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).`,
					},
					"producer_request_limits_bytes_per_request": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of a Produce request in bytes.`,
					},
					"producer_request_limits_messages_per_request": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of messages to include into a single producer request.`,
					},
					"producer_request_retries_backoff_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Backoff interval between retry attempts in milliseconds.`,
					},
					"producer_request_retries_max_attempts": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of retry attempts per single Produce request.`,
					},
					"producer_request_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Time to wait for a Produce response in milliseconds.`,
					},
					"security": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"certificate_id": schema.StringAttribute{
								Computed:    true,
								Description: `UUID of certificate entity for mTLS authentication.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Description: `Enables TLS.`,
							},
						},
					},
					"timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Socket timeout in milliseconds.`,
					},
					"topic": schema.StringAttribute{
						Computed:    true,
						Description: `The Kafka topic to publish to.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"partials": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"path": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A set of strings representing HTTP protocols.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *PluginKafkaUpstreamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongGateway)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.KongGateway, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PluginKafkaUpstreamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginKafkaUpstreamDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	request := operations.GetKafkaupstreamPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetKafkaupstreamPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.KafkaUpstreamPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKafkaUpstreamPlugin(res.KafkaUpstreamPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
