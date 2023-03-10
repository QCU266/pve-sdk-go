// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentInfoInfoRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

// Returns an object with a single `result` property.
type NodesQemuAgentInfoInfoResponse interface{}

// Execute info.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/info
func (c *Client) NodesQemuAgentInfoInfo(ctx context.Context, request *NodesQemuAgentInfoInfoRequest) (*NodesQemuAgentInfoInfoResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/agent/info"

	var response NodesQemuAgentInfoInfoResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
