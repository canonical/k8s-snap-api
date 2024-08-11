package apiv1

import "github.com/canonical/k8s-snap-api-v1/internal/util"

type WorkerJoinConfig struct {
	KubeletCert         *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	KubeletKey          *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	KubeletClientCert   *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	KubeletClientKey    *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`
	KubeProxyClientCert *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	KubeProxyClientKey  *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`

	// ExtraNodeConfigFiles will be written to /var/snap/k8s/common/args/conf.d
	ExtraNodeConfigFiles map[string]string `json:"extra-node-config-files,omitempty" yaml:"extra-node-config-files,omitempty"`

	// Extra args to add to individual services (set any arg to null to delete)
	ExtraNodeKubeProxyArgs         map[string]*string `json:"extra-node-kube-proxy-args,omitempty" yaml:"extra-node-kube-proxy-args,omitempty"`
	ExtraNodeKubeletArgs           map[string]*string `json:"extra-node-kubelet-args,omitempty" yaml:"extra-node-kubelet-args,omitempty"`
	ExtraNodeContainerdArgs        map[string]*string `json:"extra-node-containerd-args,omitempty" yaml:"extra-node-containerd-args,omitempty"`
	ExtraNodeK8sAPIServerProxyArgs map[string]*string `json:"extra-node-k8s-apiserver-proxy-args,omitempty" yaml:"extra-node-k8s-apiserver-proxy-args,omitempty"`
}

func (w *WorkerJoinConfig) GetKubeletCert() string       { return util.Deref(w.KubeletCert) }
func (w *WorkerJoinConfig) GetKubeletKey() string        { return util.Deref(w.KubeletKey) }
func (w *WorkerJoinConfig) GetKubeletClientCert() string { return util.Deref(w.KubeletClientCert) }
func (w *WorkerJoinConfig) GetKubeletClientKey() string  { return util.Deref(w.KubeletClientKey) }
func (w *WorkerJoinConfig) GetKubeProxyClientCert() string {
	return util.Deref(w.KubeProxyClientCert)
}
func (w *WorkerJoinConfig) GetKubeProxyClientKey() string {
	return util.Deref(w.KubeProxyClientKey)
}
