// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDnsUpdateDnsRequest struct {
	Dns1   *string `json:"dns1,omitempty"`   // First name server IP address.
	Dns2   *string `json:"dns2,omitempty"`   // Second name server IP address.
	Dns3   *string `json:"dns3,omitempty"`   // Third name server IP address.
	Node   string  `json:"node,omitempty"`   // The cluster node name.
	Search string  `json:"search,omitempty"` // Search domain for host-name lookup.
}

type NodesDnsUpdateDnsResponse struct{}

// Write DNS settings.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/dns
func (c *Client) NodesDnsUpdateDns(ctx context.Context, request *NodesDnsUpdateDnsRequest) (*NodesDnsUpdateDnsResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/dns"

	var response NodesDnsUpdateDnsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
