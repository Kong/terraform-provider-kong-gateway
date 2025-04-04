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
var _ datasource.DataSource = &PluginMockingDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginMockingDataSource{}

func NewPluginMockingDataSource() datasource.DataSource {
	return &PluginMockingDataSource{}
}

// PluginMockingDataSource is the data source implementation.
type PluginMockingDataSource struct {
	client *sdk.KongGateway
}

// PluginMockingDataSourceModel describes the data model.
type PluginMockingDataSourceModel struct {
	Config       *tfTypes.MockingPluginConfig       `tfsdk:"config"`
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
func (r *PluginMockingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_mocking"
}

// Schema defines the schema for the data source.
func (r *PluginMockingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginMocking DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_specification": schema.StringAttribute{
						Computed:    true,
						Description: `The contents of the specification file. You must use this option for hybrid or DB-less mode. You can include the full specification as part of the configuration. In Kong Manager, you can copy and paste the contents of the spec directly into the ` + "`" + `Config.Api Specification` + "`" + ` text field.`,
					},
					"api_specification_filename": schema.StringAttribute{
						Computed:    true,
						Description: `The path and name of the specification file loaded into Kong Gateway's database. You cannot use this option for DB-less or hybrid mode.`,
					},
					"custom_base_path": schema.StringAttribute{
						Computed:    true,
						Description: `The base path to be used for path match evaluation. This value is ignored if ` + "`" + `include_base_path` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
					},
					"include_base_path": schema.BoolAttribute{
						Computed:    true,
						Description: `Indicates whether to include the base path when performing path match evaluation.`,
					},
					"included_status_codes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.Int64Type,
						Description: `A global list of the HTTP status codes that can only be selected and returned.`,
					},
					"max_delay_time": schema.NumberAttribute{
						Computed:    true,
						Description: `The maximum value in seconds of delay time. Set this value when ` + "`" + `random_delay` + "`" + ` is enabled and you want to adjust the default. The value must be greater than the ` + "`" + `min_delay_time` + "`" + `.`,
					},
					"min_delay_time": schema.NumberAttribute{
						Computed:    true,
						Description: `The minimum value in seconds of delay time. Set this value when ` + "`" + `random_delay` + "`" + ` is enabled and you want to adjust the default. The value must be less than the ` + "`" + `max_delay_time` + "`" + `.`,
					},
					"random_delay": schema.BoolAttribute{
						Computed:    true,
						Description: `Enables a random delay in the mocked response. Introduces delays to simulate real-time response times by APIs.`,
					},
					"random_examples": schema.BoolAttribute{
						Computed:    true,
						Description: `Randomly selects one example and returns it. This parameter requires the spec to have multiple examples configured.`,
					},
					"random_status_code": schema.BoolAttribute{
						Computed:    true,
						Description: `Determines whether to randomly select an HTTP status code from the responses of the corresponding API method. The default value is ` + "`" + `false` + "`" + `, which means the minimum HTTP status code is always selected and returned.`,
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

func (r *PluginMockingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginMockingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginMockingDataSourceModel
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

	request := operations.GetMockingPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetMockingPlugin(ctx, request)
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
	if !(res.MockingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMockingPlugin(res.MockingPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
