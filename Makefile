ANSIBLE_DEPLOYER_REPO := contrail-ansible-deployer
BUILD_DIR := ../build
SRC_DIRS := cmd pkg vendor
DB_FILES := init_mysql.sql init_psql.sql init_data.yaml
ifdef ANSIBLE_DEPLOYER_REPO_DIR
  export ANSIBLE_DEPLOYER_REPO_DIR
else
  export ANSIBLE_DEPLOYER_REPO_DIR := ""
endif
ifdef ANSIBLE_DEPLOYER_BRANCH
  export ANSIBLE_DEPLOYER_BRANCH
else
  export ANSIBLE_DEPLOYER_BRANCH := master
endif

ifdef ANSIBLE_DEPLOYER_REVISION
  export ANSIBLE_DEPLOYER_REVISION
else
  export ANSIBLE_DEPLOYER_REVISION := HEAD
endif

GOPATH ?= `go env GOPATH`

all: check lint test build

deps: ## Install development dependencies
	./tools/deps.sh

check: ## Check vendored dependencies
	./tools/check.sh

lint: ## Run linters on the source code
	./tools/lint.sh

nocovtest: COVERPROFILE = none
nocovtest: test

test: ## Run tests with coverage
	./tools/test.sh $(COVERPROFILE)

build: ## Build all binaries without producing output
	go build ./cmd/...

GENERATE_DEPS := reset_gen
GENERATE_DEPS += generate_templates
GENERATE_DEPS += tmp/generate_proto_models tmp/generate_proto_baseservices tmp/generate_proto_services
GENERATE_DEPS += tmp/generate_models_doc
GENERATE_DEPS += tmp/gen_mock

gen_format:
	find ./cmd ./pkg -name 'gen_*.go' | xargs -L 1 go fmt

generate: $(GENERATE_DEPS) ## Run the source code generator
	cd extension && $(MAKE) generate

generate_templates:
	@mkdir -p public
	go run cmd/contrailschema/main.go generate \
		--schemas schemas --templates tools/templates/template_config.yaml \
		--schema-output public/schema.json --openapi-output public/openapi.json

tmp/gen_mock: pkg/types/service.go pkg/services/gen_service_interface.go pkg/types/ipam/address_manager.go
	@mkdir -p tmp
	@touch $@
	@mkdir -p pkg/types/mock
	@mkdir -p pkg/services/mock
	@mkdir -p pkg/types/ipam/mock
	mockgen -destination=pkg/types/mock/gen_types_mock.go -package=typesmock -source pkg/types/service.go
	mockgen -destination=pkg/services/mock/gen_service_mock.go -package=servicesmock -source pkg/services/gen_service_interface.go Service
	mockgen -destination=pkg/types/ipam/mock/gen_address_manager_mock.go -package=ipammock -source pkg/types/ipam/address_manager.go AddressManager

PROTO_DEPS := tmp/proto/github.com/Juniper/contrail/pkg/models/generated.tmp
PROTO_DEPS += tmp/proto/github.com/Juniper/contrail/pkg/services/baseservices/base.tmp
PROTO_DEPS += tmp/proto/github.com/Juniper/contrail/pkg/services/generated.tmp

gen_proto: $(PROTO_DEPS)

tmp/%.tmp: %.proto
	./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf \
		-I ./proto --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ \
		$<
	@mkdir -p $(shell dirname $@)
	@touch $@

tmp/generate_models_doc: proto/github.com/Juniper/contrail/pkg/models/generated.proto
	@mkdir -p tmp
	@touch $@
	./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf \
		-I ./proto --doc_out=./doc --doc_opt=markdown,proto.md proto/github.com/Juniper/contrail/pkg/services/generated.proto \
		proto/github.com/Juniper/contrail/pkg/models/generated.proto

purge_gen:
	find pkg/ -name gen_* -delete
	find pkg/ -name generated.pb.go -delete
	find proto/ -name generated.proto -delete

