// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

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
var _ datasource.DataSource = &PluginJWTDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginJWTDataSource{}

func NewPluginJWTDataSource() datasource.DataSource {
	return &PluginJWTDataSource{}
}

// PluginJWTDataSource is the data source implementation.
type PluginJWTDataSource struct {
	client *sdk.KongGateway
}

// PluginJWTDataSourceModel describes the data model.
type PluginJWTDataSourceModel struct {
	Config        *tfTypes.CreateJWTPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer           `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer           `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                    `tfsdk:"created_at"`
	Enabled       types.Bool                     `tfsdk:"enabled"`
	ID            types.String                   `tfsdk:"id"`
	InstanceName  types.String                   `tfsdk:"instance_name"`
	Ordering      types.String                   `tfsdk:"ordering"`
	Protocols     []types.String                 `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer           `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer           `tfsdk:"service"`
	Tags          []types.String                 `tfsdk:"tags"`
	UpdatedAt     types.Int64                    `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginJWTDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_jwt"
}

// Schema defines the schema for the data source.
func (r *PluginJWTDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginJWT DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.`,
					},
					"claims_to_verify": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of registered claims (according to RFC 7519) that Kong can verify as well. Accepted values: one of exp or nbf.`,
					},
					"cookie_names": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of cookie names that Kong will inspect to retrieve JWTs.`,
					},
					"header_names": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of HTTP header names that Kong will inspect to retrieve JWTs.`,
					},
					"key_claim_name": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the claim in which the key identifying the secret must be passed. The plugin will attempt to read this claim from the JWT payload and the header, in that order.`,
					},
					"maximum_expiration": schema.NumberAttribute{
						Computed:    true,
						Description: `A value between 0 and 31536000 (365 days) limiting the lifetime of the JWT to maximum_expiration seconds in the future.`,
					},
					"run_on_preflight": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that indicates whether the plugin should run (and try to authenticate) on OPTIONS preflight requests. If set to false, then OPTIONS requests will always be allowed.`,
					},
					"secret_is_base64": schema.BoolAttribute{
						Computed:    true,
						Description: `If true, the plugin assumes the credential’s secret to be base64 encoded. You will need to create a base64-encoded secret for your Consumer, and sign your JWT with the original secret.`,
					},
					"uri_param_names": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of querystring parameters that Kong will inspect to retrieve JWTs.`,
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
				Required:    true,
				Description: `ID of the Plugin to lookup`,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
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

func (r *PluginJWTDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginJWTDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginJWTDataSourceModel
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

	pluginID := data.ID.ValueString()
	request := operations.GetJwtPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetJwtPlugin(ctx, request)
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
	if !(res.JWTPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedJWTPlugin(res.JWTPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}