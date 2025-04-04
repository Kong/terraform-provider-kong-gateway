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
var _ datasource.DataSource = &TargetDataSource{}
var _ datasource.DataSourceWithConfigure = &TargetDataSource{}

func NewTargetDataSource() datasource.DataSource {
	return &TargetDataSource{}
}

// TargetDataSource is the data source implementation.
type TargetDataSource struct {
	client *sdk.KongGateway
}

// TargetDataSourceModel describes the data model.
type TargetDataSourceModel struct {
	CreatedAt  types.Number                       `tfsdk:"created_at"`
	ID         types.String                       `tfsdk:"id"`
	Tags       []types.String                     `tfsdk:"tags"`
	Target     types.String                       `tfsdk:"target"`
	UpdatedAt  types.Number                       `tfsdk:"updated_at"`
	Upstream   *tfTypes.ACLWithoutParentsConsumer `tfsdk:"upstream"`
	UpstreamID types.String                       `tfsdk:"upstream_id"`
	Weight     types.Int64                        `tfsdk:"weight"`
}

// Metadata returns the data source type name.
func (r *TargetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_target"
}

// Schema defines the schema for the data source.
func (r *TargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Target DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.NumberAttribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Target for grouping and filtering.`,
			},
			"target": schema.StringAttribute{
				Computed:    true,
				Description: `The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the ` + "`" + `port` + "`" + ` value will be overridden by the value from the DNS record.`,
			},
			"updated_at": schema.NumberAttribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
			"upstream": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"upstream_id": schema.StringAttribute{
				Required:    true,
				Description: `ID or target of the Target to lookup`,
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: `The weight this target gets within the upstream loadbalancer (` + "`" + `0` + "`" + `-` + "`" + `65535` + "`" + `). If the hostname resolves to an SRV record, the ` + "`" + `weight` + "`" + ` value will be overridden by the value from the DNS record.`,
			},
		},
	}
}

func (r *TargetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TargetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TargetDataSourceModel
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

	var upstreamID string
	upstreamID = data.UpstreamID.ValueString()

	var targetIDOrTarget string
	targetIDOrTarget = data.ID.ValueString()

	request := operations.GetTargetWithUpstreamRequest{
		UpstreamID:       upstreamID,
		TargetIDOrTarget: targetIDOrTarget,
	}
	res, err := r.client.Targets.GetTargetWithUpstream(ctx, request)
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
	if !(res.Target != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
