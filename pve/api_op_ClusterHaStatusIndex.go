// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterHaStatusIndexRequest interface{}

type ClusterHaStatusIndexResponse []ClusterHaStatusIndexResponseItem

type ClusterHaStatusIndexResponseItem struct {
}

// Directory index.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ha/status
func (c *Client) ClusterHaStatusIndex(ctx context.Context, request *ClusterHaStatusIndexRequest) (*ClusterHaStatusIndexResponse, error) {

	method := "GET"
	path := "/cluster/ha/status"

	var response ClusterHaStatusIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
