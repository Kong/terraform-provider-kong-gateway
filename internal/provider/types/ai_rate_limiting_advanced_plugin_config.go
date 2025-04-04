// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AiRateLimitingAdvancedPluginConfig struct {
	DictionaryName             types.String                `tfsdk:"dictionary_name"`
	DisablePenalty             types.Bool                  `tfsdk:"disable_penalty"`
	ErrorCode                  types.Number                `tfsdk:"error_code"`
	ErrorHideProviders         types.Bool                  `tfsdk:"error_hide_providers"`
	ErrorMessage               types.String                `tfsdk:"error_message"`
	HeaderName                 types.String                `tfsdk:"header_name"`
	HideClientHeaders          types.Bool                  `tfsdk:"hide_client_headers"`
	Identifier                 types.String                `tfsdk:"identifier"`
	LlmProviders               []LlmProviders              `tfsdk:"llm_providers"`
	Path                       types.String                `tfsdk:"path"`
	Redis                      *AiProxyAdvancedPluginRedis `tfsdk:"redis"`
	RequestPromptCountFunction types.String                `tfsdk:"request_prompt_count_function"`
	RetryAfterJitterMax        types.Number                `tfsdk:"retry_after_jitter_max"`
	Strategy                   types.String                `tfsdk:"strategy"`
	SyncRate                   types.Number                `tfsdk:"sync_rate"`
	TokensCountStrategy        types.String                `tfsdk:"tokens_count_strategy"`
	WindowType                 types.String                `tfsdk:"window_type"`
}
