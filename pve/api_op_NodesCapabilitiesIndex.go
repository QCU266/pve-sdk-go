// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCapabilitiesIndexRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCapabilitiesIndexResponseItem struct {
}

type NodesCapabilitiesIndexResponse []NodesCapabilitiesIndexResponseItem

// Node capabilities index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/capabilities
func (c *Client) NodesCapabilitiesIndex(ctx context.Context, request *NodesCapabilitiesIndexRequest) (*NodesCapabilitiesIndexResponse, error) {

	method := "GET"
	path := "/nodes/{node}/capabilities"

	var response NodesCapabilitiesIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}