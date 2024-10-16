package annotations

const (
	// AnnotationSkipCleanupKubernetesNodeOnRemove if set, only the microcluster & file cleanup is done.
	// This is useful, if an external controller (e.g. CAPI) is responsible for the Kubernetes node life cycle.
	// By default, the Kubernetes node is removed by k8sd if a node is removed from the cluster.
	AnnotationSkipCleanupKubernetesNodeOnRemove = "k8sd/v1alpha/lifecycle/skip-cleanup-kubernetes-node-on-remove"

	// AnnotationSkipStopServicesOnRemove if set, the k8s services will not be stopped on the leaving node when removing the node.
	// This is useful, if an external controller (e.g. CAPI) is responsible for the Kubernetes node life cycle.
	// By default, all services are stopped on leaving nodes.
	AnnotationSkipStopServicesOnRemove = "k8sd/v1alpha/lifecycle/skip-stop-services-on-remove"
)
