package apiv1

// SnapRefreshRPC is the path for the SnapRefresh RPC.
const SnapRefreshStatusRPC = "snap/refresh-status"

// SnapRefreshStatusRequest is the request message for the SnapRefreshStatus RPC.
type SnapRefreshStatusRequest struct {
	// The change id belonging to a snap refresh/install operation.
	ChangeID string `json:"changeId"`
}

// SnapRefreshStatusResponse is the response message for the SnapRefreshStatus RPC.
type SnapRefreshStatusResponse struct {
	// Status is the status of the snap refresh/install operation.
	Status string `json:"status"`
	// Completed is a boolean indicating if the snap refresh/install operation has completed.
	// The status should be considered final when this is true.
	Completed bool `json:"completed"`
	// ErrorMessage is the error message if the snap refresh/install operation failed.
	ErrorMessage string `json:"errorMessage"`
}
