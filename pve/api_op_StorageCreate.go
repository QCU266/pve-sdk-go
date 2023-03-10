// Code generated by gen-api. DO NOT EDIT.

package pve

import (
	"context"
)

type StorageCreateRequest struct {
	Authsupported        *string `json:"authsupported,omitempty"`         // Authsupported.
	Base                 *string `json:"base,omitempty"`                  // Base volume. This volume is automatically activated.
	Blocksize            *string `json:"blocksize,omitempty"`             // block size
	Bwlimit              *string `json:"bwlimit,omitempty"`               // Set bandwidth/io limits various operations.
	ComstarHg            *string `json:"comstar_hg,omitempty"`            // host group for comstar views
	ComstarTg            *string `json:"comstar_tg,omitempty"`            // target group for comstar views
	Content              *string `json:"content,omitempty"`               // Allowed content types.  NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	DataPool             *string `json:"data-pool,omitempty"`             // Data Pool (for erasure coding only)
	Datastore            *string `json:"datastore,omitempty"`             // Proxmox Backup Server datastore name.
	Disable              *bool   `json:"disable,omitempty"`               // Flag to disable the storage.
	Domain               *string `json:"domain,omitempty"`                // CIFS domain.
	EncryptionKey        *string `json:"encryption-key,omitempty"`        // Encryption key. Use 'autogen' to generate one automatically without passphrase.
	Export               *string `json:"export,omitempty"`                // NFS export path.
	Fingerprint          *string `json:"fingerprint,omitempty"`           // Certificate SHA 256 fingerprint.
	Format               *string `json:"format,omitempty"`                // Default image format.
	FsName               *string `json:"fs-name,omitempty"`               // The Ceph filesystem name.
	Fuse                 *bool   `json:"fuse,omitempty"`                  // Mount CephFS through FUSE.
	IsMountpoint         *string `json:"is_mountpoint,omitempty"`         // Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	Iscsiprovider        *string `json:"iscsiprovider,omitempty"`         // iscsi provider
	Keyring              *string `json:"keyring,omitempty"`               // Client keyring contents (for external clusters).
	Krbd                 *bool   `json:"krbd,omitempty"`                  // Always access rbd through krbd kernel module.
	LioTpg               *string `json:"lio_tpg,omitempty"`               // target portal group for Linux LIO targets
	MasterPubkey         *string `json:"master-pubkey,omitempty"`         // Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MaxProtectedBackups  *int64  `json:"max-protected-backups,omitempty"` // Maximal number of protected backups per guest. Use '-1' for unlimited.
	Maxfiles             *int64  `json:"maxfiles,omitempty"`              // Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Mkdir                *bool   `json:"mkdir,omitempty"`                 // Create the directory if it doesn't exist.
	Monhost              *string `json:"monhost,omitempty"`               // IP addresses of monitors (for external clusters).
	Mountpoint           *string `json:"mountpoint,omitempty"`            // mount point
	Namespace            *string `json:"namespace,omitempty"`             // Namespace.
	Nocow                *bool   `json:"nocow,omitempty"`                 // Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Nodes                *string `json:"nodes,omitempty"`                 // List of cluster node names.
	Nowritecache         *bool   `json:"nowritecache,omitempty"`          // disable write caching on the target
	Options              *string `json:"options,omitempty"`               // NFS mount options (see 'man nfs')
	Password             *string `json:"password,omitempty"`              // Password for accessing the share/datastore.
	Path                 *string `json:"path,omitempty"`                  // File system path.
	Pool                 *string `json:"pool,omitempty"`                  // Pool.
	Port                 *int64  `json:"port,omitempty"`                  // For non default port.
	Portal               *string `json:"portal,omitempty"`                // iSCSI portal (IP or DNS name with optional port).
	Preallocation        *string `json:"preallocation,omitempty"`         // Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	PruneBackups         *string `json:"prune-backups,omitempty"`         // The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	Saferemove           *bool   `json:"saferemove,omitempty"`            // Zero-out data when removing LVs.
	SaferemoveThroughput *string `json:"saferemove_throughput,omitempty"` // Wipe throughput (cstream -t parameter value).
	Server               *string `json:"server,omitempty"`                // Server IP or DNS name.
	Server2              *string `json:"server2,omitempty"`               // Backup volfile server IP or DNS name.
	Share                *string `json:"share,omitempty"`                 // CIFS share.
	Shared               *bool   `json:"shared,omitempty"`                // Mark storage as shared.
	Smbversion           *string `json:"smbversion,omitempty"`            // SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Sparse               *bool   `json:"sparse,omitempty"`                // use sparse volumes
	Storage              string  `json:"storage,omitempty"`               // The storage identifier.
	Subdir               *string `json:"subdir,omitempty"`                // Subdir to mount.
	TaggedOnly           *bool   `json:"tagged_only,omitempty"`           // Only use logical volumes tagged with 'pve-vm-ID'.
	Target               *string `json:"target,omitempty"`                // iSCSI target.
	Thinpool             *string `json:"thinpool,omitempty"`              // LVM thin pool LV name.
	Transport            *string `json:"transport,omitempty"`             // Gluster transport: tcp or rdma
	Type                 string  `json:"type,omitempty"`                  // Storage type.
	Username             *string `json:"username,omitempty"`              // RBD Id.
	Vgname               *string `json:"vgname,omitempty"`                // Volume group name.
	Volume               *string `json:"volume,omitempty"`                // Glusterfs Volume.
}

type StorageCreateResponse struct {
	Config  *StorageCreateResponseConfig `json:"config,omitempty"`  // Partial, possible server generated, configuration properties.
	Storage string                       `json:"storage,omitempty"` // The ID of the created storage.
	Type    string                       `json:"type,omitempty"`    // The type of the created storage.
}

// Partial, possible server generated, configuration properties.
type StorageCreateResponseConfig struct {
	EncryptionKey *string `json:"encryption-key,omitempty"` // The, possible auto-generated, encryption-key.
}

// Create a new storage.
// https://pve.proxmox.com/pve-docs/api-viewer/#/storage
func (c *Client) StorageCreate(ctx context.Context, request *StorageCreateRequest) (*StorageCreateResponse, error) {

	method := "POST"
	path := "/storage"

	var response StorageCreateResponse

	err := c.apiCall(ctx, method, path, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
