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
var _ datasource.DataSource = &UpstreamDataSource{}
var _ datasource.DataSourceWithConfigure = &UpstreamDataSource{}

func NewUpstreamDataSource() datasource.DataSource {
	return &UpstreamDataSource{}
}

// UpstreamDataSource is the data source implementation.
type UpstreamDataSource struct {
	// Provider configured SDK client.
	client *sdk.KongGateway
}

// UpstreamDataSourceModel describes the data model.
type UpstreamDataSourceModel struct {
	Algorithm              types.String                       `tfsdk:"algorithm"`
	ClientCertificate      *tfTypes.ACLWithoutParentsConsumer `tfsdk:"client_certificate"`
	CreatedAt              types.Int64                        `tfsdk:"created_at"`
	HashFallback           types.String                       `tfsdk:"hash_fallback"`
	HashFallbackHeader     types.String                       `tfsdk:"hash_fallback_header"`
	HashFallbackQueryArg   types.String                       `tfsdk:"hash_fallback_query_arg"`
	HashFallbackURICapture types.String                       `tfsdk:"hash_fallback_uri_capture"`
	HashOn                 types.String                       `tfsdk:"hash_on"`
	HashOnCookie           types.String                       `tfsdk:"hash_on_cookie"`
	HashOnCookiePath       types.String                       `tfsdk:"hash_on_cookie_path"`
	HashOnHeader           types.String                       `tfsdk:"hash_on_header"`
	HashOnQueryArg         types.String                       `tfsdk:"hash_on_query_arg"`
	HashOnURICapture       types.String                       `tfsdk:"hash_on_uri_capture"`
	Healthchecks           *tfTypes.Healthchecks              `tfsdk:"healthchecks"`
	HostHeader             types.String                       `tfsdk:"host_header"`
	ID                     types.String                       `tfsdk:"id"`
	Name                   types.String                       `tfsdk:"name"`
	Slots                  types.Int64                        `tfsdk:"slots"`
	Tags                   []types.String                     `tfsdk:"tags"`
	UpdatedAt              types.Int64                        `tfsdk:"updated_at"`
	UseSrvName             types.Bool                         `tfsdk:"use_srv_name"`
}

// Metadata returns the data source type name.
func (r *UpstreamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_upstream"
}

// Schema defines the schema for the data source.
func (r *UpstreamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Upstream DataSource",

		Attributes: map[string]schema.Attribute{
			"algorithm": schema.StringAttribute{
				Computed:    true,
				Description: `Which load balancing algorithm to use.`,
			},
			"client_certificate": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"hash_fallback": schema.StringAttribute{
				Computed:    true,
				Description: `What to use as hashing input if the primary ` + "`" + `hash_on` + "`" + ` does not return a hash (eg. header is missing, or no Consumer identified). Not available if ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `cookie` + "`" + `.`,
			},
			"hash_fallback_header": schema.StringAttribute{
				Computed:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `header` + "`" + `.`,
			},
			"hash_fallback_query_arg": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `.`,
			},
			"hash_fallback_uri_capture": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `.`,
			},
			"hash_on": schema.StringAttribute{
				Computed:    true,
				Description: `What to use as hashing input. Using ` + "`" + `none` + "`" + ` results in a weighted-round-robin scheme with no hashing.`,
			},
			"hash_on_cookie": schema.StringAttribute{
				Computed:    true,
				Description: `The cookie name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.`,
			},
			"hash_on_cookie_path": schema.StringAttribute{
				Computed:    true,
				Description: `The cookie path to set in the response headers. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `.`,
			},
			"hash_on_header": schema.StringAttribute{
				Computed:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `header` + "`" + `.`,
			},
			"hash_on_query_arg": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `.`,
			},
			"hash_on_uri_capture": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `.`,
			},
			"healthchecks": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Computed: true,
							},
							"headers": schema.MapAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										ElementType: types.Int64Type,
									},
									"interval": schema.Float64Attribute{
										Computed: true,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
							"http_path": schema.StringAttribute{
								Computed: true,
							},
							"https_sni": schema.StringAttribute{
								Computed: true,
							},
							"https_verify_certificate": schema.BoolAttribute{
								Computed: true,
							},
							"timeout": schema.Float64Attribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
									},
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										ElementType: types.Int64Type,
									},
									"interval": schema.Float64Attribute{
										Computed: true,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
					"passive": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										ElementType: types.Int64Type,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
									},
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										ElementType: types.Int64Type,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
					"threshold": schema.Float64Attribute{
						Computed: true,
					},
				},
			},
			"host_header": schema.StringAttribute{
				Computed:    true,
				Description: `The hostname to be used as ` + "`" + `Host` + "`" + ` header when proxying requests through Kong.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `This is a hostname, which must be equal to the ` + "`" + `host` + "`" + ` of a Service.`,
			},
			"slots": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of slots in the load balancer algorithm. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `round-robin` + "`" + `, this setting determines the maximum number of slots. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `consistent-hashing` + "`" + `, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range ` + "`" + `10` + "`" + `-` + "`" + `65536` + "`" + `.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Upstream for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
			"use_srv_name": schema.BoolAttribute{
				Computed:    true,
				Description: `If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream ` + "`" + `Host` + "`" + `.`,
			},
		},
	}
}

func (r *UpstreamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *UpstreamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UpstreamDataSourceModel
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

	request, requestDiags := data.ToOperationsGetUpstreamRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Upstreams.GetUpstream(ctx, *request)
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
	if !(res.Upstream != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedUpstream(ctx, res.Upstream)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
