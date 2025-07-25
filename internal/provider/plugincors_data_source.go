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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PluginCorsDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginCorsDataSource{}

func NewPluginCorsDataSource() datasource.DataSource {
	return &PluginCorsDataSource{}
}

// PluginCorsDataSource is the data source implementation.
type PluginCorsDataSource struct {
	// Provider configured SDK client.
	client *sdk.KongGateway
}

// PluginCorsDataSourceModel describes the data model.
type PluginCorsDataSourceModel struct {
	Config       *tfTypes.CorsPluginConfig          `tfsdk:"config"`
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
func (r *PluginCorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_cors"
}

// Schema defines the schema for the data source.
func (r *PluginCorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginCors DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allow_origin_absent": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that skip cors response headers when origin header of request is empty`,
					},
					"credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to determine whether the ` + "`" + `Access-Control-Allow-Credentials` + "`" + ` header should be sent with ` + "`" + `true` + "`" + ` as the value.`,
					},
					"exposed_headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Value for the ` + "`" + `Access-Control-Expose-Headers` + "`" + ` header. If not specified, no custom headers are exposed.`,
					},
					"headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Value for the ` + "`" + `Access-Control-Allow-Headers` + "`" + ` header.`,
					},
					"max_age": schema.Float64Attribute{
						Computed:    true,
						Description: `Indicates how long the results of the preflight request can be cached, in ` + "`" + `seconds` + "`" + `.`,
					},
					"methods": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `'Value for the ` + "`" + `Access-Control-Allow-Methods` + "`" + ` header. Available options include ` + "`" + `GET` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `PATCH` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `TRACE` + "`" + `, ` + "`" + `CONNECT` + "`" + `. By default, all options are allowed.'`,
					},
					"origins": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of allowed domains for the ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header. If you want to allow all origins, add ` + "`" + `*` + "`" + ` as a single value to this configuration field. The accepted values can either be flat strings or PCRE regexes.`,
					},
					"preflight_continue": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that instructs the plugin to proxy the ` + "`" + `OPTIONS` + "`" + ` preflight request to the Upstream service.`,
					},
					"private_network": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to determine whether the ` + "`" + `Access-Control-Allow-Private-Network` + "`" + ` header should be sent with ` + "`" + `true` + "`" + ` as the value.`,
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

func (r *PluginCorsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginCorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginCorsDataSourceModel
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

	request, requestDiags := data.ToOperationsGetCorsPluginRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Plugins.GetCorsPlugin(ctx, *request)
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
	if !(res.CorsPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedCorsPlugin(ctx, res.CorsPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
