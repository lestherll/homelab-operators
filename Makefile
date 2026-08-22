.PHONY: test run manifests generate

test:
	go test ./...

run:
	go run ./cmd

CONTROLLER_TOOLS_VERSION ?= v0.21.0
CONTROLLER_GEN = go run sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

manifests: ## Generate CRD YAML manifests.
	$(CONTROLLER_GEN) crd rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

generate: ## Generate deepcopy code (zz_generated.deepcopy.go).
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
