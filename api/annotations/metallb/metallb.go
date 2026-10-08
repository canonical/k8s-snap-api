// Package metallb defines annotation keys for the MetalLB load-balancer feature.
// Keys under v1alpha1 are experimental and may change without notice.
package metallb

const (
	// AnnotationBGPPeers holds a YAML-encoded list of BGP peer configs for multi-peer
	// MetalLB BGP mode. Each item may include: peerAddress (required), peerASN (required),
	// myASN (optional, falls back to load-balancer.bgp-local-asn), peerPort (optional,
	// default 179), nodeSelector (optional, map[string]string of matchLabels), bfdProfile
	// (optional, name of a BFDProfile, e.g. one declared in AnnotationBFDProfiles;
	// requires AnnotationBGPBackend to be set to "frr-k8s").
	//
	// Example value:
	//   - peerAddress: 10.116.3.164
	//     peerASN: 65001
	//     nodeSelector: {topology.kubernetes.io/zone: i1}
	//     bfdProfile: fast-failover
	//
	// When set, this annotation REPLACES the single-peer typed BGP keys entirely.
	// Invalid YAML causes the load-balancer feature to enter a degraded state.
	AnnotationBGPPeers = "k8sd/v1alpha1/metallb/bgp-peers"

	// AnnotationAdvertiseAllPools controls the BGPAdvertisement spec. When "true",
	// the BGPAdvertisement is emitted with an empty spec (advertises all IP pools).
	// When unset or "false", only the named IPAddressPool is advertised.
	AnnotationAdvertiseAllPools = "k8sd/v1alpha1/metallb/advertise-all-pools"

	// AnnotationBGPBackend selects the BGP implementation used by the MetalLB speaker.
	// Supported values:
	//   - "native" (default when unset): the built-in Go (GoBGP) implementation. Does
	//     not support BFD.
	//   - "frr-k8s": the FRR-based backend (github.com/metallb/frr-k8s). Required for
	//     any peer that sets a bfdProfile in AnnotationBGPPeers.
	// Any other value causes the load-balancer feature to enter a degraded state.
	AnnotationBGPBackend = "k8sd/v1alpha1/metallb/bgp-backend"

	// AnnotationBFDProfiles holds a YAML-encoded list of MetalLB BFDProfile configs
	// that k8sd keeps present in the MetalLB namespace while the load-balancer feature
	// is enabled in BGP mode. Each item may include: name (required), receiveInterval
	// (optional, ms, [10, 60000]), transmitInterval (optional, ms, [10, 60000]),
	// detectMultiplier (optional, [2, 255]), echoInterval (optional, ms, [10, 60000]),
	// echoMode (optional, bool), passiveMode (optional, bool), minimumTtl (optional,
	// multi-hop only, [1, 254]). Unset fields fall back to MetalLB defaults. Requires
	// AnnotationBGPBackend to be set to "frr-k8s".
	//
	// Example value:
	//   - name: fast-failover
	//     receiveInterval: 150
	//     transmitInterval: 150
	//     detectMultiplier: 3
	//     echoInterval: 50
	//     echoMode: false
	//     passiveMode: false
	//     minimumTtl: 254
	//
	// Invalid YAML causes the load-balancer feature to enter a degraded state.
	AnnotationBFDProfiles = "k8sd/v1alpha1/metallb/bfd-profiles"
)
