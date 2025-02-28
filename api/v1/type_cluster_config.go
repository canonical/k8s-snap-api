package apiv1

import (
	"fmt"

	"github.com/canonical/k8s-snap-api/internal/util"
	"gopkg.in/yaml.v2"
)

type UserFacingClusterConfig struct {
	// Configuration options for the network feature.
	Network NetworkConfig `json:"network,omitempty" yaml:"network,omitempty"`
	// Configuration options for the dns feature.
	DNS DNSConfig `json:"dns,omitempty" yaml:"dns,omitempty"`
	// Configuration options for the ingress feature.
	Ingress IngressConfig `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	// Configuration options for the load-balancer feature.
	LoadBalancer LoadBalancerConfig `json:"load-balancer,omitempty" yaml:"load-balancer,omitempty"`
	// Configuration options for the local-storage feature.
	LocalStorage LocalStorageConfig `json:"local-storage,omitempty" yaml:"local-storage,omitempty"`
	// Configuration options for the gateway feature.
	Gateway GatewayConfig `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	// Configuration options for the metric server feature.
	MetricsServer MetricsServerConfig `json:"metrics-server,omitempty" yaml:"metrics-server,omitempty"`
	// Sets the cloud provider to be used by the cluster.
	//
	// When this is set as `external`, node will wait for an external cloud provider to
	// do cloud specific setup and finish node initialisation.
	//
	// Possible values: `external`.
	CloudProvider *string `json:"cloud-provider,omitempty" yaml:"cloud-provider,omitempty"`
	// Annotations is a map of strings that can be used to store arbitrary metadata configuration.
	// Please refer to the annotations reference for further details on these options.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

type DNSConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `true`
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// Sets the local domain of the cluster.
	// If omitted defaults to `cluster.local`.
	ClusterDomain *string `json:"cluster-domain,omitempty" yaml:"cluster-domain,omitempty"`
	// Sets the IP address of the dns service. If omitted defaults to the IP address
	// of the Kubernetes service created by the feature.
	//
	// Can be used to point to an external dns server when feature is disabled.
	ServiceIP *string `json:"service-ip,omitempty" yaml:"service-ip,omitempty"`
	// Sets the upstream nameservers used to forward queries for out-of-cluster
	// endpoints.
	//
	// If omitted defaults to `/etc/resolv.conf` and uses the nameservers of the node.
	UpstreamNameservers *[]string `json:"upstream-nameservers,omitempty" yaml:"upstream-nameservers,omitempty"`
}

func (c DNSConfig) GetEnabled() bool                 { return util.Deref(c.Enabled) }
func (c DNSConfig) GetClusterDomain() string         { return util.Deref(c.ClusterDomain) }
func (c DNSConfig) GetServiceIP() string             { return util.Deref(c.ServiceIP) }
func (c DNSConfig) GetUpstreamNameservers() []string { return util.Deref(c.UpstreamNameservers) }

type IngressConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `false`
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// Sets the name of the secret to be used for providing default encryption to
	// ingresses.
	//
	// Ingresses can specify another TLS secret in their resource definitions,
	// in which case the default secret won't be used.
	DefaultTLSSecret *string `json:"default-tls-secret,omitempty" yaml:"default-tls-secret,omitempty"`
	// Determines if the proxy protocol should be enabled for ingresses.
	// If omitted defaults to `false`.
	EnableProxyProtocol *bool `json:"enable-proxy-protocol,omitempty" yaml:"enable-proxy-protocol,omitempty"`
}

func (c IngressConfig) GetEnabled() bool             { return util.Deref(c.Enabled) }
func (c IngressConfig) GetDefaultTLSSecret() string  { return util.Deref(c.DefaultTLSSecret) }
func (c IngressConfig) GetEnableProxyProtocol() bool { return util.Deref(c.EnableProxyProtocol) }

type LoadBalancerConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// Sets the CIDRs used for assigning IP addresses to Kubernetes services with type
	// `LoadBalancer`.
	CIDRs *[]string `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`
	// Determines if L2 mode should be enabled.
	// If omitted defaults to `true`.
	L2Mode *bool `json:"l2-mode,omitempty" yaml:"l2-mode,omitempty"`
	// Sets the interfaces to be used for announcing IP addresses through ARP.
	// If omitted all interfaces will be used.
	L2Interfaces *[]string `json:"l2-interfaces,omitempty" yaml:"l2-interfaces,omitempty"`
	// Determines if BGP mode should be enabled.
	// If omitted defaults to `false`.
	BGPMode *bool `json:"bgp-mode,omitempty" yaml:"bgp-mode,omitempty"`
	// Sets the ASN to be used for the local virtual BGP router.
	// Required if bgp-mode is true.
	BGPLocalASN *int `json:"bgp-local-asn,omitempty" yaml:"bgp-local-asn,omitempty"`
	// Sets the IP address of the BGP peer.
	// Required if bgp-mode is true.
	BGPPeerAddress *string `json:"bgp-peer-address,omitempty" yaml:"bgp-peer-address,omitempty"`
	// Sets the ASN of the BGP peer.
	// Required if bgp-mode is true.
	BGPPeerASN *int `json:"bgp-peer-asn,omitempty" yaml:"bgp-peer-asn,omitempty"`
	// Sets the port of the BGP peer.
	// Required if bgp-mode is true.
	BGPPeerPort *int `json:"bgp-peer-port,omitempty" yaml:"bgp-peer-port,omitempty"`
}

