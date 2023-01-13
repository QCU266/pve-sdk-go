// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallIpsetRemoveIpRequest struct {
	Cidr   string  `json:"cidr,omitempty"`   // Network/IP specification in CIDR format.
	Digest *string `json:"digest,omitempty"` // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name   string  `json:"name,omitempty"`   // IP set name.
	Node   string  `json:"node,omitempty"`   // The cluster node name.
	Vmid   int64   `json:"vmid,omitempty"`   // The (unique) ID of the VM.
}

type NodesQemuFirewallIpsetRemoveIpResponse struct{}

// Remove IP or Network from IPSet.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}
func (c *Client) NodesQemuFirewallIpsetRemoveIp(ctx context.Context, request *NodesQemuFirewallIpsetRemoveIpRequest) (*NodesQemuFirewallIpsetRemoveIpResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/qemu/{vmid}/firewall/ipset/{name}/{cidr}"

	var response NodesQemuFirewallIpsetRemoveIpResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}