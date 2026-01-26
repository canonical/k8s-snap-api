package apiv1

// Upgrade holds information about an upgrade of the cluster.
type Upgrade struct {
	// To is the target version.
	To Version `json:"to,omitempty" yaml:"to,omitempty"`
	// Phase is the current phase of the upgrade.
	Phase string `json:"phase,omitempty" yaml:"phase,omitempty"`
}
