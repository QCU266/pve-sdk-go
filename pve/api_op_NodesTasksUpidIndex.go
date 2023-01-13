// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesTasksUpidIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
	Upid string `query:"upid,omitempty"` //
}

type NodesTasksUpidIndexResponseItem struct {
}

type NodesTasksUpidIndexResponse []NodesTasksUpidIndexResponseItem

// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks/{upid}
func (c *Client) NodesTasksUpidIndex(ctx context.Context, request *NodesTasksUpidIndexRequest) (*NodesTasksUpidIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/tasks/{upid}"

	var response NodesTasksUpidIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}