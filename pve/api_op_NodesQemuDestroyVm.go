// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuDestroyVmRequest struct {
	DestroyUnreferencedDisks *bool  `json:"destroy-unreferenced-disks,omitempty"` // If set, destroy additionally all disks not referenced in the config but with a matching VMID from all enabled storages.
	Node                     string `json:"node,omitempty"`                       // The cluster node name.
	Purge                    *bool  `json:"purge,omitempty"`                      // Remove VMID from configurations, like backup & replication jobs and HA.
	Skiplock                 *bool  `json:"skiplock,omitempty"`                   // Ignore locks - only root is allowed to use this option.
	Vmid                     int64  `json:"vmid,omitempty"`                       // The (unique) ID of the VM.
}

type NodesQemuDestroyVmResponse string

// Destroy the VM and  all used/owned volumes. Removes any VM specific permissions and firewall rules
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}
func (c *Client) NodesQemuDestroyVm(ctx context.Context, request *NodesQemuDestroyVmRequest) (*NodesQemuDestroyVmResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/qemu/{vmid}"

	var response NodesQemuDestroyVmResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
