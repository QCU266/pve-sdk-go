// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuFirewallAliasesUpdateAliasRequest struct {
	Cidr    string  `json:"cidr,omitempty"`    // Network/IP specification in CIDR format.
	Comment *string `json:"comment,omitempty"` //
	Digest  *string `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string  `json:"name,omitempty"`    // Alias name.
	Node    string  `json:"node,omitempty"`    // The cluster node name.
	Rename  *string `json:"rename,omitempty"`  // Rename an existing alias.
	Vmid    int64   `json:"vmid,omitempty"`    // The (unique) ID of the VM.
}

type NodesQemuFirewallAliasesUpdateAliasResponse struct{}

// Update IP or Network alias.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/firewall/aliases/{name}
func (c *Client) NodesQemuFirewallAliasesUpdateAlias(ctx context.Context, request *NodesQemuFirewallAliasesUpdateAliasRequest) (*NodesQemuFirewallAliasesUpdateAliasResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/qemu/{vmid}/firewall/aliases/{name}"

	var response NodesQemuFirewallAliasesUpdateAliasResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
