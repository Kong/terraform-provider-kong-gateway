// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AddRoleToGroupRequestBody struct {
	RbacRoleID  string `json:"rbac_role_id"`
	WorkspaceID string `json:"workspace_id"`
}

func (o *AddRoleToGroupRequestBody) GetRbacRoleID() string {
	if o == nil {
		return ""
	}
	return o.RbacRoleID
}

func (o *AddRoleToGroupRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type AddRoleToGroupRequest struct {
	Groups      string                     `pathParam:"style=simple,explode=false,name=groups"`
	RequestBody *AddRoleToGroupRequestBody `request:"mediaType=application/json"`
}

func (o *AddRoleToGroupRequest) GetGroups() string {
	if o == nil {
		return ""
	}
	return o.Groups
}

func (o *AddRoleToGroupRequest) GetRequestBody() *AddRoleToGroupRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type AddRoleToGroupGroup struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *AddRoleToGroupGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddRoleToGroupGroup) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AddRoleToGroupRbacRole struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *AddRoleToGroupRbacRole) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddRoleToGroupRbacRole) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AddRoleToGroupWorkspace struct {
	ID *string `json:"id,omitempty"`
}

func (o *AddRoleToGroupWorkspace) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AddRoleToGroupResponseBody - Successfully associated the role with the group
type AddRoleToGroupResponseBody struct {
	Group     *AddRoleToGroupGroup     `json:"group,omitempty"`
	RbacRole  *AddRoleToGroupRbacRole  `json:"rbac_role,omitempty"`
	Workspace *AddRoleToGroupWorkspace `json:"workspace,omitempty"`
}

func (o *AddRoleToGroupResponseBody) GetGroup() *AddRoleToGroupGroup {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *AddRoleToGroupResponseBody) GetRbacRole() *AddRoleToGroupRbacRole {
	if o == nil {
		return nil
	}
	return o.RbacRole
}

func (o *AddRoleToGroupResponseBody) GetWorkspace() *AddRoleToGroupWorkspace {
	if o == nil {
		return nil
	}
	return o.Workspace
}

type AddRoleToGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully associated the role with the group
	Object *AddRoleToGroupResponseBody
}

func (o *AddRoleToGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddRoleToGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddRoleToGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddRoleToGroupResponse) GetObject() *AddRoleToGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}