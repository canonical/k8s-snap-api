package apiv1

// SnapRefreshRPC is the path for the SnapRefresh RPC.
const SnapRefreshRPC = "snap/refresh"

// SnapRefreshRequest is the request message for the SnapRefresh RPC.
type SnapRefreshRequest struct {
	// Channel is the channel to refresh the snap to.
	Channel string `json:"channel"`
	// Revision is the revision number to refresh the snap to.
	Revision string `json:"revision"`
	// LocalPath is the local path to use to refresh the snap.
	LocalPath string `json:"localPath"`
}

// SnapRefreshResponse is the response message for the SnapRefresh RPC.
type SnapRefreshResponse struct {
	// The change id belonging to a snap refresh/install operation.
	ChangeID string `json:"changeId"`
}
