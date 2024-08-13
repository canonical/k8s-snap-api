package apiv1

import "time"

// FeatureStatus encapsulates the deployment status of a feature.
type FeatureStatus struct {
	// Enabled shows whether or not the deployment of manifests for a status was successful.
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Message contains information about the status of a feature. It is only supposed to be human readable and informative and should not be programmatically parsed.
	Message string `json:"message" yaml:"message"`
	// Version shows the version of the deployed feature.
	Version string `json:"version" yaml:"version"`
	// UpdatedAt shows when the last update was done.
	UpdatedAt time.Time `json:"updated-at" yaml:"updated-at"`
}

func (f FeatureStatus) String() string {
	if f.Message != "" {
		return f.Message
	}
	if f.Enabled {
		return "enabled"
	}
	return "disabled"
}
