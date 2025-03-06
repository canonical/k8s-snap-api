package apiv1

// CertificatesStatusRPC is the path for the CertificatesStatus RPC.
const CertificatesStatusRPC = "k8sd/certs-status"

// CertificatesStatusRequest is the request message for the CertificatesStatus RPC.
type CertificatesStatusRequest struct{}

// CertificatesStatusResponse is the response message for the CertificatesStatus RPC.
type CertificatesStatusResponse struct {
	Certificates           []CertificateStatus          `json:"certificates"`
	CertificateAuthorities []CertificateAuthorityStatus `json:"certificate-authorities"`
}

// CertificateStatus represents the status of an individual certificate.
type CertificateStatus struct {
	// Name is the identifier of the certificate.
	Name string `json:"name"`
	// Expires is the expiration date of the certificates in RFC3339 format.
	Expires string `json:"expires"`
	// CertificateAuthority is the CN of the CA that issued this certificate.
	CertificateAuthority string `json:"certificate-authority"`
	// ExternallyManaged indicates whether the certificate is managed externally.
	ExternallyManaged bool `json:"externally-managed"`
}

// CertificateAuthorityStatus represents the status of a certificate authority (CA).
type CertificateAuthorityStatus struct {
	// Name is the identifier of the certificate authority.
	Name string `json:"name"`
	// Expires is the expiration date of the certificate authority in RFC3339 format.
	Expires string `json:"expires"`
	// ExternallyManaged indicates whether the certificate authority is managed externally.
	ExternallyManaged bool `json:"externally-managed"`
}
