// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcStatusRebootVmRebootRequest struct {
	Node    string `json:"node,omitempty"`    // The cluster node name.
	Timeout *int64 `json:"timeout,omitempty"` // Wait maximal timeout seconds for the shutdown.
	Vmid    int64  `json:"vmid,omitempty"`    // The (unique) ID of the VM.
}

type NodesLxcStatusRebootVmRebootResponse string

// Reboot the container by shutting it down, and starting it again. Applies pending changes.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/reboot
func (c *Client) NodesLxcStatusRebootVmReboot(ctx context.Context, request *NodesLxcStatusRebootVmRebootRequest) (*NodesLxcStatusRebootVmRebootResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/status/reboot"

	var response NodesLxcStatusRebootVmRebootResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
