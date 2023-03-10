// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephConfigdbConfigdbRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCephConfigdbConfigdbResponse []NodesCephConfigdbConfigdbResponseItem

type NodesCephConfigdbConfigdbResponseItem struct {
	CanUpdateAtRuntime bool   `json:"can_update_at_runtime,omitempty"` //
	Level              string `json:"level,omitempty"`                 //
	Mask               string `json:"mask,omitempty"`                  //
	Name               string `json:"name,omitempty"`                  //
	Section            string `json:"section,omitempty"`               //
	Value              string `json:"value,omitempty"`                 //
}

// Get Ceph configuration database.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/configdb
func (c *Client) NodesCephConfigdbConfigdb(ctx context.Context, request *NodesCephConfigdbConfigdbRequest) (*NodesCephConfigdbConfigdbResponse, error) {

	method := "GET"
	path := "/nodes/{node}/ceph/configdb"

	var response NodesCephConfigdbConfigdbResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
