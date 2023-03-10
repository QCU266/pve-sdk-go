// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStorageContentUpdateattributesRequest struct {
	Node      string  `json:"node,omitempty"`      // The cluster node name.
	Notes     *string `json:"notes,omitempty"`     // The new notes.
	Protected *bool   `json:"protected,omitempty"` // Protection status. Currently only supported for backups.
	Storage   *string `json:"storage,omitempty"`   // The storage identifier.
	Volume    string  `json:"volume,omitempty"`    // Volume identifier
}

type NodesStorageContentUpdateattributesResponse struct{}

// Update volume attributes
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content/{volume}
func (c *Client) NodesStorageContentUpdateattributes(ctx context.Context, request *NodesStorageContentUpdateattributesRequest) (*NodesStorageContentUpdateattributesResponse, error) {

	method := "PUT"
	path := "/nodes/{node}/storage/{storage}/content/{volume}"

	var response NodesStorageContentUpdateattributesResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
