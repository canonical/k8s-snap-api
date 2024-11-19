package apiv1

import (
	"github.com/canonical/k8s-snap-api/internal/util"
)

type ControlPlaneJoinConfig struct {
	// List of extra SANs to be added to certificates.
	ExtraSANS []string `json:"extra-sans,omitempty" yaml:"extra-sans,omitempty"`

	// Seed certificates for external CA

	// The client certificate to be used for the front proxy.
	// If omitted defaults to an auto generated certificate.
	FrontProxyClientCert            *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	// The client key to be used for the front proxy.
	// If omitted defaults to an auto generated key.
	FrontProxyClientKey             *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	// The client certificate to be used by kubelet for communicating with the kube-apiserver.
	// If omitted defaults to an auto generated certificate.
	KubeProxyClientCert             *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used by kubelet for communicating with the kube-apiserver.
	// If omitted defaults to an auto generated key.
	KubeProxyClientKey              *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`
	// The client certificate to be used for the kube-scheduler.
	// If omitted defaults to an auto generated certificate.
	KubeSchedulerClientCert         *string `json:"kube-scheduler-client-crt,omitempty" yaml:"kube-scheduler-client-crt,omitempty"`
	// The client key to be used for the kube-scheduler.
	// If omitted defaults to an auto generated key.
	KubeSchedulerClientKey          *string `json:"kube-scheduler-client-key,omitempty" yaml:"kube-scheduler-client-key,omitempty"`
	// The client certificate to be used for the Kubernetes controller manager.
	// If omitted defaults to an auto generated certificate.
	KubeControllerManagerClientCert *string `json:"kube-controller-manager-client-crt,omitempty" yaml:"kube-controller-manager-client-crt,omitempty"`
	// The client key to be used for the Kubernetes controller manager.
	// If omitted defaults to an auto generated key.
	KubeControllerManagerClientKey  *string `json:"kube-controller-manager-client-key,omitempty" yaml:"kube-ControllerManager-client-key,omitempty"`

	// The certificate to be used for the kube-apiserver.
	// If omitted defaults to an auto generated certificate.
	APIServerCert     *string `json:"apiserver-crt,omitempty" yaml:"apiserver-crt,omitempty"`
	// The key to be used for the kube-apiserver.
	// If omitted defaults to an auto generated key.
	APIServerKey      *string `json:"apiserver-key,omitempty" yaml:"apiserver-key,omitempty"`
	// The certificate to be used for the kubelet.
	// If omitted defaults to an auto generated certificate.
	KubeletCert       *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	// The key to be used for the kubelet.
	// If omitted defaults to an auto generated key.
	KubeletKey        *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	// The client certificate to be used for the kubelet.
	// If omitted defaults to an auto generated certificate.
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The client key to be used for the kubelet.
	// If omitted defaults to an auto generated key.
	KubeletClientKey  *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`

	// Additional files that are uploaded `/var/snap/k8s/common/args/conf.d/<filename>`
	// to a node on bootstrap. These files can then be referenced by Kubernetes
	// service arguments.
	//
	// The format is `map[<filename>]<filecontent>`.
	ExtraNodeConfigFiles map[string]string `json:"extra-node-config-files,omitempty" yaml:"extra-node-config-files,omitempty"`

	// Extra args to add to individual services (set any arg to null to delete).

	// Additional arguments that are passed to the `kube-apiserver` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeAPIServerArgs         map[string]*string `json:"extra-node-kube-apiserver-args,omitempty" yaml:"extra-node-kube-apiserver-args,omitempty"`
	// Additional arguments that are passed to the `kube-controller-manager` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeControllerManagerArgs map[string]*string `json:"extra-node-kube-controller-manager-args,omitempty" yaml:"extra-node-kube-controller-manager-args,omitempty"`
	// Additional arguments that are passed to the `kube-scheduler` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeSchedulerArgs         map[string]*string `json:"extra-node-kube-scheduler-args,omitempty" yaml:"extra-node-kube-scheduler-args,omitempty"`
	// Additional arguments that are passed to the `kube-proxy` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeProxyArgs             map[string]*string `json:"extra-node-kube-proxy-args,omitempty" yaml:"extra-node-kube-proxy-args,omitempty"`
	// Additional arguments that are passed to the `kubelet` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeKubeletArgs               map[string]*string `json:"extra-node-kubelet-args,omitempty" yaml:"extra-node-kubelet-args,omitempty"`
	// Additional arguments that are passed to `containerd` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeContainerdArgs            map[string]*string `json:"extra-node-containerd-args,omitempty" yaml:"extra-node-containerd-args,omitempty"`
	// Additional arguments that are passed to `k8s-dqlite` only for that specific node.
	// A parameter that is explicitly set to `null` is deleted.
	// The format is `map[<--flag-name>]<value>`.
	ExtraNodeK8sDqliteArgs             map[string]*string `json:"extra-node-k8s-dqlite-args,omitempty" yaml:"extra-node-k8s-dqlite-args,omitempty"`

	// Extra configuration for the containerd config.toml
	ExtraNodeContainerdConfig MapStringAny `json:"extra-node-containerd-config,omitempty" yaml:"extra-node-containerd-config,omitempty"`

	// The base directory in which the containerd-related files are located.
	ContainerdBaseDir string `json:"containerd-base-dir,omitempty" yaml:"containerd-base-dir,omitempty"`
}

func (c *ControlPlaneJoinConfig) GetFrontProxyClientCert() string {
	return util.Deref(c.FrontProxyClientCert)
}
func (c *ControlPlaneJoinConfig) GetFrontProxyClientKey() string {
	return util.Deref(c.FrontProxyClientKey)
}
func (b *ControlPlaneJoinConfig) GetKubeProxyClientCert() string {
	return util.Deref(b.KubeProxyClientCert)
}
func (b *ControlPlaneJoinConfig) GetKubeProxyClientKey() string {
	return util.Deref(b.KubeProxyClientKey)
}
func (b *ControlPlaneJoinConfig) GetKubeSchedulerClientCert() string {
	return util.Deref(b.KubeSchedulerClientCert)
}
func (b *ControlPlaneJoinConfig) GetKubeSchedulerClientKey() string {
	return util.Deref(b.KubeSchedulerClientKey)
}
func (b *ControlPlaneJoinConfig) GetKubeControllerManagerClientCert() string {
	return util.Deref(b.KubeControllerManagerClientCert)
}
func (b *ControlPlaneJoinConfig) GetKubeControllerManagerClientKey() string {
	return util.Deref(b.KubeControllerManagerClientKey)
}
func (c *ControlPlaneJoinConfig) GetAPIServerCert() string { return util.Deref(c.APIServerCert) }
func (c *ControlPlaneJoinConfig) GetAPIServerKey() string  { return util.Deref(c.APIServerKey) }
func (c *ControlPlaneJoinConfig) GetKubeletCert() string   { return util.Deref(c.KubeletCert) }
func (c *ControlPlaneJoinConfig) GetKubeletKey() string    { return util.Deref(c.KubeletKey) }
func (c *ControlPlaneJoinConfig) GetKubeletClientCert() string {
	return util.Deref(c.KubeletClientCert)
}
func (c *ControlPlaneJoinConfig) GetKubeletClientKey() string {
	return util.Deref(c.KubeletClientKey)
}
