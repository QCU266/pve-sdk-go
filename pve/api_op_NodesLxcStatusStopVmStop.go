// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcStatusStopVmStopRequest struct {
	Node     string `json:"node,omitempty"`     // The cluster node name.
	Skiplock *bool  `json:"skiplock,omitempty"` // Ignore locks - only root is allowed to use this option.
	Vmid     int64  `json:"vmid,omitempty"`     // The (unique) ID of the VM.
}

type NodesLxcStatusStopVmStopResponse string

// Stop the container. This will abruptly stop all processes running in the container.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/stop
func (c *Client) NodesLxcStatusStopVmStop(ctx context.Context, request *NodesLxcStatusStopVmStopRequest) (*NodesLxcStatusStopVmStopResponse, error) {

	method := "POST"
	path := "/nodes/{node}/lxc/{vmid}/status/stop"

	var response NodesLxcStatusStopVmStopResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}