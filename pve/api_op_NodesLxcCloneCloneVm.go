// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcCloneCloneVmRequest struct {
	Bwlimit     *float64 `json:"bwlimit,omitempty"`     // Override I/O bandwidth limit (in KiB/s).
	Description *string  `json:"description,omitempty"` // Description for the new CT.
	Full        *bool    `json:"full,omitempty"`        // Create a full copy of all disks. This is always done when you clone a normal CT. For CT templates, we try to create a linked clone by default.
	Hostname    *string  `json:"hostname,omitempty"`    // Set a hostname for the new CT.
	Newid       int64    `json:"newid,omitempty"`       // VMID for the clone.
	Node        string   `json:"node,omitempty"`        // The cluster node name.
	Pool        *string  `json:"pool,omitempty"`        // Add the new CT to the specified pool.
	Snapname    *string  `json:"snapname,omitempty"`    // The name of the snapshot.
	Storage     *string  `json:"storage,omitempty"`     // Target storage for full clone.
	Target      *string  `json:"target,omitempty"`      // Target node. Only allowed if the original VM is on shared storage.
	Vmid        int64    `json:"vmid,omitempty"`        // The (unique) ID of the VM.
}

type NodesLxcCloneCloneVmResponse string

// Create a container clone/copy
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/clone
func (c *Client) NodesLxcCloneCloneVm(ctx context.Context, request *NodesLxcCloneCloneVmRequest) (*NodesLxcCloneCloneVmResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/clone"

	var response NodesLxcCloneCloneVmResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
