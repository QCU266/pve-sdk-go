// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesCertificatesInfoInfoRequest struct {
	Node string `query:"node,omitempty"` // The cluster node name.
}

type NodesCertificatesInfoInfoResponse []NodesCertificatesInfoInfoResponseItem

type NodesCertificatesInfoInfoResponseItem struct {
	Filename      *string                                         `json:"filename,omitempty"`        //
	Fingerprint   *string                                         `json:"fingerprint,omitempty"`     // Certificate SHA 256 fingerprint.
	Issuer        *string                                         `json:"issuer,omitempty"`          // Certificate issuer name.
	Notafter      *int64                                          `json:"notafter,omitempty"`        // Certificate's notAfter timestamp (UNIX epoch).
	Notbefore     *int64                                          `json:"notbefore,omitempty"`       // Certificate's notBefore timestamp (UNIX epoch).
	Pem           *string                                         `json:"pem,omitempty"`             // Certificate in PEM format
	PublicKeyBits *int64                                          `json:"public-key-bits,omitempty"` // Certificate's public key size
	PublicKeyType *string                                         `json:"public-key-type,omitempty"` // Certificate's public key algorithm
	San           *[]NodesCertificatesInfoInfoResponseItemSanItem `json:"san,omitempty"`             // List of Certificate's SubjectAlternativeName entries.
	Subject       *string                                         `json:"subject,omitempty"`         // Certificate subject name.
}

type NodesCertificatesInfoInfoResponseItemSanItem string

// Get information about node's certificates.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/certificates/info
func (c *Client) NodesCertificatesInfoInfo(ctx context.Context, request *NodesCertificatesInfoInfoRequest) (*NodesCertificatesInfoInfoResponse, error) {

	method := "GET"
	path := "/nodes/{node}/certificates/info"

	var response NodesCertificatesInfoInfoResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
