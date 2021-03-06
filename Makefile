BUILD_DIR := ../build
CONTRAIL_APIDOC_PATH := public/doc/index.html
CONTRAIL_OPENAPI_PATH := public/openapi.json
DOCKER_FILE := $(BUILD_DIR)/docker/contrail_go/Dockerfile
CONTRAILSCHEMA = $(shell go list -f '{{ .Target }}' ./vendor/github.com/Juniper/asf/cmd/contrailschema)
CONTRAILUTIL := $(shell go list -f '{{ .Target }}' ./cmd/contrailutil)
GOPATH ?= $(shell go env GOPATH)
PATH := $(PATH):$(GOPATH)/bin
SOURCEDIR ?= $(realpath .)

DB_FILES := gen_init_psql.sql init_psql.sql init_data.yaml
SRC_DIRS := cmd pkg vendor

BASE_IMAGE_REGISTRY ?= opencontrailnightly
BASE_IMAGE_REPOSITORY ?= contrail-base
BASE_IMAGE_TAG ?= latest

# This is needed by generate* targets that works only sequentially
ifneq ($(filter generate,$(MAKECMDGOALS)),)
.NOTPARALLEL:
endif

all: deps govendor generate install testenv reset_db test lint check ## Perform all checks

deps: ## Install development dependencies
	./tools/deps.sh

generate: fast_generate format_gen ## Generate source code and documentation

fast_generate: generate_pb_go generate_mocks doc/proto.md ## Generate source code and documentation without formatting

generate_pb_go: generate_go pkg/models/gen_model.pb.go pkg/services/gen_service.pb.go pkg/services/gen_plugins.pb.go ## Generate *pb.go files from *.proto definitions

generate_go: install_contrailschema ## Generate source code from templates and schema
	# Generate for contrail resources.
	@mkdir -p public/
	$(CONTRAILSCHEMA) generate --no-regenerate --schemas schemas/contrail --addons schemas/addons \
		--template-config tools/templates/contrail/template_config.yaml \
		--db-import-path github.com/Juniper/contrail/pkg/db \
		--etcd-import-path github.com/Juniper/contrail/pkg/etcd \
		--models-import-path github.com/Juniper/contrail/pkg/models \
		--services-import-path github.com/Juniper/contrail/pkg/services \
		--rbac-import-path github.com/Juniper/contrail/pkg/rbac \
		--schema-output public/schema.json --openapi-output $(CONTRAIL_OPENAPI_PATH)
	# Generate for openstack api resources.
	@mkdir -p public/neutron
	$(CONTRAILSCHEMA) generate --no-regenerate --schemas schemas/neutron \
		--template-config tools/templates/neutron/template_config.yaml \
		--schema-output public/neutron/schema.json --openapi-output public/neutron/openapi.json

GOGOPROTO_PATH = ./vendor/github.com/gogo/protobuf
ASF_PATH = ./vendor/github.com/Juniper/asf
PROTO = ./bin/protoc -I $(GOGOPROTO_PATH) -I $(GOGOPROTO_PATH)/protobuf -I ./proto -I $(ASF_PATH)
PROTO_PKG_PATH := proto/github.com/Juniper/contrail/pkg

pkg/%.pb.go: $(PROTO_PKG_PATH)/%.proto
	# M options in protoc are definitions of go import substitutions in resulting go files
	# see: https://github.com/gogo/protobuf/blob/master/README#L239
	$(PROTO) --gogofaster_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mpkg/services/base.proto=github.com/Juniper/asf/pkg/services,\
