package api

// type feature_state indicates the state of a k8s cluster component.
type FeatureState string

const (
	FeatureStateEnabled  FeatureState = "enabled"
	FeatureStateDisabled FeatureState = "disabled"
	FeatureStateFailed   FeatureState = "failed"
	FeatureStateDegraded FeatureState = "degraded"
	FeatureStateWaiting  FeatureState = "waiting"
)
