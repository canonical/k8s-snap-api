package apiv1

// SnapRefreshRPC is the path for the SnapRefresh RPC.
const SnapRefreshRPC = "snap/refresh"

// SnapRefreshRequest is the request message for the SnapRefresh RPC.
type SnapRefreshRequest struct {
	// Channel is the channel to refresh the snap to.
	Channel string
	// Revision is the revision number to refresh the snap to.
	Revision string
	// LocalPath is the local path to use to refresh the snap.
	LocalPath string
}
