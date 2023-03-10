// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDisksDirectoryDeleteRequest struct {
	CleanupConfig *bool  `json:"cleanup-config,omitempty"` // Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupDisks  *bool  `json:"cleanup-disks,omitempty"`  // Also wipe disk so it can be repurposed afterwards.
	Name          string `json:"name,omitempty"`           // The storage identifier.
	Node          string `json:"node,omitempty"`           // The cluster node name.
}

type NodesDisksDirectoryDeleteResponse string

// Unmounts the storage and removes the mount unit.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/disks/directory/{name}
func (c *Client) NodesDisksDirectoryDelete(ctx context.Context, request *NodesDisksDirectoryDeleteRequest) (*NodesDisksDirectoryDeleteResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/disks/directory/{name}"

	var response NodesDisksDirectoryDeleteResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
