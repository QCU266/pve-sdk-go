// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterAcmePluginsUpdatePluginRequest struct {
	Api             *string `json:"api,omitempty"`              // API plugin name
	Data            *string `json:"data,omitempty"`             // DNS plugin data. (base64 encoded)
	Delete          *string `json:"delete,omitempty"`           // A list of settings you want to delete.
	Digest          *string `json:"digest,omitempty"`           // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Disable         *bool   `json:"disable,omitempty"`          // Flag to disable the config.
	Id              string  `json:"id,omitempty"`               // ACME Plugin ID name
	Nodes           *string `json:"nodes,omitempty"`            // List of cluster node names.
	ValidationDelay *int64  `json:"validation-delay,omitempty"` // Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
}

type ClusterAcmePluginsUpdatePluginResponse struct{}

// Update ACME plugin configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/acme/plugins/{id}
func (c *Client) ClusterAcmePluginsUpdatePlugin(ctx context.Context, request *ClusterAcmePluginsUpdatePluginRequest) (*ClusterAcmePluginsUpdatePluginResponse, error) {

	method := "PUT"
	path := "/cluster/acme/plugins/{id}"

	var response ClusterAcmePluginsUpdatePluginResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
