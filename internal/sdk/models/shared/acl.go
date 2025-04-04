// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ACLConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ACLConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ACL struct {
	Consumer *ACLConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	Group     string   `json:"group"`
	ID        *string  `json:"id,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

func (o *ACL) GetConsumer() *ACLConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ACL) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ACL) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *ACL) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ACL) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
