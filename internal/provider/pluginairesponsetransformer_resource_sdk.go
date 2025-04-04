// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginAiResponseTransformerResourceModel) ToSharedAiResponseTransformerPlugin() *shared.AiResponseTransformerPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.AiResponseTransformerPluginOrdering
	if r.Ordering != nil {
		var after *shared.AiResponseTransformerPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiResponseTransformerPluginAfter{
				Access: access,
			}
		}
		var before *shared.AiResponseTransformerPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiResponseTransformerPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiResponseTransformerPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.AiResponseTransformerPluginPartials = []shared.AiResponseTransformerPluginPartials{}
	for _, partialsItem := range r.Partials {
		id1 := new(string)
		if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
			*id1 = partialsItem.ID.ValueString()
		} else {
			id1 = nil
		}
		name := new(string)
		if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
			*name = partialsItem.Name.ValueString()
		} else {
			name = nil
		}
		path := new(string)
		if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
			*path = partialsItem.Path.ValueString()
		} else {
			path = nil
		}
		partials = append(partials, shared.AiResponseTransformerPluginPartials{
			ID:   id1,
			Name: name,
			Path: path,
		})
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.AiResponseTransformerPluginConfig
	if r.Config != nil {
		httpProxyHost := new(string)
		if !r.Config.HTTPProxyHost.IsUnknown() && !r.Config.HTTPProxyHost.IsNull() {
			*httpProxyHost = r.Config.HTTPProxyHost.ValueString()
		} else {
			httpProxyHost = nil
		}
		httpProxyPort := new(int64)
		if !r.Config.HTTPProxyPort.IsUnknown() && !r.Config.HTTPProxyPort.IsNull() {
			*httpProxyPort = r.Config.HTTPProxyPort.ValueInt64()
		} else {
			httpProxyPort = nil
		}
		httpTimeout := new(int64)
		if !r.Config.HTTPTimeout.IsUnknown() && !r.Config.HTTPTimeout.IsNull() {
			*httpTimeout = r.Config.HTTPTimeout.ValueInt64()
		} else {
			httpTimeout = nil
		}
		httpsProxyHost := new(string)
		if !r.Config.HTTPSProxyHost.IsUnknown() && !r.Config.HTTPSProxyHost.IsNull() {
			*httpsProxyHost = r.Config.HTTPSProxyHost.ValueString()
		} else {
			httpsProxyHost = nil
		}
		httpsProxyPort := new(int64)
		if !r.Config.HTTPSProxyPort.IsUnknown() && !r.Config.HTTPSProxyPort.IsNull() {
			*httpsProxyPort = r.Config.HTTPSProxyPort.ValueInt64()
		} else {
			httpsProxyPort = nil
		}
		httpsVerify := new(bool)
		if !r.Config.HTTPSVerify.IsUnknown() && !r.Config.HTTPSVerify.IsNull() {
			*httpsVerify = r.Config.HTTPSVerify.ValueBool()
		} else {
			httpsVerify = nil
		}
		var llm *shared.AiResponseTransformerPluginLlm
		if r.Config.Llm != nil {
			var auth *shared.AiResponseTransformerPluginAuth
			if r.Config.Llm.Auth != nil {
				allowOverride := new(bool)
				if !r.Config.Llm.Auth.AllowOverride.IsUnknown() && !r.Config.Llm.Auth.AllowOverride.IsNull() {
					*allowOverride = r.Config.Llm.Auth.AllowOverride.ValueBool()
				} else {
					allowOverride = nil
				}
				awsAccessKeyID := new(string)
				if !r.Config.Llm.Auth.AwsAccessKeyID.IsUnknown() && !r.Config.Llm.Auth.AwsAccessKeyID.IsNull() {
					*awsAccessKeyID = r.Config.Llm.Auth.AwsAccessKeyID.ValueString()
				} else {
					awsAccessKeyID = nil
				}
				awsSecretAccessKey := new(string)
				if !r.Config.Llm.Auth.AwsSecretAccessKey.IsUnknown() && !r.Config.Llm.Auth.AwsSecretAccessKey.IsNull() {
					*awsSecretAccessKey = r.Config.Llm.Auth.AwsSecretAccessKey.ValueString()
				} else {
					awsSecretAccessKey = nil
				}
				azureClientID := new(string)
				if !r.Config.Llm.Auth.AzureClientID.IsUnknown() && !r.Config.Llm.Auth.AzureClientID.IsNull() {
					*azureClientID = r.Config.Llm.Auth.AzureClientID.ValueString()
				} else {
					azureClientID = nil
				}
				azureClientSecret := new(string)
				if !r.Config.Llm.Auth.AzureClientSecret.IsUnknown() && !r.Config.Llm.Auth.AzureClientSecret.IsNull() {
					*azureClientSecret = r.Config.Llm.Auth.AzureClientSecret.ValueString()
				} else {
					azureClientSecret = nil
				}
				azureTenantID := new(string)
				if !r.Config.Llm.Auth.AzureTenantID.IsUnknown() && !r.Config.Llm.Auth.AzureTenantID.IsNull() {
					*azureTenantID = r.Config.Llm.Auth.AzureTenantID.ValueString()
				} else {
					azureTenantID = nil
				}
				azureUseManagedIdentity := new(bool)
				if !r.Config.Llm.Auth.AzureUseManagedIdentity.IsUnknown() && !r.Config.Llm.Auth.AzureUseManagedIdentity.IsNull() {
					*azureUseManagedIdentity = r.Config.Llm.Auth.AzureUseManagedIdentity.ValueBool()
				} else {
					azureUseManagedIdentity = nil
				}
				gcpServiceAccountJSON := new(string)
				if !r.Config.Llm.Auth.GcpServiceAccountJSON.IsUnknown() && !r.Config.Llm.Auth.GcpServiceAccountJSON.IsNull() {
					*gcpServiceAccountJSON = r.Config.Llm.Auth.GcpServiceAccountJSON.ValueString()
				} else {
					gcpServiceAccountJSON = nil
				}
				gcpUseServiceAccount := new(bool)
				if !r.Config.Llm.Auth.GcpUseServiceAccount.IsUnknown() && !r.Config.Llm.Auth.GcpUseServiceAccount.IsNull() {
					*gcpUseServiceAccount = r.Config.Llm.Auth.GcpUseServiceAccount.ValueBool()
				} else {
					gcpUseServiceAccount = nil
				}
				headerName := new(string)
				if !r.Config.Llm.Auth.HeaderName.IsUnknown() && !r.Config.Llm.Auth.HeaderName.IsNull() {
					*headerName = r.Config.Llm.Auth.HeaderName.ValueString()
				} else {
					headerName = nil
				}
				headerValue := new(string)
				if !r.Config.Llm.Auth.HeaderValue.IsUnknown() && !r.Config.Llm.Auth.HeaderValue.IsNull() {
					*headerValue = r.Config.Llm.Auth.HeaderValue.ValueString()
				} else {
					headerValue = nil
				}
				paramLocation := new(shared.AiResponseTransformerPluginParamLocation)
				if !r.Config.Llm.Auth.ParamLocation.IsUnknown() && !r.Config.Llm.Auth.ParamLocation.IsNull() {
					*paramLocation = shared.AiResponseTransformerPluginParamLocation(r.Config.Llm.Auth.ParamLocation.ValueString())
				} else {
					paramLocation = nil
				}
				paramName := new(string)
				if !r.Config.Llm.Auth.ParamName.IsUnknown() && !r.Config.Llm.Auth.ParamName.IsNull() {
					*paramName = r.Config.Llm.Auth.ParamName.ValueString()
				} else {
					paramName = nil
				}
				paramValue := new(string)
				if !r.Config.Llm.Auth.ParamValue.IsUnknown() && !r.Config.Llm.Auth.ParamValue.IsNull() {
					*paramValue = r.Config.Llm.Auth.ParamValue.ValueString()
				} else {
					paramValue = nil
				}
				auth = &shared.AiResponseTransformerPluginAuth{
					AllowOverride:           allowOverride,
					AwsAccessKeyID:          awsAccessKeyID,
					AwsSecretAccessKey:      awsSecretAccessKey,
					AzureClientID:           azureClientID,
					AzureClientSecret:       azureClientSecret,
					AzureTenantID:           azureTenantID,
					AzureUseManagedIdentity: azureUseManagedIdentity,
					GcpServiceAccountJSON:   gcpServiceAccountJSON,
					GcpUseServiceAccount:    gcpUseServiceAccount,
					HeaderName:              headerName,
					HeaderValue:             headerValue,
					ParamLocation:           paramLocation,
					ParamName:               paramName,
					ParamValue:              paramValue,
				}
			}
			var logging *shared.AiResponseTransformerPluginLogging
			if r.Config.Llm.Logging != nil {
				logPayloads := new(bool)
				if !r.Config.Llm.Logging.LogPayloads.IsUnknown() && !r.Config.Llm.Logging.LogPayloads.IsNull() {
					*logPayloads = r.Config.Llm.Logging.LogPayloads.ValueBool()
				} else {
					logPayloads = nil
				}
				logStatistics := new(bool)
				if !r.Config.Llm.Logging.LogStatistics.IsUnknown() && !r.Config.Llm.Logging.LogStatistics.IsNull() {
					*logStatistics = r.Config.Llm.Logging.LogStatistics.ValueBool()
				} else {
					logStatistics = nil
				}
				logging = &shared.AiResponseTransformerPluginLogging{
					LogPayloads:   logPayloads,
					LogStatistics: logStatistics,
				}
			}
			var model *shared.AiResponseTransformerPluginModel
			if r.Config.Llm.Model != nil {
				name1 := new(string)
				if !r.Config.Llm.Model.Name.IsUnknown() && !r.Config.Llm.Model.Name.IsNull() {
					*name1 = r.Config.Llm.Model.Name.ValueString()
				} else {
					name1 = nil
				}
				var optionsVar *shared.AiResponseTransformerPluginOptions
				if r.Config.Llm.Model.Options != nil {
					anthropicVersion := new(string)
					if !r.Config.Llm.Model.Options.AnthropicVersion.IsUnknown() && !r.Config.Llm.Model.Options.AnthropicVersion.IsNull() {
						*anthropicVersion = r.Config.Llm.Model.Options.AnthropicVersion.ValueString()
					} else {
						anthropicVersion = nil
					}
					azureAPIVersion := new(string)
					if !r.Config.Llm.Model.Options.AzureAPIVersion.IsUnknown() && !r.Config.Llm.Model.Options.AzureAPIVersion.IsNull() {
						*azureAPIVersion = r.Config.Llm.Model.Options.AzureAPIVersion.ValueString()
					} else {
						azureAPIVersion = nil
					}
					azureDeploymentID := new(string)
					if !r.Config.Llm.Model.Options.AzureDeploymentID.IsUnknown() && !r.Config.Llm.Model.Options.AzureDeploymentID.IsNull() {
						*azureDeploymentID = r.Config.Llm.Model.Options.AzureDeploymentID.ValueString()
					} else {
						azureDeploymentID = nil
					}
					azureInstance := new(string)
					if !r.Config.Llm.Model.Options.AzureInstance.IsUnknown() && !r.Config.Llm.Model.Options.AzureInstance.IsNull() {
						*azureInstance = r.Config.Llm.Model.Options.AzureInstance.ValueString()
					} else {
						azureInstance = nil
					}
					var bedrock *shared.AiResponseTransformerPluginBedrock
					if r.Config.Llm.Model.Options.Bedrock != nil {
						awsRegion := new(string)
						if !r.Config.Llm.Model.Options.Bedrock.AwsRegion.IsUnknown() && !r.Config.Llm.Model.Options.Bedrock.AwsRegion.IsNull() {
							*awsRegion = r.Config.Llm.Model.Options.Bedrock.AwsRegion.ValueString()
						} else {
							awsRegion = nil
						}
						bedrock = &shared.AiResponseTransformerPluginBedrock{
							AwsRegion: awsRegion,
						}
					}
					var gemini *shared.AiResponseTransformerPluginGemini
					if r.Config.Llm.Model.Options.Gemini != nil {
						apiEndpoint := new(string)
						if !r.Config.Llm.Model.Options.Gemini.APIEndpoint.IsUnknown() && !r.Config.Llm.Model.Options.Gemini.APIEndpoint.IsNull() {
							*apiEndpoint = r.Config.Llm.Model.Options.Gemini.APIEndpoint.ValueString()
						} else {
							apiEndpoint = nil
						}
						locationID := new(string)
						if !r.Config.Llm.Model.Options.Gemini.LocationID.IsUnknown() && !r.Config.Llm.Model.Options.Gemini.LocationID.IsNull() {
							*locationID = r.Config.Llm.Model.Options.Gemini.LocationID.ValueString()
						} else {
							locationID = nil
						}
						projectID := new(string)
						if !r.Config.Llm.Model.Options.Gemini.ProjectID.IsUnknown() && !r.Config.Llm.Model.Options.Gemini.ProjectID.IsNull() {
							*projectID = r.Config.Llm.Model.Options.Gemini.ProjectID.ValueString()
						} else {
							projectID = nil
						}
						gemini = &shared.AiResponseTransformerPluginGemini{
							APIEndpoint: apiEndpoint,
							LocationID:  locationID,
							ProjectID:   projectID,
						}
					}
					var huggingface *shared.AiResponseTransformerPluginHuggingface
					if r.Config.Llm.Model.Options.Huggingface != nil {
						useCache := new(bool)
						if !r.Config.Llm.Model.Options.Huggingface.UseCache.IsUnknown() && !r.Config.Llm.Model.Options.Huggingface.UseCache.IsNull() {
							*useCache = r.Config.Llm.Model.Options.Huggingface.UseCache.ValueBool()
						} else {
							useCache = nil
						}
						waitForModel := new(bool)
						if !r.Config.Llm.Model.Options.Huggingface.WaitForModel.IsUnknown() && !r.Config.Llm.Model.Options.Huggingface.WaitForModel.IsNull() {
							*waitForModel = r.Config.Llm.Model.Options.Huggingface.WaitForModel.ValueBool()
						} else {
							waitForModel = nil
						}
						huggingface = &shared.AiResponseTransformerPluginHuggingface{
							UseCache:     useCache,
							WaitForModel: waitForModel,
						}
					}
					inputCost := new(float64)
					if !r.Config.Llm.Model.Options.InputCost.IsUnknown() && !r.Config.Llm.Model.Options.InputCost.IsNull() {
						*inputCost, _ = r.Config.Llm.Model.Options.InputCost.ValueBigFloat().Float64()
					} else {
						inputCost = nil
					}
					llama2Format := new(shared.AiResponseTransformerPluginLlama2Format)
					if !r.Config.Llm.Model.Options.Llama2Format.IsUnknown() && !r.Config.Llm.Model.Options.Llama2Format.IsNull() {
						*llama2Format = shared.AiResponseTransformerPluginLlama2Format(r.Config.Llm.Model.Options.Llama2Format.ValueString())
					} else {
						llama2Format = nil
					}
					maxTokens := new(int64)
					if !r.Config.Llm.Model.Options.MaxTokens.IsUnknown() && !r.Config.Llm.Model.Options.MaxTokens.IsNull() {
						*maxTokens = r.Config.Llm.Model.Options.MaxTokens.ValueInt64()
					} else {
						maxTokens = nil
					}
					mistralFormat := new(shared.AiResponseTransformerPluginMistralFormat)
					if !r.Config.Llm.Model.Options.MistralFormat.IsUnknown() && !r.Config.Llm.Model.Options.MistralFormat.IsNull() {
						*mistralFormat = shared.AiResponseTransformerPluginMistralFormat(r.Config.Llm.Model.Options.MistralFormat.ValueString())
					} else {
						mistralFormat = nil
					}
					outputCost := new(float64)
					if !r.Config.Llm.Model.Options.OutputCost.IsUnknown() && !r.Config.Llm.Model.Options.OutputCost.IsNull() {
						*outputCost, _ = r.Config.Llm.Model.Options.OutputCost.ValueBigFloat().Float64()
					} else {
						outputCost = nil
					}
					temperature := new(float64)
					if !r.Config.Llm.Model.Options.Temperature.IsUnknown() && !r.Config.Llm.Model.Options.Temperature.IsNull() {
						*temperature, _ = r.Config.Llm.Model.Options.Temperature.ValueBigFloat().Float64()
					} else {
						temperature = nil
					}
					topK := new(int64)
					if !r.Config.Llm.Model.Options.TopK.IsUnknown() && !r.Config.Llm.Model.Options.TopK.IsNull() {
						*topK = r.Config.Llm.Model.Options.TopK.ValueInt64()
					} else {
						topK = nil
					}
					topP := new(float64)
					if !r.Config.Llm.Model.Options.TopP.IsUnknown() && !r.Config.Llm.Model.Options.TopP.IsNull() {
						*topP, _ = r.Config.Llm.Model.Options.TopP.ValueBigFloat().Float64()
					} else {
						topP = nil
					}
					upstreamPath := new(string)
					if !r.Config.Llm.Model.Options.UpstreamPath.IsUnknown() && !r.Config.Llm.Model.Options.UpstreamPath.IsNull() {
						*upstreamPath = r.Config.Llm.Model.Options.UpstreamPath.ValueString()
					} else {
						upstreamPath = nil
					}
					upstreamURL := new(string)
					if !r.Config.Llm.Model.Options.UpstreamURL.IsUnknown() && !r.Config.Llm.Model.Options.UpstreamURL.IsNull() {
						*upstreamURL = r.Config.Llm.Model.Options.UpstreamURL.ValueString()
					} else {
						upstreamURL = nil
					}
					optionsVar = &shared.AiResponseTransformerPluginOptions{
						AnthropicVersion:  anthropicVersion,
						AzureAPIVersion:   azureAPIVersion,
						AzureDeploymentID: azureDeploymentID,
						AzureInstance:     azureInstance,
						Bedrock:           bedrock,
						Gemini:            gemini,
						Huggingface:       huggingface,
						InputCost:         inputCost,
						Llama2Format:      llama2Format,
						MaxTokens:         maxTokens,
						MistralFormat:     mistralFormat,
						OutputCost:        outputCost,
						Temperature:       temperature,
						TopK:              topK,
						TopP:              topP,
						UpstreamPath:      upstreamPath,
						UpstreamURL:       upstreamURL,
					}
				}
				provider := new(shared.AiResponseTransformerPluginProvider)
				if !r.Config.Llm.Model.Provider.IsUnknown() && !r.Config.Llm.Model.Provider.IsNull() {
					*provider = shared.AiResponseTransformerPluginProvider(r.Config.Llm.Model.Provider.ValueString())
				} else {
					provider = nil
				}
				model = &shared.AiResponseTransformerPluginModel{
					Name:     name1,
					Options:  optionsVar,
					Provider: provider,
				}
			}
			routeType := new(shared.AiResponseTransformerPluginRouteType)
			if !r.Config.Llm.RouteType.IsUnknown() && !r.Config.Llm.RouteType.IsNull() {
				*routeType = shared.AiResponseTransformerPluginRouteType(r.Config.Llm.RouteType.ValueString())
			} else {
				routeType = nil
			}
			llm = &shared.AiResponseTransformerPluginLlm{
				Auth:      auth,
				Logging:   logging,
				Model:     model,
				RouteType: routeType,
			}
		}
		maxRequestBodySize := new(int64)
		if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
			*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
		} else {
			maxRequestBodySize = nil
		}
		parseLlmResponseJSONInstructions := new(bool)
		if !r.Config.ParseLlmResponseJSONInstructions.IsUnknown() && !r.Config.ParseLlmResponseJSONInstructions.IsNull() {
			*parseLlmResponseJSONInstructions = r.Config.ParseLlmResponseJSONInstructions.ValueBool()
		} else {
			parseLlmResponseJSONInstructions = nil
		}
		prompt := new(string)
		if !r.Config.Prompt.IsUnknown() && !r.Config.Prompt.IsNull() {
			*prompt = r.Config.Prompt.ValueString()
		} else {
			prompt = nil
		}
		transformationExtractPattern := new(string)
		if !r.Config.TransformationExtractPattern.IsUnknown() && !r.Config.TransformationExtractPattern.IsNull() {
			*transformationExtractPattern = r.Config.TransformationExtractPattern.ValueString()
		} else {
			transformationExtractPattern = nil
		}
		config = &shared.AiResponseTransformerPluginConfig{
			HTTPProxyHost:                    httpProxyHost,
			HTTPProxyPort:                    httpProxyPort,
			HTTPTimeout:                      httpTimeout,
			HTTPSProxyHost:                   httpsProxyHost,
			HTTPSProxyPort:                   httpsProxyPort,
			HTTPSVerify:                      httpsVerify,
			Llm:                              llm,
			MaxRequestBodySize:               maxRequestBodySize,
			ParseLlmResponseJSONInstructions: parseLlmResponseJSONInstructions,
			Prompt:                           prompt,
			TransformationExtractPattern:     transformationExtractPattern,
		}
	}
	var consumer *shared.AiResponseTransformerPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.AiResponseTransformerPluginConsumer{
			ID: id2,
		}
	}
	var consumerGroup *shared.AiResponseTransformerPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id3 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id3 = r.ConsumerGroup.ID.ValueString()
		} else {
			id3 = nil
		}
		consumerGroup = &shared.AiResponseTransformerPluginConsumerGroup{
			ID: id3,
		}
	}
	var protocols []shared.AiResponseTransformerPluginProtocols = []shared.AiResponseTransformerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiResponseTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiResponseTransformerPluginRoute
	if r.Route != nil {
		id4 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id4 = r.Route.ID.ValueString()
		} else {
			id4 = nil
		}
		route = &shared.AiResponseTransformerPluginRoute{
			ID: id4,
		}
	}
	var service *shared.AiResponseTransformerPluginService
	if r.Service != nil {
		id5 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id5 = r.Service.ID.ValueString()
		} else {
			id5 = nil
		}
		service = &shared.AiResponseTransformerPluginService{
			ID: id5,
		}
	}
	out := shared.AiResponseTransformerPlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Partials:      partials,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *PluginAiResponseTransformerResourceModel) RefreshFromSharedAiResponseTransformerPlugin(resp *shared.AiResponseTransformerPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiResponseTransformerPluginConfig{}
			r.Config.HTTPProxyHost = types.StringPointerValue(resp.Config.HTTPProxyHost)
			r.Config.HTTPProxyPort = types.Int64PointerValue(resp.Config.HTTPProxyPort)
			r.Config.HTTPTimeout = types.Int64PointerValue(resp.Config.HTTPTimeout)
			r.Config.HTTPSProxyHost = types.StringPointerValue(resp.Config.HTTPSProxyHost)
			r.Config.HTTPSProxyPort = types.Int64PointerValue(resp.Config.HTTPSProxyPort)
			r.Config.HTTPSVerify = types.BoolPointerValue(resp.Config.HTTPSVerify)
			if resp.Config.Llm == nil {
				r.Config.Llm = nil
			} else {
				r.Config.Llm = &tfTypes.Llm{}
				if resp.Config.Llm.Auth == nil {
					r.Config.Llm.Auth = nil
				} else {
					r.Config.Llm.Auth = &tfTypes.Auth{}
					r.Config.Llm.Auth.AllowOverride = types.BoolPointerValue(resp.Config.Llm.Auth.AllowOverride)
					r.Config.Llm.Auth.AwsAccessKeyID = types.StringPointerValue(resp.Config.Llm.Auth.AwsAccessKeyID)
					r.Config.Llm.Auth.AwsSecretAccessKey = types.StringPointerValue(resp.Config.Llm.Auth.AwsSecretAccessKey)
					r.Config.Llm.Auth.AzureClientID = types.StringPointerValue(resp.Config.Llm.Auth.AzureClientID)
					r.Config.Llm.Auth.AzureClientSecret = types.StringPointerValue(resp.Config.Llm.Auth.AzureClientSecret)
					r.Config.Llm.Auth.AzureTenantID = types.StringPointerValue(resp.Config.Llm.Auth.AzureTenantID)
					r.Config.Llm.Auth.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.Llm.Auth.AzureUseManagedIdentity)
					r.Config.Llm.Auth.GcpServiceAccountJSON = types.StringPointerValue(resp.Config.Llm.Auth.GcpServiceAccountJSON)
					r.Config.Llm.Auth.GcpUseServiceAccount = types.BoolPointerValue(resp.Config.Llm.Auth.GcpUseServiceAccount)
					r.Config.Llm.Auth.HeaderName = types.StringPointerValue(resp.Config.Llm.Auth.HeaderName)
					r.Config.Llm.Auth.HeaderValue = types.StringPointerValue(resp.Config.Llm.Auth.HeaderValue)
					if resp.Config.Llm.Auth.ParamLocation != nil {
						r.Config.Llm.Auth.ParamLocation = types.StringValue(string(*resp.Config.Llm.Auth.ParamLocation))
					} else {
						r.Config.Llm.Auth.ParamLocation = types.StringNull()
					}
					r.Config.Llm.Auth.ParamName = types.StringPointerValue(resp.Config.Llm.Auth.ParamName)
					r.Config.Llm.Auth.ParamValue = types.StringPointerValue(resp.Config.Llm.Auth.ParamValue)
				}
				if resp.Config.Llm.Logging == nil {
					r.Config.Llm.Logging = nil
				} else {
					r.Config.Llm.Logging = &tfTypes.Logging{}
					r.Config.Llm.Logging.LogPayloads = types.BoolPointerValue(resp.Config.Llm.Logging.LogPayloads)
					r.Config.Llm.Logging.LogStatistics = types.BoolPointerValue(resp.Config.Llm.Logging.LogStatistics)
				}
				if resp.Config.Llm.Model == nil {
					r.Config.Llm.Model = nil
				} else {
					r.Config.Llm.Model = &tfTypes.Model{}
					r.Config.Llm.Model.Name = types.StringPointerValue(resp.Config.Llm.Model.Name)
					if resp.Config.Llm.Model.Options == nil {
						r.Config.Llm.Model.Options = nil
					} else {
						r.Config.Llm.Model.Options = &tfTypes.OptionsObj{}
						r.Config.Llm.Model.Options.AnthropicVersion = types.StringPointerValue(resp.Config.Llm.Model.Options.AnthropicVersion)
						r.Config.Llm.Model.Options.AzureAPIVersion = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureAPIVersion)
						r.Config.Llm.Model.Options.AzureDeploymentID = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureDeploymentID)
						r.Config.Llm.Model.Options.AzureInstance = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureInstance)
						if resp.Config.Llm.Model.Options.Bedrock == nil {
							r.Config.Llm.Model.Options.Bedrock = nil
						} else {
							r.Config.Llm.Model.Options.Bedrock = &tfTypes.Bedrock{}
							r.Config.Llm.Model.Options.Bedrock.AwsRegion = types.StringPointerValue(resp.Config.Llm.Model.Options.Bedrock.AwsRegion)
						}
						if resp.Config.Llm.Model.Options.Gemini == nil {
							r.Config.Llm.Model.Options.Gemini = nil
						} else {
							r.Config.Llm.Model.Options.Gemini = &tfTypes.Gemini{}
							r.Config.Llm.Model.Options.Gemini.APIEndpoint = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.APIEndpoint)
							r.Config.Llm.Model.Options.Gemini.LocationID = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.LocationID)
							r.Config.Llm.Model.Options.Gemini.ProjectID = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.ProjectID)
						}
						if resp.Config.Llm.Model.Options.Huggingface == nil {
							r.Config.Llm.Model.Options.Huggingface = nil
						} else {
							r.Config.Llm.Model.Options.Huggingface = &tfTypes.Huggingface{}
							r.Config.Llm.Model.Options.Huggingface.UseCache = types.BoolPointerValue(resp.Config.Llm.Model.Options.Huggingface.UseCache)
							r.Config.Llm.Model.Options.Huggingface.WaitForModel = types.BoolPointerValue(resp.Config.Llm.Model.Options.Huggingface.WaitForModel)
						}
						if resp.Config.Llm.Model.Options.InputCost != nil {
							r.Config.Llm.Model.Options.InputCost = types.NumberValue(big.NewFloat(float64(*resp.Config.Llm.Model.Options.InputCost)))
						} else {
							r.Config.Llm.Model.Options.InputCost = types.NumberNull()
						}
						if resp.Config.Llm.Model.Options.Llama2Format != nil {
							r.Config.Llm.Model.Options.Llama2Format = types.StringValue(string(*resp.Config.Llm.Model.Options.Llama2Format))
						} else {
							r.Config.Llm.Model.Options.Llama2Format = types.StringNull()
						}
						r.Config.Llm.Model.Options.MaxTokens = types.Int64PointerValue(resp.Config.Llm.Model.Options.MaxTokens)
						if resp.Config.Llm.Model.Options.MistralFormat != nil {
							r.Config.Llm.Model.Options.MistralFormat = types.StringValue(string(*resp.Config.Llm.Model.Options.MistralFormat))
						} else {
							r.Config.Llm.Model.Options.MistralFormat = types.StringNull()
						}
						if resp.Config.Llm.Model.Options.OutputCost != nil {
							r.Config.Llm.Model.Options.OutputCost = types.NumberValue(big.NewFloat(float64(*resp.Config.Llm.Model.Options.OutputCost)))
						} else {
							r.Config.Llm.Model.Options.OutputCost = types.NumberNull()
						}
						if resp.Config.Llm.Model.Options.Temperature != nil {
							r.Config.Llm.Model.Options.Temperature = types.NumberValue(big.NewFloat(float64(*resp.Config.Llm.Model.Options.Temperature)))
						} else {
							r.Config.Llm.Model.Options.Temperature = types.NumberNull()
						}
						r.Config.Llm.Model.Options.TopK = types.Int64PointerValue(resp.Config.Llm.Model.Options.TopK)
						if resp.Config.Llm.Model.Options.TopP != nil {
							r.Config.Llm.Model.Options.TopP = types.NumberValue(big.NewFloat(float64(*resp.Config.Llm.Model.Options.TopP)))
						} else {
							r.Config.Llm.Model.Options.TopP = types.NumberNull()
						}
						r.Config.Llm.Model.Options.UpstreamPath = types.StringPointerValue(resp.Config.Llm.Model.Options.UpstreamPath)
						r.Config.Llm.Model.Options.UpstreamURL = types.StringPointerValue(resp.Config.Llm.Model.Options.UpstreamURL)
					}
					if resp.Config.Llm.Model.Provider != nil {
						r.Config.Llm.Model.Provider = types.StringValue(string(*resp.Config.Llm.Model.Provider))
					} else {
						r.Config.Llm.Model.Provider = types.StringNull()
					}
				}
				if resp.Config.Llm.RouteType != nil {
					r.Config.Llm.RouteType = types.StringValue(string(*resp.Config.Llm.RouteType))
				} else {
					r.Config.Llm.RouteType = types.StringNull()
				}
			}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			r.Config.ParseLlmResponseJSONInstructions = types.BoolPointerValue(resp.Config.ParseLlmResponseJSONInstructions)
			r.Config.Prompt = types.StringPointerValue(resp.Config.Prompt)
			r.Config.TransformationExtractPattern = types.StringPointerValue(resp.Config.TransformationExtractPattern)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.Ordering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.After{}
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.After{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials1 tfTypes.Partials
				partials1.ID = types.StringPointerValue(partialsItem.ID)
				partials1.Name = types.StringPointerValue(partialsItem.Name)
				partials1.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials1)
				} else {
					r.Partials[partialsCount].ID = partials1.ID
					r.Partials[partialsCount].Name = partials1.Name
					r.Partials[partialsCount].Path = partials1.Path
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
