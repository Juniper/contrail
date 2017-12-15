all: deps lint test build

deps:
	./tools/deps.sh

lint:
	./tools/lint.sh

test:
	./tools/test.sh

build:
	go build ./cmd/...

generate:
	go run cmd/contrailutil/main.go generate --schemas schemas --templates tools/templates/template_config.yaml --schema-output public/schema.json
	./tools/fmt.sh

package:
	go run cmd/contrailutil/main.go package

binaries:
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrail_{{.OS}}_{{.Arch}}" ./cmd/contrail
	gox -osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/contrailutil_{{.OS}}_{{.Arch}}" ./cmd/contrailutil