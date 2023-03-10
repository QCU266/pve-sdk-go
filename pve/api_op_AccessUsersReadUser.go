// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessUsersReadUserRequest struct {
	Userid string `query:"userid,omitempty"` // User ID
}

type AccessUsersReadUserResponse struct {
	Comment   *string                                  `json:"comment,omitempty"`   //
	Email     *string                                  `json:"email,omitempty"`     //
	Enable    *bool                                    `json:"enable,omitempty"`    // Enable the account (default). You can set this to '0' to disable the account
	Expire    *int64                                   `json:"expire,omitempty"`    // Account expiration date (seconds since epoch). '0' means no expiration date.
	Firstname *string                                  `json:"firstname,omitempty"` //
	Groups    *[]AccessUsersReadUserResponseGroupsItem `json:"groups,omitempty"`    //
	Keys      *string                                  `json:"keys,omitempty"`      // Keys for two factor auth (yubico).
	Lastname  *string                                  `json:"lastname,omitempty"`  //
	Tokens    *AccessUsersReadUserResponseTokens       `json:"tokens,omitempty"`    //
}

type AccessUsersReadUserResponseGroupsItem string

type AccessUsersReadUserResponseTokens struct {
	Comment *string `json:"comment,omitempty"` //
	Expire  *int64  `json:"expire,omitempty"`  // API token expiration date (seconds since epoch). '0' means no expiration date.
	Privsep *bool   `json:"privsep,omitempty"` // Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
}

// Get user configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/users/{userid}
func (c *Client) AccessUsersReadUser(ctx context.Context, request *AccessUsersReadUserRequest) (*AccessUsersReadUserResponse, error) {

	method := "GET"
	path := "/access/users/{userid}"

	var response AccessUsersReadUserResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
