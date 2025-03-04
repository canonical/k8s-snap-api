package apiv1

const CertificatesStatusRPC = "k8sd/certs-status"

type CertificatesStatusRequest struct{}

type CertificatesStatusResponse struct {
	Certificates           []CertificateStatus          `json:"certificates"`
	CertificateAuthorities []CertificateAuthorityStatus `json:"certificate-authorities"`
}

type CertificateStatus struct {
	Name                 string `json:"name"`
	Expires              string `json:"expires"`
	CertificateAuthority string `json:"certificat-authority"`
	ExternallyManaged    bool   `json:"externally-managed"`
}

type CertificateAuthorityStatus struct {
	Name              string `json:"name"`
	Expires           string `json:"expires"`
	ExternallyManaged bool   `json:"externally-managed"`
}
