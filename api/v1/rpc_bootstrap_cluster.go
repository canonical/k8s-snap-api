package apiv1

import "time"

// BootstrapClusterRPC is the path for the BootstrapCluster RPC.
const BootstrapClusterRPC = "k8sd/cluster"

// BootstrapClusterRequest is the request message for the BootstrapCluster RPC.
type BootstrapClusterRequest struct {
	Name    string          `json:"name"`
	Address string          `json:"address"`
	Config  BootstrapConfig `json:"config"`
	Timeout time.Duration   `json:"timeout"`
}

// BootstrapClusterResponse is the response message for the BootstrapClusterRPC.
type BootstrapClusterResponse struct{}
