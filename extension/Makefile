## Modify here for your project
PACKAGE_PATH=github.com/Juniper/contrail/extension
PROTO_PACKAGE=github.com.Juniper.contrail.extension
CONTRAIL_PATH=$(GOPATH)/src/github.com/Juniper/contrail

#TODO(nati) remove unused tables from generated sql

generate: generate_go pkg/services/generated.pb.go pkg/models/generated.pb.go ## Run the source code generator
	#TODO(nati) renable db test
	rm pkg/db/gen_db_test.go

format_gen:
	find ./cmd ./pkg -name 'gen_*.go' -exec go fmt {} \;

generate_go:
	@mkdir -p public
	go run $(CONTRAIL_PATH)/cmd/contrailschema/main.go generate --schemas $(CONTRAIL_PATH)/schemas/abstract,schemas \
		--templates $(CONTRAIL_PATH)/tools/templates/template_config.yaml \
		--package-path "$(PACKAGE_PATH)" \
		--proto-package "$(PROTO_PACKAGE)" \
		--schema-output public/schema.json --openapi-output public/openapi.json

PROTO := $(CONTRAIL_PATH)/bin/protoc -I $(CONTRAIL_PATH)/proto/ \
		-I $(CONTRAIL_PATH)/vendor/ -I $(CONTRAIL_PATH)/vendor/github.com/gogo/protobuf/protobuf \
		-I ./proto

pkg/%.pb.go: proto/$(PACKAGE_PATH)/pkg/%.proto
	$(PROTO) --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ $<
	go tool fix $@


clean_gen: ## Remove genarated files
	find pkg/ -name gen_* -delete
	find pkg/ -name generated.pb.go -delete
	rm -rf public/[^watch.html]*
	rm -rf proto/*
	rm -f tools/init_mysql.sql
	rm -f tools/init_psql.sql
	rm -f tools/cleanup.sql

install:
	go install ./cmd/...