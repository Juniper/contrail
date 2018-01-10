all: deps lint test build

deps: ## Setup the go dependencies
	./tools/deps.sh

lint: ## Runs gometalinter on the source code
	./tools/lint.sh

test: ## Run go test with race and coverage args
	./tools/test.sh

build: ## Run go build
	go build ./cmd/...

generate: ## Run the source code generator
	go run cmd/contrailutil/main.go generate --schemas schemas --templates tools/templates/template_config.yaml --schema-output public/schema.json
	./tools/fmt.sh

package: ## Generate the packages
	go run cmd/contrailutil/main.go package

reset_db: ## Reset Database with latest schema.
	./tools/reset_db.sh

binaries: ## Generate the contrail and contrailutil binaries
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
