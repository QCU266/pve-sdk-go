// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnDnsReadRequest struct {
	Dns string `query:"dns,omitempty"` // The SDN dns object identifier.
}

type ClusterSdnDnsReadResponse interface{}

// Read sdn dns configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/dns/{dns}
func (c *Client) ClusterSdnDnsRead(ctx context.Context, request *ClusterSdnDnsReadRequest) (*ClusterSdnDnsReadResponse, error) {

	method := "GET"
	path := "/cluster/sdn/dns/{dns}"

	var response ClusterSdnDnsReadResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
