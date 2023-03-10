// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesQemuAgentExecExecRequest struct {
	Command   *string `json:"command,omitempty"`    // The command as a list of program + arguments
	InputData *string `json:"input-data,omitempty"` // Data to pass as 'input-data' to the guest. Usually treated as STDIN to 'command'.
	Node      string  `json:"node,omitempty"`       // The cluster node name.
	Vmid      int64   `json:"vmid,omitempty"`       // The (unique) ID of the VM.
}

type NodesQemuAgentExecExecResponse struct {
	Pid int64 `json:"pid,omitempty"` // The PID of the process started by the guest-agent.
}

// Executes the given command in the vm via the guest-agent and returns an object with the pid.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu/{vmid}/agent/exec
func (c *Client) NodesQemuAgentExecExec(ctx context.Context, request *NodesQemuAgentExecExecRequest) (*NodesQemuAgentExecExecResponse, error) {

	method := "POST"
	path := "/nodes/{node}/qemu/{vmid}/agent/exec"

	var response NodesQemuAgentExecExecResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
