// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCephOsdInInRequest struct {
	Node  string `json:"node,omitempty"`  // The cluster node name.
	Osdid int64  `json:"osdid,omitempty"` // OSD ID
}

type NodesCephOsdInInResponse struct{}

// ceph osd in
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/ceph/osd/{osdid}/in
func (c *Client) NodesCephOsdInIn(ctx context.Context, request *NodesCephOsdInInRequest) (*NodesCephOsdInInResponse, error) {

	method := "POST"
	path := "/nodes/{node}/ceph/osd/{osdid}/in"

	var response NodesCephOsdInInResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}