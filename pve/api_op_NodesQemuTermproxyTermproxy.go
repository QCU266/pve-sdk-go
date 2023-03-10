// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuTermproxyTermproxyRequest struct {
	Node   string  `json:"node,omitempty"`   // The cluster node name.
	Serial *string `json:"serial,omitempty"` // opens a serial terminal (defaults to display)
	Vmid   int64   `json:"vmid,omitempty"`   // The (unique) ID of the VM.
}

type NodesQemuTermproxyTermproxyResponse struct {
	Port   int64  `json:"port,omitempty"`   //
	Ticket string `json:"ticket,omitempty"` //
	Upid   string `json:"upid,omitempty"`   //
	User   string `json:"user,omitempty"`   //
}

// Creates a TCP proxy connections.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/termproxy
func (c *Client) NodesQemuTermproxyTermproxy(ctx context.Context, request *NodesQemuTermproxyTermproxyRequest) (*NodesQemuTermproxyTermproxyResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/termproxy"

	var response NodesQemuTermproxyTermproxyResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
