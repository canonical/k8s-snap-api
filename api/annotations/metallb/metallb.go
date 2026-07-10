// Package metallb defines annotation keys for the MetalLB load-balancer feature.
// Keys under v1alpha1 are experimental and may change without notice.
package metallb

const (
	// AnnotationBGPPeers holds a YAML-encoded list of BGP peer configs for multi-peer
	// MetalLB BGP mode. Each item may include: peerAddress (required), peerASN (required),
	// myASN (optional, falls back to load-balancer.bgp-local-asn), peerPort (optional,
	// default 179), nodeSelector (optional, map[string]string of matchLabels).
	//
	// Example value:
	//   - peerAddress: 10.116.3.164
	//     peerASN: 65001
	//     nodeSelector: {topology.kubernetes.io/zone: i1}
	//
	// When set, this annotation REPLACES the single-peer typed BGP keys entirely.
	// Invalid YAML causes the load-balancer feature to enter a degraded state.
	AnnotationBGPPeers = "k8sd/v1alpha1/metallb/bgp-peers"

	// AnnotationAdvertiseAllPools controls the BGPAdvertisement spec. When "true",
	// the BGPAdvertisement is emitted with an empty spec (advertises all IP pools).
	// When unset or "false", only the named IPAddressPool is advertised.
	AnnotationAdvertiseAllPools = "k8sd/v1alpha1/metallb/advertise-all-pools"
)
