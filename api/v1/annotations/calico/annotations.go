package calico

// For more details check the installation reference at https://docs.tigera.io/calico-enterprise/latest/reference/installation/api
const (
	// Enable the installation of the Calico API server to enable management of Calico APIs using kubectl.
	// e.g. k8sd/v1alpha1/calico/apiserver-enabled="true"
	AnnotationAPIServerEnabled = "k8sd/v1alpha1/calico/apiserver-enabled"
	// The type of encapsulation to use on the IPv4 pool.
	// One of: IPIP, VXLAN, IPIPCrossSubnet, VXLANCrossSubnet, None
	AnnotationEncapsulationV4 = "k8sd/v1alpha1/calico/encapsulation-v4"
	// The type of encapsulation to use on the IPv6 pool.
	// One of: IPIP, VXLAN, IPIPCrossSubnet, VXLANCrossSubnet, None
	AnnotationEncapsulationV6 = "k8sd/v1alpha1/calico/encapsulation-v6"
	// Use default interface matching parameters to select an interface,
	// performing best-effort filtering based on well-known interface names.
	AnnotationAutodetectionV4FirstFound = "k8sd/v1alpha1/calico/autodetection-v4/firstFound"
	// Configure Calico to detect node addresses based on the Kubernetes API.
	AnnotationAutodetectionV4Kubernetes = "k8sd/v1alpha1/calico/autodetection-v4/kubernetes"
	// Enable IP auto-detection based on interfaces that match the given regex.
	AnnotationAutodetectionV4Interface = "k8sd/v1alpha1/calico/autodetection-v4/interface"
	// Enable IP auto-detection based on interfaces that do not match the given regex.
	AnnotationAutodetectionV4SkipInterface = "k8sd/v1alpha1/calico/autodetection-v4/skipInterface"
	// Enable IP auto-detection based on which source address on the node is used to reach the specified IP or domain.
	AnnotationAutodetectionV4CanReach = "k8sd/v1alpha1/calico/autodetection-v4/canReach"
	// Enable IP auto-detection based on which addresses on the nodes are within one of the provided CIDRs.
	AnnotationAutodetectionV4CIDRs = "k8sd/v1alpha1/calico/autodetection-v4/cidrs"
	// Use default interface matching parameters to select an interface,
	// performing best-effort filtering based on well-known interface names.
	AnnotationAutodetectionV6FirstFound = "k8sd/v1alpha1/calico/autodetection-v6/firstFound"
	// Configures Calico to detect node addresses based on the Kubernetes API.
	AnnotationAutodetectionV6Kubernetes = "k8sd/v1alpha1/calico/autodetection-v6/kubernetes"
	// Enable IP auto-detection based on interfaces that match the given regex.
	AnnotationAutodetectionV6Interface = "k8sd/v1alpha1/calico/autodetection-v6/interface"
	// Enable IP auto-detection based on interfaces that do not match the given regex.
	AnnotationAutodetectionV6SkipInterface = "k8sd/v1alpha1/calico/autodetection-v6/skipInterface"
	// Enable IP auto-detection based on which source address on the node is used to reach the specified IP or domain.
	AnnotationAutodetectionV6CanReach = "k8sd/v1alpha1/calico/autodetection-v6/canReach"
	// Enable IP auto-detection based on which addresses on the nodes are within one of the provided CIDRs.
	AnnotationAutodetectionV6CIDRs = "k8sd/v1alpha1/calico/autodetection-v6/cidrs"
)
