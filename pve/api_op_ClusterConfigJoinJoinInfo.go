// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type ClusterConfigJoinJoinInfoRequest struct {
	Node *string `query:"node,omitempty"` // The node for which the joinee gets the nodeinfo.
}

type ClusterConfigJoinJoinInfoResponse struct {
	ConfigDigest  string                                          `json:"config_digest,omitempty"`  //
	Nodelist      []ClusterConfigJoinJoinInfoResponseNodelistItem `json:"nodelist,omitempty"`       //
	PreferredNode string                                          `json:"preferred_node,omitempty"` // The cluster node name.
	Totem         ClusterConfigJoinJoinInfoResponseTotem          `json:"totem,omitempty"`          //
}

type ClusterConfigJoinJoinInfoResponseNodelistItem struct {
	Name        string  `json:"name,omitempty"`         // The cluster node name.
	Nodeid      *int64  `json:"nodeid,omitempty"`       // Node id for this node.
	PveAddr     string  `json:"pve_addr,omitempty"`     //
	PveFp       string  `json:"pve_fp,omitempty"`       // Certificate SHA 256 fingerprint.
	QuorumVotes int64   `json:"quorum_votes,omitempty"` //
	Ring0Addr   *string `json:"ring0_addr,omitempty"`   // Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
}

type ClusterConfigJoinJoinInfoResponseTotem interface{}

// Get information needed to join this cluster over the connected node.
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/config/join
func (c *Client) ClusterConfigJoinJoinInfo(ctx context.Context, request *ClusterConfigJoinJoinInfoRequest) (*ClusterConfigJoinJoinInfoResponse, error) {

	method := "GET"
	path := "/cluster/config/join"

	var response ClusterConfigJoinJoinInfoResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
