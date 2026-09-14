# ====================================================================================
# Setup Project

PROJECT_NAME ?= provider-databricks
PROJECT_REPO ?= github.com/buzz-IT/$(PROJECT_NAME)

export TERRAFORM_VERSION ?= 1.5.7

# Do not allow a version of terraform greater than 1.5.x, due to versions 1.6+ being
# licensed under BSL, which is not permitted.
TERRAFORM_VERSION_VALID := $(shell [ "$(TERRAFORM_VERSION)" = "`printf "$(TERRAFORM_VERSION)\n1.6" | sort -V | head -n1`" ] && echo 1 || echo 0)

export TERRAFORM_PROVIDER_SOURCE ?= databricks/databricks
export TERRAFORM_PROVIDER_REPO ?= https://github.com/databricks/terraform-provider-databricks
export TERRAFORM_PROVIDER_VERSION ?= 1.132.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-databricks
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/databricks/$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_NATIVE_PROVIDER_BINARY ?= $(TERRAFORM_PROVIDER_DOWNLOAD_NAME)_v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_DOCS_PATH ?= docs/resources


# Default e2e suite: cheap workspace resources (no compute / Unity Catalog).
# Override with UPTEST_EXAMPLE_LIST. Empty env values (GitHub Actions) still
# use this default. Optional suites:
#   examples/e2e/cluster/group.yaml,examples/e2e/namespaced/group.yaml
#   examples/e2e/cluster/catalog.yaml,examples/e2e/namespaced/catalog.yaml
#   examples/e2e/cluster/cluster.yaml,examples/e2e/namespaced/cluster.yaml
DEFAULT_UPTEST_EXAMPLE_LIST := examples/e2e/cluster/secretscope.yaml,examples/e2e/cluster/directory.yaml,examples/e2e/namespaced/secretscope.yaml,examples/e2e/namespaced/directory.yaml
override UPTEST_EXAMPLE_LIST := $(if $(strip $(UPTEST_EXAMPLE_LIST)),$(UPTEST_EXAMPLE_LIST),$(DEFAULT_UPTEST_EXAMPLE_LIST))

PLATFORMS ?= linux_amd64 linux_arm64

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= $(shell grep -E '^go ' go.mod | awk '{print $2}')
GOLANGCILINT_VERSION ?= 2.13.2
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis config
-include build/makelib/golang.mk

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.33.0
UPTEST_VERSION = v2.2.0
CRDDIFF_VERSION = v0.12.1
# CLI v2.4+ is published at cli.crossplane.io (binary name "crossplane"), not
# releases.crossplane.io/crank. The build submodule still downloads crank from
# releases.crossplane.io, so keep the last version that exists there.
CROSSPLANE_CLI_VERSION = v2.3.4
# for e2e testing
CROSSPLANE_VERSION = 2.4.0
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

REGISTRY_ORGS ?= ghcr.io/buzz-it
IMAGES = $(PROJECT_NAME)
-include build/makelib/imagelight.mk

# ====================================================================================
# Setup XPKG

