// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerGlobalProxyDirect undocumented
type AndroidDeviceOwnerGlobalProxyDirect struct {
	// AndroidDeviceOwnerGlobalProxy is the base model of AndroidDeviceOwnerGlobalProxyDirect
	AndroidDeviceOwnerGlobalProxy
	// Host The host name
	Host *string `json:"host,omitempty"`
	// Port The port
	Port *int `json:"port,omitempty"`
	// ExcludedHosts The excluded hosts
	ExcludedHosts []string `json:"excludedHosts,omitempty"`
}