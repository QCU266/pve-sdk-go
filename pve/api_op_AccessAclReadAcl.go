// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessAclReadAclRequest interface{}

type AccessAclReadAclResponseItem struct {
	Path      string `json:"path,omitempty"`      // Access control path
	Propagate *bool  `json:"propagate,omitempty"` // Allow to propagate (inherit) permissions.
	Roleid    string `json:"roleid,omitempty"`    //
	Type      string `json:"type,omitempty"`      //
	Ugid      string `json:"ugid,omitempty"`      //
}

type AccessAclReadAclResponse []AccessAclReadAclResponseItem

// Get Access Control List (ACLs).
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/acl
func (c *Client) AccessAclReadAcl(ctx context.Context, request *AccessAclReadAclRequest) (*AccessAclReadAclResponse, error) {

	method := "GET"
	path := "/access/acl"

	var response AccessAclReadAclResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}