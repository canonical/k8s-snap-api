package apiv1

import (
	"github.com/canonical/k8s-snap-api/internal/util"
)

// RefreshCertificatesUpdateRPC is the path for the RefreshCertificatesUpdate RPC.
const RefreshCertificatesUpdateRPC = "k8sd/refresh-certs/update"

// RefreshCertificatesUpdateRequest is the request message for the RefreshCertificatesUpdate RPC.
type RefreshCertificatesUpdateRequest struct {
	FrontProxyClientCert *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	// The client key to be used for the front proxy.
	FrontProxyClientKey *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	// The client certificate to be used by kubelet for communicating with the kube-apiserver.
	APIServerKubeletClientCert *string `json:"apiserver-kubelet-client-crt,omitempty" yaml:"apiserver-kubelet-client-crt,omitempty"`
	// The client key to be used by kubelet for communicating with the kube-apiserver.
	APIServerKubeletClientKey *string `json:"apiserver-kubelet-client-key,omitempty" yaml:"apiserver-kubelet-client-key,omitempty"`
	// The admin client certificate to be used for Kubernetes services.
	AdminClientCert *string `json:"admin-client-crt,omitempty" yaml:"admin-client-crt,omitempty"`
	// The admin client key to be used for Kubernetes services.
	AdminClientKey *string `json:"admin-client-key,omitempty" yaml:"admin-client-key,omitempty"`
	// The client certificate to be used for the kube-proxy.
	KubeProxyClientCert *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used for the kube-proxy.
	KubeProxyClientKey *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`
	// The client certificate to be used for the kube-scheduler.
	KubeSchedulerClientCert *string `json:"kube-scheduler-client-crt,omitempty" yaml:"kube-scheduler-client-crt,omitempty"`
	// The client key to be used for the kube-scheduler.
	KubeSchedulerClientKey *string `json:"kube-scheduler-client-key,omitempty" yaml:"kube-scheduler-client-key,omitempty"`
	// The client certificate to be used for the Kubernetes controller manager.
	KubeControllerManagerClientCert *string `json:"kube-controller-manager-client-crt,omitempty" yaml:"kube-controller-manager-client-crt,omitempty"`
	// The client key to be used for the Kubernetes controller manager.
	KubeControllerManagerClientKey *string `json:"kube-controller-manager-client-key,omitempty" yaml:"kube-controller-manager-client-key,omitempty"`
	// The certificate to be used for the kube-apiserver.
	APIServerCert *string `json:"apiserver-crt,omitempty" yaml:"apiserver-crt,omitempty"`
	// The key to be used for the kube-apiserver.
	APIServerKey *string `json:"apiserver-key,omitempty" yaml:"apiserver-key,omitempty"`
	// The certificate to be used for the kubelet.
	KubeletCert *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	// The key to be used for the kubelet.
	KubeletKey *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	// The certificate to be used for the kubelet client.
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The key to be used for the kubelet client.
	KubeletClientKey *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`
}

// RefreshCertificatesUpdateResponse is the request response for the RefreshCertificatesUpdate RPC.
type RefreshCertificatesUpdateResponse struct{}

func (r *RefreshCertificatesUpdateRequest) GetFrontProxyClientCert() string {
	return util.Deref(r.FrontProxyClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetFrontProxyClientKey() string {
	return util.Deref(r.FrontProxyClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetAPIServerKubeletClientCert() string {
	return util.Deref(r.APIServerKubeletClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetAPIServerKubeletClientKey() string {
	return util.Deref(r.APIServerKubeletClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetAdminClientCert() string {
	return util.Deref(r.AdminClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetAdminClientKey() string {
	return util.Deref(r.AdminClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeProxyClientCert() string {
	return util.Deref(r.KubeProxyClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeProxyClientKey() string {
	return util.Deref(r.KubeProxyClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeSchedulerClientCert() string {
	return util.Deref(r.KubeSchedulerClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeSchedulerClientKey() string {
	return util.Deref(r.KubeSchedulerClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeControllerManagerClientCert() string {
	return util.Deref(r.KubeControllerManagerClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeControllerManagerClientKey() string {
	return util.Deref(r.KubeControllerManagerClientKey)
}
func (r *RefreshCertificatesUpdateRequest) GetAPIServerCert() string {
	return util.Deref(r.APIServerCert)
}
func (r *RefreshCertificatesUpdateRequest) GetAPIServerKey() string {
	return util.Deref(r.APIServerKey)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeletCert() string { return util.Deref(r.KubeletCert) }
func (r *RefreshCertificatesUpdateRequest) GetKubeletKey() string  { return util.Deref(r.KubeletKey) }
func (r *RefreshCertificatesUpdateRequest) GetKubeletClientCert() string {
	return util.Deref(r.KubeletClientCert)
}
func (r *RefreshCertificatesUpdateRequest) GetKubeletClientKey() string {
	return util.Deref(r.KubeletClientKey)
}
