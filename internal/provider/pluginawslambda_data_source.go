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
var _ datasource.DataSource = &PluginAwsLambdaDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginAwsLambdaDataSource{}

func NewPluginAwsLambdaDataSource() datasource.DataSource {
	return &PluginAwsLambdaDataSource{}
}

// PluginAwsLambdaDataSource is the data source implementation.
type PluginAwsLambdaDataSource struct {
	// Provider configured SDK client.
	client *sdk.KongGateway
}

// PluginAwsLambdaDataSourceModel describes the data model.
type PluginAwsLambdaDataSourceModel struct {
	Config       *tfTypes.AwsLambdaPluginConfig     `tfsdk:"config"`
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
func (r *PluginAwsLambdaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_aws_lambda"
}

// Schema defines the schema for the data source.
func (r *PluginAwsLambdaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginAwsLambda DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"aws_assume_role_arn": schema.StringAttribute{
						Computed:    true,
						Description: `The target AWS IAM role ARN used to invoke the Lambda function.`,
					},
					"aws_imds_protocol_version": schema.StringAttribute{
						Computed:    true,
						Description: `Identifier to select the IMDS protocol version to use: ` + "`" + `v1` + "`" + ` or ` + "`" + `v2` + "`" + `.`,
					},
					"aws_key": schema.StringAttribute{
						Computed:    true,
						Description: `The AWS key credential to be used when invoking the function.`,
					},
					"aws_region": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"aws_role_session_name": schema.StringAttribute{
						Computed:    true,
						Description: `The identifier of the assumed role session.`,
					},
					"aws_secret": schema.StringAttribute{
						Computed:    true,
						Description: `The AWS secret credential to be used when invoking the function.`,
					},
					"aws_sts_endpoint_url": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"awsgateway_compatible": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the plugin should wrap requests into the Amazon API gateway.`,
					},
					"base64_encode_body": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that Base64-encodes the request body.`,
					},
					"disable_https": schema.BoolAttribute{
						Computed: true,
					},
					"empty_arrays_mode": schema.StringAttribute{
						Computed:    true,
						Description: `An optional value that defines whether Kong should send empty arrays (returned by Lambda function) as ` + "`" + `[]` + "`" + ` arrays or ` + "`" + `{}` + "`" + ` objects in JSON responses. The value ` + "`" + `legacy` + "`" + ` means Kong will send empty arrays as ` + "`" + `{}` + "`" + ` objects in response`,
					},
					"forward_request_body": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the request body is sent in the request_body field of the JSON-encoded request. If the body arguments can be parsed, they are sent in the separate request_body_args field of the request.`,
					},
					"forward_request_headers": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the original HTTP request headers are sent as a map in the request_headers field of the JSON-encoded request.`,
					},
					"forward_request_method": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the original HTTP request method verb is sent in the request_method field of the JSON-encoded request.`,
					},
					"forward_request_uri": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the original HTTP request URI is sent in the request_uri field of the JSON-encoded request.`,
					},
					"function_name": schema.StringAttribute{
						Computed:    true,
						Description: `The AWS Lambda function to invoke. Both function name and function ARN (including partial) are supported.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"invocation_type": schema.StringAttribute{
						Computed:    true,
						Description: `The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun.`,
					},
					"is_proxy_integration": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether the response format to receive from the Lambda to this format.`,
					},
					"keepalive": schema.Float64Attribute{
						Computed:    true,
						Description: `An optional value in milliseconds that defines how long an idle connection lives before being closed.`,
					},
					"log_type": schema.StringAttribute{
						Computed:    true,
						Description: `The LogType to use when invoking the function. By default, None and Tail are supported.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
					},
					"proxy_url": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"qualifier": schema.StringAttribute{
						Computed:    true,
						Description: `The qualifier to use when invoking the function.`,
					},
					"skip_large_bodies": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional value that defines whether Kong should send large bodies that are buffered to disk`,
					},
					"timeout": schema.Float64Attribute{
						Computed:    true,
						Description: `An optional timeout in milliseconds when invoking the function.`,
					},
					"unhandled_status": schema.Int64Attribute{
						Computed:    true,
						Description: `The response status code to use (instead of the default 200, 202, or 204) in the case of an Unhandled Function Error.`,
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

func (r *PluginAwsLambdaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginAwsLambdaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginAwsLambdaDataSourceModel
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

	request, requestDiags := data.ToOperationsGetAwslambdaPluginRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Plugins.GetAwslambdaPlugin(ctx, *request)
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
	if !(res.AwsLambdaPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAwsLambdaPlugin(ctx, res.AwsLambdaPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
