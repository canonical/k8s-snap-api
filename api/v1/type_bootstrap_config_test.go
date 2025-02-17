package apiv1_test

import (
	"testing"

	apiv1 "github.com/canonical/k8s-snap-api/api/v1"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

const (
	oldFieldYaml   = `kube-ControllerManager-client-key: "foo"`
	newFieldYaml   = `kube-controller-manager-client-key: "lish"`
	bothFieldsYaml = `
kube-ControllerManager-client-key: "foo"
kube-controller-manager-client-key: "lish"
`
)

func TestKubeControllerManagerClientKeyUnmarshal(t *testing.T) {
	for _, tc := range []struct {
		name          string
		yaml          string
		expectedValue string
	}{
		{
			name:          "kube-ControllerManager-client-key",
			yaml:          oldFieldYaml,
			expectedValue: "foo",
		},
		{
			name:          "kube-controller-manager-client-key",
			yaml:          newFieldYaml,
			expectedValue: "lish",
		},
		{
			name:          "both keys",
			yaml:          bothFieldsYaml,
			expectedValue: "foo",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			config := apiv1.BootstrapConfig{}

			err := yaml.UnmarshalStrict([]byte(tc.yaml), &config)

			g.Expect(err).To(Not(HaveOccurred()))
			g.Expect(*config.KubeControllerManagerClientKey).To(Equal(tc.expectedValue))
		})
	}
}
