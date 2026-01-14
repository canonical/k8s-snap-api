package apiv1

import "time"

// RemoveNodeRPC is the path for the RemoveNode RPC.
const RemoveNodeRPC = "k8sd/cluster/remove"

// RemoveNodeRequest is the request message for the RemoveNode RPC.
type RemoveNodeRequest struct {
	Name    string        `json:"name"`
	Force   bool          `json:"force"`
	Timeout time.Duration `json:"timeout"`
}

// RemoveNodeResponse is the response message for the RemoveNode RPC.
type RemoveNodeResponse struct {
	// Errors holds error messages from the various subsystems involved
	// (if any). This is useful when Force is true and some subsystems
	// fail to remove the node.
	Errors *RemoveNodeErrors `json:"errors,omitempty"`
}

// RemoveNodeErrors holds error messages from the various subsystems involved
type RemoveNodeErrors struct {
	Kubernetes   string `json:"kubernetes,omitempty"`
	Datastore    string `json:"datastore,omitempty"`
	Microcluster string `json:"microcluster,omitempty"`
}
