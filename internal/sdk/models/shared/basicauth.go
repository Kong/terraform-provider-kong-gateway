// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BasicAuthConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *BasicAuthConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type BasicAuth struct {
	Consumer *BasicAuthConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Password  string   `json:"password"`
	Tags      []string `json:"tags,omitempty"`
	Username  string   `json:"username"`
}

func (o *BasicAuth) GetConsumer() *BasicAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *BasicAuth) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BasicAuth) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicAuth) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *BasicAuth) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuth) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type BasicAuthInput struct {
	Consumer *BasicAuthConsumer `json:"consumer,omitempty"`
	ID       *string            `json:"id,omitempty"`
	Password string             `json:"password"`
	Tags     []string           `json:"tags,omitempty"`
	Username string             `json:"username"`
}

func (o *BasicAuthInput) GetConsumer() *BasicAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *BasicAuthInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicAuthInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *BasicAuthInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuthInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
