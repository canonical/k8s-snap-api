package api

type ClusterHealth string

const (
	ClusterHealthReady    ClusterHealth = "ready"
	ClusterHealthFailed   ClusterHealth = "failed"
	ClusterHealthDegraded ClusterHealth = "degraded"
)
