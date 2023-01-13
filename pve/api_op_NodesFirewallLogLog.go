// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesFirewallLogLogRequest struct {
	Limit *int64 `query:"limit,omitempty"` //
	Node  string `query:"node,omitempty"`  // The cluster node name.
	Start *int64 `query:"start,omitempty"` //
}

type NodesFirewallLogLogResponseItem struct {
	N int64  `json:"n,omitempty"` // Line number
	T string `json:"t,omitempty"` // Line text
}

type NodesFirewallLogLogResponse []NodesFirewallLogLogResponseItem

// Read firewall log
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/log
func (c *Client) NodesFirewallLogLog(ctx context.Context, request *NodesFirewallLogLogRequest) (*NodesFirewallLogLogResponse, error) {

	method := "GET"
	path := "/nodes/{node}/firewall/log"

	var response NodesFirewallLogLogResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}