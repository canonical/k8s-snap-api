package apiv1

import (
	"fmt"

	"github.com/canonical/k8s-snap-api-v1/internal/util"
	"gopkg.in/yaml.v2"
)

type ControlPlaneNodeJoinConfig struct {
	ExtraSANS []string `json:"extra-sans,omitempty" yaml:"extra-sans,omitempty"`

	// Seed certificates for external CA
	FrontProxyClientCert            *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	FrontProxyClientKey             *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	KubeProxyClientCert             *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	KubeProxyClientKey              *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`
	KubeSchedulerClientCert         *string `json:"kube-scheduler-client-crt,omitempty" yaml:"kube-scheduler-client-crt,omitempty"`
	KubeSchedulerClientKey          *string `json:"kube-scheduler-client-key,omitempty" yaml:"kube-scheduler-client-key,omitempty"`
	KubeControllerManagerClientCert *string `json:"kube-controller-manager-client-crt,omitempty" yaml:"kube-controller-manager-client-crt,omitempty"`
	KubeControllerManagerClientKey  *string `json:"kube-controller-manager-client-key,omitempty" yaml:"kube-ControllerManager-client-key,omitempty"`

	APIServerCert     *string `json:"apiserver-crt,omitempty" yaml:"apiserver-crt,omitempty"`
	APIServerKey      *string `json:"apiserver-key,omitempty" yaml:"apiserver-key,omitempty"`
	KubeletCert       *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	KubeletKey        *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	KubeletClientKey  *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`

	// ExtraNodeConfigFiles will be written to /var/snap/k8s/common/args/conf.d
	ExtraNodeConfigFiles map[string]string `json:"extra-node-config-files,omitempty" yaml:"extra-node-config-files,omitempty"`

	// Extra args to add to individual services (set any arg to null to delete)
	ExtraNodeKubeAPIServerArgs         map[string]*string `json:"extra-node-kube-apiserver-args,omitempty" yaml:"extra-node-kube-apiserver-args,omitempty"`
	ExtraNodeKubeControllerManagerArgs map[string]*string `json:"extra-node-kube-controller-manager-args,omitempty" yaml:"extra-node-kube-controller-manager-args,omitempty"`
	ExtraNodeKubeSchedulerArgs         map[string]*string `json:"extra-node-kube-scheduler-args,omitempty" yaml:"extra-node-kube-scheduler-args,omitempty"`
	ExtraNodeKubeProxyArgs             map[string]*string `json:"extra-node-kube-proxy-args,omitempty" yaml:"extra-node-kube-proxy-args,omitempty"`
	ExtraNodeKubeletArgs               map[string]*string `json:"extra-node-kubelet-args,omitempty" yaml:"extra-node-kubelet-args,omitempty"`
	ExtraNodeContainerdArgs            map[string]*string `json:"extra-node-containerd-args,omitempty" yaml:"extra-node-containerd-args,omitempty"`
	ExtraNodeK8sDqliteArgs             map[string]*string `json:"extra-node-k8s-dqlite-args,omitempty" yaml:"extra-node-k8s-dqlite-args,omitempty"`
}

func (c *ControlPlaneNodeJoinConfig) GetFrontProxyClientCert() string {
	return util.Deref(c.FrontProxyClientCert)
}
func (c *ControlPlaneNodeJoinConfig) GetFrontProxyClientKey() string {
	return util.Deref(c.FrontProxyClientKey)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeProxyClientCert() string {
	return util.Deref(b.KubeProxyClientCert)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeProxyClientKey() string {
	return util.Deref(b.KubeProxyClientKey)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeSchedulerClientCert() string {
	return util.Deref(b.KubeSchedulerClientCert)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeSchedulerClientKey() string {
	return util.Deref(b.KubeSchedulerClientKey)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeControllerManagerClientCert() string {
	return util.Deref(b.KubeControllerManagerClientCert)
}
func (b *ControlPlaneNodeJoinConfig) GetKubeControllerManagerClientKey() string {
	return util.Deref(b.KubeControllerManagerClientKey)
}
func (c *ControlPlaneNodeJoinConfig) GetAPIServerCert() string { return util.Deref(c.APIServerCert) }
func (c *ControlPlaneNodeJoinConfig) GetAPIServerKey() string  { return util.Deref(c.APIServerKey) }
func (c *ControlPlaneNodeJoinConfig) GetKubeletCert() string   { return util.Deref(c.KubeletCert) }
func (c *ControlPlaneNodeJoinConfig) GetKubeletKey() string    { return util.Deref(c.KubeletKey) }
func (c *ControlPlaneNodeJoinConfig) GetKubeletClientCert() string {
	return util.Deref(c.KubeletClientCert)
}
func (c *ControlPlaneNodeJoinConfig) GetKubeletClientKey() string {
	return util.Deref(c.KubeletClientKey)
}

// WorkerJoinConfigFromMicrocluster parses a microcluster map[string]string and retrieves the WorkerNodeJoinConfig.
func ControlPlaneJoinConfigFromMicrocluster(m map[string]string) (ControlPlaneNodeJoinConfig, error) {
	config := ControlPlaneNodeJoinConfig{}
	if err := yaml.UnmarshalStrict([]byte(m["controlPlaneJoinConfig"]), &config); err != nil {
		return ControlPlaneNodeJoinConfig{}, fmt.Errorf("failed to unmarshal control plane join config: %w", err)
	}
	return config, nil
}
