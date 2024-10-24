package metrics_server

const (
	// Override the default image repository for the metrics-server.
	AnnotationImageRepo = "k8sd/v1alpha1/metrics-server/image-repo"
	// Override the default image tag for the metrics-server.
	AnnotationImageTag = "k8sd/v1alpha1/metrics-server/image-tag"
)
