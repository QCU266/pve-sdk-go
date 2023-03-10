// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterCephFlagsSetFlagsRequest struct {
	Nobackfill  *bool `json:"nobackfill,omitempty"`   // Backfilling of PGs is suspended.
	NodeepScrub *bool `json:"nodeep-scrub,omitempty"` // Deep Scrubbing is disabled.
	Nodown      *bool `json:"nodown,omitempty"`       // OSD failure reports are being ignored, such that the monitors will not mark OSDs down.
	Noin        *bool `json:"noin,omitempty"`         // OSDs that were previously marked out will not be marked back in when they start.
	Noout       *bool `json:"noout,omitempty"`        // OSDs will not automatically be marked out after the configured interval.
	Norebalance *bool `json:"norebalance,omitempty"`  // Rebalancing of PGs is suspended.
	Norecover   *bool `json:"norecover,omitempty"`    // Recovery of PGs is suspended.
	Noscrub     *bool `json:"noscrub,omitempty"`      // Scrubbing is disabled.
	Notieragent *bool `json:"notieragent,omitempty"`  // Cache tiering activity is suspended.
	Noup        *bool `json:"noup,omitempty"`         // OSDs are not allowed to start.
	Pause       *bool `json:"pause,omitempty"`        // Pauses read and writes.
}

type ClusterCephFlagsSetFlagsResponse string

// Set/Unset multiple ceph flags at once.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/ceph/flags
func (c *Client) ClusterCephFlagsSetFlags(ctx context.Context, request *ClusterCephFlagsSetFlagsRequest) (*ClusterCephFlagsSetFlagsResponse, error) {

	method := "PUT"
	path := "/cluster/ceph/flags"

	var response ClusterCephFlagsSetFlagsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
