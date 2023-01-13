// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSubscriptionGetRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesSubscriptionGetResponse interface{}

// Read subscription info.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/subscription
func (c *Client) NodesSubscriptionGet(ctx context.Context, request *NodesSubscriptionGetRequest) (*NodesSubscriptionGetResponse, error) {

	method := "GET"
	path := "/nodes/{node}/subscription"

	var response NodesSubscriptionGetResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}