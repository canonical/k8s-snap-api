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

	// AnnotationDisableSeparateFeatureUpgrades if set, the separate feature upgrade is disabled.
	// This is useful, if an external controller (e.g. CAPI) is responsible for the Kubernetes node life cycle.
	// By default, the feature upgrade will be done after all nodes in a cluster are upgraded.
	AnnotationDisableSeparateFeatureUpgrades = "k8sd/v1alpha/lifecycle/disable-separate-feature-upgrades"

	// AnnotationServicesStopDelay defines a delay in seconds before stopping services on a leaving node.
	// This can be useful to allow time for services to take care of draining, etc before they stop.
	// By default, there is no delay and services are stopped immediately (unless AnnotationSkipStopServicesOnRemove is set).
	AnnotationServicesStopDelay = "k8sd/v1alpha/lifecycle/services-stop-delay"
)
