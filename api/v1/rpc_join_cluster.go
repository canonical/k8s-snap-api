package apiv1

import "time"

// JoinClusterRPC is the path for the JoinCluster RPC.
const JoinClusterRPC = "k8sd/cluster/join"

// JoinClusterRequest is the request message for the JoinCluster RPC.
type JoinClusterRequest struct {
	// Name of the node that joins.
	Name string `json:"name"`
	// Address to use for microcluster on the joining node.
	Address string `json:"address"`
	// Token is the join token.
	Token string `json:"token"`
	// Config is JSON formatted string of a ControlPlaneJoinConfig (for control plane) or a WorkerJoinConfig (for worker nodes).
	Config string `json:"config"`
	// Timeout is how long to wait until the join is complete.
	Timeout time.Duration `json:"timeout"`
}

// JoinClusterResponse is the response message for the JoinCluster RPC.
type JoinClusterResponse struct{}
