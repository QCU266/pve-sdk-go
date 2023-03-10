// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessPasswordChangePasswordRequest struct {
	Password string `json:"password,omitempty"` // The new password.
	Userid   string `json:"userid,omitempty"`   // User ID
}

type AccessPasswordChangePasswordResponse struct{}

// Change user password.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/password
func (c *Client) AccessPasswordChangePassword(ctx context.Context, request *AccessPasswordChangePasswordRequest) (*AccessPasswordChangePasswordResponse, error) {

	method := "PUT"
	path := "/access/password"

	var response AccessPasswordChangePasswordResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
