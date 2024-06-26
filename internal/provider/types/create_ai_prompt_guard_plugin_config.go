// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateAIPromptGuardPluginConfig struct {
	AllowAllConversationHistory types.Bool     `tfsdk:"allow_all_conversation_history"`
	AllowPatterns               []types.String `tfsdk:"allow_patterns"`
	DenyPatterns                []types.String `tfsdk:"deny_patterns"`
}
