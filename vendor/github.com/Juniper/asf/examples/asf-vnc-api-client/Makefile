BUILD_DIR := ../build
GENERATE_DS_REPO := contrail-api-client

GENERATE_DS_REPO_DIR ?= ""
GENERATE_DS_BRANCH ?= master
GENERATE_DS_REVISION ?= HEAD

YAML_SCHEMA_DIR := $(BUILD_DIR)/schemas
CONTRAILSCHEMA = $(shell go list -f '{{ .Target }}' github.com/Juniper/asf/cmd/contrailschema)

all: generate install test check ## Perform all targets

check: ## Ensure the go mod dependencies are correct
	## TODO(ijohnson): Remove GOPRIVATE variable once asf repo is made public
	GOPRIVATE=github.com/Juniper/asf go mod tidy

generateds_prepare: ## Prepare generateDS repo
	rm -rf $(BUILD_DIR)/$(GENERATE_DS_REPO)
ifeq ($(GENERATE_DS_REPO_DIR),"")
		git clone -b $(GENERATE_DS_BRANCH) https://github.com/Juniper/$(GENERATE_DS_REPO).git $(BUILD_DIR)/$(GENERATE_DS_REPO)
		cd $(BUILD_DIR)/$(GENERATE_DS_REPO) && git checkout $(GENERATE_DS_REVISION)
else
		cp -r $(GENERATE_DS_REPO_DIR) $(BUILD_DIR)/$(GENERATE_DS_REPO)
endif

generate_yaml_schema: generateds_prepare
	rm -rf $(YAML_SCHEMA_DIR)
	mkdir -p $(YAML_SCHEMA_DIR)
	## https://github.com/Juniper/contrail-api-client/blob/master/generateds/README_jsonschema.md
	$(BUILD_DIR)/$(GENERATE_DS_REPO)/generateds/generateDS.py -f -o $(YAML_SCHEMA_DIR) -g contrail-json-schema $(BUILD_DIR)/$(GENERATE_DS_REPO)/schema/all_cfg.xsd

install_contrailschema: ## Install Contrailschema binary
	go install github.com/Juniper/asf/cmd/contrailschema

generate_client: generate_yaml_schema install_contrailschema 
	# Generate client library for contrail resources.
	@mkdir -p public/
	$(CONTRAILSCHEMA) generate --no-regenerate --schemas $(YAML_SCHEMA_DIR) \
		--template-config templates/template_config.yaml \
                --models-import-path github.com/Juniper/asf/examples/asf-vnc-api-client/pkg/models \
                --services-import-path github.com/Juniper/asf/examples/asf-vnc-api-client/pkg/services

format_gen: ## Format generated source code
	find ./cmd ./pkg -name 'gen_*.go' -exec go fmt {} \;

generate: generate_client format_gen
