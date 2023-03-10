// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesHardwareUsbUsbscanRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesHardwareUsbUsbscanResponse []NodesHardwareUsbUsbscanResponseItem

type NodesHardwareUsbUsbscanResponseItem struct {
	Busnum       int64   `json:"busnum,omitempty"`       //
	Class        int64   `json:"class,omitempty"`        //
	Devnum       int64   `json:"devnum,omitempty"`       //
	Level        int64   `json:"level,omitempty"`        //
	Manufacturer *string `json:"manufacturer,omitempty"` //
	Port         int64   `json:"port,omitempty"`         //
	Prodid       string  `json:"prodid,omitempty"`       //
	Product      *string `json:"product,omitempty"`      //
	Serial       *string `json:"serial,omitempty"`       //
	Speed        string  `json:"speed,omitempty"`        //
	Usbpath      *string `json:"usbpath,omitempty"`      //
	Vendid       string  `json:"vendid,omitempty"`       //
}

// List local USB devices.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/hardware/usb
func (c *Client) NodesHardwareUsbUsbscan(ctx context.Context, request *NodesHardwareUsbUsbscanRequest) (*NodesHardwareUsbUsbscanResponse, error) {

	method := "GET"
	path := "/nodes/{node}/hardware/usb"

	var response NodesHardwareUsbUsbscanResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
