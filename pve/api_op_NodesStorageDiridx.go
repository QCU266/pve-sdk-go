// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStorageDiridxRequest struct {
	Node    string `query:"node,omitempty"`    // The cluster node name.
	Storage string `query:"storage,omitempty"` // The storage identifier.
}

type NodesStorageDiridxResponse []NodesStorageDiridxResponseItem

type NodesStorageDiridxResponseItem struct {
	Subdir string `json:"subdir,omitempty"` //
}

// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}
func (c *Client) NodesStorageDiridx(ctx context.Context, request *NodesStorageDiridxRequest) (*NodesStorageDiridxResponse, error) {

	method := "GET"
	path := "/nodes/{node}/storage/{storage}"

	var response NodesStorageDiridxResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
