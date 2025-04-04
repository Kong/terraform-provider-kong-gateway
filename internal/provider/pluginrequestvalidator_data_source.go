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
var _ datasource.DataSource = &PluginRequestValidatorDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginRequestValidatorDataSource{}

func NewPluginRequestValidatorDataSource() datasource.DataSource {
	return &PluginRequestValidatorDataSource{}
}

// PluginRequestValidatorDataSource is the data source implementation.
type PluginRequestValidatorDataSource struct {
	client *sdk.KongGateway
}

// PluginRequestValidatorDataSourceModel describes the data model.
type PluginRequestValidatorDataSourceModel struct {
	Config       *tfTypes.RequestValidatorPluginConfig `tfsdk:"config"`
	Consumer     *tfTypes.ACLWithoutParentsConsumer    `tfsdk:"consumer"`
	CreatedAt    types.Int64                           `tfsdk:"created_at"`
	Enabled      types.Bool                            `tfsdk:"enabled"`
	ID           types.String                          `tfsdk:"id"`
	InstanceName types.String                          `tfsdk:"instance_name"`
	Ordering     *tfTypes.Ordering                     `tfsdk:"ordering"`
	Partials     []tfTypes.Partials                    `tfsdk:"partials"`
	Protocols    []types.String                        `tfsdk:"protocols"`
	Route        *tfTypes.ACLWithoutParentsConsumer    `tfsdk:"route"`
	Service      *tfTypes.ACLWithoutParentsConsumer    `tfsdk:"service"`
	Tags         []types.String                        `tfsdk:"tags"`
	UpdatedAt    types.Int64                           `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginRequestValidatorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_request_validator"
}

// Schema defines the schema for the data source.
func (r *PluginRequestValidatorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginRequestValidator DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed_content_types": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of allowed content types. The value can be configured with the ` + "`" + `charset` + "`" + ` parameter. For example, ` + "`" + `application/json; charset=UTF-8` + "`" + `.`,
					},
					"body_schema": schema.StringAttribute{
						Computed:    true,
						Description: `The request body schema specification. One of ` + "`" + `body_schema` + "`" + ` or ` + "`" + `parameter_schema` + "`" + ` must be specified.`,
					},
					"content_type_parameter_validation": schema.BoolAttribute{
						Computed:    true,
						Description: `Determines whether to enable parameters validation of request content-type.`,
					},
					"parameter_schema": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"explode": schema.BoolAttribute{
									Computed:    true,
									Description: `Required when ` + "`" + `schema` + "`" + ` and ` + "`" + `style` + "`" + ` are set. When ` + "`" + `explode` + "`" + ` is ` + "`" + `true` + "`" + `, parameter values of type ` + "`" + `array` + "`" + ` or ` + "`" + `object` + "`" + ` generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters, this property has no effect.`,
								},
								"in": schema.StringAttribute{
									Computed:    true,
									Description: `The location of the parameter.`,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `The name of the parameter. Parameter names are case-sensitive, and correspond to the parameter name used by the ` + "`" + `in` + "`" + ` property. If ` + "`" + `in` + "`" + ` is ` + "`" + `path` + "`" + `, the ` + "`" + `name` + "`" + ` field MUST correspond to the named capture group from the configured ` + "`" + `route` + "`" + `.`,
								},
								"required": schema.BoolAttribute{
									Computed:    true,
									Description: `Determines whether this parameter is mandatory.`,
								},
								"schema": schema.StringAttribute{
									Computed:    true,
									Description: `Requred when ` + "`" + `style` + "`" + ` and ` + "`" + `explode` + "`" + ` are set. This is the schema defining the type used for the parameter. It is validated using ` + "`" + `draft4` + "`" + ` for JSON Schema draft 4 compliant validator. In addition to being a valid JSON Schema, the parameter schema MUST have a top-level ` + "`" + `type` + "`" + ` property to enable proper deserialization before validating.`,
								},
								"style": schema.StringAttribute{
									Computed:    true,
									Description: `Required when ` + "`" + `schema` + "`" + ` and ` + "`" + `explode` + "`" + ` are set. Describes how the parameter value will be deserialized depending on the type of the parameter value.`,
								},
							},
						},
						Description: `Array of parameter validator specification. One of ` + "`" + `body_schema` + "`" + ` or ` + "`" + `parameter_schema` + "`" + ` must be specified.`,
					},
					"verbose_response": schema.BoolAttribute{
						Computed:    true,
						Description: `If enabled, the plugin returns more verbose and detailed validation errors.`,
					},
					"version": schema.StringAttribute{
						Computed:    true,
						Description: `Which validator to use. Supported values are ` + "`" + `kong` + "`" + ` (default) for using Kong's own schema validator, or ` + "`" + `draft4` + "`" + ` for using a JSON Schema Draft 4-compliant validator.`,
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

func (r *PluginRequestValidatorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginRequestValidatorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginRequestValidatorDataSourceModel
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

	request := operations.GetRequestvalidatorPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetRequestvalidatorPlugin(ctx, request)
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
	if !(res.RequestValidatorPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRequestValidatorPlugin(res.RequestValidatorPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
