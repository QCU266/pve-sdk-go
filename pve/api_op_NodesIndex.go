// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesIndexResponse []NodesIndexResponseItem

type NodesIndexResponseItem struct {
}

// Node index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}
func (c *Client) NodesIndex(ctx context.Context, request *NodesIndexRequest) (*NodesIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}"

	var response NodesIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
