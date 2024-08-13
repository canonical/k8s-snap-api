package apiv1

// ClusterStatus holds information about the cluster, e.g. its current members
type ClusterStatus struct {
	// Ready is true if at least one node in the cluster is in READY state.
	Ready     bool                    `json:"ready,omitempty"`
	Members   []NodeStatus            `json:"members,omitempty"`
	Config    UserFacingClusterConfig `json:"config,omitempty"`
	Datastore Datastore               `json:"datastore,omitempty"`

	DNS           FeatureStatus `json:"dns,omitempty" yaml:"dns,omitempty"`
	Network       FeatureStatus `json:"network,omitempty" yaml:"network,omitempty"`
	LoadBalancer  FeatureStatus `json:"load-balancer,omitempty" yaml:"load-balancer,omitempty"`
	Ingress       FeatureStatus `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Gateway       FeatureStatus `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	MetricsServer FeatureStatus `json:"metrics-server,omitempty" yaml:"metrics-server,omitempty"`
	LocalStorage  FeatureStatus `json:"local-storage,omitempty" yaml:"local-storage,omitempty"`
}
