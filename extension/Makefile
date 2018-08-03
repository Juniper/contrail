## Modify here for your project
PACKAGE_PATH=github.com/Juniper/contrail/extension
PROTO_PACKAGE=github.com.Juniper.contrail.extension
CONTRAIL_PATH=$(GOPATH)/src/github.com/Juniper/contrail

#TODO(nati) remove unused tables from generated sql

generate: reset_gen ## Run the source code generator
	mkdir -p public

	find pkg -type f -name "*.go" -print0 | xargs -0 sed -i -e "s|github.com/Juniper/contrail/extension|$(PACKAGE_PATH)|"
	go run $(CONTRAIL_PATH)/cmd/contrailschema/main.go generate --schemas $(CONTRAIL_PATH)/schemas/abstract,schemas \
		--templates $(CONTRAIL_PATH)/tools/templates/template_config.yaml \
		--package-path "$(PACKAGE_PATH)" \
		--proto-package "$(PROTO_PACKAGE)" \
		--schema-output public/schema.json --openapi-output public/openapi.json

	$(CONTRAIL_PATH)/bin/protoc -I $(CONTRAIL_PATH)/vendor/ \
		-I $(CONTRAIL_PATH)/vendor/github.com/gogo/protobuf/protobuf -I ./proto \
		--gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ \
		proto/$(PACKAGE_PATH)/pkg/models/generated.proto
	$(CONTRAIL_PATH)/bin/protoc -I $(CONTRAIL_PATH)/proto/ \
		-I $(CONTRAIL_PATH)/vendor/ -I $(CONTRAIL_PATH)/vendor/github.com/gogo/protobuf/protobuf \
		-I ./proto --gogo_out=Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,plugins=grpc:$(GOPATH)/src/ \
		proto/$(PACKAGE_PATH)/pkg/services/generated.proto
	go tool fix ./pkg/services/generated.pb.go
	#TODO(nati) renable db test
	rm pkg/db/gen_db_test.go
	go fmt $(PACKAGE_PATH)/pkg/...

reset_gen: ## Remove genarated files
	find pkg/ -name gen_* -delete
	find pkg/ -name generated.pb.go -delete
	rm -rf public/[^watch.html]*
	rm -rf proto/*
	rm -f tools/init_mysql.sql
	rm -f tools/init_psql.sql
	rm -f tools/cleanup.sql

install:
	go install ./cmd/...