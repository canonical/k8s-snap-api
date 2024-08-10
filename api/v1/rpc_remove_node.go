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
type RemoveNodeResponse struct{}
