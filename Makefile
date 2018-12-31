ANSIBLE_DEPLOYER_REPO := contrail-ansible-deployer
BUILD_DIR := ../build
SRC_DIRS := cmd pkg vendor
DB_FILES := gen_init_mysql.sql gen_init_psql.sql init_data.yaml
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

# This is needed by generate* targets that works only sequentially
.NOTPARALLEL:

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

format_gen:
	find ./cmd ./pkg -name 'gen_*.go' -exec go fmt {} \;

fast_generate: generate_pb_go generate_mocks doc/proto.md

generate_pb_go: generate_go pkg/models/gen_model.pb.go pkg/services/baseservices/base.pb.go pkg/services/gen_service.pb.go

generate: fast_generate format_gen

generate_go:
	# Generate for contrail resources.
	@mkdir -p public/
	go run cmd/contrailschema/main.go generate \
		--schemas schemas/contrail --templates tools/templates/contrail/template_config.yaml \
		--schema-output public/schema.json --openapi-output public/openapi.json
	# Generate for openstack api resources.
	@mkdir -p public/neutron
	go run  cmd/contrailschema/main.go generate \
	    --schemas schemas/neutron --templates tools/templates/neutron/template_config.yaml \
		--schema-output public/neutron/schema.json --openapi-output public/neutron/openapi.json

TYPES_MOCK := pkg/types/mock/gen_service_mock.go
SERVICES_MOCK := pkg/services/mock/gen_service_mock.go
IPAM_MOCK := pkg/types/ipam/mock/gen_address_manager_mock.go

generate_mocks: $(TYPES_MOCK) $(SERVICES_MOCK) $(IPAM_MOCK)

$(TYPES_MOCK): pkg/types/service.go
	mkdir -p $(@D)
	mockgen -destination=$@ -package=typesmock -source $<

$(SERVICES_MOCK): pkg/services/gen_service_interface.go
	mkdir -p $(@D)
	mockgen -destination=$@ -package=servicesmock -source $<

$(IPAM_MOCK): pkg/types/ipam/address_manager.go
	mkdir -p $(@D)
	mockgen -destination=$@ -package=ipammock -source $<

PROTO := ./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf -I ./proto
PROTO_PKG_PATH := proto/github.com/Juniper/contrail/pkg

pkg/%.pb.go: $(PROTO_PKG_PATH)/%.proto
	$(PROTO) --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
plugins=grpc:$(GOPATH)/src/ $<
	go tool fix $@

doc/proto.md: $(PROTO_PKG_PATH)/models/gen_model.proto $(PROTO_PKG_PATH)/services/gen_service.proto
	$(PROTO) --doc_out=./doc --doc_opt=markdown,proto.md $^

clean_gen:
	rm -rf public/[^watch.html]*
	find tools/ prot/ pkg/ -name gen_* -delete

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

install:
	go install ./cmd/contrail
	go install ./cmd/contrailcli
	go install ./cmd/contrailutil

testenv: ## Setup docker based test environment
	./tools/patroni/build_patroni.sh
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
	docker exec -i contrail_mysql mysql -uroot -pcontrail123 contrail_test < tools/gen_cleanup_mysql.sql

clean_psql:
	docker exec -i contrail_postgres psql -U postgres -d contrail_test < tools/gen_cleanup_psql.sql

init_db: init_mysql init_psql ## Load initial data to databases

init_mysql:
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml

init_psql:
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail_postgres.yml

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailcli_{{.OS}}_{{.Arch}}" ./cmd/contrailcli
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

docker_prepare: ## Prepare common data to generate Docker files (use target `docker` or `docker_config_api` instead)
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

docker: docker_prepare ## Generate Docker files
	docker build -t "contrail-go" $(BUILD_DIR)/docker/contrail_go

# This target creates contrail-go docker that is able to work as a drop-in replacement to original config-api.
# It depends on 'docker' target to inherit all the necesary steps with minimal changes
docker_config_api: docker_prepare ## Create contrail-go docker as a drop-in replacement to original config-api
	## Copy dockerfile because it must be in a build context dir
	cp -f docker/contrail_go/Dockerfile-config_api $(BUILD_DIR)/docker/contrail_go
	cp -f sample/contrail-config_api.yml $(BUILD_DIR)/docker/contrail_go/etc/
	docker build -t "contrail-go-config" -f $(BUILD_DIR)/docker/contrail_go/Dockerfile-config_api $(BUILD_DIR)/docker/contrail_go

help: ## Display help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.PHONY: docker_prepare docker_config_api docker generate_go
.SUFFIXES: .go .proto
