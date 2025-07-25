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
var _ datasource.DataSource = &JwtDataSource{}
var _ datasource.DataSourceWithConfigure = &JwtDataSource{}

func NewJwtDataSource() datasource.DataSource {
	return &JwtDataSource{}
}

// JwtDataSource is the data source implementation.
type JwtDataSource struct {
	// Provider configured SDK client.
	client *sdk.KongGateway
}

// JwtDataSourceModel describes the data model.
type JwtDataSourceModel struct {
	Algorithm    types.String                       `tfsdk:"algorithm"`
	Consumer     *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
	ConsumerID   types.String                       `tfsdk:"consumer_id"`
	CreatedAt    types.Int64                        `tfsdk:"created_at"`
	ID           types.String                       `tfsdk:"id"`
	Key          types.String                       `tfsdk:"key"`
	RsaPublicKey types.String                       `tfsdk:"rsa_public_key"`
	Secret       types.String                       `tfsdk:"secret"`
	Tags         []types.String                     `tfsdk:"tags"`
}

// Metadata returns the data source type name.
func (r *JwtDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jwt"
}

// Schema defines the schema for the data source.
func (r *JwtDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Jwt DataSource",

		Attributes: map[string]schema.Attribute{
			"algorithm": schema.StringAttribute{
				Computed: true,
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"consumer_id": schema.StringAttribute{
				Required:    true,
				Description: `Consumer ID for nested entities`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"key": schema.StringAttribute{
				Computed: true,
			},
			"rsa_public_key": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *JwtDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *JwtDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *JwtDataSourceModel
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

	request, requestDiags := data.ToOperationsGetJwtWithConsumerRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.JWTs.GetJwtWithConsumer(ctx, *request)
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
	if !(res.Jwt != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedJwt(ctx, res.Jwt)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
