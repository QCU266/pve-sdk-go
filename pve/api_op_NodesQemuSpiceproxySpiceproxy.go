// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuSpiceproxySpiceproxyRequest struct {
	Node  string  `json:"node,omitempty"`  // The cluster node name.
	Proxy *string `json:"proxy,omitempty"` // SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Vmid  int64   `json:"vmid,omitempty"`  // The (unique) ID of the VM.
}

// Returned values can be directly passed to the 'remote-viewer' application.
type NodesQemuSpiceproxySpiceproxyResponse struct {
	Host     string `json:"host,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Proxy    string `json:"proxy,omitempty"`    //
	TlsPort  int64  `json:"tls-port,omitempty"` //
	Type     string `json:"type,omitempty"`     //
}

// Returns a SPICE configuration to connect to the VM.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/spiceproxy
func (c *Client) NodesQemuSpiceproxySpiceproxy(ctx context.Context, request *NodesQemuSpiceproxySpiceproxyRequest) (*NodesQemuSpiceproxySpiceproxyResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/spiceproxy"

	var response NodesQemuSpiceproxySpiceproxyResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}