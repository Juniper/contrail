all: deps lint test build

deps:
	./tools/deps.sh

lint:
	./tools/lint.sh

test:
	go test $(glide novendor)

build:
	go build ./cmd/...

generate:
	go run cmd/contrail_util/main.go generate --schema-dir schemas --template-config cmd/contrail_util/templates/template_config.yaml
	go fmt ./pkg/generated/...