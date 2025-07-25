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
var _ datasource.DataSource = &RouteExpressionDataSource{}
var _ datasource.DataSourceWithConfigure = &RouteExpressionDataSource{}

func NewRouteExpressionDataSource() datasource.DataSource {
	return &RouteExpressionDataSource{}
}

// RouteExpressionDataSource is the data source implementation.
type RouteExpressionDataSource struct {
	// Provider configured SDK client.
	client *sdk.KongGateway
}

// RouteExpressionDataSourceModel describes the data model.
type RouteExpressionDataSourceModel struct {
	CreatedAt               types.Int64                        `tfsdk:"created_at"`
	Expression              types.String                       `tfsdk:"expression"`
	HTTPSRedirectStatusCode types.Int64                        `tfsdk:"https_redirect_status_code"`
	ID                      types.String                       `tfsdk:"id"`
	Name                    types.String                       `tfsdk:"name"`
	PathHandling            types.String                       `tfsdk:"path_handling"`
	PreserveHost            types.Bool                         `tfsdk:"preserve_host"`
	Priority                types.Int64                        `tfsdk:"priority"`
	Protocols               []types.String                     `tfsdk:"protocols"`
	RequestBuffering        types.Bool                         `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool                         `tfsdk:"response_buffering"`
	Service                 *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service"`
	StripPath               types.Bool                         `tfsdk:"strip_path"`
	Tags                    []types.String                     `tfsdk:"tags"`
	UpdatedAt               types.Int64                        `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *RouteExpressionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_expression"
}

// Schema defines the schema for the data source.
func (r *RouteExpressionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RouteExpression DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"expression": schema.StringAttribute{
				Computed:    true,
				Description: `Use Router Expression to perform route match. This option is only available when ` + "`" + `router_flavor` + "`" + ` is set to ` + "`" + `expressions` + "`" + `.`,
			},
			"https_redirect_status_code": schema.Int64Attribute{
				Computed:    true,
				Description: `The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is ` + "`" + `HTTP` + "`" + ` instead of ` + "`" + `HTTPS` + "`" + `. ` + "`" + `Location` + "`" + ` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the ` + "`" + `https` + "`" + ` protocol.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".`,
			},
			"path_handling": schema.StringAttribute{
				Computed:    true,
				Description: `Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.`,
			},
			"preserve_host": schema.BoolAttribute{
				Computed:    true,
				Description: `When matching a Route via one of the ` + "`" + `hosts` + "`" + ` domain names, use the request ` + "`" + `Host` + "`" + ` header in the upstream request headers. If set to ` + "`" + `false` + "`" + `, the upstream ` + "`" + `Host` + "`" + ` header will be that of the Service's ` + "`" + `host` + "`" + `.`,
			},
			"priority": schema.Int64Attribute{
				Computed: true,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An array of the protocols this Route should allow. See the [Route Object](#route-object) section for a list of accepted protocols. When set to only ` + "`" + `"https"` + "`" + `, HTTP requests are answered with an upgrade error. When set to only ` + "`" + `"http"` + "`" + `, HTTPS requests are answered with an error.`,
			},
			"request_buffering": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.`,
			},
			"response_buffering": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `The Service this Route is associated to. This is where the Route proxies traffic to.`,
			},
			"strip_path": schema.BoolAttribute{
				Computed:    true,
				Description: `When matching a Route via one of the ` + "`" + `paths` + "`" + `, strip the matching prefix from the upstream request URL.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Route for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *RouteExpressionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RouteExpressionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RouteExpressionDataSourceModel
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

	request, requestDiags := data.ToOperationsGetRouteRouteExpressionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Routes.GetRouteRouteExpression(ctx, *request)
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
	if !(res.RouteExpression != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRouteExpression(ctx, res.RouteExpression)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
