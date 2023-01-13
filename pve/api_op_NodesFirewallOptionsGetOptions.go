// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesFirewallOptionsGetOptionsRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesFirewallOptionsGetOptionsResponse struct {
	Enable                           *bool   `json:"enable,omitempty"`                               // Enable host firewall rules.
	LogLevelIn                       *string `json:"log_level_in,omitempty"`                         // Log level for incoming traffic.
	LogLevelOut                      *string `json:"log_level_out,omitempty"`                        // Log level for outgoing traffic.
	LogNfConntrack                   *bool   `json:"log_nf_conntrack,omitempty"`                     // Enable logging of conntrack information.
	Ndp                              *bool   `json:"ndp,omitempty"`                                  // Enable NDP (Neighbor Discovery Protocol).
	NfConntrackAllowInvalid          *bool   `json:"nf_conntrack_allow_invalid,omitempty"`           // Allow invalid packets on connection tracking.
	NfConntrackMax                   *int64  `json:"nf_conntrack_max,omitempty"`                     // Maximum number of tracked connections.
	NfConntrackTcpTimeoutEstablished *int64  `json:"nf_conntrack_tcp_timeout_established,omitempty"` // Conntrack established timeout.
	NfConntrackTcpTimeoutSynRecv     *int64  `json:"nf_conntrack_tcp_timeout_syn_recv,omitempty"`    // Conntrack syn recv timeout.
	Nosmurfs                         *bool   `json:"nosmurfs,omitempty"`                             // Enable SMURFS filter.
	ProtectionSynflood               *bool   `json:"protection_synflood,omitempty"`                  // Enable synflood protection
	ProtectionSynfloodBurst          *int64  `json:"protection_synflood_burst,omitempty"`            // Synflood protection rate burst by ip src.
	ProtectionSynfloodRate           *int64  `json:"protection_synflood_rate,omitempty"`             // Synflood protection rate syn/sec by ip src.
	SmurfLogLevel                    *string `json:"smurf_log_level,omitempty"`                      // Log level for SMURFS filter.
	TcpFlagsLogLevel                 *string `json:"tcp_flags_log_level,omitempty"`                  // Log level for illegal tcp flags filter.
	Tcpflags                         *bool   `json:"tcpflags,omitempty"`                             // Filter illegal combinations of TCP flags.
}

// Get host firewall options.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/firewall/options
func (c *Client) NodesFirewallOptionsGetOptions(ctx context.Context, request *NodesFirewallOptionsGetOptionsRequest) (*NodesFirewallOptionsGetOptionsResponse, error) {

	method := "GET"
	path := "/nodes/{node}/firewall/options"

	var response NodesFirewallOptionsGetOptionsResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}