// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessUsersTokenGenerateTokenRequest struct {
	Comment *string `json:"comment,omitempty"` //
	Expire  *int64  `json:"expire,omitempty"`  // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *bool   `json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Tokenid string  `json:"tokenid,omitempty"` // User-specific token identifier.
	Userid  string  `json:"userid,omitempty"`  // User ID
}

type AccessUsersTokenGenerateTokenResponse struct {
	FullTokenid string                                    `json:"full-tokenid,omitempty"` // The full token id.
	Info        AccessUsersTokenGenerateTokenResponseInfo `json:"info,omitempty"`         //
	Value       string                                    `json:"value,omitempty"`        // API token value used for authentication.
}

type AccessUsersTokenGenerateTokenResponseInfo struct {
	Comment *string `json:"comment,omitempty"` //
	Expire  *int64  `json:"expire,omitempty"`  // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *bool   `json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}

// Generate a new API token for a specific user. NOTE: returns API token value, which needs to be stored as it cannot be retrieved afterwards!
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/users/{userid}/token/{tokenid}
func (c *Client) AccessUsersTokenGenerateToken(ctx context.Context, request *AccessUsersTokenGenerateTokenRequest) (*AccessUsersTokenGenerateTokenResponse, error) {

	method := "POST"
	path := "/access/users/{userid}/token/{tokenid}"

	var response AccessUsersTokenGenerateTokenResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
