ANSIBLE_DEPLOYER_REPO := contrail-ansible-deployer
CONTRAIL_API_CLIENT_REPO := contrail-api-client
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

ifdef CONTRAIL_API_CLIENT_REPO_DIR
  export CONTRAIL_API_CLIENT_REPO_DIR
else
  export CONTRAIL_API_CLIENT_REPO_DIR := ""
endif
ifdef CONTRAIL_API_CLIENT_BRANCH
  export CONTRAIL_API_CLIENT_BRANCH
else
  export CONTRAIL_API_CLIENT_BRANCH := master
endif

ifdef ANSIBLE_DEPLOYER_REVISION
  export ANSIBLE_DEPLOYER_REVISION
else
  export ANSIBLE_DEPLOYER_REVISION := HEAD
endif

GOPATH ?= `go env GOPATH`

all: check lint test build

deps: ## Setup the go dependencies
	./tools/deps.sh

check: ## Check vendored dependencies
	./tools/check.sh

lint: ## Run linters on the source code
	./tools/lint.sh

nocovtest: COVERPROFILE = none
nocovtest: test

test: ## Run go test with race and coverage args
	./tools/test.sh $(COVERPROFILE)

build: ## Run go build
	go build ./cmd/...

generate: reset_gen ## Run the source code generator
	mkdir -p public
	go run cmd/contrailschema/main.go generate --schemas schemas --templates tools/templates/template_config.yaml --schema-output public/schema.json --openapi-output public/openapi.json
	./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf -I ./proto --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ proto/github.com/Juniper/contrail/pkg/models/generated.proto
	./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf -I ./proto --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ proto/github.com/Juniper/contrail/pkg/services/generated.proto
	./bin/protoc -I ./vendor/ -I ./vendor/github.com/gogo/protobuf/protobuf -I ./proto --doc_out=./doc --doc_opt=markdown,proto.md proto/github.com/Juniper/contrail/pkg/services/generated.proto proto/github.com/Juniper/contrail/pkg/models/generated.proto
	go tool fix ./pkg/services/generated.pb.go
	go fmt github.com/Juniper/contrail/pkg/db
	go fmt github.com/Juniper/contrail/pkg/models
	go fmt github.com/Juniper/contrail/pkg/services
	go fmt github.com/Juniper/contrail/pkg/compilationif
	mkdir -p pkg/types/mock
	mockgen -destination=pkg/types/mock/gen_in_transaction_doer_mock.go -package=typesmock -source pkg/types/service.go InTransactionDoer
	mkdir -p pkg/services/mock
	mockgen -destination=pkg/services/mock/gen_service_mock.go -package=servicesmock -source pkg/services/gen_service_interface.go Service
	mkdir -p pkg/types/ipam/mock
	mockgen -destination=pkg/types/ipam/mock/gen_address_manager_mock.go -package=ipammock -source pkg/types/ipam/address_manager.go AddressManager
	mockgen -destination=pkg/types/ipam/mock/gen_int_pool_allocator_mock.go -package=ipammock -source pkg/types/ipam/interface.go IntPoolAllocator

reset_gen: ## Remove genarated files
	find pkg/ -name gen_* -delete
	find pkg/ -name generated.pb.go -delete
	rm -rf public/[^watch.html]*
	rm -rf proto/*
	rm -f tools/init_mysql.sql
	rm -f tools/init_psql.sql
	rm -f tools/cleanup.sql
	rm -rf pkg/types/mock
	rm -rf pkg/services/mock
	rm -rf pkg/types/ipam/mock

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

install:
	go install ./cmd/contrail
	go install ./cmd/contrailcli
	go install ./cmd/contrailutil

testenv: ## Setup docker based test environment. (You need docker)
	./tools/testenv.sh

reset_db: ## Reset Database with latest schema.
	./tools/reset_db_mysql.sh
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
	./tools/reset_db_psql.sh
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail_postgres.yml

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailcli_{{.OS}}_{{.Arch}}" ./cmd/contrailcli
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

.PHONY: docker
docker: ## Generate docker files
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
	cp pkg/cloud/configs/onprem_cloud_topology.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	cp pkg/cloud/configs/public_cloud_topology.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	cp pkg/cloud/configs/secret.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	cp pkg/cluster/configs/contrail_common.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	cp pkg/cluster/configs/gateway_common.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	cp pkg/cluster/configs/tor_common.tmpl $(BUILD_DIR)/docker/contrail_go/templates/
	mkdir -p $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && rm -rf $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)/
ifeq ($(ANSIBLE_DEPLOYER_REPO_DIR),"")
		git clone -b $(ANSIBLE_DEPLOYER_BRANCH) https://github.com/Juniper/$(ANSIBLE_DEPLOYER_REPO).git $(BUILD_DIR)/docker/contrail_go/contrail-ansible-deployer
		cd $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && git checkout $(ANSIBLE_DEPLOYER_REVISION)
else
		cp -r $(ANSIBLE_DEPLOYER_REPO_DIR) $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)
endif

ifeq ($(CONTRAIL_API_CLIENT_REPO_DIR),"")
		git clone -b $(CONTRAIL_API_CLIENT_BRANCH) https://github.com/Juniper/$(CONTRAIL_API_CLIENT_REPO).git $(BUILD_DIR)/docker/contrail_go/$(CONTRAIL_API_CLIENT_REPO)
else
		cp -r $(CONTRAIL_API_CLIENT_REPO_DIR) $(BUILD_DIR)/docker/contrail_go/$(CONTRAIL_API_CLIENT_REPO)
endif
	docker build -t "contrail-go" $(BUILD_DIR)/docker/contrail_go

help: ## Display help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
