package main

import (
	"fmt"
	"time"

	"github.com/canonical/k8s-snap-api/v2/api"
	"gopkg.in/yaml.v2"
)

func main() {
	enabled := true

	status := api.ClusterStatus{
		Ready: true,
		Members: []api.NodeStatus{
			{
				Name:          "node-1",
				Address:       "10.0.0.1",
				ClusterRole:   api.ClusterRoleControlPlane,
				DatastoreRole: api.DatastoreRoleVoter,
			},
		},
		Config: api.UserFacingClusterConfig{
			Network: api.NetworkConfig{Enabled: &enabled},
			DNS:     api.DNSConfig{Enabled: &enabled},
		},
		Datastore: api.Datastore{
			Type:    "etcd",
			Servers: []string{"10.0.0.1:6443"},
		},
		DNS: api.FeatureStatus{
			Enabled:   true,
			Version:   "1.11.1",
			Message:   "enabled at 10.152.183.10",
			UpdatedAt: time.Now(),
		},
	}

	// b, _ := json.MarshalIndent(status, "", " ")
	// fmt.Println(string(b))
	b, _ := yaml.Marshal(status)
	fmt.Println(string(b))
}
