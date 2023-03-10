// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessRolesIndexRequest interface{}

type AccessRolesIndexResponse []AccessRolesIndexResponseItem

type AccessRolesIndexResponseItem struct {
	Privs   *string `json:"privs,omitempty"`   //
	Roleid  string  `json:"roleid,omitempty"`  //
	Special *bool   `json:"special,omitempty"` //
}

// Role index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/roles
func (c *Client) AccessRolesIndex(ctx context.Context, request *AccessRolesIndexRequest) (*AccessRolesIndexResponse, error) {

	method := "GET"
	path := "/access/roles"

	var response AccessRolesIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