plugins=grpc:. $<
	cp -a github.com/Juniper/contrail/* .
	rm -r github.com
	go tool fix $@

MOCKGEN = $(shell go list -f "{{ .Target }}" ./vendor/github.com/golang/mock/mockgen)
MOCKS := pkg/types/mock/service.go \
	pkg/services/mock/gen_service_interface.go \
	pkg/services/mock/useragent_kv.go \
	pkg/types/ipam/mock/address_manager.go \

define create-generate-mock-target
  $1: $(shell dirname $(shell dirname $1))/$(shell basename $1)
	mkdir -p $(shell dirname $1)
	$(MOCKGEN) -destination=$1 \
	-package=$(shell basename $(shell dirname $(shell dirname $1)))mock \
	-source $(shell dirname $(shell dirname $1))/$(shell basename $1)
endef

$(foreach mock,$(MOCKS),$(eval $(call create-generate-mock-target,$(mock))))

generate_mocks: $(MOCKS) ## Generate source code of mocks

doc/proto.md: $(PROTO_PKG_PATH)/models/gen_model.proto $(PROTO_PKG_PATH)/services/gen_service.proto ## Generate Protobuf definitions documentation
	$(PROTO) --doc_out=./doc --doc_opt=markdown,proto.md $^

format_gen: ## Format generated source code
	find ./cmd ./pkg -name 'gen_*.go' -exec go fmt {} \;

clean_gen: ## Remove generated source code and documentation
	rm -rf public/[^watch.html]* doc/proto.md
	find tools/ proto/ pkg/ -name gen_* -delete

build: ## Build all binaries without producing output
	go build ./cmd/...

install: install_contrail install_contrailcli install_contrailschema install_contrailutil ## Install all binaries

install_contrail: ## Install Contrail binary
	go install ./cmd/contrail

install_contrailcli:  ## Install Contrailcli binary
	go install ./cmd/contrailcli

install_contrailschema: ## Install Contrailschema binary
	go install ./vendor/github.com/Juniper/asf/cmd/contrailschema/

install_contrailutil: ## Install Contrailutil binary
	go install ./cmd/contrailutil

testenv: ## Setup Docker based test environment
	./tools/testenv.sh

reset_db: zero_db init_db ## Reset database with the latest schema and load initial data

zero_db: ## Drop and recreate test database
	./tools/reset_db_psql.sh

init_db: install_contrailutil ## Load initial data to databases
	$(CONTRAILUTIL) convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml

clean_db: truncate_db init_db ## Truncate all database tables and load initial data

truncate_db: ## Remove test database data
	docker exec -i contrail_postgres psql -U postgres -d contrail_test < tools/gen_cleanup_psql.sql

nocovtest: COVERPROFILE = none
nocovtest: test

test: ## Run tests with coverage
	./tools/test.sh $(COVERPROFILE)

lint: ## Run linters on the source code
	./tools/lint.sh

govendor: ## vendor all dependencies
	./tools/vendor.sh

check: ## Check if dependencies are locked
	./tools/check.sh

format: ## Format source code
	./tools/fmt.sh

docker: apidoc docker_prepare docker_build ## Build contrail-go Docker image

$(CONTRAIL_OPENAPI_PATH):
	$(MAKE) generate_go

DOCKER_GO_SRC_DIR := /go/src/github.com/Juniper/contrail
$(CONTRAIL_APIDOC_PATH): $(CONTRAIL_OPENAPI_PATH)
ifeq (, $(shell which spectacle))
	$(info No spectacle in $(PATH) consider installing it. Running in docker.)
	docker run --rm -v $(SOURCEDIR):/go node:13-alpine sh -c \
		"npm install --unsafe-perm -g spectacle-docs@1.0.7 && spectacle -1 -t $(DOCKER_GO_SRC_DIR)/$(dir $(CONTRAIL_APIDOC_PATH)) $(DOCKER_GO_SRC_DIR)/$(CONTRAIL_OPENAPI_PATH)"
else
	mkdir -p $(dir $(CONTRAIL_APIDOC_PATH))
	spectacle -1 -t $(dir $(CONTRAIL_APIDOC_PATH)) $(CONTRAIL_OPENAPI_PATH)
endif

apidoc: $(CONTRAIL_APIDOC_PATH) ## Generate OpenAPI html documentation

docker_prepare: ## Prepare common data to generate Docker files (use target `docker` or `docker_config_api` instead)
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)/docker/contrail_go/src/contrail && cp -r docker/* $(BUILD_DIR)/docker/
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrail" ./cmd/contrail
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrailcli" ./cmd/contrailcli
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrailutil" ./cmd/contrailutil
	cp -r sample $(BUILD_DIR)/docker/contrail_go/etc
	$(foreach db_file, $(DB_FILES), cp tools/$(db_file) $(BUILD_DIR)/docker/contrail_go/etc;)
	cp -r public $(BUILD_DIR)/docker/contrail_go/public
	$(foreach src, $(SRC_DIRS), cp -a ../contrail/$(src) $(BUILD_DIR)/docker/contrail_go/src/contrail;)

docker_build: ## Build Docker image with Contrail binary
	# Remove ARG and modify FROM (workaround for bug https://bugzilla.redhat.com/show_bug.cgi?id=1572019)
	sed -e '/FROM/,$$!d' \
		-e 's/FROM $${BASE_IMAGE_REGISTRY}\/$${BASE_IMAGE_REPOSITORY}:$${BASE_IMAGE_TAG}/FROM ${BASE_IMAGE_REGISTRY}\/${BASE_IMAGE_REPOSITORY}:${BASE_IMAGE_TAG}/' ${DOCKER_FILE} > ${DOCKER_FILE}.patched
	docker build \
		--build-arg BASE_IMAGE_REGISTRY=$(BASE_IMAGE_REGISTRY) \
		--build-arg BASE_IMAGE_REPOSITORY=$(BASE_IMAGE_REPOSITORY) \
		--build-arg BASE_IMAGE_TAG=$(BASE_IMAGE_TAG) \
		--build-arg GOPATH=$(GOPATH) \
		--file ${DOCKER_FILE}.patched \
		-t "contrail-go" $(BUILD_DIR)/docker/contrail_go

package: install_contrailutil ## Generate the packages
	$(CONTRAILUTIL) package

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailcli_{{.OS}}_{{.Arch}}" ./cmd/contrailcli
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

help: ## Display help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: docker_prepare docker generate_go
.SUFFIXES: .go .proto
