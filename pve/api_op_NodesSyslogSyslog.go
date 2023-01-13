// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesSyslogSyslogRequest struct {
	Limit   *int64  `query:"limit,omitempty"`   //
	Node    string  `query:"node,omitempty"`    // The cluster node name.
	Service *string `query:"service,omitempty"` // Service ID
	Since   *string `query:"since,omitempty"`   // Display all log since this date-time string.
	Start   *int64  `query:"start,omitempty"`   //
	Until   *string `query:"until,omitempty"`   // Display all log until this date-time string.
}

type NodesSyslogSyslogResponseItem struct {
	N int64  `json:"n,omitempty"` // Line number
	T string `json:"t,omitempty"` // Line text
}

type NodesSyslogSyslogResponse []NodesSyslogSyslogResponseItem

// Read system log
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/syslog
func (c *Client) NodesSyslogSyslog(ctx context.Context, request *NodesSyslogSyslogRequest) (*NodesSyslogSyslogResponse, error) {

	method := "GET"
	path := "/nodes/{node}/syslog"

	var response NodesSyslogSyslogResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}