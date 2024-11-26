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
var _ datasource.DataSource = &PluginXMLThreatProtectionDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginXMLThreatProtectionDataSource{}

func NewPluginXMLThreatProtectionDataSource() datasource.DataSource {
	return &PluginXMLThreatProtectionDataSource{}
}

// PluginXMLThreatProtectionDataSource is the data source implementation.
type PluginXMLThreatProtectionDataSource struct {
	client *sdk.KongGateway
}

// PluginXMLThreatProtectionDataSourceModel describes the data model.
type PluginXMLThreatProtectionDataSourceModel struct {
	Config        tfTypes.XMLThreatProtectionPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                    `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                    `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                             `tfsdk:"created_at"`
	Enabled       types.Bool                              `tfsdk:"enabled"`
	ID            types.String                            `tfsdk:"id"`
	InstanceName  types.String                            `tfsdk:"instance_name"`
	Ordering      *tfTypes.ACLPluginOrdering              `tfsdk:"ordering"`
	Protocols     []types.String                          `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                    `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                    `tfsdk:"service"`
	Tags          []types.String                          `tfsdk:"tags"`
	UpdatedAt     types.Int64                             `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginXMLThreatProtectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_xml_threat_protection"
}

// Schema defines the schema for the data source.
func (r *PluginXMLThreatProtectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginXMLThreatProtection DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allow_dtd": schema.BoolAttribute{
						Computed:    true,
						Description: `Indicates whether an XML Document Type Definition (DTD) section is allowed.`,
					},
					"allowed_content_types": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of Content-Type values with payloads that are allowed, but aren't validated.`,
					},
					"attribute": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the attribute value.`,
					},
					"bla_max_amplification": schema.NumberAttribute{
						Computed:    true,
						Description: `Sets the maximum allowed amplification. This protects against the Billion Laughs Attack.`,
					},
					"bla_threshold": schema.Int64Attribute{
						Computed:    true,
						Description: `Sets the threshold after which the protection starts. This protects against the Billion Laughs Attack.`,
					},
					"buffer": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the unparsed buffer (see below).`,
					},
					"checked_content_types": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of Content-Type values with payloads that must be validated.`,
					},
					"comment": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of comments.`,
					},
					"document": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the entire document.`,
					},
					"entity": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of entity values in EntityDecl.`,
					},
					"entityname": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of entity names in EntityDecl.`,
					},
					"entityproperty": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of systemId, publicId, or notationName in EntityDecl.`,
					},
					"localname": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the localname. This applies to tags and attributes.`,
					},
					"max_attributes": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of attributes allowed on a tag, including default ones. Note: If namespace-aware parsing is disabled, then the namespaces definitions are counted as attributes.`,
					},
					"max_children": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of children allowed (Element, Text, Comment, ProcessingInstruction, CDATASection). Note: Adjacent text and CDATA sections are counted as one. For example, text-cdata-text-cdata is one child.`,
					},
					"max_depth": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum depth of tags. Child elements such as Text or Comments are not counted as another level.`,
					},
					"max_namespaces": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of namespaces defined on a tag. This value is required if parsing is namespace-aware.`,
					},
					"namespace_aware": schema.BoolAttribute{
						Computed:    true,
						Description: `If not parsing namespace aware, all prefixes and namespace attributes will be counted as regular attributes and element names, and validated as such.`,
					},
					"namespaceuri": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the namespace URI. This value is required if parsing is namespace-aware.`,
					},
					"pidata": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of processing instruction data.`,
					},
					"pitarget": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of processing instruction targets.`,
					},
					"prefix": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of the prefix. This applies to tags and attributes. This value is required if parsing is namespace-aware.`,
					},
					"text": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum text inside tags (counted over all adjacent text/CDATA elements combined).`,
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

func (r *PluginXMLThreatProtectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginXMLThreatProtectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginXMLThreatProtectionDataSourceModel
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

	request := operations.GetXmlthreatprotectionPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetXmlthreatprotectionPlugin(ctx, request)
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
	if !(res.XMLThreatProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedXMLThreatProtectionPlugin(res.XMLThreatProtectionPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
