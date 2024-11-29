package apiv1

import "github.com/canonical/k8s-snap-api/internal/util"

// BootstrapConfig is used to seed cluster configuration when bootstrapping a new cluster.
type BootstrapConfig struct {
	// ClusterConfig

	ClusterConfig UserFacingClusterConfig `json:"cluster-config,omitempty" yaml:"cluster-config,omitempty"`

	// Seed configuration for the control plane (flat on purpose). Empty values are ignored

	// List of taints to be applied to control plane nodes.
	ControlPlaneTaints  []string `json:"control-plane-taints,omitempty" yaml:"control-plane-taints,omitempty"`
	// The CIDR to be used for assigning pod addresses.
	// If omitted defaults to `10.1.0.0/16`.
	PodCIDR             *string  `json:"pod-cidr,omitempty" yaml:"pod-cidr,omitempty"`
	// The CIDR to be used for assigning service addresses.
	// If omitted defaults to `10.152.183.0/24`.
	ServiceCIDR         *string  `json:"service-cidr,omitempty" yaml:"service-cidr,omitempty"`
	// Determines if RBAC should be disabled.
	// If omitted defaults to `false`.
	DisableRBAC         *bool    `json:"disable-rbac,omitempty" yaml:"disable-rbac,omitempty"`
	// The port number for kube-apiserver to use.
	// If omitted defaults to `6443`.
	SecurePort          *int     `json:"secure-port,omitempty" yaml:"secure-port,omitempty"`
	// The port number for k8s-dqlite to use.
	// If omitted defaults to `9000`.
	K8sDqlitePort       *int     `json:"k8s-dqlite-port,omitempty" yaml:"k8s-dqlite-port,omitempty"`
	// The type of datastore to be used.
	// If omitted defaults to `k8s-dqlite`.
	//
	// Can be used to point to an external datastore like etcd.
	//
	// Possible Values: `k8s-dqlite | external`.
	DatastoreType       *string  `json:"datastore-type,omitempty" yaml:"datastore-type,omitempty"`
	// The server addresses to be used when `datastore-type` is set to `external`.
	DatastoreServers    []string `json:"datastore-servers,omitempty" yaml:"datastore-servers,omitempty"`
	// The CA certificate to be used when communicating with the external datastore.
	DatastoreCACert     *string  `json:"datastore-ca-crt,omitempty" yaml:"datastore-ca-crt,omitempty"`
	// The client certificate to be used when communicating with the external
	// datastore.
	DatastoreClientCert *string  `json:"datastore-client-crt,omitempty" yaml:"datastore-client-crt,omitempty"`
	// The client key to be used when communicating with the external datastore.
	DatastoreClientKey  *string  `json:"datastore-client-key,omitempty" yaml:"datastore-client-key,omitempty"`

	// Seed configuration for certificates

	// List of extra SANs to be added to certificates.
	ExtraSANs []string `json:"extra-sans,omitempty" yaml:"extra-sans,omitempty"`

	// Seed configuration for external certificates (cluster-wide)

	// The CA certificate to be used for Kubernetes services.
	// If omitted defaults to an auto generated certificate.
	CACert                          *string `json:"ca-crt,omitempty" yaml:"ca-crt,omitempty"`
	// The CA key to be used for Kubernetes services.
	// If omitted defaults to an auto generated key.
	CAKey                           *string `json:"ca-key,omitempty" yaml:"ca-key,omitempty"`
	// The client CA certificate to be used for Kubernetes services.
	// If omitted defaults to an auto generated certificate.
	ClientCACert                    *string `json:"client-ca-crt,omitempty" yaml:"client-ca-crt,omitempty"`
	// The client CA key to be used for Kubernetes services.
	// If omitted defaults to an auto generated key.
	ClientCAKey                     *string `json:"client-ca-key,omitempty" yaml:"client-ca-key,omitempty"`
	// The CA certificate to be used for the front proxy.
	// If omitted defaults to an auto generated certificate.
	FrontProxyCACert                *string `json:"front-proxy-ca-crt,omitempty" yaml:"front-proxy-ca-crt,omitempty"`
	// The CA key to be used for the front proxy.
	// If omitted defaults to an auto generated key.
	FrontProxyCAKey                 *string `json:"front-proxy-ca-key,omitempty" yaml:"front-proxy-ca-key,omitempty"`
	// The client certificate to be used for the front proxy.
	// If omitted defaults to an auto generated certificate.
	FrontProxyClientCert            *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	// The client key to be used for the front proxy.
	// If omitted defaults to an auto generated key.
	FrontProxyClientKey             *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	// The client certificate to be used by kubelet for communicating with the kube-apiserver.
	// If omitted defaults to an auto generated certificate.
	APIServerKubeletClientCert      *string `json:"apiserver-kubelet-client-crt,omitempty" yaml:"apiserver-kubelet-client-crt,omitempty"`
	// The client key to be used by kubelet for communicating with the kube-apiserver.
	// If omitted defaults to an auto generated key.
	APIServerKubeletClientKey       *string `json:"apiserver-kubelet-client-key,omitempty" yaml:"apiserver-kubelet-client-key,omitempty"`

	// The admin client certificate to be used for Kubernetes services.
	// If omitted defaults to an auto generated certificate.
	AdminClientCert                 *string `json:"admin-client-crt,omitempty" yaml:"admin-client-crt,omitempty"`
	// The admin client key to be used for Kubernetes services.
	// If omitted defaults to an auto generated key.
	AdminClientKey                  *string `json:"admin-client-key,omitempty" yaml:"admin-client-key,omitempty"`
	// The client certificate to be used for the kube-proxy.
	// If omitted defaults to an auto generated certificate.
	KubeProxyClientCert             *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used for the kube-proxy.
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
	// The key to be used by the default service account.
	// If omitted defaults to an auto generated key.
	ServiceAccountKey               *string `json:"service-account-key,omitempty" yaml:"service-account-key,omitempty"`

	// Seed configuration for external certificates (node-specific)

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
	// The certificate to be used for the kubelet client.
	// If omitted defaults to an auto generated certificate.
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The key to be used for the kubelet client.
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

func (b *BootstrapConfig) GetDatastoreType() string        { return util.Deref(b.DatastoreType) }
func (b *BootstrapConfig) GetDatastoreCACert() string      { return util.Deref(b.DatastoreCACert) }
func (b *BootstrapConfig) GetDatastoreClientCert() string  { return util.Deref(b.DatastoreClientCert) }
func (b *BootstrapConfig) GetDatastoreClientKey() string   { return util.Deref(b.DatastoreClientKey) }
func (b *BootstrapConfig) GetK8sDqlitePort() int           { return util.Deref(b.K8sDqlitePort) }
func (b *BootstrapConfig) GetCACert() string               { return util.Deref(b.CACert) }
func (b *BootstrapConfig) GetCAKey() string                { return util.Deref(b.CAKey) }
func (b *BootstrapConfig) GetClientCACert() string         { return util.Deref(b.ClientCACert) }
func (b *BootstrapConfig) GetClientCAKey() string          { return util.Deref(b.ClientCAKey) }
func (b *BootstrapConfig) GetFrontProxyCACert() string     { return util.Deref(b.FrontProxyCACert) }
func (b *BootstrapConfig) GetFrontProxyCAKey() string      { return util.Deref(b.FrontProxyCAKey) }
func (b *BootstrapConfig) GetFrontProxyClientCert() string { return util.Deref(b.FrontProxyClientCert) }
func (b *BootstrapConfig) GetFrontProxyClientKey() string  { return util.Deref(b.FrontProxyClientKey) }
func (b *BootstrapConfig) GetAPIServerKubeletClientCert() string {
	return util.Deref(b.APIServerKubeletClientCert)
}
func (b *BootstrapConfig) GetAPIServerKubeletClientKey() string {
	return util.Deref(b.APIServerKubeletClientKey)
}
func (b *BootstrapConfig) GetAdminClientCert() string     { return util.Deref(b.AdminClientCert) }
func (b *BootstrapConfig) GetAdminClientKey() string      { return util.Deref(b.AdminClientKey) }
func (b *BootstrapConfig) GetKubeProxyClientCert() string { return util.Deref(b.KubeProxyClientCert) }
func (b *BootstrapConfig) GetKubeProxyClientKey() string  { return util.Deref(b.KubeProxyClientKey) }
func (b *BootstrapConfig) GetKubeSchedulerClientCert() string {
	return util.Deref(b.KubeSchedulerClientCert)
}
func (b *BootstrapConfig) GetKubeSchedulerClientKey() string {
	return util.Deref(b.KubeSchedulerClientKey)
}
func (b *BootstrapConfig) GetKubeControllerManagerClientCert() string {
	return util.Deref(b.KubeControllerManagerClientCert)
}
func (b *BootstrapConfig) GetKubeControllerManagerClientKey() string {
	return util.Deref(b.KubeControllerManagerClientKey)
}
func (b *BootstrapConfig) GetServiceAccountKey() string { return util.Deref(b.ServiceAccountKey) }
func (b *BootstrapConfig) GetAPIServerCert() string     { return util.Deref(b.APIServerCert) }
func (b *BootstrapConfig) GetAPIServerKey() string      { return util.Deref(b.APIServerKey) }
func (b *BootstrapConfig) GetKubeletCert() string       { return util.Deref(b.KubeletCert) }
func (b *BootstrapConfig) GetKubeletKey() string        { return util.Deref(b.KubeletKey) }
func (b *BootstrapConfig) GetKubeletClientCert() string { return util.Deref(b.KubeletClientCert) }
func (b *BootstrapConfig) GetKubeletClientKey() string  { return util.Deref(b.KubeletClientKey) }
