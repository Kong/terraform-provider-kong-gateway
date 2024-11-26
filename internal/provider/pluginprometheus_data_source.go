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
var _ datasource.DataSource = &PluginPrometheusDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginPrometheusDataSource{}

func NewPluginPrometheusDataSource() datasource.DataSource {
	return &PluginPrometheusDataSource{}
}

// PluginPrometheusDataSource is the data source implementation.
type PluginPrometheusDataSource struct {
	client *sdk.KongGateway
}

// PluginPrometheusDataSourceModel describes the data model.
type PluginPrometheusDataSourceModel struct {
	Config        tfTypes.PrometheusPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer           `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer           `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                    `tfsdk:"created_at"`
	Enabled       types.Bool                     `tfsdk:"enabled"`
	ID            types.String                   `tfsdk:"id"`
	InstanceName  types.String                   `tfsdk:"instance_name"`
	Ordering      *tfTypes.ACLPluginOrdering     `tfsdk:"ordering"`
	Protocols     []types.String                 `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer           `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer           `tfsdk:"service"`
	Tags          []types.String                 `tfsdk:"tags"`
	UpdatedAt     types.Int64                    `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginPrometheusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_prometheus"
}

// Schema defines the schema for the data source.
func (r *PluginPrometheusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginPrometheus DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"ai_metrics": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if ai metrics should be collected. If enabled, the ` + "`" + `ai_llm_requests_total` + "`" + `, ` + "`" + `ai_llm_cost_total` + "`" + ` and ` + "`" + `ai_llm_tokens_total` + "`" + ` metrics will be exported.`,
					},
					"bandwidth_metrics": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if bandwidth metrics should be collected. If enabled, ` + "`" + `bandwidth_bytes` + "`" + ` and ` + "`" + `stream_sessions_total` + "`" + ` metrics will be exported.`,
					},
					"latency_metrics": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if latency metrics should be collected. If enabled, ` + "`" + `kong_latency_ms` + "`" + `, ` + "`" + `upstream_latency_ms` + "`" + ` and ` + "`" + `request_latency_ms` + "`" + ` metrics will be exported.`,
					},
					"per_consumer": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if per-consumer metrics should be collected. If enabled, the ` + "`" + `kong_http_requests_total` + "`" + ` and ` + "`" + `kong_bandwidth_bytes` + "`" + ` metrics fill in the consumer label when available.`,
					},
					"status_code_metrics": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if status code metrics should be collected. If enabled, ` + "`" + `http_requests_total` + "`" + `, ` + "`" + `stream_sessions_total` + "`" + ` metrics will be exported.`,
					},
					"upstream_health_metrics": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that determines if upstream metrics should be collected. If enabled, ` + "`" + `upstream_target_health` + "`" + ` metric will be exported.`,
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
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
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
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
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

func (r *PluginPrometheusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginPrometheusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginPrometheusDataSourceModel
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

	request := operations.GetPrometheusPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetPrometheusPlugin(ctx, request)
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
	if !(res.PrometheusPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrometheusPlugin(res.PrometheusPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
