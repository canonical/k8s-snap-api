package apiv1

// GetClusterConfigRPC is the path for the GetClusterConfig RPC.
const GetClusterConfigRPC = "k8sd/cluster/config"

// GetClusterConfigRequest is the request message for the GetClusterConfig RPC.
type GetClusterConfigRequest struct{}

// GetClusterConfigResponse is the response message for the GetClusterConfig RPC.
type GetClusterConfigResponse struct {
	// Config is the cluster configuration.
	Config UserFacingClusterConfig `json:"status"`
	// Datastore is the datastore configuration.
	Datastore UserFacingDatastoreConfig `json:"datastore,omitempty" yaml:"datastore,omitempty"`
}
