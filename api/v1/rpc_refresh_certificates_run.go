package apiv1

// RefreshCertificatesRunRPC is the path for the RefreshCertificatesRun RPC.
const RefreshCertificatesRunRPC = "k8sd/refresh-certs/run"

// RefreshCertificatesRunRequest is the request message for the RefreshCertificatesRun RPC.
type RefreshCertificatesRunRequest struct {
	// Seed must match the value returned by the RefreshCertificatesPlan RPC.
	Seed int `json:"seed"`
	// ExpirationSeconds is the desired duration of the new certificates.
	ExpirationSeconds int `json:"expiration-seconds"`
	// ExtraSANs is a list of extra SANs (DNS names or IP addresses) to add to the kube-apiserver certificates.
	ExtraSANs []string `json:"extra-sans"`
}

// RefreshCertificatesRunResponse is the response message for the RefreshCertificatesRun RPC.
type RefreshCertificatesRunResponse struct {
	// ExpirationSeconds is the duration of the new certificates (might not match the requested value).
	ExpirationSeconds int `json:"expiration-seconds"`
}
