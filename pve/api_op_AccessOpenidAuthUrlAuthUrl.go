// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessOpenidAuthUrlAuthUrlRequest struct {
	Realm       string `json:"realm,omitempty"`        // Authentication domain ID
	RedirectUrl string `json:"redirect-url,omitempty"` // Redirection Url. The client should set this to the used server url (location.origin).
}

// Redirection URL.
type AccessOpenidAuthUrlAuthUrlResponse string

// Get the OpenId Authorization Url for the specified realm.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/openid/auth-url
func (c *Client) AccessOpenidAuthUrlAuthUrl(ctx context.Context, request *AccessOpenidAuthUrlAuthUrlRequest) (*AccessOpenidAuthUrlAuthUrlResponse, error) {

	method := "POST"
	path := "/access/openid/auth-url"

	var response AccessOpenidAuthUrlAuthUrlResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}