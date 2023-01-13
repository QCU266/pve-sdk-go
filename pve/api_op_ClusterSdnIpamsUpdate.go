// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnIpamsUpdateRequest struct {
	Delete  *string `json:"delete,omitempty"`  // A list of settings you want to delete.
	Digest  *string `json:"digest,omitempty"`  // Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Ipam    string  `json:"ipam,omitempty"`    // The SDN ipam object identifier.
	Section *int64  `json:"section,omitempty"` //
	Token   *string `json:"token,omitempty"`   //
	Url     *string `json:"url,omitempty"`     //
}

type ClusterSdnIpamsUpdateResponse struct{}

// Update sdn ipam object configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/ipams/{ipam}
func (c *Client) ClusterSdnIpamsUpdate(ctx context.Context, request *ClusterSdnIpamsUpdateRequest) (*ClusterSdnIpamsUpdateResponse, error) {

	method := "PUT"
	path := "/cluster/sdn/ipams/{ipam}"

	var response ClusterSdnIpamsUpdateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}