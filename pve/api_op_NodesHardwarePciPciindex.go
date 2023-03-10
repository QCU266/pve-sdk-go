// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesHardwarePciPciindexRequest struct {
	Node  string `query:"node,omitempty"`  // The cluster node name.
	Pciid string `query:"pciid,omitempty"` //
}

type NodesHardwarePciPciindexResponse []NodesHardwarePciPciindexResponseItem

type NodesHardwarePciPciindexResponseItem struct {
	Method string `json:"method,omitempty"` //
}

// Index of available pci methods
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/hardware/pci/{pciid}
func (c *Client) NodesHardwarePciPciindex(ctx context.Context, request *NodesHardwarePciPciindexRequest) (*NodesHardwarePciPciindexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/hardware/pci/{pciid}"

	var response NodesHardwarePciPciindexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
