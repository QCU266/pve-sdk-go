// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentGetHostNameGetHostNameRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Vmid int64  `query:"vmid,omitempty"` // The (unique) ID of the VM.
}

// Returns an object with a single `result` property.
type NodesQemuAgentGetHostNameGetHostNameResponse interface{}

// Execute get-host-name.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/get-host-name
func (c *Client) NodesQemuAgentGetHostNameGetHostName(ctx context.Context, request *NodesQemuAgentGetHostNameGetHostNameRequest) (*NodesQemuAgentGetHostNameGetHostNameResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/agent/get-host-name"

	var response NodesQemuAgentGetHostNameGetHostNameResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
