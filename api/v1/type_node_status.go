package apiv1

// NodeStatus holds information about a node in the k8s cluster.
type NodeStatus struct {
	// Name is the name for this cluster member that was when joining the cluster.
	// This is typically the hostname of the node.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Address is the IP address of the node.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
	// ClusterRole is the role that the node has within the k8s cluster.
	ClusterRole ClusterRole `json:"cluster-role,omitempty" yaml:"cluster-role,omitempty"`
	// DatastoreRole is the role that the node has within the datastore cluster.
	// Only applicable for control-plane nodes, empty for workers.
	DatastoreRole DatastoreRole `json:"datastore-role,omitempty" yaml:"datastore-role,omitempty"`
}
