package apiv1

const (
	// K8sdAPIVersion is the path prefix that will be used for the k8sd endpoints for this api version.
	K8sdAPIVersion = "1.0"

	// AnnotationSkipCleanupKubernetesNodeOnRemove if set, only the microcluster & file cleanup is done.
	// This is useful, if an external controller (e.g. CAPI) is responsible for the Kubernetes node life cycle.
	// By default, the Kubernetes node is removed by k8sd if a node is removed from the cluster.
	AnnotationSkipCleanupKubernetesNodeOnRemove = "k8sd/v1alpha/lifecycle/skip-cleanup-kubernetes-node-on-remove"
)
