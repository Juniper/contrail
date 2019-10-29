.PHONY: bench setup test build dist docker examples

EXAMPLE_DIR=$(shell pwd)/examples
DOCS_DIR=$(EXAMPLE_DIR)/doc
PROTOS_DIR=$(EXAMPLE_DIR)/proto

EXAMPLE_CMD=protoc --plugin=protoc-gen-doc -Iexamples/proto --doc_out=examples/doc
DOCKER_CMD=docker run --rm -v $(DOCS_DIR):/out:rw -v $(PROTOS_DIR):/protos:ro -v $(EXAMPLE_DIR)/templates:/templates:ro pseudomuto/protoc-gen-doc

setup:
	$(info Synching dev tools and dependencies...)
	@if test -z $(which retool); then go get github.com/twitchtv/retool; fi
	@retool sync
	@retool do dep ensure

resources.go: resources/*.tmpl resources/*.json
	$(info Generating resources...)
	@go run resources/main.go -in resources -out resources.go -pkg gendoc

fixtures/fileset.pb: fixtures/*.proto fixtures/generate.go
	$(info Generating fixtures...)
	@cd fixtures && go generate

test: fixtures/fileset.pb resources.go
	@go test -cover -race ./ ./cmd/...

bench:
	@go test -bench=.

build: setup resources.go
	@go build ./cmd/...

dist:
	@script/dist.sh

docker:
	@script/push_to_docker.sh

docker_test: build docker
	@rm -f examples/doc/*
	@$(DOCKER_CMD) --doc_opt=docbook,example.docbook:Ignore*
	@$(DOCKER_CMD) --doc_opt=html,example.html:Ignore*
	@$(DOCKER_CMD) --doc_opt=json,example.json:Ignore*
	@$(DOCKER_CMD) --doc_opt=markdown,example.md:Ignore*
	@$(DOCKER_CMD) --doc_opt=/templates/asciidoc.tmpl,example.txt:Ignore*

examples: build examples/proto/*.proto examples/templates/*.tmpl
	$(info Making examples...)
	@rm -f examples/doc/*
	@$(EXAMPLE_CMD) --doc_opt=docbook,example.docbook:Ignore* examples/proto/*.proto
	@$(EXAMPLE_CMD) --doc_opt=html,example.html:Ignore* examples/proto/*.proto
	@$(EXAMPLE_CMD) --doc_opt=json,example.json:Ignore* examples/proto/*.proto
	@$(EXAMPLE_CMD) --doc_opt=markdown,example.md:Ignore* examples/proto/*.proto
	@$(EXAMPLE_CMD) --doc_opt=examples/templates/asciidoc.tmpl,example.txt:Ignore* examples/proto/*.proto
