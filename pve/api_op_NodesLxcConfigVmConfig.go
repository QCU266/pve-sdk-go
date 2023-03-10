// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type NodesLxcConfigVmConfigRequest struct {
	Current  *bool   `query:"current,omitempty"`  // Get current values (instead of pending values).
	Node     string  `query:"node,omitempty"`     // The cluster node name.
	Snapshot *string `query:"snapshot,omitempty"` // Fetch config values from given snapshot.
	Vmid     int64   `query:"vmid,omitempty"`     // The (unique) ID of the VM.
}

type NodesLxcConfigVmConfigResponse struct {
	Arch         *string                                  `json:"arch,omitempty"`         // OS architecture type.
	Cmode        *string                                  `json:"cmode,omitempty"`        // Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Console      *bool                                    `json:"console,omitempty"`      // Attach a console device (/dev/console) to the container.
	Cores        *int64                                   `json:"cores,omitempty"`        // The number of cores assigned to the container. A container can use all available cores by default.
	Cpulimit     *float64                                 `json:"cpulimit,omitempty"`     // Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Cpuunits     *int64                                   `json:"cpuunits,omitempty"`     // CPU weight for a container, will be clamped to [1, 10000] in cgroup v2.
	Debug        *bool                                    `json:"debug,omitempty"`        // Try to be more verbose. For now this only enables debug log-level on start.
	Description  *string                                  `json:"description,omitempty"`  // Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Digest       string                                   `json:"digest,omitempty"`       // SHA1 digest of configuration file. This can be used to prevent concurrent modifications.
	Features     *string                                  `json:"features,omitempty"`     // Allow containers access to advanced features.
	Hookscript   *string                                  `json:"hookscript,omitempty"`   // Script that will be exectued during various steps in the containers lifetime.
	Hostname     *string                                  `json:"hostname,omitempty"`     // Set a host name for the container.
	Lock         *string                                  `json:"lock,omitempty"`         // Lock/unlock the container.
	Lxc          *[]NodesLxcConfigVmConfigResponseLxcItem `json:"lxc,omitempty"`          // Array of lxc low-level configurations ([[key1, value1], [key2, value2] ...]).
	Memory       *int64                                   `json:"memory,omitempty"`       // Amount of RAM for the container in MB.
	Mpn          *string                                  `json:"mp[n],omitempty"`        // Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	Nameserver   *string                                  `json:"nameserver,omitempty"`   // Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Netn         *string                                  `json:"net[n],omitempty"`       // Specifies network interfaces for the container.
	Onboot       *bool                                    `json:"onboot,omitempty"`       // Specifies whether a container will be started during system bootup.
	Ostype       *string                                  `json:"ostype,omitempty"`       // OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Protection   *bool                                    `json:"protection,omitempty"`   // Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Rootfs       *string                                  `json:"rootfs,omitempty"`       // Use volume as container root.
	Searchdomain *string                                  `json:"searchdomain,omitempty"` // Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Startup      *string                                  `json:"startup,omitempty"`      // Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Swap         *int64                                   `json:"swap,omitempty"`         // Amount of SWAP for the container in MB.
	Tags         *string                                  `json:"tags,omitempty"`         // Tags of the Container. This is only meta information.
	Template     *bool                                    `json:"template,omitempty"`     // Enable/disable Template.
	Timezone     *string                                  `json:"timezone,omitempty"`     // Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Tty          *int64                                   `json:"tty,omitempty"`          // Specify the number of tty available to the container
	Unprivileged *bool                                    `json:"unprivileged,omitempty"` // Makes the container run as unprivileged user. (Should not be modified manually.)
	Unusedn      *string                                  `json:"unused[n],omitempty"`    // Reference to unused volumes. This is used internally, and should not be modified manually.
}

type NodesLxcConfigVmConfigResponseLxcItem []string

// Get container configuration.
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/config
func (c *Client) NodesLxcConfigVmConfig(ctx context.Context, request *NodesLxcConfigVmConfigRequest) (*NodesLxcConfigVmConfigResponse, error) {

	method := "GET"
	path := "/nodes/{node}/lxc/{vmid}/config"

	var response NodesLxcConfigVmConfigResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
