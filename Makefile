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
	go run cmd/contrailutil/main.go generate --schemas schemas --templates tools/templates/template_config.yaml
	./tools/fmt.sh