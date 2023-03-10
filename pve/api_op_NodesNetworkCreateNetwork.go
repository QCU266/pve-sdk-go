// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesNetworkCreateNetworkRequest struct {
	Address            *string `json:"address,omitempty"`               // IP address.
	Address6           *string `json:"address6,omitempty"`              // IP address.
	Autostart          *bool   `json:"autostart,omitempty"`             // Automatically start interface on boot.
	BondPrimary        *string `json:"bond-primary,omitempty"`          // Specify the primary interface for active-backup bond.
	BondMode           *string `json:"bond_mode,omitempty"`             // Bonding mode.
	BondXmitHashPolicy *string `json:"bond_xmit_hash_policy,omitempty"` // Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	BridgePorts        *string `json:"bridge_ports,omitempty"`          // Specify the interfaces you want to add to your bridge.
	BridgeVlanAware    *bool   `json:"bridge_vlan_aware,omitempty"`     // Enable bridge vlan support.
	Cidr               *string `json:"cidr,omitempty"`                  // IPv4 CIDR.
	Cidr6              *string `json:"cidr6,omitempty"`                 // IPv6 CIDR.
	Comments           *string `json:"comments,omitempty"`              // Comments
	Comments6          *string `json:"comments6,omitempty"`             // Comments
	Gateway            *string `json:"gateway,omitempty"`               // Default gateway address.
	Gateway6           *string `json:"gateway6,omitempty"`              // Default ipv6 gateway address.
	Iface              string  `json:"iface,omitempty"`                 // Network interface name.
	Mtu                *int64  `json:"mtu,omitempty"`                   // MTU.
	Netmask            *string `json:"netmask,omitempty"`               // Network mask.
	Netmask6           *int64  `json:"netmask6,omitempty"`              // Network mask.
	Node               string  `json:"node,omitempty"`                  // The cluster node name.
	OvsBonds           *string `json:"ovs_bonds,omitempty"`             // Specify the interfaces used by the bonding device.
	OvsBridge          *string `json:"ovs_bridge,omitempty"`            // The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	OvsOptions         *string `json:"ovs_options,omitempty"`           // OVS interface options.
	OvsPorts           *string `json:"ovs_ports,omitempty"`             // Specify the interfaces you want to add to your bridge.
	OvsTag             *int64  `json:"ovs_tag,omitempty"`               // Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	Slaves             *string `json:"slaves,omitempty"`                // Specify the interfaces used by the bonding device.
	Type               string  `json:"type,omitempty"`                  // Network interface type
	VlanId             *int64  `json:"vlan-id,omitempty"`               // vlan-id for a custom named vlan interface (ifupdown2 only).
	VlanRawDevice      *string `json:"vlan-raw-device,omitempty"`       // Specify the raw interface for the vlan interface.
}

type NodesNetworkCreateNetworkResponse struct{}

// Create network device configuration
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/network
func (c *Client) NodesNetworkCreateNetwork(ctx context.Context, request *NodesNetworkCreateNetworkRequest) (*NodesNetworkCreateNetworkResponse, error) {

	method := "POST"
	path := "/nodes/{node}/network"

	var response NodesNetworkCreateNetworkResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
