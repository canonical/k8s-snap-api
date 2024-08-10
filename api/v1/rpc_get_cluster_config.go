package apiv1

// GetClusterConfigRPC is the path for the GetClusterConfig RPC.
const GetClusterConfigRPC = "k8sd/cluster/config"

// GetClusterConfigRequest is the request message for the GetClusterConfig RPC.
type GetClusterConfigRequest struct{}

// GetClusterConfigResponse is the response message for the GetClusterConfig RPC.
type GetClusterConfigResponse struct {
	Config UserFacingClusterConfig `json:"status"`
}
