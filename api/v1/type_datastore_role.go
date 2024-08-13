package apiv1

// DatastoreRole as provided by dqlite
type DatastoreRole string

const (
	DatastoreRoleVoter   DatastoreRole = "voter"
	DatastoreRoleStandBy DatastoreRole = "stand-by"
	DatastoreRoleSpare   DatastoreRole = "spare"
	DatastoreRolePending DatastoreRole = "PENDING"
	DatastoreRoleUnknown DatastoreRole = "unknown"
)
