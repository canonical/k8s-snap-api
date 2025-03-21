package apiv1

// ReconcileFeaturesRPC is the path for the ReconcileFeatures RPC.
const ReconcileFeaturesRPC = "k8sd/cluster/features/reconcile"

// ReconcileFeaturesRequest is the request message for the ReconcileFeatures RPC.
type ReconcileFeaturesRequest struct {
	Network       bool `json:"network"`
	Ingress       bool `json:"ingress"`
	LoadBalancer  bool `json:"load-balancer"`
	LocalStorage  bool `json:"local-storage"`
	Gateway       bool `json:"gateway"`
	MetricsServer bool `json:"metrics-server"`
	DNS           bool `json:"dns"`
}

// ReconcileFeaturesResponse is the response message for the ReconcileFeatures RPC.
type ReconcileFeaturesResponse struct{}
