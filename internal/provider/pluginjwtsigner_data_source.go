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
var _ datasource.DataSource = &PluginJWTSignerDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginJWTSignerDataSource{}

func NewPluginJWTSignerDataSource() datasource.DataSource {
	return &PluginJWTSignerDataSource{}
}

// PluginJWTSignerDataSource is the data source implementation.
type PluginJWTSignerDataSource struct {
	client *sdk.KongGateway
}

// PluginJWTSignerDataSourceModel describes the data model.
type PluginJWTSignerDataSourceModel struct {
	Config        *tfTypes.CreateJWTSignerPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                 `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                 `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                          `tfsdk:"created_at"`
	Enabled       types.Bool                           `tfsdk:"enabled"`
	ID            types.String                         `tfsdk:"id"`
	InstanceName  types.String                         `tfsdk:"instance_name"`
	Ordering      types.String                         `tfsdk:"ordering"`
	Protocols     []types.String                       `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                 `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                 `tfsdk:"service"`
	Tags          []types.String                       `tfsdk:"tags"`
	UpdatedAt     types.Int64                          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginJWTSignerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_jwt_signer"
}

// Schema defines the schema for the data source.
func (r *PluginJWTSignerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginJWTSigner DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_token_consumer_by": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When the plugin tries to apply an access token to a Kong consumer mapping, it tries to find a matching Kong consumer from properties defined using this configuration parameter. The parameter can take an array of alues. Valid values are ` + "`" + `id` + "`" + `, ` + "`" + `username` + "`" + `, and ` + "`" + `custom_id` + "`" + `.`,
					},
					"access_token_consumer_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When you set a value for this parameter, the plugin tries to map an arbitrary claim specified with this configuration parameter (for example, ` + "`" + `sub` + "`" + ` or ` + "`" + `username` + "`" + `) in an access token to Kong consumer entity.`,
					},
					"access_token_introspection_authorization": schema.StringAttribute{
						Computed:    true,
						Description: `If the introspection endpoint requires client authentication (client being the JWT Signer plugin), you can specify the ` + "`" + `Authorization` + "`" + ` header's value with this configuration parameter.`,
					},
					"access_token_introspection_body_args": schema.StringAttribute{
						Computed:    true,
						Description: `This parameter allows you to pass URL encoded request body arguments. For example: ` + "`" + `resource=` + "`" + ` or ` + "`" + `a=1&b=&c` + "`" + `.`,
					},
					"access_token_introspection_consumer_by": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When the plugin tries to do access token introspection results to Kong consumer mapping, it tries to find a matching Kong consumer from properties defined using this configuration parameter. The parameter can take an array of values.`,
					},
					"access_token_introspection_consumer_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When you set a value for this parameter, the plugin tries to map an arbitrary claim specified with this configuration parameter (such as ` + "`" + `sub` + "`" + ` or ` + "`" + `username` + "`" + `) in access token introspection results to the Kong consumer entity.`,
					},
					"access_token_introspection_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `When you use ` + "`" + `opaque` + "`" + ` access tokens and you want to turn on access token introspection, you need to specify the OAuth 2.0 introspection endpoint URI with this configuration parameter.`,
					},
					"access_token_introspection_hint": schema.StringAttribute{
						Computed:    true,
						Description: `If you need to give ` + "`" + `hint` + "`" + ` parameter when introspecting an access token, use this parameter to specify the value. By default, the plugin sends ` + "`" + `hint=access_token` + "`" + `.`,
					},
					"access_token_introspection_jwt_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `If your introspection endpoint returns an access token in one of the keys (or claims) within the introspection results (` + "`" + `JSON` + "`" + `). If the key cannot be found, the plugin responds with ` + "`" + `401 Unauthorized` + "`" + `. Also if the key is found but cannot be decoded as JWT, it also responds with ` + "`" + `401 Unauthorized` + "`" + `.`,
					},
					"access_token_introspection_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `Adjusts clock skew between the token issuer introspection results and Kong. The value is added to introspection results (` + "`" + `JSON` + "`" + `) ` + "`" + `exp` + "`" + ` claim/property before checking token expiry against Kong servers current time in seconds. You can disable access token introspection ` + "`" + `expiry` + "`" + ` verification altogether with ` + "`" + `config.verify_access_token_introspection_expiry` + "`" + `.`,
					},
					"access_token_introspection_scopes_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the claim/property in access token introspection results (` + "`" + `JSON` + "`" + `) to be verified against values of ` + "`" + `config.access_token_introspection_scopes_required` + "`" + `. This supports nested claims. For example, with Keycloak you could use ` + "`" + `[ "realm_access", "roles" ]` + "`" + `, hich can be given as ` + "`" + `realm_access,roles` + "`" + ` (form post). If the claim is not found in access token introspection results, and you have specified ` + "`" + `config.access_token_introspection_scopes_required` + "`" + `, the plugin responds with ` + "`" + `403 Forbidden` + "`" + `.`,
					},
					"access_token_introspection_scopes_required": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the required values (or scopes) that are checked by an introspection claim/property specified by ` + "`" + `config.access_token_introspection_scopes_claim` + "`" + `.`,
					},
					"access_token_introspection_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `Timeout in milliseconds for an introspection request. The plugin tries to introspect twice if the first request fails for some reason. If both requests timeout, then the plugin runs two times the ` + "`" + `config.access_token_introspection_timeout` + "`" + ` on access token introspection.`,
					},
					"access_token_issuer": schema.StringAttribute{
						Computed:    true,
						Description: `The ` + "`" + `iss` + "`" + ` claim of a signed or re-signed access token is set to this value. Original ` + "`" + `iss` + "`" + ` claim of the incoming token (possibly introspected) is stored in ` + "`" + `original_iss` + "`" + ` claim of the newly signed access token.`,
					},
					"access_token_jwks_uri": schema.StringAttribute{
						Computed:    true,
						Description: `Specify the URI where the plugin can fetch the public keys (JWKS) to verify the signature of the access token.`,
					},
					"access_token_jwks_uri_client_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The client certificate that will be used to authenticate Kong if ` + "`" + `access_token_jwks_uri` + "`" + ` is an https uri that requires mTLS Auth.`,
					},
					"access_token_jwks_uri_client_password": schema.StringAttribute{
						Computed:    true,
						Description: `The client password that will be used to authenticate Kong if ` + "`" + `access_token_jwks_uri` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `access_token_jwks_uri_client_username` + "`" + ``,
					},
					"access_token_jwks_uri_client_username": schema.StringAttribute{
						Computed:    true,
						Description: `The client username that will be used to authenticate Kong if ` + "`" + `access_token_jwks_uri` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `access_token_jwks_uri_client_password` + "`" + ``,
					},
					"access_token_jwks_uri_rotate_period": schema.NumberAttribute{
						Computed:    true,
						Description: `Specify the period (in seconds) to auto-rotate the jwks for ` + "`" + `access_token_jwks_uri` + "`" + `. The default value 0 means no auto-rotation.`,
					},
					"access_token_keyset": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the keyset containing signing keys.`,
					},
					"access_token_keyset_client_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The client certificate that will be used to authenticate Kong if ` + "`" + `access_token_keyset` + "`" + ` is an https uri that requires mTLS Auth.`,
					},
					"access_token_keyset_client_password": schema.StringAttribute{
						Computed:    true,
						Description: `The client password that will be used to authenticate Kong if ` + "`" + `access_token_keyset` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `access_token_keyset_client_username` + "`" + ``,
					},
					"access_token_keyset_client_username": schema.StringAttribute{
						Computed:    true,
						Description: `The client username that will be used to authenticate Kong if ` + "`" + `access_token_keyset` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `access_token_keyset_client_password` + "`" + ``,
					},
					"access_token_keyset_rotate_period": schema.NumberAttribute{
						Computed:    true,
						Description: `Specify the period (in seconds) to auto-rotate the jwks for ` + "`" + `access_token_keyset` + "`" + `. The default value 0 means no auto-rotation.`,
					},
					"access_token_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `Adjusts clock skew between the token issuer and Kong. The value is added to the token's ` + "`" + `exp` + "`" + ` claim before checking token expiry against Kong servers' current time in seconds. You can disable access token ` + "`" + `expiry` + "`" + ` verification altogether with ` + "`" + `config.verify_access_token_expiry` + "`" + `.`,
					},
					"access_token_optional": schema.BoolAttribute{
						Computed:    true,
						Description: `If an access token is not provided or no ` + "`" + `config.access_token_request_header` + "`" + ` is specified, the plugin cannot verify the access token. In that case, the plugin normally responds with ` + "`" + `401 Unauthorized` + "`" + ` (client didn't send a token) or ` + "`" + `500 Unexpected` + "`" + ` (a configuration error). Use this parameter to allow the request to proceed even when there is no token to check. If the token is provided, then this parameter has no effect`,
					},
					"access_token_request_header": schema.StringAttribute{
						Computed:    true,
						Description: `This parameter tells the name of the header where to look for the access token.`,
					},
					"access_token_scopes_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the claim in an access token to verify against values of ` + "`" + `config.access_token_scopes_required` + "`" + `.`,
					},
					"access_token_scopes_required": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the required values (or scopes) that are checked by a claim specified by ` + "`" + `config.access_token_scopes_claim` + "`" + `.`,
					},
					"access_token_signing_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `When this plugin sets the upstream header as specified with ` + "`" + `config.access_token_upstream_header` + "`" + `, re-signs the original access token using the private keys of the JWT Signer plugin. Specify the algorithm that is used to sign the token. The ` + "`" + `config.access_token_issuer` + "`" + ` specifies which ` + "`" + `keyset` + "`" + ` is used to sign the new token issued by Kong using the specified signing algorithm. must be one of ["HS256", "HS384", "HS512", "RS256", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]`,
					},
					"access_token_upstream_header": schema.StringAttribute{
						Computed:    true,
						Description: `Removes the ` + "`" + `config.access_token_request_header` + "`" + ` from the request after reading its value. With ` + "`" + `config.access_token_upstream_header` + "`" + `, you can specify the upstream header where the plugin adds the Kong signed token. If you don't specify a value, such as use ` + "`" + `null` + "`" + ` or ` + "`" + `""` + "`" + ` (empty string), the plugin does not even try to sign or re-sign the token.`,
					},
					"access_token_upstream_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `If you want to add or subtract (using a negative value) expiry time (in seconds) of the original access token, you can specify a value that is added to the original access token's ` + "`" + `exp` + "`" + ` claim.`,
					},
					"add_access_token_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Add customized claims if they are not present yet. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"add_channel_token_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Add customized claims if they are not present yet. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"add_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Add customized claims to both tokens if they are not present yet. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"cache_access_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `Whether to cache access token introspection results.`,
					},
					"cache_channel_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `Whether to cache channel token introspection results.`,
					},
					"channel_token_consumer_by": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When the plugin tries to do channel token to Kong consumer mapping, it tries to find a matching Kong consumer from properties defined using this configuration parameter. The parameter can take an array of valid values: ` + "`" + `id` + "`" + `, ` + "`" + `username` + "`" + `, and ` + "`" + `custom_id` + "`" + `.`,
					},
					"channel_token_consumer_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When you set a value for this parameter, the plugin tries to map an arbitrary claim specified with this configuration parameter. Kong consumers have an ` + "`" + `id` + "`" + `, a ` + "`" + `username` + "`" + `, and a ` + "`" + `custom_id` + "`" + `. If this parameter is enabled but the mapping fails, such as when there's a non-existent Kong consumer, the plugin responds with ` + "`" + `403 Forbidden` + "`" + `.`,
					},
					"channel_token_introspection_authorization": schema.StringAttribute{
						Computed:    true,
						Description: `When using ` + "`" + `opaque` + "`" + ` channel tokens, and you want to turn on channel token introspection, you need to specify the OAuth 2.0 introspection endpoint URI with this configuration parameter. Otherwise the plugin will not try introspection, and instead returns ` + "`" + `401 Unauthorized` + "`" + ` when using opaque channel tokens.`,
					},
					"channel_token_introspection_body_args": schema.StringAttribute{
						Computed:    true,
						Description: `If you need to pass additional body arguments to introspection endpoint when the plugin introspects the opaque channel token, you can use this config parameter to specify them. You should URL encode the value. For example: ` + "`" + `resource=` + "`" + ` or ` + "`" + `a=1&b=&c` + "`" + `.`,
					},
					"channel_token_introspection_consumer_by": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When the plugin tries to do channel token introspection results to Kong consumer mapping, it tries to find a matching Kong consumer from properties defined using this configuration parameter. The parameter can take an array of values. Valid values are ` + "`" + `id` + "`" + `, ` + "`" + `username` + "`" + ` and ` + "`" + `custom_id` + "`" + `.`,
					},
					"channel_token_introspection_consumer_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `When you set a value for this parameter, the plugin tries to map an arbitrary claim specified with this configuration parameter (such as ` + "`" + `sub` + "`" + ` or ` + "`" + `username` + "`" + `) in channel token introspection results to Kong consumer entity`,
					},
					"channel_token_introspection_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `When you use ` + "`" + `opaque` + "`" + ` access tokens and you want to turn on access token introspection, you need to specify the OAuth 2.0 introspection endpoint URI with this configuration parameter. Otherwise, the plugin does not try introspection and returns ` + "`" + `401 Unauthorized` + "`" + ` instead.`,
					},
					"channel_token_introspection_hint": schema.StringAttribute{
						Computed:    true,
						Description: `If you need to give ` + "`" + `hint` + "`" + ` parameter when introspecting a channel token, you can use this parameter to specify the value of such parameter. By default, a ` + "`" + `hint` + "`" + ` isn't sent with channel token introspection.`,
					},
					"channel_token_introspection_jwt_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `If your introspection endpoint returns a channel token in one of the keys (or claims) in the introspection results (` + "`" + `JSON` + "`" + `), the plugin can use that value instead of the introspection results when doing expiry verification and signing of the new token issued by Kong.`,
					},
					"channel_token_introspection_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `You can use this parameter to adjust clock skew between the token issuer introspection results and Kong. The value will be added to introspection results (` + "`" + `JSON` + "`" + `) ` + "`" + `exp` + "`" + ` claim/property before checking token expiry against Kong servers current time (in seconds). You can disable channel token introspection ` + "`" + `expiry` + "`" + ` verification altogether with ` + "`" + `config.verify_channel_token_introspection_expiry` + "`" + `.`,
					},
					"channel_token_introspection_scopes_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Use this parameter to specify the claim/property in channel token introspection results (` + "`" + `JSON` + "`" + `) to be verified against values of ` + "`" + `config.channel_token_introspection_scopes_required` + "`" + `. This supports nested claims.`,
					},
					"channel_token_introspection_scopes_required": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Use this parameter to specify the required values (or scopes) that are checked by an introspection claim/property specified by ` + "`" + `config.channel_token_introspection_scopes_claim` + "`" + `.`,
					},
					"channel_token_introspection_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `Timeout in milliseconds for an introspection request. The plugin tries to introspect twice if the first request fails for some reason. If both requests timeout, then the plugin runs two times the ` + "`" + `config.access_token_introspection_timeout` + "`" + ` on channel token introspection.`,
					},
					"channel_token_issuer": schema.StringAttribute{
						Computed:    true,
						Description: `The ` + "`" + `iss` + "`" + ` claim of the re-signed channel token is set to this value, which is ` + "`" + `kong` + "`" + ` by default. The original ` + "`" + `iss` + "`" + ` claim of the incoming token (possibly introspected) is stored in the ` + "`" + `original_iss` + "`" + ` claim of the newly signed channel token.`,
					},
					"channel_token_jwks_uri": schema.StringAttribute{
						Computed:    true,
						Description: `If you want to use ` + "`" + `config.verify_channel_token_signature` + "`" + `, you must specify the URI where the plugin can fetch the public keys (JWKS) to verify the signature of the channel token. If you don't specify a URI and you pass a JWT token to the plugin, then the plugin responds with ` + "`" + `401 Unauthorized` + "`" + `.`,
					},
					"channel_token_jwks_uri_client_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The client certificate that will be used to authenticate Kong if ` + "`" + `access_token_jwks_uri` + "`" + ` is an https uri that requires mTLS Auth.`,
					},
					"channel_token_jwks_uri_client_password": schema.StringAttribute{
						Computed:    true,
						Description: `The client password that will be used to authenticate Kong if ` + "`" + `channel_token_jwks_uri` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `channel_token_jwks_uri_client_username` + "`" + ``,
					},
					"channel_token_jwks_uri_client_username": schema.StringAttribute{
						Computed:    true,
						Description: `The client username that will be used to authenticate Kong if ` + "`" + `channel_token_jwks_uri` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `channel_token_jwks_uri_client_password` + "`" + ``,
					},
					"channel_token_jwks_uri_rotate_period": schema.NumberAttribute{
						Computed:    true,
						Description: `Specify the period (in seconds) to auto-rotate the jwks for ` + "`" + `channel_token_jwks_uri` + "`" + `. The default value 0 means no auto-rotation.`,
					},
					"channel_token_keyset": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the keyset containing signing keys.`,
					},
					"channel_token_keyset_client_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The client certificate that will be used to authenticate Kong if ` + "`" + `channel_token_keyset` + "`" + ` is an https uri that requires mTLS Auth.`,
					},
					"channel_token_keyset_client_password": schema.StringAttribute{
						Computed:    true,
						Description: `The client password that will be used to authenticate Kong if ` + "`" + `channel_token_keyset` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `channel_token_keyset_client_username` + "`" + ``,
					},
					"channel_token_keyset_client_username": schema.StringAttribute{
						Computed:    true,
						Description: `The client username that will be used to authenticate Kong if ` + "`" + `channel_token_keyset` + "`" + ` is a uri that requires Basic Auth. Should be configured together with ` + "`" + `channel_token_keyset_client_password` + "`" + ``,
					},
					"channel_token_keyset_rotate_period": schema.NumberAttribute{
						Computed:    true,
						Description: `Specify the period (in seconds) to auto-rotate the jwks for ` + "`" + `channel_token_keyset` + "`" + `. The default value 0 means no auto-rotation.`,
					},
					"channel_token_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `Adjusts clock skew between the token issuer and Kong. The value will be added to token's ` + "`" + `exp` + "`" + ` claim before checking token expiry against Kong servers current time in seconds. You can disable channel token ` + "`" + `expiry` + "`" + ` verification altogether with ` + "`" + `config.verify_channel_token_expiry` + "`" + `.`,
					},
					"channel_token_optional": schema.BoolAttribute{
						Computed:    true,
						Description: `If a channel token is not provided or no ` + "`" + `config.channel_token_request_header` + "`" + ` is specified, the plugin cannot verify the channel token. In that case, the plugin normally responds with ` + "`" + `401 Unauthorized` + "`" + ` (client didn't send a token) or ` + "`" + `500 Unexpected` + "`" + ` (a configuration error). Enable this parameter to allow the request to proceed even when there is no channel token to check. If the channel token is provided, then this parameter has no effect`,
					},
					"channel_token_request_header": schema.StringAttribute{
						Computed:    true,
						Description: `This parameter tells the name of the header where to look for the channel token. If you don't want to do anything with the channel token, then you can set this to ` + "`" + `null` + "`" + ` or ` + "`" + `""` + "`" + ` (empty string).`,
					},
					"channel_token_scopes_claim": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the claim in a channel token to verify against values of ` + "`" + `config.channel_token_scopes_required` + "`" + `. This supports nested claims.`,
					},
					"channel_token_scopes_required": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Specify the required values (or scopes) that are checked by a claim specified by ` + "`" + `config.channel_token_scopes_claim` + "`" + `.`,
					},
					"channel_token_signing_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `When this plugin sets the upstream header as specified with ` + "`" + `config.channel_token_upstream_header` + "`" + `, it also re-signs the original channel token using private keys of this plugin. Specify the algorithm that is used to sign the token. must be one of ["HS256", "HS384", "HS512", "RS256", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]`,
					},
					"channel_token_upstream_header": schema.StringAttribute{
						Computed:    true,
						Description: `This plugin removes the ` + "`" + `config.channel_token_request_header` + "`" + ` from the request after reading its value.`,
					},
					"channel_token_upstream_leeway": schema.NumberAttribute{
						Computed:    true,
						Description: `If you want to add or perhaps subtract (using negative value) expiry time of the original channel token, you can specify a value that is added to the original channel token's ` + "`" + `exp` + "`" + ` claim.`,
					},
					"enable_access_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `If you don't want to support opaque access tokens, change this configuration parameter to ` + "`" + `false` + "`" + ` to disable introspection.`,
					},
					"enable_channel_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `If you don't want to support opaque channel tokens, disable introspection by changing this configuration parameter to ` + "`" + `false` + "`" + `.`,
					},
					"enable_hs_signatures": schema.BoolAttribute{
						Computed:    true,
						Description: `Tokens signed with HMAC algorithms such as ` + "`" + `HS256` + "`" + `, ` + "`" + `HS384` + "`" + `, or ` + "`" + `HS512` + "`" + ` are not accepted by default. If you need to accept such tokens for verification, enable this setting.`,
					},
					"enable_instrumentation": schema.BoolAttribute{
						Computed:    true,
						Description: `Writes log entries with some added information using ` + "`" + `ngx.CRIT` + "`" + ` (CRITICAL) level.`,
					},
					"original_access_token_upstream_header": schema.StringAttribute{
						Computed:    true,
						Description: `The HTTP header name used to store the original access token.`,
					},
					"original_channel_token_upstream_header": schema.StringAttribute{
						Computed:    true,
						Description: `The HTTP header name used to store the original channel token.`,
					},
					"realm": schema.StringAttribute{
						Computed:    true,
						Description: `When authentication or authorization fails, or there is an unexpected error, the plugin sends an ` + "`" + `WWW-Authenticate` + "`" + ` header with the ` + "`" + `realm` + "`" + ` attribute value.`,
					},
					"remove_access_token_claims": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `remove claims. It should be an array, and each element is a claim key string.`,
					},
					"remove_channel_token_claims": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `remove claims. It should be an array, and each element is a claim key string.`,
					},
					"set_access_token_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Set customized claims. If a claim is already present, it will be overwritten. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"set_channel_token_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Set customized claims. If a claim is already present, it will be overwritten. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"set_claims": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Set customized claims to both tokens. If a claim is already present, it will be overwritten. Value can be a regular or JSON string; if JSON, decoded data is used as the claim's value.`,
					},
					"trust_access_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `Use this parameter to enable and disable further checks on a payload before the new token is signed. If you set this to ` + "`" + `true` + "`" + `, the expiry or scopes are not checked on a payload.`,
					},
					"trust_channel_token_introspection": schema.BoolAttribute{
						Computed:    true,
						Description: `Providing an opaque channel token for plugin introspection, and verifying expiry and scopes on introspection results may make further payload checks unnecessary before the plugin signs a new token. This also applies when using a JWT token with introspection JSON as per config.channel_token_introspection_jwt_claim. Use this parameter to manage additional payload checks before signing a new token. With true (default), payload's expiry or scopes aren't checked.`,
					},
					"verify_access_token_expiry": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn access token expiry verification off and on as needed.`,
					},
					"verify_access_token_introspection_expiry": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn access token introspection expiry verification off and on as needed.`,
					},
					"verify_access_token_introspection_scopes": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn off and on the access token introspection scopes verification, specified with ` + "`" + `config.access_token_introspection_scopes_required` + "`" + `.`,
					},
					"verify_access_token_scopes": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn off and on the access token required scopes verification, specified with ` + "`" + `config.access_token_scopes_required` + "`" + `.`,
					},
					"verify_access_token_signature": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn access token signature verification off and on as needed.`,
					},
					"verify_channel_token_expiry": schema.BoolAttribute{
						Computed: true,
					},
					"verify_channel_token_introspection_expiry": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn on/off the channel token introspection expiry verification.`,
					},
					"verify_channel_token_introspection_scopes": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn on/off the channel token introspection scopes verification specified with ` + "`" + `config.channel_token_introspection_scopes_required` + "`" + `.`,
					},
					"verify_channel_token_scopes": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn on/off the channel token required scopes verification specified with ` + "`" + `config.channel_token_scopes_required` + "`" + `.`,
					},
					"verify_channel_token_signature": schema.BoolAttribute{
						Computed:    true,
						Description: `Quickly turn on/off the channel token signature verification.`,
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

func (r *PluginJWTSignerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginJWTSignerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginJWTSignerDataSourceModel
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
	request := operations.GetJwtsignerPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetJwtsignerPlugin(ctx, request)
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
	if !(res.JWTSignerPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedJWTSignerPlugin(res.JWTSignerPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
