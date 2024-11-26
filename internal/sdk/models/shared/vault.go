// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Vault entities are used to configure different Vault connectors. Examples of Vaults are Environment Variables, Hashicorp Vault and AWS Secrets Manager. Configuring a Vault allows referencing the secrets with other entities. For example a certificate entity can store a reference to a certificate and key, stored in a vault, instead of storing the certificate and key within the entity. This allows a proper separation of secrets and configuration and prevents secret sprawl.
type Vault struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page.
	Config any `json:"config"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The description of the Vault entity.
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name string `json:"name"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix string `json:"prefix"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *Vault) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Vault) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Vault) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Vault) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Vault) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Vault) GetPrefix() string {
	if o == nil {
		return ""
	}
	return o.Prefix
}

func (o *Vault) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Vault) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
