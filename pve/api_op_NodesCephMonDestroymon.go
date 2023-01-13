// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephMonDestroymonRequest struct {
	Monid string `json:"monid,omitempty"` // Monitor ID
	Node  string `json:"node,omitempty"`  // The cluster node name.
}

type NodesCephMonDestroymonResponse string

// Destroy Ceph Monitor and Manager.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/mon/{monid}
func (c *Client) NodesCephMonDestroymon(ctx context.Context, request *NodesCephMonDestroymonRequest) (*NodesCephMonDestroymonResponse, error) {

	method := "DELETE"
	path := "/nodes/{node}/ceph/mon/{monid}"

	var response NodesCephMonDestroymonResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}