package apiv1

import "github.com/canonical/k8s-snap-api/internal/util"

type WorkerJoinConfig struct {
	// The certificate to be used for the kubelet.
	// If omitted defaults to an auto generated certificate.
	KubeletCert         *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	// The key to be used for the kubelet.
	// If omitted defaults to an auto generated key.
	KubeletKey          *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	// The client certificate to be used for the kubelet.
	// If omitted defaults to an auto generated certificate.
	KubeletClientCert   *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The client key to be used for the kubelet.
	// If omitted defaults to an auto generated key.
	KubeletClientKey    *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`
	// The client certificate to be used for the kube-proxy.
	// If omitted defaults to an auto generated certificate.
	KubeProxyClientCert *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used for the kube-proxy.
	// If omitted defaults to an auto generated key.
	KubeProxyClientKey  *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`

	// Additional files that are uploaded `/var/snap/k8s/common/args/conf.d/<filename>`
	// to a node on bootstrap. These files can then be referenced by Kubernetes
	// service arguments.
	//
	// The format is `map[<filename>]<filecontent>`.
	ExtraNodeConfigFiles map[string]string `json:"extra-node-config-files,omitempty" yaml:"extra-node-config-files,omitempty"`

	// Extra args to add to individual services (set any arg to null to delete)

	// Additional arguments that are passed to the `kube-proxy` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeProxyArgs         map[string]*string `json:"extra-node-kube-proxy-args,omitempty" yaml:"extra-node-kube-proxy-args,omitempty"`
	// Additional arguments that are passed to the `kubelet` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeletArgs           map[string]*string `json:"extra-node-kubelet-args,omitempty" yaml:"extra-node-kubelet-args,omitempty"`
	// Additional arguments that are passed to `containerd` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeContainerdArgs        map[string]*string `json:"extra-node-containerd-args,omitempty" yaml:"extra-node-containerd-args,omitempty"`
	// Additional arguments that are passed to `k8s-api-server-proxy` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeK8sAPIServerProxyArgs map[string]*string `json:"extra-node-k8s-apiserver-proxy-args,omitempty" yaml:"extra-node-k8s-apiserver-proxy-args,omitempty"`

	// Extra configuration for the containerd config.toml
	ExtraNodeContainerdConfig MapStringAny `json:"extra-node-containerd-config,omitempty" yaml:"extra-node-containerd-config,omitempty"`

	// The base directory in which the containerd-related files are located.
	ContainerdBaseDir string `json:"containerd-base-dir,omitempty" yaml:"containerd-base-dir,omitempty"`
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
