// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
)

func (r *GatewayConsumerGroupMemberResourceModel) ToOperationsAddConsumerToGroupRequestBody() *operations.AddConsumerToGroupRequestBody {
	consumerID := new(string)
	if !r.ConsumerID.IsUnknown() && !r.ConsumerID.IsNull() {
		*consumerID = r.ConsumerID.ValueString()
	} else {
		consumerID = nil
	}
	out := operations.AddConsumerToGroupRequestBody{
		ConsumerID: consumerID,
	}
	return &out
}

func (r *GatewayConsumerGroupMemberResourceModel) RefreshFromOperationsAddConsumerToGroupResponseBody(resp *operations.AddConsumerToGroupResponseBody) {
	if resp != nil {
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ConsumerGroup{}
			r.ConsumerGroup.CreatedAt = types.Int64PointerValue(resp.ConsumerGroup.CreatedAt)
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
			r.ConsumerGroup.Name = types.StringValue(resp.ConsumerGroup.Name)
			r.ConsumerGroup.Tags = make([]types.String, 0, len(resp.ConsumerGroup.Tags))
			for _, v := range resp.ConsumerGroup.Tags {
				r.ConsumerGroup.Tags = append(r.ConsumerGroup.Tags, types.StringValue(v))
			}
			r.ConsumerGroup.UpdatedAt = types.Int64PointerValue(resp.ConsumerGroup.UpdatedAt)
		}
		r.Consumers = []tfTypes.Consumer{}
		if len(r.Consumers) > len(resp.Consumers) {
			r.Consumers = r.Consumers[:len(resp.Consumers)]
		}
		for consumersCount, consumersItem := range resp.Consumers {
			var consumers1 tfTypes.Consumer
			consumers1.CreatedAt = types.Int64PointerValue(consumersItem.CreatedAt)
			consumers1.CustomID = types.StringPointerValue(consumersItem.CustomID)
			consumers1.ID = types.StringPointerValue(consumersItem.ID)
			consumers1.Tags = make([]types.String, 0, len(consumersItem.Tags))
			for _, v := range consumersItem.Tags {
				consumers1.Tags = append(consumers1.Tags, types.StringValue(v))
			}
			consumers1.UpdatedAt = types.Int64PointerValue(consumersItem.UpdatedAt)
			consumers1.Username = types.StringPointerValue(consumersItem.Username)
			if consumersCount+1 > len(r.Consumers) {
				r.Consumers = append(r.Consumers, consumers1)
			} else {
				r.Consumers[consumersCount].CreatedAt = consumers1.CreatedAt
				r.Consumers[consumersCount].CustomID = consumers1.CustomID
				r.Consumers[consumersCount].ID = consumers1.ID
				r.Consumers[consumersCount].Tags = consumers1.Tags
				r.Consumers[consumersCount].UpdatedAt = consumers1.UpdatedAt
				r.Consumers[consumersCount].Username = consumers1.Username
			}
		}
	}
}
