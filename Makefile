LOCALBIN ?= $(shell pwd)/bin
CRDS_DIR ?= $(shell pwd)/api
CRDS_OUTPUT_DIR ?= $(shell pwd)/crds
CONTROLLER_TOOLS_VERSION ?= v0.17.2
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

### CRDs
.PHONY: crds
crds: generate manifests ## Build all targets.

.PHONY: manifests
manifests: controller-gen ## Generate CustomResourceDefinition objects.
	$(CONTROLLER_GEN) crd paths="$(CRDS_DIR)/..." output:crd:artifacts:config=$(CRDS_OUTPUT_DIR)

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object paths="$(CRDS_DIR)/..."

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.

$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef
