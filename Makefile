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
	mockgen -destination=pkg/types/mock/gen_types_mock.go -package=typesmock -source pkg/types/service.go
	mkdir -p pkg/services/mock
	mockgen -destination=pkg/services/mock/gen_service_mock.go -package=servicesmock -source pkg/services/gen_service_interface.go Service
	mkdir -p pkg/types/ipam/mock
	mockgen -destination=pkg/types/ipam/mock/gen_address_manager_mock.go -package=ipammock -source pkg/types/ipam/address_manager.go AddressManager

reset_gen: ## Remove genarated files
	find pkg/ -name gen_* -delete
	find pkg/ -name generated.pb.go -delete
	rm -rf public/[^watch.html]*
	rm -rf proto/*
	rm -f tools/init_mysql.sql
	rm -f tools/init_psql.sql
	rm -f tools/cleanup_mysql.sql
	rm -f tools/cleanup_psql.sql
	rm -rf pkg/types/mock
	rm -rf pkg/services/mock
	rm -rf pkg/types/ipam/mock

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

install:
	go install ./cmd/contrail
	go install ./cmd/contrailcli
	go install ./cmd/contrailutil

testenv: ## Setup docker based test environment
	./tools/testenv.sh

reset_db: ## Reset databases with latest schema and load initial data
	./tools/reset_db_mysql.sh
	./tools/reset_db_psql.sh
	make init_db

clean_db: ## Truncate all database tables and load initial data
	docker exec -i contrail_mysql mysql -uroot -pcontrail123 contrail_test < tools/cleanup_mysql.sql
	docker exec -i contrail_postgres psql -U postgres -d contrail_test < tools/cleanup_psql.sql
	make init_db

init_db: ## Load initial data to databases
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail.yml
	go run cmd/contrailutil/main.go convert --intype yaml --in tools/init_data.yaml --outtype rdbms -c sample/contrail_postgres.yml

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailcli_{{.OS}}_{{.Arch}}" ./cmd/contrailcli
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

.PHONY: docker
docker: ## Generate Docker files
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
	mkdir -p $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && rm -rf $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)/
ifeq ($(ANSIBLE_DEPLOYER_REPO_DIR),"")
		git clone -b $(ANSIBLE_DEPLOYER_BRANCH) https://github.com/Juniper/$(ANSIBLE_DEPLOYER_REPO).git $(BUILD_DIR)/docker/contrail_go/contrail-ansible-deployer
		cd $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO) && git checkout $(ANSIBLE_DEPLOYER_REVISION)
else
		cp -r $(ANSIBLE_DEPLOYER_REPO_DIR) $(BUILD_DIR)/docker/contrail_go/$(ANSIBLE_DEPLOYER_REPO)
endif
	docker build -t "contrail-go" $(BUILD_DIR)/docker/contrail_go

help: ## Display help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