func (c LoadBalancerConfig) GetEnabled() bool          { return util.Deref(c.Enabled) }
func (c LoadBalancerConfig) GetCIDRs() []string        { return util.Deref(c.CIDRs) }
func (c LoadBalancerConfig) GetL2Mode() bool           { return util.Deref(c.L2Mode) }
func (c LoadBalancerConfig) GetL2Interfaces() []string { return util.Deref(c.L2Interfaces) }
func (c LoadBalancerConfig) GetBGPMode() bool          { return util.Deref(c.BGPMode) }
func (c LoadBalancerConfig) GetBGPLocalASN() int       { return util.Deref(c.BGPLocalASN) }
func (c LoadBalancerConfig) GetBGPPeerAddress() string { return util.Deref(c.BGPPeerAddress) }
func (c LoadBalancerConfig) GetBGPPeerASN() int        { return util.Deref(c.BGPPeerASN) }
func (c LoadBalancerConfig) GetBGPPeerPort() int       { return util.Deref(c.BGPPeerPort) }

type LocalStorageConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// Sets the path to be used for storing volume data.
	// If omitted defaults to `/var/snap/k8s/common/rawfile-storage`
	LocalPath *string `json:"local-path,omitempty" yaml:"local-path,omitempty"`
	// Sets the reclaim policy of the storage class.
	// If omitted defaults to `Delete`.
	// Possible values: `Retain | Recycle | Delete`
	ReclaimPolicy *string `json:"reclaim-policy,omitempty" yaml:"reclaim-policy,omitempty"`
	// Determines if the storage class should be set as default.
	// If omitted defaults to `true`
	Default *bool `json:"default,omitempty" yaml:"default,omitempty"`
}

func (c LocalStorageConfig) GetEnabled() bool         { return util.Deref(c.Enabled) }
func (c LocalStorageConfig) GetLocalPath() string     { return util.Deref(c.LocalPath) }
func (c LocalStorageConfig) GetReclaimPolicy() string { return util.Deref(c.ReclaimPolicy) }
func (c LocalStorageConfig) GetDefault() bool         { return util.Deref(c.Default) }

type NetworkConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `true`
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// PodCIDR is the CIDR range for the pods in the cluster.
	PodCIDR *string `json:"pod-cidr,omitempty" yaml:"pod-cidr,omitempty"`
	// ServiceCIDR is the CIDR range for the services in the cluster.
	ServiceCIDR *string `json:"service-cidr,omitempty" yaml:"service-cidr,omitempty"`
}

func (c NetworkConfig) GetEnabled() bool { return util.Deref(c.Enabled) }

type GatewayConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

func (c GatewayConfig) GetEnabled() bool { return util.Deref(c.Enabled) }

type MetricsServerConfig struct {
	// Determines if the feature should be enabled.
	// If omitted defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

func (c MetricsServerConfig) GetEnabled() bool { return util.Deref(c.Enabled) }

type UserFacingDatastoreConfig struct {
	// Type of the datastore.
	Type *string `json:"type,omitempty" yaml:"type,omitempty"`
	// Datastore server addresses.
	Servers *[]string `json:"servers,omitempty" yaml:"servers,omitempty"`
	// Datastore CA certificate.
	CACert *string `json:"ca-crt,omitempty" yaml:"ca-crt,omitempty"`
	// Datastore client certificate.
	ClientCert *string `json:"client-crt,omitempty" yaml:"client-crt,omitempty"`
	// Datastore client key.
	ClientKey *string `json:"client-key,omitempty" yaml:"client-key,omitempty"`
}

func (c UserFacingDatastoreConfig) GetType() string       { return util.Deref(c.Type) }
func (c UserFacingDatastoreConfig) GetServers() []string  { return util.Deref(c.Servers) }
func (c UserFacingDatastoreConfig) GetCACert() string     { return util.Deref(c.CACert) }
func (c UserFacingDatastoreConfig) GetClientCert() string { return util.Deref(c.ClientCert) }
func (c UserFacingDatastoreConfig) GetClientKey() string  { return util.Deref(c.ClientKey) }

func (c UserFacingClusterConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c NetworkConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c DNSConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c IngressConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c LoadBalancerConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c LocalStorageConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c GatewayConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}

func (c MetricsServerConfig) String() string {
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("%#v\n", c)
	}
	return string(b)
}
