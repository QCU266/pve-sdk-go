// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnControllersReadRequest struct {
	Controller string `query:"controller,omitempty"` // The SDN controller object identifier.
	Pending    *bool  `query:"pending,omitempty"`    // Display pending config.
	Running    *bool  `query:"running,omitempty"`    // Display running config.
}

type ClusterSdnControllersReadResponse interface{}

// Read sdn controller configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/controllers/{controller}
func (c *Client) ClusterSdnControllersRead(ctx context.Context, request *ClusterSdnControllersReadRequest) (*ClusterSdnControllersReadResponse, error) {

	method := "GET"
	path := "/cluster/sdn/controllers/{controller}"

	var response ClusterSdnControllersReadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
