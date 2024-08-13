package apiv1

// KubeConfigRPC is the path for the KubeConfig RPC.
const KubeConfigRPC = "k8sd/kubeconfig"

// KubeConfigRequest is the request message for the KubeConfig RPC.
type KubeConfigRequest struct {
	// Server is the server URL to use (e.g. in case of an external LoadBalancer endpoint).
	Server string `json:"server"`
}

// KubeConfigResponse is the response message for the KubeConfig RPC.
type KubeConfigResponse struct {
	// KubeConfig is an admin kubeconfig that can be used to access the cluster.
	KubeConfig string `json:"kubeconfig"`
}
