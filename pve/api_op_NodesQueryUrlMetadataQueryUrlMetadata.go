// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQueryUrlMetadataQueryUrlMetadataRequest struct {
	Node               string `query:"node,omitempty"`                // The cluster node name.
	Url                string `query:"url,omitempty"`                 // The URL to query the metadata from.
	VerifyCertificates *bool  `query:"verify-certificates,omitempty"` // If false, no SSL/TLS certificates will be verified.
}

type NodesQueryUrlMetadataQueryUrlMetadataResponse struct {
	Filename *string `json:"filename,omitempty"` //
	Mimetype *string `json:"mimetype,omitempty"` //
	Size     *int64  `json:"size,omitempty"`     //
}

// Query metadata of an URL: file size, file name and mime type.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/query-url-metadata
func (c *Client) NodesQueryUrlMetadataQueryUrlMetadata(ctx context.Context, request *NodesQueryUrlMetadataQueryUrlMetadataRequest) (*NodesQueryUrlMetadataQueryUrlMetadataResponse, error) {

	method := "GET"
	path := "/nodes/{node}/query-url-metadata"

	var response NodesQueryUrlMetadataQueryUrlMetadataResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
