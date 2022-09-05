// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// FlexVolumeSource FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
//
// swagger:model FlexVolumeSource
type FlexVolumeSource struct {

	// driver is the name of the driver to use for this volume.
	// Required: true
	Driver *string `json:"driver"`

	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	FsType string `json:"fsType,omitempty"`

	// options is Optional: this field holds extra command options if any.
	Options map[string]string `json:"options,omitempty"`

	// readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
}
