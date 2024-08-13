package apiv1

// SetClusterConfigRPC is the path for the SetClusterConfig RPC.
const SetClusterConfigRPC = "k8sd/cluster/config"

// SetClusterConfigRequest is the request message for the SetClusterConfig RPC.
type SetClusterConfigRequest struct {
	Config    UserFacingClusterConfig   `json:"config,omitempty" yaml:"config,omitempty"`
	Datastore UserFacingDatastoreConfig `json:"datastore,omitempty" yaml:"datastore,omitempty"`
}

// SetClusterConfigResponse is the response message for the SetClusterConfig RPC.
type SetClusterConfigResponse struct{}
