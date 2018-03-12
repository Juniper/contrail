all: deps lint test build integration

deps: ## Setup the go dependencies
	./tools/deps.sh

lint: ## Runs gometalinter on the source code
	./tools/lint.sh

test: ## Run go test with race and coverage args
	./tools/test.sh

build: ## Run go build
	go build ./cmd/...

generate: ## Run the source code generator
	rm -rf pkg/models/gen_*
	rm -rf pkg/services/gen_*
	rm -rf pkg/db/gen_*
	rm -rf terraform/terraform-provider-contrail/contrail/gen_*
	mkdir -p public || echo "ok"
	go run cmd/contrailutil/main.go generate --schemas schemas --templates tools/templates/template_config.yaml --schema-output public/schema.json --openapi-output public/openapi.json
	./bin/protoc -I $(GOPATH)/src/ -I $(GOPATH)/smakrc/github.com/gogo/protobuf/protobuf -I ./proto --doc_out=./doc --doc_opt=markdown,proto_model.md  --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ proto/github.com/Juniper/contrail/pkg/models/generated.proto
	./bin/protoc -I $(GOPATH)/src/ -I $(GOPATH)/src/github.com/gogo/protobuf/protobuf -I ./proto --doc_out=./doc --doc_opt=markdown,proto_service.md --gogo_out=plugins=grpc:$(GOPATH)/src/ proto/github.com/Juniper/contrail/pkg/services/generated.proto
	go fmt github.com/Juniper/contrail/pkg/db
	go fmt github.com/Juniper/contrail/pkg/models
	go fmt github.com/Juniper/contrail/pkg/services
	go fmt github.com/Juniper/contrail/terraform/terraform-provider-contrail/contrail

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

reset_gen:
	rm pkg/models/gen*
	rm pkg/services/gen*
	rm pkg/db/gen_*
	rm doc/proto_model.md
	rm doc/proto_service.md
	rm -rf public/*
	rm -rf proto/*
	rm tools/init_mysql.sql
	rm tools/init_psql.sql
	rm tools/cleanup.sql
	rm pkg/serviceif/serviceif.go

install:
	go install ./cmd/contrail
	go install ./cmd/contrailcli
	go install ./cmd/contrailutil

reset_db: ## Reset Database with latest schema.
	./tools/reset_db.sh

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
