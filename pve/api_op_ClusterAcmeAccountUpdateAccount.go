// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterAcmeAccountUpdateAccountRequest struct {
	Contact *string `json:"contact,omitempty"` // Contact email addresses.
	Name    *string `json:"name,omitempty"`    // ACME account config file name.
}

type ClusterAcmeAccountUpdateAccountResponse string

// Update existing ACME account information with CA. Note: not specifying any new account information triggers a refresh.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/acme/account/{name}
func (c *Client) ClusterAcmeAccountUpdateAccount(ctx context.Context, request *ClusterAcmeAccountUpdateAccountRequest) (*ClusterAcmeAccountUpdateAccountResponse, error) {

	method := "PUT"
	path := "/cluster/acme/account/{name}"

	var response ClusterAcmeAccountUpdateAccountResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}