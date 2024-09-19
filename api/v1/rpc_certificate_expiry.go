package apiv1

// ClusterAPICertificatesExpiryRPC is the path for the ClusterAPICertificatesExpiry RPC.
const ClusterAPICertificatesExpiryRPC = "x/capi/certificates-expiry"

// CertificatesExpiryRequest is the request message for the CertificatesExpiry RPC.
type CertificatesExpiryRequest struct{}

// CertificatesExpiryResponse is the response message for the CertificatesExpiry RPC.
type CertificatesExpiryResponse struct {
	// ExpiryDate is the expiry date of the certificates on the node in RFC3339 format.
	ExpiryDate string `json:"expiry-date"`
}
