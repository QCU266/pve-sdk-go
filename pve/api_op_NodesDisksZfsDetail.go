// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesDisksZfsDetailRequest struct {
	Name string `query:"name,omitempty"` // The storage identifier.
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesDisksZfsDetailResponse struct {
	Action   *string                                   `json:"action,omitempty"`   // Information about the recommended action to fix the state.
	Children []NodesDisksZfsDetailResponseChildrenItem `json:"children,omitempty"` // The pool configuration information, including the vdevs for each section (e.g. spares, cache), may be nested.
	Errors   string                                    `json:"errors,omitempty"`   // Information about the errors on the zpool.
	Name     string                                    `json:"name,omitempty"`     // The name of the zpool.
	Scan     *string                                   `json:"scan,omitempty"`     // Information about the last/current scrub.
	State    string                                    `json:"state,omitempty"`    // The state of the zpool.
	Status   *string                                   `json:"status,omitempty"`   // Information about the state of the zpool.
}

type NodesDisksZfsDetailResponseChildrenItem struct {
	Cksum *float64 `json:"cksum,omitempty"` //
	Msg   string   `json:"msg,omitempty"`   // An optional message about the vdev.
	Name  string   `json:"name,omitempty"`  // The name of the vdev or section.
	Read  *float64 `json:"read,omitempty"`  //
	State *string  `json:"state,omitempty"` // The state of the vdev.
	Write *float64 `json:"write,omitempty"` //
}

// Get details about a zpool.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/disks/zfs/{name}
func (c *Client) NodesDisksZfsDetail(ctx context.Context, request *NodesDisksZfsDetailRequest) (*NodesDisksZfsDetailResponse, error) {

	method := "GET"
	path := "/nodes/{node}/disks/zfs/{name}"

	var response NodesDisksZfsDetailResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
