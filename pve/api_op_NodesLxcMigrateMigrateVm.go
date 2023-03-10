// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcMigrateMigrateVmRequest struct {
	Bwlimit       *float64 `json:"bwlimit,omitempty"`        // Override I/O bandwidth limit (in KiB/s).
	Node          string   `json:"node,omitempty"`           // The cluster node name.
	Online        *bool    `json:"online,omitempty"`         // Use online/live migration.
	Restart       *bool    `json:"restart,omitempty"`        // Use restart migration
	Target        string   `json:"target,omitempty"`         // Target node.
	TargetStorage *string  `json:"target-storage,omitempty"` // Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Timeout       *int64   `json:"timeout,omitempty"`        // Timeout in seconds for shutdown for restart migration
	Vmid          int64    `json:"vmid,omitempty"`           // The (unique) ID of the VM.
}

// the task ID.
type NodesLxcMigrateMigrateVmResponse string

// Migrate the container to another node. Creates a new migration task.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/migrate
func (c *Client) NodesLxcMigrateMigrateVm(ctx context.Context, request *NodesLxcMigrateMigrateVmRequest) (*NodesLxcMigrateMigrateVmResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/migrate"

	var response NodesLxcMigrateMigrateVmResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
