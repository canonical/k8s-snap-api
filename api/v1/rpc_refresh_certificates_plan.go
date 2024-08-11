package apiv1

// RefreshCertificatesPlanRPC is the path for the RefreshCertificatesPlan RPC.
const RefreshCertificatesPlanRPC = "k8sd/refresh-certs/plan"

// RefreshCertificatesPlanRequest is the request message for the RefreshCertificatesPlan RPC.
type RefreshCertificatesPlanRequest struct{}

// RefreshCertificatesPlanResponse is the response message for the RefreshCertificatesPlan RPC.
type RefreshCertificatesPlanResponse struct {
	// Seed should be passed by clients to the RefreshCertificatesRun RPC.
	Seed int `json:"seconds"`
	// CertificateSigningRequests is a list of names of the CertificateSigningRequests that need to be signed externally (for worker nodes).
	CertificateSigningRequests []string `json:"certificate-signing-requests"`
}
