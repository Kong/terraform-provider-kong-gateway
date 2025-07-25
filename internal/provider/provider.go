// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

var _ provider.Provider = (*KongGatewayProvider)(nil)
var _ provider.ProviderWithEphemeralResources = (*KongGatewayProvider)(nil)

type KongGatewayProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KongGatewayProviderModel describes the provider data model.
type KongGatewayProviderModel struct {
	AdminToken types.String `tfsdk:"admin_token"`
	ServerURL  types.String `tfsdk:"server_url"`
}

func (p *KongGatewayProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "kong-gateway"
	resp.Version = p.version
}

func (p *KongGatewayProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Description: `Server URL (defaults to {protocol}://{hostname}:{port}{path})`,
				Optional:    true,
			},
		},
		MarkdownDescription: `Kong Gateway Admin API: OpenAPI 3.0 spec for Kong Gateway's Admin API.` + "\n" +
			`` + "\n" +
			`You can lean more about Kong Gateway at [developer.konghq.com](https://developer.konghq.com)` + "\n" +
			`.Give Kong a star at [Kong/kong](https://github.com/kong/kong) repository.`,
	}
}

func (p *KongGatewayProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KongGatewayProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "{protocol}://{hostname}:{port}{path}"
	}

	security := shared.Security{}

	if !data.AdminToken.IsUnknown() {
		security.AdminToken = data.AdminToken.ValueStringPointer()
	}

	providerHTTPTransportOpts := ProviderHTTPTransportOpts{
		SetHeaders: make(map[string]string),
		Transport:  http.DefaultTransport,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewProviderHTTPTransport(providerHTTPTransportOpts)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}

	client := sdk.New(opts...)
	resp.DataSourceData = client
	resp.EphemeralResourceData = client
	resp.ResourceData = client
}

