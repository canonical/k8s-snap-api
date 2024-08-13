package apiv1

// ClusterStatusRPC is the path for the ClusterStatus RPC.
const ClusterStatusRPC = "k8sd/cluster"

// ClusterStatusRequest is the request message for the ClusterStatus RPC.
type ClusterStatusRequest struct{}

// ClusterStatusResponse is the response message for the ClusterStatus RPC.
type ClusterStatusResponse struct {
	ClusterStatus ClusterStatus `json:"status"`
}
