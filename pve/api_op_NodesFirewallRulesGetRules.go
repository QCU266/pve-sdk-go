// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesFirewallRulesGetRulesRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesFirewallRulesGetRulesResponse []NodesFirewallRulesGetRulesResponseItem

type NodesFirewallRulesGetRulesResponseItem struct {
	Pos int64 `json:"pos,omitempty"` //
}

// List rules.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/rules
func (c *Client) NodesFirewallRulesGetRules(ctx context.Context, request *NodesFirewallRulesGetRulesRequest) (*NodesFirewallRulesGetRulesResponse, error) {

	method := "GET"
	path := "/nodes/{node}/firewall/rules"

	var response NodesFirewallRulesGetRulesResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
