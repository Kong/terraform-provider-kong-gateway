// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-kong-gateway/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PluginServiceProtectionResource{}
var _ resource.ResourceWithImportState = &PluginServiceProtectionResource{}

func NewPluginServiceProtectionResource() resource.Resource {
	return &PluginServiceProtectionResource{}
}

// PluginServiceProtectionResource defines the resource implementation.
type PluginServiceProtectionResource struct {
	client *sdk.KongGateway
}

// PluginServiceProtectionResourceModel describes the resource data model.
type PluginServiceProtectionResourceModel struct {
	Config       *tfTypes.ServiceProtectionPluginConfig `tfsdk:"config"`
	CreatedAt    types.Int64                            `tfsdk:"created_at"`
	Enabled      types.Bool                             `tfsdk:"enabled"`
	ID           types.String                           `tfsdk:"id"`
	InstanceName types.String                           `tfsdk:"instance_name"`
	Ordering     *tfTypes.Ordering                      `tfsdk:"ordering"`
	Partials     []tfTypes.Partials                     `tfsdk:"partials"`
	Protocols    []types.String                         `tfsdk:"protocols"`
	Service      *tfTypes.ACLWithoutParentsConsumer     `tfsdk:"service"`
	Tags         []types.String                         `tfsdk:"tags"`
	UpdatedAt    types.Int64                            `tfsdk:"updated_at"`
}

func (r *PluginServiceProtectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_service_protection"
}

func (r *PluginServiceProtectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginServiceProtection Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"dictionary_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is ` + "`" + `config.strategy` + "`" + ` is ` + "`" + `cluster` + "`" + ` or ` + "`" + `redis` + "`" + ` and ` + "`" + `config.sync_rate` + "`" + ` isn't ` + "`" + `-1` + "`" + `), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.`,
					},
					"disable_penalty": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `If set to ` + "`" + `true` + "`" + `, this doesn't count denied requests (status = ` + "`" + `429` + "`" + `). If set to ` + "`" + `false` + "`" + `, all requests, including denied ones, are counted. This parameter only affects the ` + "`" + `sliding` + "`" + ` window_type.`,
					},
					"error_code": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Set a custom error code to return when the rate limit is exceeded.`,
					},
					"error_message": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Set a custom error message to return when the rate limit is exceeded.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.`,
					},
					"limit": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.`,
					},
					"lock_dictionary_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The shared dictionary where concurrency control locks are stored. The default shared dictionary is ` + "`" + `kong_locks` + "`" + `. The shared dictionary should be declared in nginx-kong.conf.`,
					},
					"namespace": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. ` + "`" + `strategy` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `sync_rate` + "`" + `, ` + "`" + `dictionary_name` + "`" + `, need to be the same.`,
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"cluster_max_redirections": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum retry attempts for redirection.`,
							},
							"cluster_nodes": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
											Validators: []validator.Int64{
												int64validator.AtMost(65535),
											},
										},
									},
								},
								Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.`,
							},
							"connect_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"connection_is_proxied": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If the connection to Redis is proxied (e.g. Envoy), set it ` + "`" + `true` + "`" + `. Set the ` + "`" + `host` + "`" + ` and ` + "`" + `port` + "`" + ` to point to the proxy address.`,
							},
							"database": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"keepalive_backlog": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"keepalive_pool_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.`,
								Validators: []validator.Int64{
									int64validator.Between(1, 2147483646),
								},
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
								Validators: []validator.Int64{
									int64validator.AtMost(65535),
								},
							},
							"read_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"send_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"sentinel_master": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_nodes": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
											Validators: []validator.Int64{
												int64validator.AtMost(65535),
											},
										},
									},
								},
								Description: `Sentinel node addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.`,
							},
							"sentinel_password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
							},
							"sentinel_role": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel. must be one of ["any", "master", "slave"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"any",
										"master",
										"slave",
									),
								},
							},
							"sentinel_username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, uses SSL to connect to Redis.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
					},
					"retry_after_jitter_max": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The upper bound of a jitter (random delay) in seconds to be added to the ` + "`" + `Retry-After` + "`" + ` header of denied requests (status = ` + "`" + `429` + "`" + `) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is ` + "`" + `0` + "`" + `; in this case, the ` + "`" + `Retry-After` + "`" + ` header is equal to the ` + "`" + `RateLimit-Reset` + "`" + ` header.`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: ` + "`" + `local` + "`" + ` and ` + "`" + `cluster` + "`" + `. must be one of ["cluster", "local", "redis"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"cluster",
								"local",
								"redis",
							),
						},
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).`,
					},
					"window_size": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.`,
					},
					"window_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Sets the time window type to either ` + "`" + `sliding` + "`" + ` (default) or ` + "`" + `fixed` + "`" + `. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters. must be one of ["fixed", "sliding"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fixed",
								"sliding",
							),
						},
					},
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"partials": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"path": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A set of strings representing HTTP protocols.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *PluginServiceProtectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongGateway)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.KongGateway, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PluginServiceProtectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PluginServiceProtectionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedServiceProtectionPlugin()
	res, err := r.client.Plugins.CreateServiceprotectionPlugin(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ServiceProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedServiceProtectionPlugin(res.ServiceProtectionPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginServiceProtectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PluginServiceProtectionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request := operations.GetServiceprotectionPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetServiceprotectionPlugin(ctx, request)
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
	if !(res.ServiceProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedServiceProtectionPlugin(res.ServiceProtectionPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginServiceProtectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PluginServiceProtectionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	serviceProtectionPlugin := *data.ToSharedServiceProtectionPlugin()
	request := operations.UpdateServiceprotectionPluginRequest{
		PluginID:                pluginID,
		ServiceProtectionPlugin: serviceProtectionPlugin,
	}
	res, err := r.client.Plugins.UpdateServiceprotectionPlugin(ctx, request)
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
	if !(res.ServiceProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedServiceProtectionPlugin(res.ServiceProtectionPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginServiceProtectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PluginServiceProtectionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request := operations.DeleteServiceprotectionPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.DeleteServiceprotectionPlugin(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *PluginServiceProtectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
