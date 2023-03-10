// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDisksLvmDeleteRequest struct {
	CleanupConfig *bool  `json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *bool  `json:"cleanup-disks,omitempty"`  // Also wipe disks so they can be repurposed afterwards.
	Name          string `json:"name,omitempty"`           // The storage identifier.
	Node          string `json:"node,omitempty"`           // The cluster node name.
}

type NodesDisksLvmDeleteResponse string

// Remove an LVM Volume Group.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/disks/lvm/{name}
func (c *Client) NodesDisksLvmDelete(ctx context.Context, request *NodesDisksLvmDeleteRequest) (*NodesDisksLvmDeleteResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/disks/lvm/{name}"

	var response NodesDisksLvmDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
