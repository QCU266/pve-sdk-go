// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnVnetsReadRequest struct {
	Pending *bool  `query:"pending,omitempty"` // Display pending config.
	Running *bool  `query:"running,omitempty"` // Display running config.
	Vnet    string `query:"vnet,omitempty"`    // The SDN vnet object identifier.
}

type ClusterSdnVnetsReadResponse interface{}

// Read sdn vnet configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/vnets/{vnet}
func (c *Client) ClusterSdnVnetsRead(ctx context.Context, request *ClusterSdnVnetsReadRequest) (*ClusterSdnVnetsReadResponse, error) {

	method := "GET"
	path := "/cluster/sdn/vnets/{vnet}"

	var response ClusterSdnVnetsReadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}