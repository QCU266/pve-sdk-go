// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterConfigCreateRequest struct {
	Clustername string  `json:"clustername,omitempty"` // The name of the cluster.
	Linkn       *string `json:"link[n],omitempty"`     // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	Nodeid      *int64  `json:"nodeid,omitempty"`      // Node id for this node.
	Votes       *int64  `json:"votes,omitempty"`       // Number of votes for this node.
}

type ClusterConfigCreateResponse string

// Generate new cluster configuration. If no links given, default to local IP address as link0.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/config
func (c *Client) ClusterConfigCreate(ctx context.Context, request *ClusterConfigCreateRequest) (*ClusterConfigCreateResponse, error) {

	method := "POST"
	path := "/cluster/config"

	var response ClusterConfigCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}