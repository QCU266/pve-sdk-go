// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesStorageContentCreateRequest struct {
	Filename string  `json:"filename,omitempty"` // The name of the file to create.
	Format   *string `json:"format,omitempty"`   //
	Node     string  `json:"node,omitempty"`     // The cluster node name.
	Size     string  `json:"size,omitempty"`     // Size in kilobyte (1024 bytes). Optional suffixes 'M' (megabyte, 1024K) and 'G' (gigabyte, 1024M)
	Storage  string  `json:"storage,omitempty"`  // The storage identifier.
	Vmid     int64   `json:"vmid,omitempty"`     // Specify owner VM
}

// Volume identifier
type NodesStorageContentCreateResponse string

// Allocate disk images.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content
func (c *Client) NodesStorageContentCreate(ctx context.Context, request *NodesStorageContentCreateRequest) (*NodesStorageContentCreateResponse, error) {

	method := "POST"
	path := "/nodes/{node}/storage/{storage}/content"

	var response NodesStorageContentCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}