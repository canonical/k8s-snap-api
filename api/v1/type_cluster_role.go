package apiv1

type ClusterRole string

const (
	ClusterRoleControlPlane ClusterRole = "control-plane"
	ClusterRoleWorker       ClusterRole = "worker"
	// The role of a node is unknown if it has not yet joined a cluster,
	// currently joining or is about to leave.
	ClusterRoleUnknown ClusterRole = "unknown"
)
