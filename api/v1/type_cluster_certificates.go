package apiv1

// CertificateName represents the name of a certificate used by components in the cluster.
type CertificateName string

const (
	// CertificateFrontProxyClient is used by the front proxy client to authenticate with the API server.
	CertificateFrontProxyClient CertificateName = "front-proxy-client"

	// CertificateAPIServerKubeletClient is used by the API server to authenticate with kubelets.
	CertificateAPIServerKubeletClient CertificateName = "apiserver-kubelet-client"

	// CertificateAPIServer is the serving certificate for the Kubernetes API server.
	CertificateAPIServer CertificateName = "apiserver"

	// CertificateKubelet is the serving certificate for the kubelet.
	CertificateKubelet CertificateName = "kubelet"

	// CertificateAdminClient is the client certificate for cluster administrators, used in admin.conf.
	CertificateAdminClient CertificateName = "admin.conf"

	// CertificateSchedulerClient is the client certificate for the kube-scheduler, used in scheduler.conf.
	CertificateSchedulerClient CertificateName = "scheduler.conf"

	// CertificateControllerManagerClient is the client certificate for the kube-controller-manager, used in controller.conf.
	CertificateControllerManagerClient CertificateName = "controller.conf"

	// CertificateKubeletClient is the client certificate used by the kubelet to authenticate to the API server, used in kubelet.conf.
	CertificateKubeletClient CertificateName = "kubelet.conf"

	// CertificateProxyClient is the client certificate for the kube-proxy, used in proxy.conf.
	CertificateProxyClient CertificateName = "proxy.conf"
)

// Map roles to the set of valid certificate names for that role.
var CertificatesByRole = map[ClusterRole]map[CertificateName]struct{}{
	ClusterRoleWorker: {
		CertificateKubelet:       {},
		CertificateKubeletClient: {},
		CertificateProxyClient:   {},
	},
	ClusterRoleControlPlane: {
		CertificateAdminClient:             {},
		CertificateFrontProxyClient:        {},
		CertificateAPIServerKubeletClient:  {},
		CertificateSchedulerClient:         {},
		CertificateControllerManagerClient: {},
		CertificateAPIServer:               {},
		CertificateKubeletClient:           {},
		CertificateKubelet:                 {},
		CertificateProxyClient:             {},
	},
	// NOTE: Required due to exhaustive check
	ClusterRoleUnknown: {},
}
