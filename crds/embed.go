package crds

import _ "embed"

//go:embed k8sd.io_upgrades.yaml
var UpgradesCRDYaml []byte
