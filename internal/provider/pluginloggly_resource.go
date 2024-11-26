// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PluginLogglyResource{}
var _ resource.ResourceWithImportState = &PluginLogglyResource{}

func NewPluginLogglyResource() resource.Resource {
	return &PluginLogglyResource{}
}

// PluginLogglyResource defines the resource implementation.
type PluginLogglyResource struct {
	client *sdk.KongGateway
}

// PluginLogglyResourceModel describes the resource data model.
type PluginLogglyResourceModel struct {
	Config        tfTypes.LogglyPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer       `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer       `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                `tfsdk:"created_at"`
	Enabled       types.Bool                 `tfsdk:"enabled"`
	ID            types.String               `tfsdk:"id"`
	InstanceName  types.String               `tfsdk:"instance_name"`
	Ordering      *tfTypes.ACLPluginOrdering `tfsdk:"ordering"`
	Protocols     []types.String             `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer       `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer       `tfsdk:"service"`
	Tags          []types.String             `tfsdk:"tags"`
	UpdatedAt     types.Int64                `tfsdk:"updated_at"`
}

func (r *PluginLogglyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_loggly"
}

func (r *PluginLogglyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginLoggly Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client_errors_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["debug", "info", "notice", "warning", "err", "crit", "alert", "emerg"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"debug",
								"info",
								"notice",
								"warning",
								"err",
								"crit",
								"alert",
								"emerg",
							),
						},
					},
					"custom_fields_by_lua": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `Lua code as a key-value map`,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"log_level": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["debug", "info", "notice", "warning", "err", "crit", "alert", "emerg"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"debug",
								"info",
								"notice",
								"warning",
								"err",
								"crit",
								"alert",
								"emerg",
							),
						},
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
						Validators: []validator.Int64{
							int64validator.AtMost(65535),
						},
					},
					"server_errors_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["debug", "info", "notice", "warning", "err", "crit", "alert", "emerg"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"debug",
								"info",
								"notice",
								"warning",
								"err",
								"crit",
								"alert",
								"emerg",
							),
						},
					},
					"successful_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["debug", "info", "notice", "warning", "err", "crit", "alert", "emerg"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"debug",
								"info",
								"notice",
								"warning",
								"err",
								"crit",
								"alert",
								"emerg",
							),
						},
					},
					"tags": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"timeout": schema.NumberAttribute{
						Computed: true,
						Optional: true,
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
				Computed: true,
				Optional: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
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

func (r *PluginLogglyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PluginLogglyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PluginLogglyResourceModel
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

	request := data.ToSharedLogglyPluginInput()
	res, err := r.client.Plugins.CreateLogglyPlugin(ctx, request)
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
	if !(res.LogglyPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLogglyPlugin(res.LogglyPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginLogglyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PluginLogglyResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	request := operations.GetLogglyPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetLogglyPlugin(ctx, request)
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
	if !(res.LogglyPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLogglyPlugin(res.LogglyPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginLogglyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PluginLogglyResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	logglyPlugin := data.ToSharedLogglyPluginInput()
	request := operations.UpdateLogglyPluginRequest{
		PluginID:     pluginID,
		LogglyPlugin: logglyPlugin,
	}
	res, err := r.client.Plugins.UpdateLogglyPlugin(ctx, request)
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
	if !(res.LogglyPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLogglyPlugin(res.LogglyPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginLogglyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PluginLogglyResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	request := operations.DeleteLogglyPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.DeleteLogglyPlugin(ctx, request)
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

func (r *PluginLogglyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
