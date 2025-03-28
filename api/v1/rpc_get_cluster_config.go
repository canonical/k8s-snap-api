package apiv1

// GetClusterConfigRPC is the path for the GetClusterConfig RPC.
const GetClusterConfigRPC = "k8sd/cluster/config"

// GetClusterConfigRequest is the request message for the GetClusterConfig RPC.
type GetClusterConfigRequest struct{}

// GetClusterConfigResponse is the response message for the GetClusterConfig RPC.
type GetClusterConfigResponse struct {
	Config UserFacingClusterConfig `json:"status"`
	// Datastore is the datastore configuration.
	Datastore UserFacingDatastoreConfig `json:"datastore,omitempty" yaml:"datastore,omitempty"`
	// PodCIDR is the CIDR range for the pods in the cluster.
	PodCIDR *string `json:"pod-cidr,omitempty" yaml:"pod-cidr,omitempty"`
	// ServiceCIDR is the CIDR range for the services in the cluster.
	ServiceCIDR *string `json:"service-cidr,omitempty" yaml:"service-cidr,omitempty"`
}
