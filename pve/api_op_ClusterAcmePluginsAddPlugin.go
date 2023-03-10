// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterAcmePluginsAddPluginRequest struct {
	Api             *string `json:"api,omitempty"`              // API plugin name
	Data            *string `json:"data,omitempty"`             // DNS plugin data. (base64 encoded)
	Disable         *bool   `json:"disable,omitempty"`          // Flag to disable the config.
	Id              string  `json:"id,omitempty"`               // ACME Plugin ID name
	Nodes           *string `json:"nodes,omitempty"`            // List of cluster node names.
	Type            string  `json:"type,omitempty"`             // ACME challenge type.
	ValidationDelay *int64  `json:"validation-delay,omitempty"` // Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
}

type ClusterAcmePluginsAddPluginResponse struct{}

// Add ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/acme/plugins
func (c *Client) ClusterAcmePluginsAddPlugin(ctx context.Context, request *ClusterAcmePluginsAddPluginRequest) (*ClusterAcmePluginsAddPluginResponse, error) {

	method := "POST"
	path := "/cluster/acme/plugins"

	var response ClusterAcmePluginsAddPluginResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
