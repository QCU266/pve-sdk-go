// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallLogLogRequest struct {
	Limit *int64 `query:"limit,omitempty"` //
	Node  string `query:"node,omitempty"`  // The cluster node name.
	Start *int64 `query:"start,omitempty"` //
	Vmid  int64  `query:"vmid,omitempty"`  // The (unique) ID of the VM.
}

type NodesQemuFirewallLogLogResponse []NodesQemuFirewallLogLogResponseItem

type NodesQemuFirewallLogLogResponseItem struct {
	N int64  `json:"n,omitempty"` // Line number
	T string `json:"t,omitempty"` // Line text
}

// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/log
func (c *Client) NodesQemuFirewallLogLog(ctx context.Context, request *NodesQemuFirewallLogLogRequest) (*NodesQemuFirewallLogLogResponse, error) {

	method := "GET"
	path := "/nodes/{node}/qemu/{vmid}/firewall/log"

	var response NodesQemuFirewallLogLogResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
