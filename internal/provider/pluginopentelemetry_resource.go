// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-gateway/internal/validators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-kong-gateway/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PluginOpentelemetryResource{}
var _ resource.ResourceWithImportState = &PluginOpentelemetryResource{}

func NewPluginOpentelemetryResource() resource.Resource {
	return &PluginOpentelemetryResource{}
}

// PluginOpentelemetryResource defines the resource implementation.
type PluginOpentelemetryResource struct {
	client *sdk.KongGateway
}

// PluginOpentelemetryResourceModel describes the resource data model.
type PluginOpentelemetryResourceModel struct {
	Config        *tfTypes.CreateOpentelemetryPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                     `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                     `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                              `tfsdk:"created_at"`
	Enabled       types.Bool                               `tfsdk:"enabled"`
	ID            types.String                             `tfsdk:"id"`
	InstanceName  types.String                             `tfsdk:"instance_name"`
	Ordering      types.String                             `tfsdk:"ordering"`
	Protocols     []types.String                           `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                     `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                     `tfsdk:"service"`
	Tags          []types.String                           `tfsdk:"tags"`
	UpdatedAt     types.Int64                              `tfsdk:"updated_at"`
}

func (r *PluginOpentelemetryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_opentelemetry"
}

func (r *PluginOpentelemetryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginOpentelemetry Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"batch_flush_delay": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The delay, in seconds, between two consecutive batches.`,
					},
					"batch_span_count": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of spans to be sent in a single batch.`,
					},
					"connect_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
					},
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"header_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["preserve", "ignore", "b3", "b3-single", "w3c", "jaeger", "ot", "aws", "gcp", "datadog"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"preserve",
								"ignore",
								"b3",
								"b3-single",
								"w3c",
								"jaeger",
								"ot",
								"aws",
								"gcp",
								"datadog",
							),
						},
					},
					"headers": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.`,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"http_response_header_for_traceid": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"propagation": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"clear": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Header names to clear after context extraction. This allows to extract the context from a certain header and then remove it from the request, useful when extraction and injection are performed on different header formats and the original header should not be sent to the upstream. If left empty, no headers are cleared.`,
							},
							"default_format": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The default header format to use when extractors did not match any format in the incoming headers and ` + "`" + `inject` + "`" + ` is configured with the value: ` + "`" + `preserve` + "`" + `. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in ` + "`" + `extract` + "`" + `. Not Null; must be one of ["b3", "gcp", "b3-single", "jaeger", "aws", "ot", "w3c", "datadog"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"b3",
										"gcp",
										"b3-single",
										"jaeger",
										"aws",
										"ot",
										"w3c",
										"datadog",
									),
								},
							},
							"extract": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Header formats used to extract tracing context from incoming requests. If multiple values are specified, the first one found will be used for extraction. If left empty, Kong will not extract any tracing context information from incoming requests and generate a trace with no parent and a new trace ID.`,
							},
							"inject": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Header formats used to inject tracing context. The value ` + "`" + `preserve` + "`" + ` will use the same header format as the incoming request. If multiple values are specified, all of them will be used during injection. If left empty, Kong will not inject any tracing context information in outgoing requests.`,
							},
						},
					},
					"queue": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"initial_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Time in seconds before the initial retry is made for a failing batch.`,
							},
							"max_batch_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of entries that can be processed at a time.`,
							},
							"max_bytes": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of bytes that can be waiting on a queue, requires string content.`,
							},
							"max_coalescing_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of entries that can be waiting on the queue.`,
							},
							"max_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum time in seconds between retries, caps exponential backoff.`,
							},
							"max_retry_time": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch.`,
							},
						},
					},
					"read_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
					},
					"resource_attributes": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"sampling_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global ` + "`" + `tracing_sampling_rate` + "`" + ` setting from kong.conf.`,
					},
					"send_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Plugin to lookup`,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
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

func (r *PluginOpentelemetryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongGateway)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.KongGateway, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PluginOpentelemetryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PluginOpentelemetryResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedCreateOpentelemetryPlugin()
	res, err := r.client.Plugins.CreateOpentelemetryPlugin(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.OpentelemetryPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginOpentelemetryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PluginOpentelemetryResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	pluginID := data.ID.ValueString()
	request := operations.GetOpentelemetryPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetOpentelemetryPlugin(ctx, request)
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
	if !(res.OpentelemetryPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginOpentelemetryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PluginOpentelemetryResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	pluginID := data.ID.ValueString()
	createOpentelemetryPlugin := data.ToSharedCreateOpentelemetryPlugin()
	request := operations.UpdateOpentelemetryPluginRequest{
		PluginID:                  pluginID,
		CreateOpentelemetryPlugin: createOpentelemetryPlugin,
	}
	res, err := r.client.Plugins.UpdateOpentelemetryPlugin(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.OpentelemetryPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginOpentelemetryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PluginOpentelemetryResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	pluginID := data.ID.ValueString()
	request := operations.DeleteOpentelemetryPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.DeleteOpentelemetryPlugin(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *PluginOpentelemetryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}