// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type AccessTfaGetTfaEntryRequest struct {
	Id     string `query:"id,omitempty"`     // A TFA entry id.
	Userid string `query:"userid,omitempty"` // User ID
}

// TFA Entry.
type AccessTfaGetTfaEntryResponse struct {
	Created     int64  `json:"created,omitempty"`     // Creation time of this entry as unix epoch.
	Description string `json:"description,omitempty"` // User chosen description for this entry.
	Enable      *bool  `json:"enable,omitempty"`      // Whether this TFA entry is currently enabled.
	Id          string `json:"id,omitempty"`          // The id used to reference this entry.
	Type        string `json:"type,omitempty"`        // TFA Entry Type.
}

// Fetch a requested TFA entry if present.
// https://pve.proxmox.com/pve-docs/api-viewer/#/access/tfa/{userid}/{id}
func (c *Client) AccessTfaGetTfaEntry(ctx context.Context, request *AccessTfaGetTfaEntryRequest) (*AccessTfaGetTfaEntryResponse, error) {

	method := "GET"
	path := "/access/tfa/{userid}/{id}"

	var response AccessTfaGetTfaEntryResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
