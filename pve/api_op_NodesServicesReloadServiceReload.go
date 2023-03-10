// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesServicesReloadServiceReloadRequest struct {
	Node    string `json:"node,omitempty"`    // The cluster node name.
	Service string `json:"service,omitempty"` // Service ID
}

type NodesServicesReloadServiceReloadResponse string

// Reload service. Falls back to restart if service cannot be reloaded.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/services/{service}/reload
func (c *Client) NodesServicesReloadServiceReload(ctx context.Context, request *NodesServicesReloadServiceReloadRequest) (*NodesServicesReloadServiceReloadResponse, error) {

	method := "POST"
	path := "/nodes/{node}/services/{service}/reload"

	var response NodesServicesReloadServiceReloadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
