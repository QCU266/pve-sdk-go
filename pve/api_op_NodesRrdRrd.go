// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesRrdRrdRequest struct {
	Cf        *string `query:"cf,omitempty"`        // The RRD consolidation function
	Ds        string  `query:"ds,omitempty"`        // The list of datasources you want to display.
	Node      string  `query:"node,omitempty"`      // The cluster node name.
	Timeframe string  `query:"timeframe,omitempty"` // Specify the time frame you are interested in.
}

type NodesRrdRrdResponse struct {
	Filename string `json:"filename,omitempty"` //
}

// Read node RRD statistics (returns PNG)
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/rrd
func (c *Client) NodesRrdRrd(ctx context.Context, request *NodesRrdRrdRequest) (*NodesRrdRrdResponse, error) {

	method := "GET"
	path := "/nodes/{node}/rrd"

	var response NodesRrdRrdResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}