XPKG_REG_ORGS ?= ghcr.io/buzz-it
# NOTE(hasheddan): skip promoting on xpkg.crossplane.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= ghcr.io/buzz-it
XPKGS = $(PROJECT_NAME)
XPKG_DIR = $(OUTPUT_DIR)/package/$(PLATFORM)
XPKG_IGNORE = kustomize/*,crds/kustomization.yaml
-include build/makelib/xpkg.mk

package.prepare: $(YQ)
	@$(INFO) preparing package manifests
	@rm -rf $(XPKG_DIR)
	@mkdir -p $(XPKG_DIR)
	@cp -a package/. $(XPKG_DIR)/
	@for crd in $(XPKG_DIR)/crds/*.yaml; do \
		$(YQ) eval -i '.spec.conversion = {"strategy":"Webhook","webhook":{"clientConfig":{"service":{"path":"/convert"}},"conversionReviewVersions":["v1"]}}' "$${crd}" || exit 1; \
	done
	@$(OK) preparing package manifests

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# NOTE(hasheddan): we force image building to happen prior to xpkg build so that
# we ensure image is present in daemon.
xpkg.build.provider-databricks: package.prepare do.build.images

# NOTE(hasheddan): we ensure host tools are installed prior to running
# platform-specific build steps in parallel to avoid encountering installation
# race conditions.
build.init: $(UP) $(CROSSPLANE_CLI) $(YQ) check-terraform-version

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json
TF_PROVIDER_MODULE := hack/terraform-provider-databricks

$(TF_PROVIDER_MODULE):
	@if [ ! -f $(TF_PROVIDER_MODULE)/xpprovider/xpprovider.go ]; then \
		$(INFO) preparing terraform-provider-databricks v$(TERRAFORM_PROVIDER_VERSION) with xpprovider; \
		rm -rf $(TF_PROVIDER_MODULE); \
		git clone -c advice.detachedHead=false --depth 1 --branch "v$(TERRAFORM_PROVIDER_VERSION)" "$(TERRAFORM_PROVIDER_REPO)" $(TF_PROVIDER_MODULE); \
		mkdir -p $(TF_PROVIDER_MODULE)/xpprovider; \
		cp hack/xpprovider.go.src $(TF_PROVIDER_MODULE)/xpprovider/xpprovider.go; \
		$(OK) prepared terraform-provider-databricks v$(TERRAFORM_PROVIDER_VERSION); \
	fi

check-terraform-version:
ifneq ($(TERRAFORM_VERSION_VALID),1)
	$(error invalid TERRAFORM_VERSION $(TERRAFORM_VERSION), must be less than 1.6.0 since that version introduced a not permitted BSL license))
endif

$(TERRAFORM): check-terraform-version
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)

pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" ]; then \
  		mkdir -p "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" && \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)"; \
	fi
	@git -C "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: $(TF_PROVIDER_MODULE) $(TERRAFORM_PROVIDER_SCHEMA) pull-docs
	python ./scripts/docs_fix.py

generate.clean-apis:
	@find ./apis -type f -name 'zz_*' -delete
	@find ./apis -type d -empty -delete

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs check-terraform-version generate.clean-apis package.prepare $(TF_PROVIDER_MODULE)
# ====================================================================================
# Targets

# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

go.mod.cachedir:
	@go env GOMODCACHE

# go.mod replace points at this gitignored checkout.
go.modules.download: $(TF_PROVIDER_MODULE)
go.modules.check: $(TF_PROVIDER_MODULE)

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

examples.sync:
	@$(INFO) syncing examples from examples-generated
	@mkdir -p ./examples/cluster ./examples/namespaced
	@cp -a ./examples-generated/cluster/. ./examples/cluster/
	@cp -a ./examples-generated/namespaced/. ./examples/namespaced/
	@$(OK) synced examples from examples-generated

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	$(GO_OUT_DIR)/provider --debug

# ====================================================================================
# End to End Testing
CROSSPLANE_NAMESPACE = crossplane-system
-include build/makelib/local.xpkg.mk
-include build/makelib/controlplane.mk

# End-to-end testing (uptest). Required credentials, either:
#   UPTEST_CLOUD_CREDENTIALS  raw Databricks provider JSON
# or individual vars (also used as GitHub Actions secrets/vars):
#   DATABRICKS_HOST, DATABRICKS_TOKEN
#   DATABRICKS_ACCOUNT_ID, DATABRICKS_CLIENT_ID, DATABRICKS_CLIENT_SECRET
#   DATABRICKS_AUTH_TYPE
#   DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID, DATABRICKS_AZURE_CLIENT_ID,
#   DATABRICKS_AZURE_CLIENT_SECRET, DATABRICKS_AZURE_TENANT_ID,
#   DATABRICKS_AZURE_ENVIRONMENT, DATABRICKS_AZURE_USE_MSI
#   DATABRICKS_GOOGLE_CREDENTIALS, DATABRICKS_GOOGLE_SERVICE_ACCOUNT
# Optional resource parameters:
#   DATABRICKS_NODE_TYPE_ID, DATABRICKS_SPARK_VERSION, DATABRICKS_CATALOG_NAME,
#   DATABRICKS_SCHEMA_NAME, DATABRICKS_DIRECTORY_PATH, DATABRICKS_GROUP_DISPLAY_NAME
#   UPTEST_EXAMPLE_LIST, UPTEST_SKIP_DELETE=true
UPTEST_DATASOURCE_PATH ?= $(WORK_DIR)/uptest-datasource.yaml
UPTEST_DELETE_FLAG := $(if $(filter true,$(UPTEST_SKIP_DELETE)),--skip-delete,)

uptest: $(TF_PROVIDER_MODULE) $(UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@$(INFO) running automated tests
	@UPTEST_DATASOURCE_PATH="$(UPTEST_DATASOURCE_PATH)" ./cluster/test/render-datasource.sh
	@KUBECTL=$(KUBECTL) CHAINSAW=$(CHAINSAW) CROSSPLANE_CLI=$(CROSSPLANE_CLI) CROSSPLANE_NAMESPACE=$(CROSSPLANE_NAMESPACE) UPTEST_DATASOURCE_PATH="$(UPTEST_DATASOURCE_PATH)" $(UPTEST) e2e "${UPTEST_EXAMPLE_LIST}" --data-source="${UPTEST_DATASOURCE_PATH}" --setup-script=cluster/test/setup.sh --teardown-script=cluster/test/dump-managed.sh --default-conditions="Ready" $(UPTEST_DELETE_FLAG) || $(FAIL)
	@$(OK) running automated tests

uptest-debug: UPTEST_SKIP_DELETE := true
uptest-debug: uptest

local-deploy: $(TF_PROVIDER_MODULE) build controlplane.up local.xpkg.deploy.provider.$(PROJECT_NAME)
	@$(INFO) running locally built provider
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 10m
	@$(KUBECTL) -n crossplane-system wait --for=condition=Available deployment --all --timeout=10m
	@$(OK) running locally built provider

e2e: local-deploy uptest

crddiff: $(UPTEST)
	@$(INFO) Checking breaking CRD schema changes
	@for crd in $${MODIFIED_CRD_LIST}; do \
		if ! git cat-file -e "$${GITHUB_BASE_REF}:$${crd}" 2>/dev/null; then \
			echo "CRD $${crd} does not exist in the $${GITHUB_BASE_REF} branch. Skipping..." ; \
			continue ; \
		fi ; \
		echo "Checking $${crd} for breaking API changes..." ; \
		changes_detected=$$(go run github.com/crossplane/uptest/cmd/crddiff@$(CRDDIFF_VERSION) revision --enable-upjet-extensions <(git cat-file -p "$${GITHUB_BASE_REF}:$${crd}") "$${crd}" 2>&1) ; \
		if [[ $$? != 0 ]] ; then \
			printf "\033[31m"; echo "Breaking change detected!"; printf "\033[0m" ; \
			echo "$${changes_detected}" ; \
			echo ; \
		fi ; \
	done
	@$(OK) Checking breaking CRD schema changes

crd-breaking-check:
	@$(INFO) Checking CRD backward compatibility against branch base
	@./scripts/check_crd_breaking_changes.sh "$${BASE_REF:-main}" "$${HEAD_REF:-HEAD}"
	@$(OK) Checking CRD backward compatibility against branch base

schema-version-diff:
	@$(INFO) Checking for native state schema version changes
	@export PREV_PROVIDER_VERSION=$$(git cat-file -p "${GITHUB_BASE_REF}:Makefile" | sed -nr 's/^export[[:space:]]*TERRAFORM_PROVIDER_VERSION[[:space:]]*:=[[:space:]]*(.+)/\1/p'); \
	echo Detected previous Terraform provider version: $${PREV_PROVIDER_VERSION}; \
	echo Current Terraform provider version: $${TERRAFORM_PROVIDER_VERSION}; \
	mkdir -p $(WORK_DIR); \
	git cat-file -p "$${GITHUB_BASE_REF}:config/schema.json" > "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}"; \
	./scripts/version_diff.py config/generated.lst "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}" config/schema.json
	@$(OK) Checking for native state schema version changes

.PHONY: cobertura examples.sync submodules fallthrough run crds.clean

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

endef
# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

.PHONY: crossplane.help help-special

# TODO(negz): Update CI to use these targets.
vendor: modules.download
vendor.check: modules.check

.PHONY: fmt
fmt:
	@echo "✓ Formatting source code with goimports ..."
	@go tool goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*" -not -name "zz_generated.*.go")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*" -not -name "zz_generated.*.go")

