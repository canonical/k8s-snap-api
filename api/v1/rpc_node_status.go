package apiv1

// NodeStatusRPC is the path for the NodeStatus RPC.
const NodeStatusRPC = "k8sd/node"

// NodeStatusRequest is the request message for the NodeStatus RPC.
type NodeStatusRequest struct{}

// NodeStatusResponse is the response message for the NodeStatus RPC.
type NodeStatusResponse struct {
	NodeStatus NodeStatus `json:"status"`
	// Taints is a list of taints applied to the node.
	Taints []string `json:"taints,omitempty" yaml:"taints,omitempty"`
}
