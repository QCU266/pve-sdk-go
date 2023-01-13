// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSubscriptionUpdateRequest struct {
	Force *bool  `json:"force,omitempty"` // Always connect to server, even if we have up to date info inside local cache.
	Node  string `json:"node,omitempty"`  // The cluster node name.
}

type NodesSubscriptionUpdateResponse struct{}

// Update subscription info.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/subscription
func (c *Client) NodesSubscriptionUpdate(ctx context.Context, request *NodesSubscriptionUpdateRequest) (*NodesSubscriptionUpdateResponse, error) {

	method := "POST"
	path := "/nodes/{node}/subscription"

	var response NodesSubscriptionUpdateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}