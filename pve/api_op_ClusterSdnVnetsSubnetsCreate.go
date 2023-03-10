// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterSdnVnetsSubnetsCreateRequest struct {
	Dnszoneprefix *string `json:"dnszoneprefix,omitempty"` // dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Gateway       *string `json:"gateway,omitempty"`       // Subnet Gateway: Will be assign on vnet for layer3 zones
	Snat          *bool   `json:"snat,omitempty"`          // enable masquerade for this subnet if pve-firewall
	Subnet        string  `json:"subnet,omitempty"`        // The SDN subnet object identifier.
	Type          string  `json:"type,omitempty"`          //
	Vnet          string  `json:"vnet,omitempty"`          // associated vnet
}

type ClusterSdnVnetsSubnetsCreateResponse struct{}

// Create a new sdn subnet object.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/sdn/vnets/{vnet}/subnets
func (c *Client) ClusterSdnVnetsSubnetsCreate(ctx context.Context, request *ClusterSdnVnetsSubnetsCreateRequest) (*ClusterSdnVnetsSubnetsCreateResponse, error) {

	method := "POST"
	path := "/cluster/sdn/vnets/{vnet}/subnets"

	var response ClusterSdnVnetsSubnetsCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
