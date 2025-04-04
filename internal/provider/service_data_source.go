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
var _ datasource.DataSource = &ServiceDataSource{}
var _ datasource.DataSourceWithConfigure = &ServiceDataSource{}

func NewServiceDataSource() datasource.DataSource {
	return &ServiceDataSource{}
}

// ServiceDataSource is the data source implementation.
type ServiceDataSource struct {
	client *sdk.KongGateway
}

// ServiceDataSourceModel describes the data model.
type ServiceDataSourceModel struct {
	CaCertificates    []types.String                     `tfsdk:"ca_certificates"`
	ClientCertificate *tfTypes.ACLWithoutParentsConsumer `tfsdk:"client_certificate"`
	ConnectTimeout    types.Int64                        `tfsdk:"connect_timeout"`
	CreatedAt         types.Int64                        `tfsdk:"created_at"`
	Enabled           types.Bool                         `tfsdk:"enabled"`
	Host              types.String                       `tfsdk:"host"`
	ID                types.String                       `tfsdk:"id"`
	Name              types.String                       `tfsdk:"name"`
	Path              types.String                       `tfsdk:"path"`
	Port              types.Int64                        `tfsdk:"port"`
	Protocol          types.String                       `tfsdk:"protocol"`
	ReadTimeout       types.Int64                        `tfsdk:"read_timeout"`
	Retries           types.Int64                        `tfsdk:"retries"`
	Tags              []types.String                     `tfsdk:"tags"`
	TLSVerify         types.Bool                         `tfsdk:"tls_verify"`
	TLSVerifyDepth    types.Int64                        `tfsdk:"tls_verify_depth"`
	UpdatedAt         types.Int64                        `tfsdk:"updated_at"`
	WriteTimeout      types.Int64                        `tfsdk:"write_timeout"`
}

// Metadata returns the data source type name.
func (r *ServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

// Schema defines the schema for the data source.
func (r *ServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Service DataSource",

		Attributes: map[string]schema.Attribute{
			"ca_certificates": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `Array of ` + "`" + `CA Certificate` + "`" + ` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to ` + "`" + `null` + "`" + ` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).`,
			},
			"client_certificate": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Certificate to be used as client certificate while TLS handshaking to the upstream server.`,
			},
			"connect_timeout": schema.Int64Attribute{
				Computed:    true,
				Description: `The timeout in milliseconds for establishing a connection to the upstream server.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the Service is active. If set to ` + "`" + `false` + "`" + `, the proxy behavior will be as if any routes attached to it do not exist (404). Default: ` + "`" + `true` + "`" + `.`,
			},
			"host": schema.StringAttribute{
				Computed:    true,
				Description: `The host of the upstream server. Note that the host value is case sensitive.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The Service name.`,
			},
			"path": schema.StringAttribute{
				Computed:    true,
				Description: `The path to be used in requests to the upstream server.`,
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: `The upstream server port.`,
			},
			"protocol": schema.StringAttribute{
				Computed:    true,
				Description: `The protocol used to communicate with the upstream.`,
			},
			"read_timeout": schema.Int64Attribute{
				Computed:    true,
				Description: `The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.`,
			},
			"retries": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of retries to execute upon failure to proxy.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Service for grouping and filtering.`,
			},
			"tls_verify": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to enable verification of upstream server TLS certificate. If set to ` + "`" + `null` + "`" + `, then the Nginx default is respected.`,
			},
			"tls_verify_depth": schema.Int64Attribute{
				Computed:    true,
				Description: `Maximum depth of chain while verifying Upstream server's TLS certificate. If set to ` + "`" + `null` + "`" + `, then the Nginx default is respected.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
			"write_timeout": schema.Int64Attribute{
				Computed:    true,
				Description: `The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.`,
			},
		},
	}
}

func (r *ServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ServiceDataSourceModel
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

	var serviceIDOrName string
	serviceIDOrName = data.ID.ValueString()

	request := operations.GetServiceRequest{
		ServiceIDOrName: serviceIDOrName,
	}
	res, err := r.client.Services.GetService(ctx, request)
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
	if !(res.Service != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedServiceOutput(res.Service)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
