package cilium

const (
	// Make Cilium take ownership over the `/etc/cni/net.d` directory on the
	// node, renaming all non-Cilium CNI configurations to `*.cilium_bak`.
	// This ensures no Pods can be scheduled using other CNI plugins during Cilium
	// agent downtime.
	// Leave this unset if you wish to use other CNIs such as Multus.
	AnnotationCNIExclusive = "k8sd/v1alpha1/cilium/cni-exclusive"

	// List of devices facing cluster/external network (used for BPF
	// NodePort, BPF masquerading and host firewall); supports '+' as
	// wildcard in device name, e.g. 'eth+'
	// e.g. k8sd/v1alpha1/cilium/devices="eth+ lxdbr+"
	AnnotationDevices = "k8sd/v1alpha1/cilium/devices"

	// Device name used to connect nodes in direct routing mode (used by
	// BPF NodePort, BPF host routing; if empty, automatically set to a
	// device with k8s InternalIP/ExternalIP or with a default route
	// bridge type devices are ignored in automatic selection
	AnnotationDirectRoutingDevice = "k8sd/v1alpha1/cilium/direct-routing-device"

	// Comma separated list of VLAN tags to bypass eBPF filtering on native
	// devices. Cilium enables firewalling on native devices and filters
	// all unknown traffic, including VLAN 802.1q packets, which pass
	// through the main device with the associated tag (e.g., VLAN device
	// eth0.4000 and its main interface eth0). Supports '0' as wildcard for
	// bypassing all VLANs.
	// e.g., k8sd/v1alpha1/cilium/vlan-bpf-bypass="4001,4002"
	AnnotationVLANBPFBypass = "k8sd/v1alpha1/cilium/vlan-bpf-bypass"
)
