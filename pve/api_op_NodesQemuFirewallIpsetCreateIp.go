// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIpsetCreateIpRequest struct {
	Cidr    string  `json:"cidr,omitempty"`    // Network/IP specification in CIDR format.
	Comment *string `json:"comment,omitempty"` //
	Name    string  `json:"name,omitempty"`    // IP set name.
	Node    string  `json:"node,omitempty"`    // The cluster node name.
	Nomatch *bool   `json:"nomatch,omitempty"` //
	Vmid    int64   `json:"vmid,omitempty"`    // The (unique) ID of the VM.
}

type NodesQemuFirewallIpsetCreateIpResponse struct{}

// Add IP or Network to IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}
func (c *Client) NodesQemuFirewallIpsetCreateIp(ctx context.Context, request *NodesQemuFirewallIpsetCreateIpRequest) (*NodesQemuFirewallIpsetCreateIpResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}"

	var response NodesQemuFirewallIpsetCreateIpResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
