// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesAplinfoAplinfoRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesAplinfoAplinfoResponse []NodesAplinfoAplinfoResponseItem

type NodesAplinfoAplinfoResponseItem struct {
}

// Get list of appliances.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/aplinfo
func (c *Client) NodesAplinfoAplinfo(ctx context.Context, request *NodesAplinfoAplinfoRequest) (*NodesAplinfoAplinfoResponse, error) {

	method := "GET"
	path := "/nodes/{node}/aplinfo"

	var response NodesAplinfoAplinfoResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