reset_gen: ## Remove genarated files
	rm -rf public/[^watch.html]*
	rm -f tools/init_mysql.sql
	rm -f tools/init_psql.sql
	rm -f tools/cleanup_mysql.sql
	rm -f tools/cleanup_psql.sql
	cd extension && $(MAKE) reset_gen

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

install:
	go install ./cmd/contrail
	go install ./cmd/contrailcli
	go install ./cmd/contrailutil

testenv: ## Setup docker based test environment
	./tools/testenv.sh

reset_db: zero_db init_db ## Reset databases with latest schema and load initial data

reset_mysql: zero_mysql init_mysql

reset_psql: zero_psql init_psql

zero_db: zero_mysql zero_psql

zero_mysql:
	./tools/reset_db_mysql.sh

zero_psql:
	./tools/reset_db_psql.sh

clean_db: clean_mysql clean_psql init_db ## Truncate all database tables and load initial data

clean_mysql:
	docker exec -i contrail_mysql mysql -uroot -pcontrail123 contrail_test < tools/cleanup_mysql.sql

clean_psql:
	docker exec -i contrail_postgres psql -U postgres -d contrail_test < tools/cleanup_psql.sql

init_db: init_mysql init_psql ## Load initial data to databases

init_mysql:
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml

init_psql:
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail_postgres.yml

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailcli_{{.OS}}_{{.Arch}}" ./cmd/contrailcli
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

.PHONY: docker_prepare
docker_prepare: ## Prepare common data to generate Docker files (use target `docker` or `docker_k8s` instead)
	rm -rf $(BUILD_DIR) && mkdir -p $(BUILD_DIR)/contrail
	cp -r docker $(BUILD_DIR)
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrail" ./cmd/contrail
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrailcli" ./cmd/contrailcli
	CGO_ENABLED=0 gox -osarch="linux/amd64" --output "$(BUILD_DIR)/docker/contrail_go/contrailutil" ./cmd/contrailutil
	cp -r sample $(BUILD_DIR)/docker/contrail_go/etc
	$(foreach db_file, $(DB_FILES), cp tools/$(db_file) $(BUILD_DIR)/docker/contrail_go/etc;)
	cp -r public $(BUILD_DIR)/docker/contrail_go/public
	$(foreach src, $(SRC_DIRS), cp -r ../contrail/$(src) $(BUILD_DIR)/contrail;)
	mkdir -p $(BUILD_DIR)/docker/contrail_go/templates/ && cp pkg/cluster/configs/instances.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	mkdir -p $(BUILD_DIR)/docker/contrail_go/templates/ && cp pkg/cluster/configs/inventory.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	mkdir -p $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && rm -rf $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)/
ifeq ($(ANSIBLE_DEPLOYER_REPO_DIR),"")
		git clone -b $(ANSIBLE_DEPLOYER_BRANCH) https://github.com/Juniper/$(ANSIBLE_DEPLOYER_REPO).git $(BUILD_DIR)/docker/contrail_go/contrail-ansible-deployer
		cd $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && git checkout $(ANSIBLE_DEPLOYER_REVISION)
else
		cp -r $(ANSIBLE_DEPLOYER_REPO_DIR) $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)
endif

.PHONY: docker
docker: docker_prepare ## Generate Docker files
	docker build -t "contrail-go" $(BUILD_DIR)/docker/contrail_go

# This target creates contrail-go docker that is able to work as a drop-in replacement to original config-api.
# It depends on 'docker' target to inherit all the necesary steps with minimal changes
.PHONY: docker_k8s
docker_k8s: docker_prepare ## Create contrail-go docker as a drop-in replacement to original config-api
	## Copy dockerfile because it must be in a build context dir
	cp -f docker/contrail_go/Dockerfile-k8s $(BUILD_DIR)/docker/contrail_go
	cp -f docker/contrail_go/etc/contrail-k8s.yml $(BUILD_DIR)/docker/contrail_go/etc/
	docker build -t "contrail-go-config" -f $(BUILD_DIR)/docker/contrail_go/Dockerfile-k8s $(BUILD_DIR)/docker/contrail_go

help: ## Display help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