func (p *KongGatewayProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewACLResource,
		NewBasicAuthResource,
		NewCACertificateResource,
		NewCertificateResource,
		NewConsumerResource,
		NewConsumerGroupResource,
		NewGatewayConsumerGroupMemberResource,
		NewHMACAuthResource,
		NewJwtResource,
		NewKeyResource,
		NewKeyAuthResource,
		NewKeySetResource,
		NewMTLSAuthResource,
		NewOidcJwkResource,
		NewPartialResource,
		NewPluginACLResource,
		NewPluginAcmeResource,
		NewPluginAiAzureContentSafetyResource,
		NewPluginAiPromptDecoratorResource,
		NewPluginAiPromptGuardResource,
		NewPluginAiPromptTemplateResource,
		NewPluginAiProxyResource,
		NewPluginAiProxyAdvancedResource,
		NewPluginAiRagInjectorResource,
		NewPluginAiRateLimitingAdvancedResource,
		NewPluginAiRequestTransformerResource,
		NewPluginAiResponseTransformerResource,
		NewPluginAiSanitizerResource,
		NewPluginAiSemanticCacheResource,
		NewPluginAiSemanticPromptGuardResource,
		NewPluginAwsLambdaResource,
		NewPluginAzureFunctionsResource,
		NewPluginBasicAuthResource,
		NewPluginBotDetectionResource,
		NewPluginCanaryResource,
		NewPluginConfluentResource,
		NewPluginConfluentConsumeResource,
		NewPluginCorrelationIDResource,
		NewPluginCorsResource,
		NewPluginDatadogResource,
		NewPluginDatadogTracingResource,
		NewPluginDegraphqlResource,
		NewPluginExitTransformerResource,
		NewPluginFileLogResource,
		NewPluginForwardProxyResource,
		NewPluginGraphqlProxyCacheAdvancedResource,
		NewPluginGraphqlRateLimitingAdvancedResource,
		NewPluginGrpcGatewayResource,
		NewPluginGrpcWebResource,
		NewPluginHeaderCertAuthResource,
		NewPluginHmacAuthResource,
		NewPluginHTTPLogResource,
		NewPluginInjectionProtectionResource,
		NewPluginIPRestrictionResource,
		NewPluginJqResource,
		NewPluginJSONThreatProtectionResource,
		NewPluginJweDecryptResource,
		NewPluginJwtResource,
		NewPluginJwtSignerResource,
		NewPluginKafkaConsumeResource,
		NewPluginKafkaLogResource,
		NewPluginKafkaUpstreamResource,
		NewPluginKeyAuthResource,
		NewPluginKeyAuthEncResource,
		NewPluginLdapAuthResource,
		NewPluginLdapAuthAdvancedResource,
		NewPluginLogglyResource,
		NewPluginMockingResource,
		NewPluginMtlsAuthResource,
		NewPluginOasValidationResource,
		NewPluginOauth2Resource,
		NewPluginOauth2IntrospectionResource,
		NewPluginOpaResource,
		NewPluginOpenidConnectResource,
		NewPluginOpentelemetryResource,
		NewPluginPostFunctionResource,
		NewPluginPreFunctionResource,
		NewPluginPrometheusResource,
		NewPluginProxyCacheResource,
		NewPluginProxyCacheAdvancedResource,
		NewPluginRateLimitingResource,
		NewPluginRateLimitingAdvancedResource,
		NewPluginRedirectResource,
		NewPluginRequestCalloutResource,
		NewPluginRequestSizeLimitingResource,
		NewPluginRequestTerminationResource,
		NewPluginRequestTransformerResource,
		NewPluginRequestTransformerAdvancedResource,
		NewPluginRequestValidatorResource,
		NewPluginResponseRatelimitingResource,
		NewPluginResponseTransformerResource,
		NewPluginResponseTransformerAdvancedResource,
		NewPluginRouteByHeaderResource,
		NewPluginRouteTransformerAdvancedResource,
		NewPluginSamlResource,
		NewPluginServiceProtectionResource,
		NewPluginSessionResource,
		NewPluginStandardWebhooksResource,
		NewPluginStatsdResource,
		NewPluginStatsdAdvancedResource,
		NewPluginSyslogResource,
		NewPluginTCPLogResource,
		NewPluginTLSHandshakeModifierResource,
		NewPluginTLSMetadataHeadersResource,
		NewPluginUDPLogResource,
		NewPluginUpstreamOauthResource,
		NewPluginUpstreamTimeoutResource,
		NewPluginVaultAuthResource,
		NewPluginWebsocketSizeLimitResource,
		NewPluginWebsocketValidatorResource,
		NewPluginXMLThreatProtectionResource,
		NewPluginZipkinResource,
		NewRouteResource,
		NewRouteExpressionResource,
		NewServiceResource,
		NewSniResource,
		NewTargetResource,
		NewUpstreamResource,
		NewVaultResource,
	}
}

