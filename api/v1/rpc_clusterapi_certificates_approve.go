package apiv1

const ClusterAPIApproveWorkerCSRRPC = "x/capi/refresh-certs/approve"

// ClusterAPIApproveWorkerCSRRequest is the request message for the ClusterAPIApproveWorkerCSR RPC.
type ClusterAPIApproveWorkerCSRRequest struct {
	Seed int `json:"seed"`
}

// ClusterAPIApproveWorkerCSRResponse is the response message for the ClusterAPIApproveWorkerCSR RPC.
type ClusterAPIApproveWorkerCSRResponse struct{}
