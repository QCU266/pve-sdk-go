// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesReplicationStatusJobStatusRequest struct {
	Id   string `query:"id,omitempty"`   // Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesReplicationStatusJobStatusResponse interface{}

// Get replication job status.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/replication/{id}/status
func (c *Client) NodesReplicationStatusJobStatus(ctx context.Context, request *NodesReplicationStatusJobStatusRequest) (*NodesReplicationStatusJobStatusResponse, error) {

	method := "GET"
	path := "/nodes/{node}/replication/{id}/status"

	var response NodesReplicationStatusJobStatusResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