func (p *KongGatewayProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewACLDataSource,
		NewBasicAuthDataSource,
		NewCACertificateDataSource,
		NewCertificateDataSource,
		NewConsumerDataSource,
		NewConsumerGroupDataSource,
		NewHMACAuthDataSource,
		NewJwtDataSource,
		NewKeyDataSource,
		NewKeyAuthDataSource,
		NewKeySetDataSource,
		NewMTLSAuthDataSource,
		NewOidcJwkDataSource,
		NewPartialDataSource,
		NewPluginACLDataSource,
		NewPluginAcmeDataSource,
		NewPluginAiAzureContentSafetyDataSource,
		NewPluginAiPromptDecoratorDataSource,
		NewPluginAiPromptGuardDataSource,
		NewPluginAiPromptTemplateDataSource,
		NewPluginAiProxyDataSource,
		NewPluginAiProxyAdvancedDataSource,
		NewPluginAiRagInjectorDataSource,
		NewPluginAiRateLimitingAdvancedDataSource,
		NewPluginAiRequestTransformerDataSource,
		NewPluginAiResponseTransformerDataSource,
		NewPluginAiSanitizerDataSource,
		NewPluginAiSemanticCacheDataSource,
		NewPluginAiSemanticPromptGuardDataSource,
		NewPluginAwsLambdaDataSource,
		NewPluginAzureFunctionsDataSource,
		NewPluginBasicAuthDataSource,
		NewPluginBotDetectionDataSource,
		NewPluginCanaryDataSource,
		NewPluginConfluentDataSource,
		NewPluginConfluentConsumeDataSource,
		NewPluginCorrelationIDDataSource,
		NewPluginCorsDataSource,
		NewPluginDatadogDataSource,
		NewPluginDatadogTracingDataSource,
		NewPluginDegraphqlDataSource,
		NewPluginExitTransformerDataSource,
		NewPluginFileLogDataSource,
		NewPluginForwardProxyDataSource,
		NewPluginGraphqlProxyCacheAdvancedDataSource,
		NewPluginGraphqlRateLimitingAdvancedDataSource,
		NewPluginGrpcGatewayDataSource,
		NewPluginGrpcWebDataSource,
		NewPluginHeaderCertAuthDataSource,
		NewPluginHmacAuthDataSource,
		NewPluginHTTPLogDataSource,
		NewPluginInjectionProtectionDataSource,
		NewPluginIPRestrictionDataSource,
		NewPluginJqDataSource,
		NewPluginJSONThreatProtectionDataSource,
		NewPluginJweDecryptDataSource,
		NewPluginJwtDataSource,
		NewPluginJwtSignerDataSource,
		NewPluginKafkaConsumeDataSource,
		NewPluginKafkaLogDataSource,
		NewPluginKafkaUpstreamDataSource,
		NewPluginKeyAuthDataSource,
		NewPluginKeyAuthEncDataSource,
		NewPluginLdapAuthDataSource,
		NewPluginLdapAuthAdvancedDataSource,
		NewPluginLogglyDataSource,
		NewPluginMockingDataSource,
		NewPluginMtlsAuthDataSource,
		NewPluginOasValidationDataSource,
		NewPluginOauth2DataSource,
		NewPluginOauth2IntrospectionDataSource,
		NewPluginOpaDataSource,
		NewPluginOpenidConnectDataSource,
		NewPluginOpentelemetryDataSource,
		NewPluginPostFunctionDataSource,
		NewPluginPreFunctionDataSource,
		NewPluginPrometheusDataSource,
		NewPluginProxyCacheDataSource,
		NewPluginProxyCacheAdvancedDataSource,
		NewPluginRateLimitingDataSource,
		NewPluginRateLimitingAdvancedDataSource,
		NewPluginRedirectDataSource,
		NewPluginRequestCalloutDataSource,
		NewPluginRequestSizeLimitingDataSource,
		NewPluginRequestTerminationDataSource,
		NewPluginRequestTransformerDataSource,
		NewPluginRequestTransformerAdvancedDataSource,
		NewPluginRequestValidatorDataSource,
		NewPluginResponseRatelimitingDataSource,
		NewPluginResponseTransformerDataSource,
		NewPluginResponseTransformerAdvancedDataSource,
		NewPluginRouteByHeaderDataSource,
		NewPluginRouteTransformerAdvancedDataSource,
		NewPluginSamlDataSource,
		NewPluginServiceProtectionDataSource,
		NewPluginSessionDataSource,
		NewPluginStandardWebhooksDataSource,
		NewPluginStatsdDataSource,
		NewPluginStatsdAdvancedDataSource,
		NewPluginSyslogDataSource,
		NewPluginTCPLogDataSource,
		NewPluginTLSHandshakeModifierDataSource,
		NewPluginTLSMetadataHeadersDataSource,
		NewPluginUDPLogDataSource,
		NewPluginUpstreamOauthDataSource,
		NewPluginUpstreamTimeoutDataSource,
		NewPluginVaultAuthDataSource,
		NewPluginWebsocketSizeLimitDataSource,
		NewPluginWebsocketValidatorDataSource,
		NewPluginXMLThreatProtectionDataSource,
		NewPluginZipkinDataSource,
		NewRouteDataSource,
		NewRouteExpressionDataSource,
		NewServiceDataSource,
		NewSniDataSource,
		NewTargetDataSource,
		NewUpstreamDataSource,
		NewVaultDataSource,
	}
}

func (p *KongGatewayProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KongGatewayProvider{
			version: version,
		}
	}
}
