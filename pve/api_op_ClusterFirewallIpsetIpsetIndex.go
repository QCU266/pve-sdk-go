// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterFirewallIpsetIpsetIndexRequest interface{}

type ClusterFirewallIpsetIpsetIndexResponse []ClusterFirewallIpsetIpsetIndexResponseItem

type ClusterFirewallIpsetIpsetIndexResponseItem struct {
	Comment *string `json:"comment,omitempty"` //
	Digest  string  `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Name    string  `json:"name,omitempty"`    // IP set name.
}

// List IPSets
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/firewall/ipset
func (c *Client) ClusterFirewallIpsetIpsetIndex(ctx context.Context, request *ClusterFirewallIpsetIpsetIndexRequest) (*ClusterFirewallIpsetIpsetIndexResponse, error) {

	method := "GET"
	path := "/cluster/firewall/ipset"

	var response ClusterFirewallIpsetIpsetIndexResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
