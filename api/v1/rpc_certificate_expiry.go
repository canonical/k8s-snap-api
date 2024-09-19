package apiv1

// ClusterAPICertificatesExpiryRPC is the path for the ClusterAPICertificatesExpiry RPC.
const ClusterAPICertificatesExpiryRPC = "x/capi/certificates-expiry"

type CertificatesExpiryResponse struct {
	// ExpiryDate is the expiry date of the certificates on the node.
	ExpiryDate string `json:"expiry-date"`
}
