package calico

const (
	AnnotationAPIServerEnabled             = "k8sd/v1alpha1/calico/apiserver-enabled"
	AnnotationEncapsulationV4              = "k8sd/v1alpha1/calico/encapsulation-v4"
	AnnotationEncapsulationV6              = "k8sd/v1alpha1/calico/encapsulation-v6"
	AnnotationAutodetectionV4FirstFound    = "k8sd/v1alpha1/calico/autodetection-v4/firstFound"
	AnnotationAutodetectionV4Kubernetes    = "k8sd/v1alpha1/calico/autodetection-v4/kubernetes"
	AnnotationAutodetectionV4Interface     = "k8sd/v1alpha1/calico/autodetection-v4/interface"
	AnnotationAutodetectionV4SkipInterface = "k8sd/v1alpha1/calico/autodetection-v4/skipInterface"
	AnnotationAutodetectionV4CanReach      = "k8sd/v1alpha1/calico/autodetection-v4/canReach"
	AnnotationAutodetectionV4CIDRs         = "k8sd/v1alpha1/calico/autodetection-v4/cidrs"
	AnnotationAutodetectionV6FirstFound    = "k8sd/v1alpha1/calico/autodetection-v6/firstFound"
	AnnotationAutodetectionV6Kubernetes    = "k8sd/v1alpha1/calico/autodetection-v6/kubernetes"
	AnnotationAutodetectionV6Interface     = "k8sd/v1alpha1/calico/autodetection-v6/interface"
	AnnotationAutodetectionV6SkipInterface = "k8sd/v1alpha1/calico/autodetection-v6/skipInterface"
	AnnotationAutodetectionV6CanReach      = "k8sd/v1alpha1/calico/autodetection-v6/canReach"
	AnnotationAutodetectionV6CIDRs         = "k8sd/v1alpha1/calico/autodetection-v6/cidrs"
)
