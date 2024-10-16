package cilium

const (
	// List of devices facing cluster/external network (used for BPF NodePort, BPF masquerading and host firewall); supports '+' as wildcard in device name, e.g. 'eth+'
	// e.g. k8sd/v1alpha1/cilium/devices="eth+ lxdbr+"
	AnnotationDevices = "k8sd/v1alpha1/cilium/devices"

	// Device name used to connect nodes in direct routing mode (used by BPF NodePort, BPF host routing; if empty, automatically set to a device with k8s InternalIP/ExternalIP or with a default route
	// bridge type devices are ingored in automatic selection
	AnnotationDirectRoutingDevice = "k8sd/v1alpha1/cilium/direct-routing-device"
)
