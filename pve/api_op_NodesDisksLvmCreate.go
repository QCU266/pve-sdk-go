// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDisksLvmCreateRequest struct {
	AddStorage *bool  `json:"add_storage,omitempty"` // Configure storage using the Volume Group
	Device     string `json:"device,omitempty"`      // The block device you want to create the volume group on
	Name       string `json:"name,omitempty"`        // The storage identifier.
	Node       string `json:"node,omitempty"`        // The cluster node name.
}

type NodesDisksLvmCreateResponse string

// Create an LVM Volume Group
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/disks/lvm
func (c *Client) NodesDisksLvmCreate(ctx context.Context, request *NodesDisksLvmCreateRequest) (*NodesDisksLvmCreateResponse, error) {

	method := "POST"
	path := "/nodes/{node}/disks/lvm"

	var response NodesDisksLvmCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
